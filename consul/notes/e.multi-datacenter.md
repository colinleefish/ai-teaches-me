# Multi-Datacenter & Advanced Features

## Multi-Datacenter Overview

Consul natively supports multiple datacenters with WAN federation, enabling global service discovery and communication.

```
┌─────────────────────────────┐     ┌─────────────────────────────┐
│      Datacenter: dc1        │     │      Datacenter: dc2        │
│         (Primary)           │     │       (Secondary)           │
│  ┌────┐  ┌────┐  ┌────┐   │     │  ┌────┐  ┌────┐  ┌────┐   │
│  │ S1 │◄─┤ S2 │─▶│ S3 │   │◄───▶│  │ S1 │◄─┤ S2 │─▶│ S3 │   │
│  └────┘  └────┘  └────┘   │     │  └────┘  └────┘  └────┘   │
│         LAN Gossip         │     │         LAN Gossip         │
│                            │     │                            │
│  Services: web, api, db    │     │  Services: cache, worker   │
└─────────────────────────────┘     └─────────────────────────────┘
              │                                   │
              └────────── WAN Gossip ─────────────┘
```

**Benefits:**
- Disaster recovery and high availability
- Geo-distributed services
- Latency-based routing
- Regional data compliance
- Cross-datacenter service mesh

## WAN Federation (Traditional)

### Configure Primary Datacenter

**consul-server1-dc1.json:**

```json
{
  "datacenter": "dc1",
  "primary_datacenter": "dc1",
  "server": true,
  "bootstrap_expect": 3,
  "retry_join": ["consul-server2-dc1", "consul-server3-dc1"],
  "retry_join_wan": ["consul-server1-dc2"],
  "ports": {
    "serf_wan": 8302
  },
  "bind_addr": "{{ GetInterfaceIP \"eth0\" }}",
  "advertise_addr_wan": "203.0.113.1"
}
```

### Configure Secondary Datacenter

**consul-server1-dc2.json:**

```json
{
  "datacenter": "dc2",
  "primary_datacenter": "dc1",
  "server": true,
  "bootstrap_expect": 3,
  "retry_join": ["consul-server2-dc2", "consul-server3-dc2"],
  "retry_join_wan": ["consul-server1-dc1"],
  "ports": {
    "serf_wan": 8302
  },
  "bind_addr": "{{ GetInterfaceIP \"eth0\" }}",
  "advertise_addr_wan": "198.51.100.1"
}
```

### Join WAN

```bash
# From dc1 server
consul join -wan consul-server1-dc2

# Verify WAN members
consul members -wan

# Output:
# Node                Address              Status  Type    Build   Protocol  DC   Partition  Segment
# consul-server1-dc1  203.0.113.1:8302    alive   server  1.17.0  2         dc1  default    <all>
# consul-server1-dc2  198.51.100.1:8302   alive   server  1.17.0  2         dc2  default    <all>
```

## Mesh Gateways (Modern Approach)

Mesh gateways enable secure cross-datacenter service mesh communication.

### Enable Mesh Gateway

**mesh-gateway.json:**

```json
{
  "service": {
    "name": "mesh-gateway",
    "kind": "mesh-gateway",
    "port": 8443,
    "proxy": {
      "config": {
        "envoy_gateway_bind_addresses": {
          "all-interfaces": {
            "address": "0.0.0.0",
            "port": 8443
          }
        }
      }
    }
  }
}
```

### Run Mesh Gateway

```bash
# Register gateway service
consul services register mesh-gateway.json

# Start Envoy as mesh gateway
consul connect envoy \
  -gateway=mesh \
  -register \
  -service=mesh-gateway \
  -address='{{ GetInterfaceIP "eth0" }}:8443'
```

### Configure Cross-DC Communication

```bash
# Set proxy defaults
consul config write - <<EOF
Kind = "proxy-defaults"
Name = "global"
MeshGateway {
  Mode = "local"
}
EOF

# Service defaults for cross-DC service
consul config write - <<EOF
Kind = "service-defaults"
Name = "api"
Protocol = "http"
MeshGateway {
  Mode = "local"
}
EOF
```

## Cross-Datacenter Queries

### Query Service in Another DC

```bash
# DNS query
dig @localhost -p 8600 api.service.dc2.consul

# HTTP API
curl http://localhost:8500/v1/catalog/service/api?dc=dc2

# CLI
consul catalog services -datacenter=dc2
```

### Connect to Service in Another DC

```json
{
  "service": {
    "name": "web",
    "connect": {
      "sidecar_service": {
        "proxy": {
          "upstreams": [{
            "destination_name": "api",
            "destination_namespace": "default",
            "destination_partition": "default",
            "datacenter": "dc2",
            "local_bind_port": 9001
          }]
        }
      }
    }
  }
}
```

