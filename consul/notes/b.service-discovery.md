# Service Discovery & Registration

## Service Discovery Overview

Service discovery allows applications to find and connect to services without hard-coded hostnames or IP addresses. Consul provides multiple interfaces for service discovery.

```
Traditional Approach:
┌─────────┐
│ App A   │──▶ http://10.0.1.5:8080  (Hard-coded IP)
└─────────┘

Consul Service Discovery:
┌─────────┐
│ App A   │──▶ web.service.consul  (Dynamic lookup)
└─────────┘          │
                     ▼
              ┌──────────────┐
              │ Consul finds │
              │ healthy      │
              │ instances    │
              └──────────────┘
                     │
         ┌───────────┴───────────┐
         ▼                       ▼
    10.0.1.5:8080          10.0.1.6:8080
```

## Service Registration

### Method 1: Configuration File

Create `/consul/config/web.json`:

```json
{
  "service": {
    "name": "web",
    "tags": ["v1", "production"],
    "port": 8080,
    "check": {
      "http": "http://localhost:8080/health",
      "interval": "10s",
      "timeout": "1s"
    }
  }
}
```

Load it when starting Consul:

```bash
consul agent -config-dir=/consul/config
```

### Method 2: API Registration

```bash
# Register service via API
curl -X PUT http://localhost:8500/v1/agent/service/register \
  -d '{
    "name": "web",
    "id": "web-1",
    "port": 8080,
    "tags": ["production", "v1"],
    "address": "192.168.1.10",
    "check": {
      "http": "http://192.168.1.10:8080/health",
      "interval": "10s"
    }
  }'

# Deregister service
curl -X PUT http://localhost:8500/v1/agent/service/deregister/web-1
```

### Method 3: CLI Registration

```bash
# Create service definition file
cat > web-service.json << EOF
{
  "Name": "web",
  "Port": 8080,
  "Tags": ["production"],
  "Check": {
    "HTTP": "http://localhost:8080/health",
    "Interval": "10s"
  }
}
EOF

# Register using CLI
consul services register web-service.json

# List all services
consul catalog services

# List specific service instances
consul catalog service web
```

## Service Discovery Methods

### 1. DNS Interface

Consul provides a DNS server on port 8600.

```bash
# Query service by name
dig @localhost -p 8600 web.service.consul

# Output:
# web.service.consul. 0 IN A 10.0.1.5
# web.service.consul. 0 IN A 10.0.1.6

# Query with SRV records (includes port)
dig @localhost -p 8600 web.service.consul SRV

# Output:
# web.service.consul. 0 IN SRV 1 1 8080 node1.node.dc1.consul.

# Query by tag
dig @localhost -p 8600 production.web.service.consul

# Prepared query
dig @localhost -p 8600 nearest.web.query.consul
```

**Configure system DNS:**

```bash
# /etc/resolv.conf
nameserver 127.0.0.1
options ndots:1
```

### 2. HTTP API

```bash
# Get all healthy instances
curl http://localhost:8500/v1/health/service/web?passing

# Response:
[
  {
    "Node": {
      "Node": "node1",
      "Address": "10.0.1.5"
    },
    "Service": {
      "ID": "web-1",
      "Service": "web",
      "Port": 8080,
      "Address": "10.0.1.5"
    },
    "Checks": [
      {
        "Status": "passing",
        "Output": "HTTP GET http://10.0.1.5:8080/health: 200 OK"
      }
    ]
  }
]

# Get service with specific tag
curl http://localhost:8500/v1/health/service/web?tag=production

# Get service in specific datacenter
curl http://localhost:8500/v1/health/service/web?dc=dc2
```

### 3. Consul Template

Dynamic configuration file generation:

```bash
# Install consul-template
consul-template \
  -template "nginx.conf.ctmpl:nginx.conf:nginx -s reload"
```

**nginx.conf.ctmpl:**

```nginx
upstream backend {
{{range service "web"}}
  server {{.Address}}:{{.Port}};
{{end}}
}

server {
  location / {
    proxy_pass http://backend;
  }
}
```

## Health Checks

### HTTP Health Check

```json
{
  "check": {
    "http": "http://localhost:8080/health",
    "interval": "10s",
    "timeout": "1s"
  }
}
```

### TCP Health Check

```json
{
  "check": {
    "tcp": "localhost:8080",
    "interval": "10s",
    "timeout": "1s"
  }
}
```

