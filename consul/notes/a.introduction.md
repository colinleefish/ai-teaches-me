# Consul Introduction

## What is Consul

HashiCorp Consul is a service networking platform that provides service discovery, configuration, and segmentation functionality. It's a distributed, highly available system with built-in service mesh capabilities for microservices architectures.

**Core Concepts:**

- **Service Discovery**: Automatically detect and register services in your infrastructure
- **Health Checking**: Monitor service health and route traffic only to healthy instances
- **Service Mesh**: Secure service-to-service communication with automatic TLS encryption
- **Key-Value Store**: Distributed configuration and metadata storage
- **Multi-Datacenter**: Native support for multiple datacenters and regions

## What is a Service Mesh

A service mesh is a dedicated infrastructure layer for handling service-to-service communication. It provides:

- **Traffic Management**: Load balancing, routing, retries
- **Security**: mTLS, authentication, authorization between services
- **Observability**: Metrics, logs, distributed tracing
- **Resilience**: Circuit breakers, timeouts, retries

```
Without Service Mesh:
┌─────────┐  HTTP   ┌─────────┐
│Service A│────────▶│Service B│
└─────────┘         └─────────┘
  - Manual TLS setup
  - Hard-coded endpoints
  - No automatic retries

With Consul Service Mesh:
┌─────────┐         ┌─────────┐
│Service A│         │Service B│
└────┬────┘         └────┬────┘
     │                   │
     ▼                   ▼
┌─────────┐  mTLS  ┌─────────┐
│ Proxy A │───────▶│ Proxy B │
└─────────┘         └─────────┘
  - Automatic mTLS
  - Service discovery
  - Health checking
  - Traffic management
```

## Scenarios to Use Consul

1. **Microservices Service Discovery**: Replace hard-coded service URLs with dynamic discovery
2. **Service Mesh for Security**: Add mTLS encryption between all services without code changes
3. **Multi-Cloud Networking**: Connect services across AWS, GCP, Azure, on-prem
4. **Dynamic Configuration**: Store and distribute configuration across distributed systems
5. **Health Monitoring**: Automatic health checks and failover
6. **API Gateway Integration**: Backend for service routing and load balancing
7. **Zero-Trust Networking**: Enforce service-to-service authorization policies

## Architecture Overview

### Consul Components

```
┌────────────────────────────────────────────────────────┐
│                   Consul Cluster                       │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐            │
│  │ Server 1 │  │ Server 2 │  │ Server 3 │ (Leaders)  │
│  │ (Leader) │  │          │  │          │            │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘            │
│       │             │             │                   │
│       └─────────────┴─────────────┘                   │
│              Raft Consensus                           │
└───────────────────┬────────────────────────────────────┘
                    │
        ┌───────────┼───────────┐
        │           │           │
    ┌───▼───┐   ┌───▼───┐   ┌───▼───┐
    │Client │   │Client │   │Client │  (Agents)
    │Agent 1│   │Agent 2│   │Agent 3│
    └───┬───┘   └───┬───┘   └───┬───┘
        │           │           │
    ┌───▼───┐   ┌───▼───┐   ┌───▼───┐
    │Service│   │Service│   │Service│  (Applications)
    │   A   │   │   B   │   │   C   │
    └───────┘   └───────┘   └───────┘
```

### Server vs Client Agents

**Server Agents:**
- Maintain cluster state (Raft consensus)
- Handle queries and persist data
- Elect a leader
- Typically 3-5 servers per datacenter

**Client Agents:**
- Forward requests to servers
- Run health checks
- Register local services
- Lightweight and scalable

### Data Flow Example

```
1. Service Registration:
   App ──register──▶ Client Agent ──▶ Server Cluster

2. Service Discovery:
   App ──query──▶ Client Agent ──▶ Server ──▶ return endpoints

3. Service Mesh Communication:
   App A ──▶ Sidecar Proxy A ──mTLS──▶ Sidecar Proxy B ──▶ App B
```

## Setup Architectures

### Development: Single Server

```
┌─────────────────────────┐
│  Consul Server (dev)    │
│  - All features enabled │
│  - Non-production       │
└─────────────────────────┘

docker run -d --name=consul \
  -p 8500:8500 \
  hashicorp/consul:latest agent -dev -ui -client=0.0.0.0
```

### Production: Multi-Server Cluster

```
                 Load Balancer
                      │
         ┌────────────┼────────────┐
         │            │            │
    ┌────▼───┐   ┌────▼───┐   ┌────▼───┐
    │Server 1│◄─▶│Server 2│◄─▶│Server 3│
    │(Leader)│   │        │   │        │
    └────────┘   └────────┘   └────────┘
         │            │            │
    ┌────┴────────────┴────────────┴────┐
    │          Client Agents             │
    │   ┌──────┐  ┌──────┐  ┌──────┐   │
    │   │Node 1│  │Node 2│  │Node N│   │
    │   └──────┘  └──────┘  └──────┘   │
    └────────────────────────────────────┘
```

### Multi-Datacenter Setup

```
    Datacenter 1 (Primary)           Datacenter 2 (Secondary)
┌──────────────────────────┐    ┌──────────────────────────┐
│  ┌────┐  ┌────┐  ┌────┐ │    │  ┌────┐  ┌────┐  ┌────┐ │
│  │ S1 │◄▶│ S2 │◄▶│ S3 │ │◄──▶│  │ S1 │◄▶│ S2 │◄▶│ S3 │ │
│  └────┘  └────┘  └────┘ │    │  └────┘  └────┘  └────┘ │
│          WAN             │    │          WAN             │
└──────────────────────────┘    └──────────────────────────┘
          │                                  │
      Services                           Services
```

## Key Ports

- **8500**: HTTP API and Web UI
- **8600**: DNS interface (UDP/TCP)
- **8300**: Server RPC (cluster communication)
- **8301**: Serf LAN (gossip within datacenter)
- **8302**: Serf WAN (gossip between datacenters)
- **8443**: HTTPS API (when TLS enabled)
- **21000-21255**: Envoy proxy ports for service mesh

## Quick Commands

```bash
# Start dev server
consul agent -dev -ui

# Check cluster members
consul members

# Register a service
consul services register service.json

# Query DNS
dig @localhost -p 8600 web.service.consul

# Check service health
consul catalog services
consul health service web

# KV operations
consul kv put config/app/database "postgres://..."
consul kv get config/app/database

# Intentions (service ACLs)
consul intention create -allow web api
consul intention check web api
```

## Consul vs Alternatives

| Feature | Consul | Kubernetes | Istio | Eureka |
|---------|--------|------------|-------|--------|
| Service Discovery | ✓ | ✓ | ✓ | ✓ |
| Health Checks | ✓ | ✓ | ✓ | ✓ |
| Service Mesh | ✓ | - | ✓ | - |
| Multi-DC | ✓ Native | Complex | Limited | - |
| KV Store | ✓ | ✓ (etcd) | - | - |
| Platform Agnostic | ✓ | - | K8s only | JVM only |
| DNS Interface | ✓ | ✓ | - | - |