## Prepared Queries

Prepared queries enable advanced service discovery patterns including failover and geo-routing.

### Create Prepared Query

```bash
# Simple query
curl -X POST http://localhost:8500/v1/query \
  -d '{
    "Name": "api-query",
    "Service": {
      "Service": "api",
      "Failover": {
        "NearestN": 3,
        "Datacenters": ["dc2", "dc3"]
      },
      "OnlyPassing": true
    }
  }'
```

### Failover Query

```bash
# Query with automatic failover
curl -X POST http://localhost:8500/v1/query \
  -d '{
    "Name": "critical-service",
    "Service": {
      "Service": "database",
      "Failover": {
        "Datacenters": ["dc2", "dc3"]
      },
      "OnlyPassing": true,
      "Near": "_agent"
    },
    "DNS": {
      "TTL": "10s"
    }
  }'
```

### Geographic Failover

```bash
# Prefer nearest, failover to specific regions
curl -X POST http://localhost:8500/v1/query \
  -d '{
    "Name": "geo-api",
    "Service": {
      "Service": "api",
      "Failover": {
        "Datacenters": ["us-east-1", "us-west-2", "eu-west-1"]
      }
    }
  }'
```

### Execute Prepared Query

```bash
# DNS
dig @localhost -p 8600 api-query.query.consul

# HTTP API
curl http://localhost:8500/v1/query/api-query/execute

# CLI
consul catalog service api-query -query
```

## ACLs (Access Control Lists)

ACLs provide fine-grained security for Consul resources.

### Bootstrap ACL System

```bash
# Bootstrap ACLs (generates master token)
consul acl bootstrap

# Output:
# AccessorID:       12345678-1234-1234-1234-123456789012
# SecretID:         abcdefgh-abcd-abcd-abcd-abcdefghijkl
# Description:      Bootstrap Token (Global Management)
```

### Enable ACLs in Configuration

**acl-config.json:**

```json
{
  "acl": {
    "enabled": true,
    "default_policy": "deny",
    "enable_token_persistence": true,
    "tokens": {
      "master": "abcdefgh-abcd-abcd-abcd-abcdefghijkl"
    }
  }
}
```

### Create ACL Policy

```bash
# Define policy
consul acl policy create \
  -name "read-only" \
  -description "Read-only access to services" \
  -rules @policy.hcl

# policy.hcl
service_prefix "" {
  policy = "read"
}
node_prefix "" {
  policy = "read"
}
```

### Create ACL Token

```bash
# Create token with policy
consul acl token create \
  -description "Service discovery token" \
  -policy-name "read-only"

# Output:
# AccessorID:       87654321-4321-4321-4321-210987654321
# SecretID:         zyxwvuts-zyxw-zyxw-zyxw-zyxwvutsrqpo
```

### Use ACL Token

```bash
# With CLI
consul catalog services -token="zyxwvuts-zyxw-zyxw-zyxw-zyxwvutsrqpo"

# With HTTP API
curl -H "X-Consul-Token: zyxwvuts-zyxw-zyxw-zyxw-zyxwvutsrqpo" \
  http://localhost:8500/v1/catalog/services

# Set default token
export CONSUL_HTTP_TOKEN="zyxwvuts-zyxw-zyxw-zyxw-zyxwvutsrqpo"
```

### Advanced ACL Policies

**service-specific.hcl:**

```hcl
# Allow write to specific service
service "web" {
  policy = "write"
}

# Read all other services
service_prefix "" {
  policy = "read"
}

# Write to specific KV prefix
key_prefix "config/web/" {
  policy = "write"
}

# Read all KV
key_prefix "" {
  policy = "read"
}

# Create intentions for web service
service "web" {
  policy = "write"
  intentions = "write"
}

# Session management
session_prefix "" {
  policy = "write"
}

# Node registration
node_prefix "" {
  policy = "write"
}
```

## Namespaces (Enterprise)

Namespaces provide multi-tenancy and resource isolation.

```bash
# Create namespace
consul namespace create -name "team-a"

# Register service in namespace
consul services register -namespace="team-a" service.json

# Query service in namespace
consul catalog service web -namespace="team-a"

# Cross-namespace intentions
consul intention create \
  -source "team-a/web" \
  -destination "team-b/api" \
  -allow
```

## Hands-On Lab

### Exercise 1: Simulate Multi-DC Setup

