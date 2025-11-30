# Advanced Topics & Real-World Patterns

## Network Segments (Enterprise)

Network segments allow partitioning agents into isolated gossip pools within a datacenter.

```
┌─────────────────────────────────────────────┐
│           Datacenter: dc1                   │
│  ┌──────────────────┐  ┌─────────────────┐ │
│  │  Segment: alpha  │  │  Segment: beta  │ │
│  │  ┌────┐  ┌────┐ │  │  ┌────┐  ┌────┐ │ │
│  │  │ C1 │  │ C2 │ │  │  │ C3 │  │ C4 │ │ │
│  │  └────┘  └────┘ │  │  └────┘  └────┘ │ │
│  └──────────────────┘  └─────────────────┘ │
│           │                     │           │
│           └────────┬────────────┘           │
│                    │                        │
│          ┌─────────▼────────┐               │
│          │  Server Cluster  │               │
│          └──────────────────┘               │
└─────────────────────────────────────────────┘
```

**Benefits:**
- Reduce gossip overhead in large clusters
- Isolate network traffic
- Support very large datacenters (10,000+ nodes)

## Admin Partitions (Enterprise)

Partitions provide strong isolation between different tenants/teams.

```bash
# Create partition
consul partition create -name "team-a"

# Register service in partition
consul services register \
  -partition="team-a" \
  -namespace="default" \
  service.json

# Cross-partition communication via exported services
consul config write - <<EOF
Kind = "exported-services"
Name = "default"
Partition = "team-a"
Services = [
  {
    Name = "api"
    Consumers = [
      {
        Partition = "team-b"
      }
    ]
  }
]
EOF
```

## Service Mesh Patterns

### Sidecar Pattern

```yaml
# Docker Compose example
services:
  app:
    image: my-app
    network_mode: "service:app-proxy"
    depends_on:
      - app-proxy
  
  app-proxy:
    image: envoyproxy/envoy-alpine
    command: consul connect envoy -sidecar-for app-1
    ports:
      - "8080:8080"
    environment:
      - CONSUL_HTTP_ADDR=consul:8500
```

### API Gateway Pattern

```bash
# Register API gateway
consul config write - <<EOF
Kind = "api-gateway"
Name = "main-gateway"
Listeners = [
  {
    Port = 8080
    Protocol = "http"
    Name = "http-listener"
  }
]
EOF

# Create routes
consul config write - <<EOF
Kind = "http-route"
Name = "api-route"
Parents = [
  {
    Name = "main-gateway"
  }
]
Rules = [
  {
    Services = [
      {
        Name = "api"
      }
    ]
  }
]
EOF
```

### Circuit Breaker Pattern

```bash
# Configure circuit breaker
consul config write - <<EOF
Kind = "service-defaults"
Name = "api"
UpstreamConfig = {
  Defaults = {
    Limits = {
      MaxConnections = 100
      MaxPendingRequests = 50
      MaxConcurrentRequests = 50
    }
    PassiveHealthCheck = {
      MaxFailures = 5
      Interval = "10s"
    }
  }
}
EOF
```

## Kubernetes Integration

### Consul on Kubernetes

```yaml
# consul-values.yaml
global:
  name: consul
  datacenter: dc1
  image: "hashicorp/consul:1.17.0"

server:
  replicas: 3
  storage: 10Gi
  storageClass: fast-ssd

client:
  enabled: true

connectInject:
  enabled: true
  default: true

ui:
  enabled: true
  service:
    type: LoadBalancer
```

```bash
# Install via Helm
helm repo add hashicorp https://helm.releases.hashicorp.com
helm install consul hashicorp/consul -f consul-values.yaml
```

### Service Mesh on Kubernetes

```yaml
# app-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web
spec:
  replicas: 3
  template:
    metadata:
      annotations:
        "consul.hashicorp.com/connect-inject": "true"
        "consul.hashicorp.com/connect-service-upstreams": "api:9001"
    spec:
      containers:
      - name: web
        image: my-web-app
        ports:
        - containerPort: 8080
```

## Consul on AWS/Cloud

### AWS Auto-Join

```json
{
  "retry_join": [
    "provider=aws tag_key=consul-server tag_value=true"
  ],
  "cloud_auto_join": {
    "provider": "aws",
    "tag_key": "consul-server",
    "tag_value": "true"
  }
}
```

### AWS Secrets Manager Integration

```bash
# Store Consul gossip key in Secrets Manager
aws secretsmanager create-secret \
  --name consul-gossip-key \
  --secret-string "$(consul keygen)"

# Retrieve in startup script
GOSSIP_KEY=$(aws secretsmanager get-secret-value \
  --secret-id consul-gossip-key \
  --query SecretString --output text)

consul agent -encrypt="$GOSSIP_KEY" ...
```

