# Observability & Operations

## Monitoring & Metrics

Consul exposes metrics in multiple formats for comprehensive observability.

### Metrics Endpoints

```bash
# Prometheus format
curl http://localhost:8500/v1/agent/metrics?format=prometheus

# JSON format
curl http://localhost:8500/v1/agent/metrics

# Sample output:
{
  "Timestamp": "2025-11-30 12:00:00 +0000 UTC",
  "Gauges": [
    {"Name": "consul.serf.member.status", "Value": 1},
    {"Name": "consul.raft.state.leader", "Value": 1},
    {"Name": "consul.runtime.alloc_bytes", "Value": 12345678}
  ],
  "Counters": [...],
  "Samples": [...]
}
```

### Key Metrics to Monitor

**Raft Consensus:**
- `consul.raft.leader` - Is this node the leader? (1 or 0)
- `consul.raft.leader.lastContact` - Time since last contact with followers
- `consul.raft.commitTime` - Time to commit log entries
- `consul.raft.replication.appendEntries` - Raft replication rate

**Health:**
- `consul.health.service.query-tag` - Service health query timing
- `consul.health.service.not-found` - Failed health checks

**KV Store:**
- `consul.kvs.apply` - KV write operations
- `consul.txn.apply` - Transaction operations

**Network:**
- `consul.serf.member.flap` - Network instability
- `consul.memberlist.gossip` - Gossip protocol metrics

**Client:**
- `consul.client.rpc` - RPC call rate
- `consul.client.rpc.failed` - Failed RPC calls

### Prometheus Integration

**prometheus.yml:**

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'consul-servers'
    metrics_path: '/v1/agent/metrics'
    params:
      format: ['prometheus']
    static_configs:
      - targets:
        - 'consul-server1:8500'
        - 'consul-server2:8500'
        - 'consul-server3:8500'

  - job_name: 'consul-services'
    consul_sd_configs:
      - server: 'consul-server1:8500'
        datacenter: 'dc1'
    relabel_configs:
      - source_labels: [__meta_consul_service]
        target_label: service
      - source_labels: [__meta_consul_node]
        target_label: node
```

### Grafana Dashboards

Import official Consul dashboard:

```bash
# Dashboard ID: 12049 (HashiCorp Consul)
# Dashboard ID: 13396 (Consul Exporter)

# Or import JSON from:
# https://grafana.com/grafana/dashboards/12049
```

**Key panels:**
- Cluster size and health
- Raft leader elections
- Service registration rate
- KV store operations
- Network latency
- Memory and CPU usage

## Logging

### Log Levels

```bash
# Start with specific log level
consul agent -log-level=DEBUG

# Available levels: TRACE, DEBUG, INFO, WARN, ERR
```

### JSON Logging

```bash
# Enable JSON logging
consul agent -log-json

# Output:
{"@level":"info","@message":"Consul agent running!","@timestamp":"2025-11-30T12:00:00.000000Z"}
```

### Log to File

```bash
# Log to file
consul agent -log-file=/var/log/consul/consul.log

# With rotation
consul agent \
  -log-file=/var/log/consul/consul.log \
  -log-rotate-duration=24h \
  -log-rotate-max-files=7
```

### Important Log Patterns

**Health Issues:**
```
[WARN] agent: Check is now critical: check=service:web-1
[ERR] agent: failed to sync remote state: error="connection refused"
```

**Raft Issues:**
```
[WARN] raft: Heartbeat timeout reached, starting election
[INFO] raft: Node at 10.0.1.5:8300 [Follower] entering Leader state
```

**Network Issues:**
```
[WARN] memberlist: Was able to connect to consul-server2 but other probes failed
[ERR] agent: Coordinate update error: error="No cluster leader"
```

## Consul UI

The web UI provides visual monitoring and management.

**Access:** http://localhost:8500/ui

**Features:**
- Service topology view
- Health check status
- Node list and details
- KV store browser
- Intentions management
- ACL token management

### Service Topology

```
Web UI → Services → [Service Name] → Topology

Visual graph showing:
- Upstream dependencies
- Downstream consumers  
- Health status
- Intention rules
- Metrics integration
```

## CLI Tools

### Health Checks

```bash
# Check cluster health
consul operator raft list-peers

# Check specific service
consul health service web

# Watch service health
consul watch -type=health -service=web -state=critical \
  sh -c 'echo "Service critical!" | mail -s Alert ops@example.com'
```

### Debug Commands

```bash
# Collect debug bundle
consul debug -duration=30s -interval=5s -output=debug.tar.gz

