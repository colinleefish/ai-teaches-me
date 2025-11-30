# Key-Value Store & Configuration

## KV Store Overview

Consul's Key-Value (KV) store is a distributed, highly available storage system for configuration, feature flags, coordination, and leader election.

```
┌──────────────────────────────────────┐
│         Consul KV Store              │
│                                      │
│  config/                             │
│  ├── app/                            │
│  │   ├── database/url               │
│  │   ├── database/pool-size         │
│  │   └── features/new-ui            │
│  ├── shared/                         │
│  │   ├── api-keys/stripe            │
│  │   └── limits/rate                │
│  └── service/                        │
│      └── endpoints/                  │
└──────────────────────────────────────┘
```

**Use Cases:**
- Dynamic configuration
- Feature flags
- Service coordination
- Distributed locking
- Secrets management (basic)

## Basic KV Operations

### Set/Put Values

```bash
# Set a simple value
consul kv put config/app/name "MyApp"

# Set with hierarchical key
consul kv put config/app/database/host "postgres.example.com"
consul kv put config/app/database/port "5432"

# Set from file
consul kv put config/app/certificate @cert.pem

# Set with flags
consul kv put -flags=42 config/app/version "1.2.3"

# Set multiple values
consul kv put config/env "production"
consul kv put config/region "us-west-2"
consul kv put config/replicas "3"
```

### Get/Read Values

```bash
# Get single value
consul kv get config/app/name
# Output: MyApp

# Get with metadata
consul kv get -detailed config/app/name

# Get recursively (all keys under path)
consul kv get -recurse config/app/

# Export as JSON
consul kv get -recurse -format=json config/app/ | jq
```

### Delete Values

```bash
# Delete single key
consul kv delete config/app/name

# Delete recursively
consul kv delete -recurse config/app/database/

# Delete with check-and-set
consul kv delete -cas -modify-index=123 config/app/name
```

### List Keys

```bash
# List all keys under path
consul kv get -keys config/app/

# Output:
# config/app/database/host
# config/app/database/port
# config/app/name
```

## HTTP API Operations

### PUT (Create/Update)

```bash
# Set value
curl -X PUT -d 'postgres://...' \
  http://localhost:8500/v1/kv/config/database/url

# Set with flags
curl -X PUT -d '{"value": "MTIzNDU=", "flags": 42}' \
  http://localhost:8500/v1/kv/config/feature-flag
```

### GET (Read)

```bash
# Get single key
curl http://localhost:8500/v1/kv/config/database/url

# Response:
[
  {
    "LockIndex": 0,
    "Key": "config/database/url",
    "Flags": 0,
    "Value": "cG9zdGdyZXM6Ly8uLi4=",  # Base64 encoded
    "CreateIndex": 100,
    "ModifyIndex": 200
  }
]

# Get and decode
curl http://localhost:8500/v1/kv/config/database/url?raw
# Output: postgres://...

# Get recursively
curl http://localhost:8500/v1/kv/config/?recurse

# List keys only
curl http://localhost:8500/v1/kv/config/?keys
```

### DELETE

```bash
# Delete key
curl -X DELETE http://localhost:8500/v1/kv/config/old-key

# Delete recursively
curl -X DELETE http://localhost:8500/v1/kv/config/temp/?recurse

# Check-and-Set delete
curl -X DELETE http://localhost:8500/v1/kv/config/key?cas=123
```

## Atomic Operations (CAS)

Check-and-Set ensures atomic updates using modify indexes.

```bash
# Get current modify index
RESPONSE=$(curl -s http://localhost:8500/v1/kv/config/counter)
MODIFY_INDEX=$(echo $RESPONSE | jq -r '.[0].ModifyIndex')
VALUE=$(echo $RESPONSE | jq -r '.[0].Value' | base64 -d)

# Increment counter
NEW_VALUE=$((VALUE + 1))

# Atomic update with CAS
curl -X PUT -d "$NEW_VALUE" \
  "http://localhost:8500/v1/kv/config/counter?cas=$MODIFY_INDEX"

# Returns: true (success) or false (conflict)
```

## Watches

Watches provide real-time notifications when data changes.

### Watch a Key

```bash
# Watch single key
consul watch -type=key -key=config/app/feature-flag \
  sh -c 'echo "Value changed: $(consul kv get config/app/feature-flag)"'

# Watch runs command whenever key changes
```

### Watch a Key Prefix

```bash
# Watch all keys under prefix
consul watch -type=keyprefix -prefix=config/app/ \
  /usr/local/bin/reload-config.sh
```

