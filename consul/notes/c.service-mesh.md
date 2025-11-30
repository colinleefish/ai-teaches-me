# Service Mesh & Security

## Consul Connect Overview

Consul Connect is the built-in service mesh feature providing secure service-to-service communication with automatic mutual TLS (mTLS) encryption, authorization, and traffic management.

```
Without Connect:
┌─────────┐  HTTP   ┌─────────┐
│Service A│────────▶│Service B│
└─────────┘         └─────────┘
   ❌ No encryption
   ❌ No authentication
   ❌ Manual endpoints

With Connect:
┌─────────┐         ┌─────────┐
│Service A│         │Service B│
└────┬────┘         └────┬────┘
     │ localhost          │ localhost
     ▼                    ▼
┌─────────┐  mTLS   ┌─────────┐
│ Proxy A │────────▶│ Proxy B │
└─────────┘ verify  └─────────┘
   ✅ Automatic mTLS
   ✅ Identity-based auth
   ✅ Dynamic discovery
```

## Enabling Connect

### Enable Connect in Service Definition

```json
{
  "service": {
    "name": "web",
    "port": 8080,
    "connect": {
      "sidecar_service": {
        "port": 21000,
        "proxy": {
          "upstreams": [
            {
              "destination_name": "database",
              "local_bind_port": 5432
            }
          ]
        }
      }
    }
  }
}
```

### Register Connect-Enabled Service

```bash
# Register web service with sidecar
consul services register web-with-connect.json

# Verify registration
consul catalog service web
consul catalog service web-sidecar-proxy
```

## Sidecar Proxies

Consul Connect uses Envoy as the default sidecar proxy.

### Native Sidecar (Built-in)

```bash
# Start service
./my-app &

# Start sidecar proxy
consul connect proxy \
  -sidecar-for web-1 \
  -listen 0.0.0.0:21000
```

### Envoy Sidecar (Recommended)

```bash
# Start Envoy sidecar
consul connect envoy \
  -sidecar-for web-1 \
  -admin-bind 127.0.0.1:19000
```

### Envoy in Docker

```yaml
services:
  web:
    image: my-web-app
    network_mode: "service:web-sidecar"
  
  web-sidecar:
    image: envoyproxy/envoy-alpine:v1.28-latest
    command:
      - consul
      - connect
      - envoy
      - -sidecar-for=web
    environment:
      - CONSUL_HTTP_ADDR=consul-server1:8500
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

## Intentions (Service Authorization)

Intentions control which services can communicate. They're deny-by-default once enabled.

### Create Intentions via CLI

```bash
# Allow web to call api
consul intention create -allow web api

# Deny web to call database
consul intention create -deny web database

# List intentions
consul intention list

# Check if connection allowed
consul intention check web api
# Allowed: true

consul intention check web database
# Denied
```

### Intentions via API

```bash
# Create allow intention
curl -X PUT http://localhost:8500/v1/connect/intentions \
  -d '{
    "SourceName": "web",
    "DestinationName": "api",
    "Action": "allow"
  }'

# Create deny intention
curl -X PUT http://localhost:8500/v1/connect/intentions \
  -d '{
    "SourceName": "web",
    "DestinationName": "database",
    "Action": "deny"
  }'

# List all intentions
curl http://localhost:8500/v1/connect/intentions
```

### Layer 7 Intentions (Advanced)

```bash
# Allow only GET requests from web to api
consul config write - <<EOF
Kind = "service-intentions"
Name = "api"
Sources = [
  {
    Name = "web"
    Permissions = [
      {
        Action = "allow"
        HTTP {
          PathPrefix = "/public"
          Methods = ["GET"]
        }
      }
    ]
  }
]
EOF
```

## Mutual TLS (mTLS)

Consul automatically manages certificates for all Connect-enabled services.

### Certificate Management

```bash
# View CA configuration
consul connect ca get-config

# Get CA root certificate
curl http://localhost:8500/v1/connect/ca/roots

# Get service leaf certificate
curl http://localhost:8500/v1/agent/connect/ca/leaf/web
```

### Certificate Rotation

```bash
# Certificates auto-rotate before expiry (default 72 hours)
# Force rotation:
consul connect ca set-config -config-file ca.json
```

### Custom CA Integration

```json
{
  "Provider": "vault",
  "Config": {
    "Address": "https://vault.example.com",
    "Token": "s.abc123...",
    "RootPKIPath": "connect-root",
    "IntermediatePKIPath": "connect-intermediate"
  }
}
```

```bash
consul connect ca set-config -config-file vault-ca.json
```

## Upstream Services

Configure which services your service needs to call:

```json
{
  "service": {
    "name": "web",
    "port": 8080,
    "connect": {
      "sidecar_service": {
        "proxy": {
          "upstreams": [
            {
              "destination_name": "api",
              "local_bind_port": 9001
            },
            {
              "destination_name": "cache",
              "local_bind_port": 9002,
              "datacenter": "dc2"
            }
          ]
        }
      }
    }
  }
}
```

Application connects to localhost:
- `localhost:9001` → routes to `api` service via mTLS
- `localhost:9002` → routes to `cache` service in dc2 via mTLS

## Traffic Management

### Service Defaults

```bash
# Set protocol for service
consul config write - <<EOF
Kind = "service-defaults"
Name = "api"
Protocol = "http"
EOF
```

### Service Splitter (Traffic Splitting)

```bash
# Split traffic 90/10 between versions
consul config write - <<EOF
Kind = "service-splitter"
Name = "api"
Splits = [
  {
    Weight = 90
    ServiceSubset = "v1"
  },
  {
    Weight = 10
    ServiceSubset = "v2"
  }
]
EOF
```

### Service Router (Advanced Routing)

```bash
# Route based on HTTP path
consul config write - <<EOF
Kind = "service-router"
Name = "api"
Routes = [
  {
    Match {
      HTTP {
        PathPrefix = "/v2"
      }
    }
    Destination {
      ServiceSubset = "v2"
    }
  }
]
EOF
```

### Service Resolver (Subsets)

```bash
# Define service subsets by tag
consul config write - <<EOF
Kind = "service-resolver"
Name = "api"
Subsets = {
  "v1" = {
    Filter = "Service.Tags contains \"v1\""
  }
  "v2" = {
    Filter = "Service.Tags contains \"v2\""
  }
}
EOF
```

## Hands-On Lab

### Exercise 1: Set Up Service Mesh

```bash
# Start Consul cluster
docker-compose up -d