### Multi-Region on AWS

```
Region: us-east-1          Region: us-west-2
┌─────────────────┐        ┌─────────────────┐
│  VPC: 10.0.0.0  │        │  VPC: 10.1.0.0  │
│  ┌──────────┐   │        │  ┌──────────┐   │
│  │ Consul   │   │◄──VPN─▶│  │ Consul   │   │
│  │ Cluster  │   │        │  │ Cluster  │   │
│  └──────────┘   │        │  └──────────┘   │
└─────────────────┘        └─────────────────┘
```

## Service Catalog Integration

### Terraform Integration

```hcl
# Configure Consul provider
provider "consul" {
  address    = "consul.example.com:8500"
  datacenter = "dc1"
  token      = var.consul_token
}

# Register service
resource "consul_service" "api" {
  name = "api"
  port = 8080
  tags = ["v1", "production"]

  check {
    http     = "http://api.example.com/health"
    interval = "10s"
    timeout  = "1s"
  }
}

# Create KV entry
resource "consul_keys" "app" {
  key {
    path  = "config/app/database/url"
    value = "postgres://..."
  }
}

# Create intention
resource "consul_intention" "web_to_api" {
  source_name      = "web"
  destination_name = "api"
  action           = "allow"
}
```

### Ansible Integration

```yaml
# playbook.yml
- hosts: app_servers
  tasks:
    - name: Register service with Consul
      consul:
        service_name: web
        service_port: 8080
        tags:
          - production
          - v2
        checks:
          - http: http://localhost:8080/health
            interval: 10s

    - name: Set configuration in KV
      consul_kv:
        key: config/app/feature-flag
        value: "true"
```

## Performance Tuning

### Raft Performance

```json
{
  "performance": {
    "raft_multiplier": 1,
    "leave_drain_time": "5s",
    "rpc_hold_timeout": "7s"
  }
}
```

Lower `raft_multiplier` = faster but more CPU. Range: 1-10

### Network Performance

```json
{
  "performance": {
    "rpc_rate": -1,
    "rpc_max_burst": 1000
  },
  "limits": {
    "http_max_conns_per_client": 200,
    "https_handshake_timeout": "10s",
    "rpc_max_conns_per_client": 100
  }
}
```

### Cache Tuning

```bash
# Client-side caching
consul agent -config-file=consul.json

# consul.json
{
  "cache": {
    "entry_fetch_max_burst": 4,
    "entry_fetch_rate": -1
  }
}
```

## Security Hardening

### TLS Configuration

```bash
# Generate CA
consul tls ca create

# Generate server cert
consul tls cert create -server -dc dc1

# Generate client cert
consul tls cert create -client

# Configure Consul with TLS
consul agent \
  -encrypt="$(consul keygen)" \
  -ca-file=consul-agent-ca.pem \
  -cert-file=dc1-server-consul-0.pem \
  -key-file=dc1-server-consul-0-key.pem \
  -verify-incoming=true \
  -verify-outgoing=true \
  -verify-server-hostname=true
```

**Full TLS config:**

```json
{
  "encrypt": "encryption-key-here",
  "verify_incoming": true,
  "verify_outgoing": true,
  "verify_server_hostname": true,
  "ca_file": "/etc/consul.d/consul-agent-ca.pem",
  "cert_file": "/etc/consul.d/server.pem",
  "key_file": "/etc/consul.d/server-key.pem",
  "ports": {
    "https": 8501,
    "http": -1
  }
}
```

### Auto-Encrypt

Automatically distribute certificates from servers to clients:

```json
{
  "auto_encrypt": {
    "allow_tls": true
  }
}
```

Client config:

```json
{
  "auto_encrypt": {
    "tls": true
  },
  "verify_incoming": false,
  "verify_outgoing": true,
  "verify_server_hostname": true,
  "ca_file": "/etc/consul.d/consul-agent-ca.pem"
}
```

## Real-World Use Cases

### Microservices Platform

```
┌─────────────────────────────────────────────────┐
│              API Gateway                        │
│         (Consul API Gateway)                    │
└──────────────┬──────────────────────────────────┘
               │
    ┌──────────┼──────────┐
    │          │          │
    ▼          ▼          ▼
┌────────┐ ┌────────┐ ┌────────┐
│  Web   │ │  API   │ │ Worker │
│Service │ │Service │ │Service │
└───┬────┘ └───┬────┘ └───┬────┘
    │          │          │
    └──────────┼──────────┘
               ▼
       ┌──────────────┐
       │   Database   │
       │  (via proxy) │
       └──────────────┘

Features:
- Service discovery via Consul DNS
- mTLS between all services
- Intentions for zero-trust
- KV store for configuration
- Health checks and failover
```