### Watch via HTTP (Blocking Query)

```bash
# Initial request
RESPONSE=$(curl -s http://localhost:8500/v1/kv/config/app?index=0)
INDEX=$(echo $RESPONSE | jq -r '.[0].ModifyIndex')

# Blocking query (waits for change)
curl "http://localhost:8500/v1/kv/config/app?index=$INDEX&wait=5m"
# Returns immediately when value changes or after 5min timeout
```

## Sessions & Locking

Distributed locks for coordination and leader election.

### Create Session

```bash
# Create session
SESSION=$(curl -X PUT http://localhost:8500/v1/session/create \
  -d '{
    "Name": "my-lock",
    "TTL": "10s",
    "Behavior": "release",
    "LockDelay": "15s"
  }' | jq -r '.ID')

echo "Session ID: $SESSION"
```

### Acquire Lock

```bash
# Acquire lock
curl -X PUT -d "owner-data" \
  "http://localhost:8500/v1/kv/locks/my-resource?acquire=$SESSION"

# Returns: true (acquired) or false (already locked)
```

### Check Lock Status

```bash
# Get lock info
curl http://localhost:8500/v1/kv/locks/my-resource

# Check session
curl http://localhost:8500/v1/session/info/$SESSION
```

### Release Lock

```bash
# Release lock
curl -X PUT \
  "http://localhost:8500/v1/kv/locks/my-resource?release=$SESSION"

# Destroy session
curl -X PUT http://localhost:8500/v1/session/destroy/$SESSION
```

## Consul Template

Automatically generate config files from KV data.

### Install

```bash
# Download from releases
wget https://releases.hashicorp.com/consul-template/0.34.0/consul-template_0.34.0_linux_amd64.zip
unzip consul-template_0.34.0_linux_amd64.zip
mv consul-template /usr/local/bin/
```

### Template Example

**config.tmpl:**

```go-template
# Application Configuration
APP_NAME={{ key "config/app/name" }}
DATABASE_URL={{ key "config/app/database/url" }}
DATABASE_POOL={{ keyOrDefault "config/app/database/pool" "10" }}

# Feature Flags
{{ range ls "config/features" }}
FEATURE_{{ .Key | toUpper }}={{ .Value }}
{{ end }}

# Services
{{ range service "api" }}
API_SERVER={{ .Address }}:{{ .Port }}
{{ end }}
```

### Run Template

```bash
# One-shot mode
consul-template -template "config.tmpl:config.env" -once

# Daemon mode (watches for changes)
consul-template \
  -template "config.tmpl:config.env:systemctl reload myapp"

# Multiple templates
consul-template \
  -template "nginx.conf.tmpl:nginx.conf:nginx -s reload" \
  -template "app.conf.tmpl:app.conf:systemctl reload app"
```

## Hands-On Lab

### Exercise 1: Basic KV Operations

```bash
# Set configuration values
docker exec consul-server1 consul kv put config/app/name "MyWebApp"
docker exec consul-server1 consul kv put config/app/version "1.0.0"
docker exec consul-server1 consul kv put config/app/debug "true"
docker exec consul-server1 consul kv put config/db/host "postgres.local"
docker exec consul-server1 consul kv put config/db/port "5432"

# Read values
docker exec consul-server1 consul kv get config/app/name

# List all config keys
docker exec consul-server1 consul kv get -keys -recurse config/

# Get all values
docker exec consul-server1 consul kv get -recurse config/

# Export as JSON
docker exec consul-server1 sh -c \
  'consul kv get -recurse -format=json config/ | jq'

# Delete a key
docker exec consul-server1 consul kv delete config/app/debug
```

### Exercise 2: Atomic Operations

```bash
# Set initial counter
docker exec consul-server1 consul kv put counters/visits "0"

# Get current value and index
docker exec consul-server1 sh -c '
  RESPONSE=$(consul kv get -detailed counters/visits)
  echo "$RESPONSE"
'

# Atomic increment using CAS via API
docker exec consul-server1 sh -c '
  # Get current state
  DATA=$(curl -s http://localhost:8500/v1/kv/counters/visits)
  INDEX=$(echo $DATA | jq -r ".[0].ModifyIndex")
  VALUE=$(echo $DATA | jq -r ".[0].Value" | base64 -d)
  
  # Increment
  NEW_VALUE=$((VALUE + 1))
  
  # CAS update
  RESULT=$(curl -s -X PUT -d "$NEW_VALUE" \
    "http://localhost:8500/v1/kv/counters/visits?cas=$INDEX")
  
  echo "CAS Result: $RESULT"
'
```