# Create web service with Connect
docker exec consul-server1 sh -c 'cat > /tmp/web.json << EOF
{
  "service": {
    "name": "web",
    "id": "web-1",
    "port": 8080,
    "connect": {
      "sidecar_service": {
        "proxy": {
          "upstreams": [{
            "destination_name": "api",
            "local_bind_port": 9001
          }]
        }
      }
    }
  }
}
EOF'

docker exec consul-server1 consul services register /tmp/web.json

# Create api service with Connect
docker exec consul-server1 sh -c 'cat > /tmp/api.json << EOF
{
  "service": {
    "name": "api",
    "id": "api-1",
    "port": 9000,
    "connect": {
      "sidecar_service": {}
    }
  }
}
EOF'

docker exec consul-server1 consul services register /tmp/api.json

# Verify services registered
docker exec consul-server1 consul catalog services
```

### Exercise 2: Test Intentions

```bash
# Without intention (deny by default)
docker exec consul-server1 consul intention check web api
# Denied (implicit)

# Allow web to call api
docker exec consul-server1 consul intention create -allow web api

# Verify
docker exec consul-server1 consul intention check web api
# Allowed: true

# List all intentions
docker exec consul-server1 consul intention list

# Delete intention
docker exec consul-server1 consul intention delete web api
```

### Exercise 3: Full Service Mesh Demo

```bash
# Start simple API server
docker run -d --name api-service \
  --network consul_consul-network \
  hashicorp/http-echo \
  -text="API v1 Response" \
  -listen=:9000

# Start Envoy sidecar for API
docker run -d --name api-sidecar \
  --network container:api-service \
  envoyproxy/envoy-alpine:v1.28-latest \
  consul connect envoy \
  -sidecar-for api-1 \
  -http-addr=consul-server1:8500

# Create intention
docker exec consul-server1 consul intention create -allow web api

# Test connection through proxy
docker exec consul-server1 curl localhost:9001
# Returns: API v1 Response
```

### Exercise 4: Traffic Splitting

```bash
# Register two versions of API
docker exec consul-server1 sh -c 'cat > /tmp/api-v1.json << EOF
{
  "service": {
    "name": "api",
    "id": "api-v1",
    "port": 9001,
    "tags": ["v1"],
    "connect": { "sidecar_service": {} }
  }
}
EOF'

docker exec consul-server1 sh -c 'cat > /tmp/api-v2.json << EOF
{
  "service": {
    "name": "api",
    "id": "api-v2",
    "port": 9002,
    "tags": ["v2"],
    "connect": { "sidecar_service": {} }
  }
}
EOF'

docker exec consul-server1 consul services register /tmp/api-v1.json
docker exec consul-server1 consul services register /tmp/api-v2.json

# Create traffic split
docker exec consul-server1 consul config write - <<EOF
Kind = "service-resolver"
Name = "api"
Subsets = {
  "v1" = { Filter = "Service.Tags contains \"v1\"" }
  "v2" = { Filter = "Service.Tags contains \"v2\"" }
}
EOF

docker exec consul-server1 consul config write - <<EOF
Kind = "service-splitter"
Name = "api"
Splits = [
  { Weight = 80, ServiceSubset = "v1" },
  { Weight = 20, ServiceSubset = "v2" }
]
EOF

# Test - 80% goes to v1, 20% to v2
for i in {1..10}; do
  docker exec consul-server1 curl -s localhost:9001
done
```

## Security Best Practices

1. **Enable intentions immediately**: Deny-by-default after Connect is enabled
2. **Use Layer 7 intentions**: Fine-grained control based on HTTP paths/methods
3. **Regular cert rotation**: Ensure automatic rotation is working
4. **Integrate with external CA**: Use Vault for enterprise PKI
5. **Monitor proxy metrics**: Watch Envoy metrics for connection issues
6. **Use service mesh gradually**: Start with critical services first
7. **Test intention changes**: Verify before deploying to production

## Troubleshooting

```bash
# Check proxy status
consul connect proxy -sidecar-for web-1 -log-level debug

# View Envoy admin interface
curl localhost:19000/clusters
curl localhost:19000/config_dump

# Verify certificates
curl http://localhost:8500/v1/agent/connect/ca/leaf/web | jq

# Debug intentions
consul intention check -verbose web api

# View proxy logs
docker logs api-sidecar
```