# View member status
consul members -detailed

# Check raft status
consul operator raft list-peers
consul operator raft read-config

# Verify service connectivity
consul intention check web api
```

### Monitor Commands

```bash
# Stream logs from agent
consul monitor -log-level=debug

# Watch KV changes
consul watch -type=key -key=config/app/feature

# Watch service changes  
consul watch -type=service -service=web

# Watch health checks
consul watch -type=checks -service=web
```

## Telemetry Integration

### Datadog

```bash
# Install Datadog agent with Consul integration

# Enable telemetry in Consul
consul agent \
  -config-file=consul.json \
  -telemetry-dogstatsd-addr=127.0.0.1:8125
```

**consul.json:**

```json
{
  "telemetry": {
    "dogstatsd_addr": "127.0.0.1:8125",
    "dogstatsd_tags": ["env:production", "datacenter:dc1"]
  }
}
```

### StatsD

```json
{
  "telemetry": {
    "statsd_address": "127.0.0.1:8125",
    "metrics_prefix": "consul"
  }
}
```

### Circonus

```json
{
  "telemetry": {
    "circonus_api_token": "...",
    "circonus_api_app": "consul",
    "circonus_submission_url": "https://..."
  }
}
```

## Backup & Restore

### Snapshot Backup

```bash
# Create snapshot
consul snapshot save backup.snap

# With ACL token
consul snapshot save -token="$CONSUL_HTTP_TOKEN" backup.snap

# Automated backup script
#!/bin/bash
DATE=$(date +%Y%m%d-%H%M%S)
consul snapshot save /backups/consul-$DATE.snap
find /backups -name "consul-*.snap" -mtime +7 -delete
```

### Snapshot Restore

```bash
# Restore from snapshot
consul snapshot restore backup.snap

# Force restore (dangerous)
consul snapshot restore -force backup.snap
```

### KV Backup

```bash
# Export KV store
consul kv export > kv-backup.json

# Restore KV store
consul kv import @kv-backup.json

# Export specific prefix
consul kv export config/ > config-backup.json
```

### Automated Backup

```bash
# Cron job for daily backups
0 2 * * * /usr/local/bin/consul-backup.sh

# consul-backup.sh
#!/bin/bash
set -e

BACKUP_DIR="/var/backups/consul"
DATE=$(date +\%Y\%m\%d)
CONSUL_TOKEN="your-token"

# Create backup directory
mkdir -p $BACKUP_DIR

# Snapshot backup
consul snapshot save \
  -token="$CONSUL_TOKEN" \
  "$BACKUP_DIR/snapshot-$DATE.snap"

# KV backup
consul kv export \
  -token="$CONSUL_TOKEN" \
  > "$BACKUP_DIR/kv-$DATE.json"

# Upload to S3
aws s3 cp "$BACKUP_DIR/snapshot-$DATE.snap" \
  s3://my-consul-backups/

# Cleanup old backups (keep 30 days)
find $BACKUP_DIR -name "*.snap" -mtime +30 -delete
find $BACKUP_DIR -name "*.json" -mtime +30 -delete
```

## Disaster Recovery

### Leader Election Failure

```bash
# Check raft status
consul operator raft list-peers

# If quorum lost, bootstrap new cluster
consul operator raft bootstrap -recover -datacenter=dc1
```

### Split Brain Recovery

```bash
# Identify partition
consul members -detailed | grep -v alive

# Rejoin nodes
consul join <leader-ip>

# Force remove dead nodes
consul force-leave <node-name>
```

### Data Corruption

```bash
# Restore from snapshot
consul snapshot restore backup.snap

# Verify cluster health
consul operator raft list-peers
consul members
```

## Hands-On Lab

### Exercise 1: Metrics Collection

```bash
# Configure Prometheus
cat > monitoring/prometheus.yml << EOF
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'consul'
    metrics_path: '/v1/agent/metrics'
    params:
      format: ['prometheus']
    static_configs:
      - targets:
        - 'consul-server1:8500'
        - 'consul-server2:8500'
        - 'consul-server3:8500'
EOF

# Start monitoring stack
docker-compose up -d prometheus grafana

# Query metrics
curl http://localhost:8500/v1/agent/metrics?format=prometheus | grep consul_raft