### Exercise 3: Distributed Lock

```bash
# Create a lock acquisition script
docker exec consul-server1 sh -c 'cat > /tmp/acquire-lock.sh << "SCRIPT"
#!/bin/sh

# Create session
SESSION=$(curl -s -X PUT http://localhost:8500/v1/session/create \
  -d "{\"Name\": \"worker-lock\", \"TTL\": \"30s\"}" | jq -r ".ID")

echo "Session: $SESSION"

# Try to acquire lock
ACQUIRED=$(curl -s -X PUT -d "$HOSTNAME" \
  "http://localhost:8500/v1/kv/locks/job-processor?acquire=$SESSION")

if [ "$ACQUIRED" = "true" ]; then
  echo "Lock acquired! Processing job..."
  sleep 10
  echo "Job complete"
  
  # Release lock
  curl -s -X PUT \
    "http://localhost:8500/v1/kv/locks/job-processor?release=$SESSION"
  echo "Lock released"
else
  echo "Failed to acquire lock"
fi

# Cleanup session
curl -s -X PUT http://localhost:8500/v1/session/destroy/$SESSION > /dev/null
SCRIPT
chmod +x /tmp/acquire-lock.sh'

# Run from multiple nodes simultaneously
docker exec consul-server1 /tmp/acquire-lock.sh &
docker exec consul-server2 /tmp/acquire-lock.sh &
docker exec consul-server3 /tmp/acquire-lock.sh &

# Only one will acquire the lock
```

### Exercise 4: Watch for Changes

```bash
# In terminal 1: Watch a key
docker exec -it consul-server1 sh -c '
  consul watch -type=key -key=config/feature/beta \
    sh -c "echo Feature flag changed: \$(date); consul kv get config/feature/beta"
'

# In terminal 2: Change the value
docker exec consul-server1 consul kv put config/feature/beta "true"
# Watch in terminal 1 triggers

docker exec consul-server1 consul kv put config/feature/beta "false"
# Watch triggers again
```

### Exercise 5: Configuration Management Pattern

```bash
# Set application configuration
docker exec consul-server1 sh -c '
consul kv put config/myapp/database/host "db.example.com"
consul kv put config/myapp/database/port "5432"
consul kv put config/myapp/database/name "myapp_prod"
consul kv put config/myapp/cache/redis "redis.example.com:6379"
consul kv put config/myapp/features/new-ui "true"
consul kv put config/myapp/features/beta-api "false"
'

# Application reads config on startup
docker exec consul-server1 sh -c '
#!/bin/bash
echo "Loading configuration..."

DB_HOST=$(consul kv get config/myapp/database/host)
DB_PORT=$(consul kv get config/myapp/database/port)
DB_NAME=$(consul kv get config/myapp/database/name)
REDIS=$(consul kv get config/myapp/cache/redis)

echo "Database: $DB_HOST:$DB_PORT/$DB_NAME"
echo "Cache: $REDIS"

# Export for application
export DATABASE_URL="postgres://$DB_HOST:$DB_PORT/$DB_NAME"
export REDIS_URL="redis://$REDIS"
'
```

## Best Practices

1. **Use hierarchical keys**: Organize with `/` like a file system
2. **Avoid large values**: KV store not for big data (max 512KB per key)
3. **Use flags for metadata**: Store type/version info in flags field
4. **Watch instead of poll**: Use watches for real-time updates
5. **CAS for atomicity**: Use check-and-set for race-free updates
6. **Session TTLs**: Keep session TTLs reasonable (10-30s)
7. **Lock prefix convention**: Use `locks/` prefix for coordination keys
8. **Secrets warning**: KV store is not encrypted at rest by default
9. **Backup regularly**: Export critical config with `kv export`
10. **Version config changes**: Track who changed what with flags/metadata

## KV Store vs Other Solutions

| Feature | Consul KV | etcd | Redis | ZooKeeper |
|---------|-----------|------|-------|-----------|
| Distributed | ✓ | ✓ | ✓ | ✓ |
| Watch/Subscribe | ✓ | ✓ | ✓ | ✓ |
| Transactions | Limited | ✓ | ✓ | ✓ |
| HTTP API | ✓ | ✓ | - | - |
| Sessions/Locks | ✓ | ✓ | ✓ | ✓ |
| Max Value Size | 512KB | 1MB | 512MB | 1MB |
| Use Case | Service config | K8s backend | Cache/queue | Coordination |

