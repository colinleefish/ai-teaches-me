# Consul Service Mesh

HashiCorp Consul is a service networking solution providing service discovery, health checking, service mesh with mutual TLS, and a distributed key-value store.

## Knowledge Structure

### a. Introduction & Core Concepts
- What is Consul & Service Mesh
- Use cases and scenarios
- Architecture overview
- Setup examples

### b. Service Discovery & Registration
- Agent modes (server vs client)
- Service registration (manual, API, sidecar)
- DNS and HTTP interfaces
- Health checks

### c. Service Mesh & Security
- Connect overview (sidecar proxies)
- Intentions (service-to-service authorization)
- Mutual TLS (mTLS) encryption
- Certificate management

### d. Key-Value Store & Configuration
- KV store operations
- Configuration management
- Watches and blocking queries
- Use cases

### e. Multi-Datacenter & Advanced Features
- WAN federation
- Mesh gateways
- Prepared queries
- ACLs and security policies

### f. Observability & Operations
- Monitoring and metrics
- UI and CLI tools
- Backup and restore
- Production best practices

## Quick Start

```bash
# Start the 3-node cluster
docker-compose up -d

# Check cluster members
docker exec consul-server1 consul members

# Access UI
open http://localhost:8500

# Create a test service
docker exec consul-server1 consul services register -name=web -port=8080
```

## Ports & Endpoints

- **8500**: HTTP API & UI
- **8600**: DNS interface
- **8300**: Server RPC
- **8301**: Serf LAN
- **8302**: Serf WAN
- **21000-21255**: Sidecar proxies