### Script Health Check

```json
{
  "check": {
    "args": ["/bin/check-service.sh"],
    "interval": "30s",
    "timeout": "5s"
  }
}
```

### TTL Health Check (Application-driven)

```json
{
  "check": {
    "ttl": "30s",
    "deregister_critical_service_after": "90s"
  }
}
```

Application updates status:

```bash
# Mark as passing
curl -X PUT http://localhost:8500/v1/agent/check/pass/service:web-1

# Mark as failing
curl -X PUT http://localhost:8500/v1/agent/check/fail/service:web-1
```

### Docker Health Check

```json
{
  "check": {
    "docker_container_id": "abcd1234",
    "shell": "/bin/sh",
    "args": ["/check.sh"],
    "interval": "10s"
  }
}
```

## Health Check Statuses

- **passing**: Service is healthy
- **warning**: Service degraded but still available
- **critical**: Service is unhealthy

```bash
# Query only passing services
curl http://localhost:8500/v1/health/service/web?passing

# Query with warnings included
curl http://localhost:8500/v1/health/service/web?passing&warning
```

## Hands-On Lab

### Exercise 1: Register Multiple Service Instances

```bash
# Start Consul cluster
docker-compose up -d

# Register instance 1
docker exec consul-server1 sh -c 'cat > /tmp/web1.json << EOF
{
  "service": {
    "name": "web",
    "id": "web-1",
    "port": 8080,
    "tags": ["v1"],
    "check": {
      "tcp": "localhost:8080",
      "interval": "5s"
    }
  }
}
EOF'

docker exec consul-server1 consul services register /tmp/web1.json

# Register instance 2
docker exec consul-server1 sh -c 'cat > /tmp/web2.json << EOF
{
  "service": {
    "name": "web",
    "id": "web-2",
    "port": 8081,
    "tags": ["v2"],
    "check": {
      "tcp": "localhost:8081",
      "interval": "5s"
    }
  }
}
EOF'

docker exec consul-server1 consul services register /tmp/web2.json

# Query all instances
docker exec consul-server1 consul catalog service web

# DNS lookup
docker exec consul-server1 dig @localhost -p 8600 web.service.consul
```

### Exercise 2: Test Health Checks

```bash
# Start a simple HTTP server
docker exec -d consul-client1 sh -c \
  'while true; do echo -e "HTTP/1.1 200 OK\n\n$(date)" | nc -l -p 8080; done'

# Register with health check
docker exec consul-client1 sh -c 'cat > /tmp/api.json << EOF
{
  "service": {
    "name": "api",
    "port": 8080,
    "check": {
      "http": "http://localhost:8080",
      "interval": "5s"
    }
  }
}
EOF'

docker exec consul-client1 consul services register /tmp/api.json

# Watch health status
docker exec consul-client1 consul watch -type=service -service=api

# Kill the server and watch it go critical
docker exec consul-client1 pkill nc

# Check health status
curl http://localhost:8500/v1/health/service/api
```

### Exercise 3: Service Discovery via DNS

```bash
# Configure container to use Consul DNS
docker exec consul-client1 sh -c \
  'echo "nameserver 127.0.0.1" > /etc/resolv.conf'

# Test DNS resolution
docker exec consul-client1 getent hosts web.service.consul

# Use in application
docker exec consul-client1 wget -O- http://web.service.consul:8080
```

## Service Tags and Metadata

```json
{
  "service": {
    "name": "web",
    "tags": [
      "production",
      "v2.1.0",
      "us-west-1"
    ],
    "meta": {
      "version": "2.1.0",
      "git-sha": "abc123",
      "team": "platform"
    }
  }
}
```

Query by tag:

```bash
# DNS with tag filter
dig @localhost -p 8600 production.web.service.consul

# HTTP API with tag
curl http://localhost:8500/v1/catalog/service/web?tag=production
```

## Best Practices

1. **Use health checks**: Always define health checks for services
2. **Unique service IDs**: Use unique IDs when registering multiple instances
3. **Semantic tags**: Use tags for versions, environments, regions
4. **Graceful deregistration**: Deregister services on shutdown
5. **Check intervals**: Balance between responsiveness and overhead (10-30s typical)
6. **DNS caching**: Be aware of DNS TTLs when using DNS interface
7. **Metadata**: Use service metadata for additional context without affecting queries