```bash
# Create dc2 docker-compose
cat > docker-compose-dc2.yml << EOF
version: '3.8'
networks:
  consul-dc2:
    driver: bridge

services:
  consul-server1-dc2:
    image: hashicorp/consul:latest
    container_name: consul-server1-dc2
    networks:
      - consul-dc2
    ports:
      - "8510:8500"
    command: >
      agent -server -ui
      -bootstrap-expect=1
      -datacenter=dc2
      -node=consul-server1-dc2
      -bind=0.0.0.0
      -client=0.0.0.0
EOF

docker-compose -f docker-compose-dc2.yml up -d

# Join WANs
docker exec consul-server1 \
  consul join -wan $(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' consul-server1-dc2)

# Verify
docker exec consul-server1 consul members -wan
```

### Exercise 2: Cross-DC Service Discovery

```bash
# Register service in dc1
docker exec consul-server1 sh -c 'cat > /tmp/web-dc1.json << EOF
{
  "service": {
    "name": "web",
    "tags": ["dc1"],
    "port": 8080
  }
}
EOF'
docker exec consul-server1 consul services register /tmp/web-dc1.json

# Register service in dc2
docker exec consul-server1-dc2 sh -c 'cat > /tmp/web-dc2.json << EOF
{
  "service": {
    "name": "web",
    "tags": ["dc2"],
    "port": 8080
  }
}
EOF'
docker exec consul-server1-dc2 consul services register /tmp/web-dc2.json

# Query from dc1
docker exec consul-server1 consul catalog service web -datacenter=dc1
docker exec consul-server1 consul catalog service web -datacenter=dc2
```

### Exercise 3: Prepared Query with Failover

```bash
# Create prepared query
docker exec consul-server1 sh -c "
curl -X POST http://localhost:8500/v1/query -d '{
  \"Name\": \"web-failover\",
  \"Service\": {
    \"Service\": \"web\",
    \"Failover\": {
      \"Datacenters\": [\"dc2\"]
    },
    \"OnlyPassing\": true
  }
}'
"

# Execute query
docker exec consul-server1 \
  curl -s http://localhost:8500/v1/query/web-failover/execute | jq

# Test DNS
docker exec consul-server1 \
  dig @localhost -p 8600 web-failover.query.consul
```

### Exercise 4: ACL Setup

```bash
# Bootstrap ACLs
MASTER_TOKEN=$(docker exec consul-server1 \
  consul acl bootstrap -format=json | jq -r '.SecretID')

echo "Master Token: $MASTER_TOKEN"

# Create read-only policy
docker exec consul-server1 sh -c "cat > /tmp/read-policy.hcl << EOF
service_prefix \"\" {
  policy = \"read\"
}
node_prefix \"\" {
  policy = \"read\"
}
key_prefix \"\" {
  policy = \"read\"
}
EOF"

docker exec consul-server1 \
  consul acl policy create \
  -token="$MASTER_TOKEN" \
  -name="read-only" \
  -rules=@/tmp/read-policy.hcl

# Create read-only token
READ_TOKEN=$(docker exec consul-server1 \
  consul acl token create \
  -token="$MASTER_TOKEN" \
  -description="Read-only token" \
  -policy-name="read-only" \
  -format=json | jq -r '.SecretID')

echo "Read Token: $READ_TOKEN"

# Test with token
docker exec consul-server1 \
  consul catalog services -token="$READ_TOKEN"

# Try to register without write permissions (should fail)
docker exec consul-server1 sh -c "
  consul services register \
  -token='$READ_TOKEN' \
  /tmp/web-dc1.json 2>&1
"
```

## Best Practices

1. **Primary datacenter**: Designate one DC as primary for ACL/config replication
2. **Mesh gateways**: Use mesh gateways instead of WAN join for service mesh
3. **Prepared queries**: Use for automatic failover and geo-routing
4. **ACL default deny**: Set `default_policy = "deny"` for security
5. **Namespace isolation**: Use namespaces (Enterprise) for team separation
6. **Token rotation**: Regularly rotate ACL tokens
7. **Least privilege**: Grant minimum necessary permissions
8. **Monitor replication**: Watch ACL and config replication lag
9. **Network segmentation**: Use separate networks for WAN traffic
10. **Disaster recovery**: Test DC failover procedures regularly

## Troubleshooting

```bash
# Check WAN members
consul members -wan

# Check replication status
consul operator raft list-peers -datacenter=dc1
consul operator raft list-peers -datacenter=dc2

# Debug ACL issues
consul acl token read -self -token="$TOKEN"

# Check mesh gateway status
consul catalog service mesh-gateway

# View intentions
consul intention list

# Verify certificates for cross-DC
curl http://localhost:8500/v1/connect/ca/roots
```