# Access Grafana
open http://localhost:3000
# Login: admin/admin
# Add Prometheus datasource: http://prometheus:9090
```

### Exercise 2: Health Monitoring

```bash
# Register service with health check
docker exec consul-server1 sh -c 'cat > /tmp/monitored-service.json << EOF
{
  "service": {
    "name": "api",
    "port": 8080,
    "checks": [
      {
        "http": "http://localhost:8080/health",
        "interval": "10s",
        "timeout": "1s"
      },
      {
        "tcp": "localhost:8080",
        "interval": "5s"
      }
    ]
  }
}
EOF'

docker exec consul-server1 consul services register /tmp/monitored-service.json

# Watch health status
docker exec consul-server1 consul watch -type=checks -service=api

# Simulate failure
docker exec consul-server1 sh -c 'pkill -f "nc.*8080"'

# Check critical services
docker exec consul-server1 \
  curl -s http://localhost:8500/v1/health/state/critical | jq
```

### Exercise 3: Backup & Restore

```bash
# Create test data
docker exec consul-server1 consul kv put test/key1 "value1"
docker exec consul-server1 consul kv put test/key2 "value2"

# Create snapshot
docker exec consul-server1 consul snapshot save /tmp/backup.snap

# Export to host
docker cp consul-server1:/tmp/backup.snap ./backup.snap

# Delete data
docker exec consul-server1 consul kv delete -recurse test/

# Verify deleted
docker exec consul-server1 consul kv get test/key1
# Error: key not found

# Restore snapshot
docker cp ./backup.snap consul-server1:/tmp/backup.snap
docker exec consul-server1 consul snapshot restore /tmp/backup.snap

# Verify restored
docker exec consul-server1 consul kv get test/key1
# Output: value1
```

### Exercise 4: Debug Bundle Collection

```bash
# Collect debug information
docker exec consul-server1 consul debug \
  -duration=30s \
  -interval=5s \
  -output=/tmp/debug.tar.gz

# Export bundle
docker cp consul-server1:/tmp/debug.tar.gz ./consul-debug.tar.gz

# Extract and inspect
tar -xzf consul-debug.tar.gz
cd consul-debug-*

# Contents:
# - agent.json (agent configuration)
# - host.json (host information)
# - members.json (cluster members)
# - metrics.json (telemetry)
# - profile.prof (CPU profile)
# - goroutine.prof (goroutine profile)
```

## Production Best Practices

### Deployment

1. **Use 3 or 5 servers**: Odd numbers for raft quorum
2. **Separate server and client**: Dedicated server nodes
3. **Resource allocation**: 2-4 CPU cores, 4-8GB RAM per server
4. **Persistent storage**: Use volumes for data directories
5. **Network security**: Firewall, TLS encryption

### Monitoring

1. **Alert on leader changes**: Frequent elections indicate issues
2. **Monitor raft lag**: High commitTime indicates performance issues
3. **Track service health**: Alert on critical services
4. **Watch memory usage**: Prevent OOM kills
5. **Network monitoring**: Track gossip and RPC latency

### Operations

1. **Automated backups**: Daily snapshots to S3/GCS
2. **Rolling updates**: Update one server at a time
3. **Capacity planning**: Monitor growth trends
4. **Disaster recovery plan**: Test restore procedures
5. **Documentation**: Maintain runbooks for common issues

### Security

1. **Enable ACLs**: Default deny policy
2. **Enable TLS**: Encrypt all traffic
3. **Rotate tokens**: Regular credential rotation
4. **Audit logs**: Track configuration changes
5. **Network segmentation**: Isolate Consul traffic

## Troubleshooting Guide

### High CPU Usage

```bash
# Collect CPU profile
consul debug -duration=60s

# Check for:
# - Excessive log writes
# - High query rate
# - Large KV values
# - Gossip storms
```

### Memory Leaks

```bash
# Collect heap profile
consul debug -duration=60s

# Check for:
# - Large number of registered services
# - KV store growth
# - Connection leaks
```

### Network Issues

```bash
# Check connectivity
consul members -detailed

# Test ports
nc -zv consul-server1 8300
nc -zv consul-server1 8301
nc -zv consul-server1 8500

# Check iptables/firewall
iptables -L -n
```

### Raft Issues

```bash
# Check raft state
consul operator raft list-peers

# View raft configuration
consul operator raft read-config

# Check logs for raft errors
consul monitor | grep raft
```

## Useful Dashboards

### Grafana Queries

**Service Health:**
```promql
consul_health_service_query_tag{status="passing"}
```

**Raft Performance:**
```promql
consul_raft_commitTime{quantile="0.99"}
```

**Memory Usage:**
```promql
consul_runtime_alloc_bytes / 1024 / 1024
```

**Active Services:**
```promql
count(consul_health_service_query_tag)
```