### Multi-Cloud Deployment

```
AWS (us-east-1)              Azure (eastus)
┌─────────────────┐          ┌─────────────────┐
│ Consul DC: aws  │◄────────▶│ Consul DC: azure│
│                 │  Mesh    │                 │
│ Services:       │  Gateway │ Services:       │
│ - frontend      │          │ - analytics     │
│ - api           │          │ - ml-service    │
└─────────────────┘          └─────────────────┘

Features:
- Cross-cloud service discovery
- Mesh gateways for secure comms
- Prepared queries for failover
- Centralized configuration
```

### Legacy Integration

```
┌────────────────────────────────────────┐
│         Consul Cluster                 │
└───────────┬────────────────────────────┘
            │
    ┌───────┼───────┐
    │       │       │
    ▼       ▼       ▼
┌────────┐ ┌────────┐ ┌────────┐
│External│ │Consul  │ │ Native │
│Service │ │Connect │ │Consul  │
│(registered│Service │ │Service │
│ manually)│        │ │        │
└────────┘ └────────┘ └────────┘

# Register external service
{
  "service": {
    "name": "legacy-db",
    "address": "db.legacy.com",
    "port": 5432,
    "check": {
      "tcp": "db.legacy.com:5432",
      "interval": "10s"
    }
  }
}
```

## Migration Strategies

### From Eureka to Consul

```python
# Before (Eureka)
eureka_client.register(
    app_name="my-service",
    port=8080
)
instances = eureka_client.get_instances("api-service")

# After (Consul)
import consul

c = consul.Consul()
c.agent.service.register(
    name="my-service",
    service_id="my-service-1",
    port=8080
)

# Service discovery
_, instances = c.health.service("api-service", passing=True)
```

### From etcd to Consul KV

```bash
# Export from etcd
etcdctl get --prefix "/config" > etcd-export.txt

# Import to Consul (with transformation)
cat etcd-export.txt | while read key value; do
  consul kv put "$key" "$value"
done
```

## Troubleshooting Patterns

### Debug Connection Issues

```bash
# Test service connectivity
consul intention check web api

# View proxy config
consul connect proxy -sidecar-for web-1 -admin-bind=:19000 &
curl localhost:19000/config_dump | jq

# Check certificates
curl http://localhost:8500/v1/agent/connect/ca/leaf/web | \
  jq -r '.CertPEM' | openssl x509 -text -noout
```

### Performance Diagnosis

```bash
# CPU profile
consul debug -duration=60s -output=debug.tar.gz

# Check slow queries
consul monitor | grep -i slow

# Network latency test
consul rtt node1 node2
```

### Split Brain Detection

```bash
# Check all servers
for server in server1 server2 server3; do
  echo "=== $server ==="
  consul operator raft list-peers -http-addr=$server:8500
done

# All should show same leader
```

## Production Checklist

### Pre-Deployment

- [ ] Consul version tested in staging
- [ ] TLS certificates generated and distributed
- [ ] ACL policies defined and tested
- [ ] Backup procedure established
- [ ] Monitoring configured (Prometheus/Grafana)
- [ ] Disaster recovery plan documented
- [ ] Network firewall rules configured
- [ ] Resource limits set (memory, connections)

### Deployment

- [ ] Deploy servers first (1 at a time)
- [ ] Verify cluster health between upgrades
- [ ] Deploy clients gradually
- [ ] Monitor metrics during rollout
- [ ] Test service discovery functionality
- [ ] Verify service mesh connectivity
- [ ] Check ACL enforcement working

### Post-Deployment

- [ ] All nodes showing as healthy
- [ ] Leader elected and stable
- [ ] Services registering correctly
- [ ] Health checks passing
- [ ] Metrics flowing to monitoring
- [ ] Backup job successful
- [ ] Documentation updated

## Further Resources

**Official Documentation:**
- https://www.consul.io/docs
- https://learn.hashicorp.com/consul

**Community:**
- https://discuss.hashicorp.com/c/consul
- https://github.com/hashicorp/consul

**Tools:**
- Consul Template: https://github.com/hashicorp/consul-template
- Consul ESM: https://github.com/hashicorp/consul-esm
- Consul K8s: https://github.com/hashicorp/consul-k8s
- Consul Terraform Sync: https://github.com/hashicorp/consul-terraform-sync

**Books & Courses:**
- "Consul: Up and Running" by Luke Kysow
- HashiCorp Learn Consul Tutorials
- Linux Academy Consul Courses


