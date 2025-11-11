# Kafka Configuration Guide

## KRaft Mode Configurations

### `KAFKA_NODE_ID`
- **Purpose**: Unique identifier for this broker/node
- **Value**: Integer (1, 2, 3, etc.)
- **Impact**: Must be unique across all nodes in cluster
- **Example**: `KAFKA_NODE_ID: 1`

### `KAFKA_PROCESS_ROLES`
- **Purpose**: Defines what roles this node performs
- **Values**:
  - `broker` - Handles client requests and stores data
  - `controller` - Manages cluster metadata and leader elections
  - `broker,controller` - Combined mode (both roles)
- **Impact**: Combined mode = simpler for single node, separate = better for large clusters
- **Example**: `KAFKA_PROCESS_ROLES: 'broker,controller'`

### `KAFKA_CONTROLLER_QUORUM_VOTERS`
- **Purpose**: List of controller nodes that participate in consensus/voting
- **Format**: `NODE_ID@HOST:PORT,NODE_ID@HOST:PORT,...`
- **Impact**: Defines the controller quorum for leader election
- **Single node**: `'1@kafka:9093'`
- **3-node cluster**: `'1@kafka1:9093,2@kafka2:9093,3@kafka3:9093'`

## Network Configurations

### `KAFKA_LISTENERS`
- **Purpose**: Network interfaces Kafka binds to internally
- **Format**: `LISTENER_NAME://HOST:PORT,LISTENER_NAME://HOST:PORT`
- **Impact**: Defines what interfaces/ports Kafka listens on
- **Example**: `'PLAINTEXT://0.0.0.0:9092,CONTROLLER://0.0.0.0:9093'`
  - `0.0.0.0` = listen on all network interfaces
  - Port 9092 = client traffic
  - Port 9093 = controller traffic (KRaft)

### `KAFKA_ADVERTISED_LISTENERS`
- **Purpose**: Address that clients use to connect to Kafka
- **Impact**: Must be reachable from client's perspective
- **Example**: `'PLAINTEXT://localhost:9092'`
- **Important**: 
  - Use `localhost` for local dev only
  - Use actual hostname/IP for production
  - If wrong, clients won't be able to connect

### `KAFKA_CONTROLLER_LISTENER_NAMES`
- **Purpose**: Declares which listener is for controller-to-controller traffic
- **Value**: Must match a listener name from `KAFKA_LISTENERS`
- **Example**: `'CONTROLLER'`

### `KAFKA_LISTENER_SECURITY_PROTOCOL_MAP`
- **Purpose**: Maps listener names to security protocols
- **Format**: `LISTENER_NAME:PROTOCOL,LISTENER_NAME:PROTOCOL`
- **Protocols**:
  - `PLAINTEXT` - No encryption (dev/testing)
  - `SSL` - TLS encryption
  - `SASL_PLAINTEXT` - Authentication without encryption
  - `SASL_SSL` - Authentication + encryption (production)
- **Example**: `'CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT'`

## Replication Configurations

### `KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR`
- **Purpose**: How many replicas of consumer offset data (`__consumer_offsets` topic)
- **Impact**: Consumer offset durability
- **Values**:
  - `1` = single node (dev/testing)
  - `3` = production cluster (recommended)
- **Example**: `KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1`

### `KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR`
- **Purpose**: Replication factor for transaction state log
- **Impact**: Transactional guarantee durability
- **Values**:
  - `1` = single node
  - `3` = production (minimum recommended)
- **Example**: `KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1`

### `KAFKA_TRANSACTION_STATE_LOG_MIN_ISR`
- **Purpose**: Minimum In-Sync Replicas for transaction log
- **Impact**: How many replicas must acknowledge before transaction commits
- **Values**:
  - `1` = single node (less safe)
  - `2` = production cluster (ensures data safety)
- **Example**: `KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1`

## Storage Configuration

### `KAFKA_LOG_DIRS`
- **Purpose**: Directory where Kafka stores topic data and partitions
- **Impact**: All message data is persisted here
- **Example**: `KAFKA_LOG_DIRS: '/tmp/kraft-combined-logs'`
- **Important**: 
  - Should be on fast disk (SSD)
  - Mount to host for persistence
  - Monitor disk space

### `CLUSTER_ID`
- **Purpose**: Unique cluster identifier (base64-encoded UUID)
- **Impact**: 
  - Required for KRaft mode
  - Prevents accidental cluster merges
  - All nodes in same cluster must have same ID
- **Example**: `CLUSTER_ID: 'MkU3OEVBNTcwNTJENDM2Qk'`
- **Generate**: `kafka-storage random-uuid`

## Important Configuration Concepts

### Cluster Sizing - Odd Numbers Required

**Controllers need odd numbers (3, 5, 7) for Raft consensus:**

| Nodes | Quorum Needed | Tolerate Failures | Recommended |
|-------|---------------|-------------------|-------------|
| 1     | 1             | 0 failures        | Dev only |
| 2     | 2             | 0 failures ❌     | Don't use! |
| 3     | 2             | 1 failure ✅      | Standard production |
| 4     | 3             | 1 failure (waste!)| Don't use! |
| 5     | 3             | 2 failures ✅     | Large production |

**Why odd?** Even numbers provide same fault tolerance as N-1 but cost more.

**Bad setup example:**
```yaml
node_1: broker,controller  # Only 1 controller = single point of failure!
node_2: broker
node_3: broker
```
If node_1 dies, entire cluster fails. **Always use ≥3 controllers.**

### Combined vs Separated Modes

**Combined mode** (most common):
```yaml
KAFKA_PROCESS_ROLES: 'broker,controller'
```
- Each node does both data + metadata
- Simpler setup
- 3 nodes = 3 brokers + 3 controllers
- Good for small-medium clusters

**Separated mode** (large scale):
```yaml
# 3 dedicated controller nodes
KAFKA_PROCESS_ROLES: 'controller'

# 10+ dedicated broker nodes (different machines)
KAFKA_PROCESS_ROLES: 'broker'
```
- Controllers only manage metadata
- Brokers only handle data
- Better for large-scale production (50+ nodes)

### KAFKA_CONTROLLER_QUORUM_VOTERS - Configure Everywhere

**Must be identical on ALL nodes** (brokers and controllers need to know the voter list):

```yaml
# Node 1
KAFKA_NODE_ID: 1
KAFKA_CONTROLLER_QUORUM_VOTERS: '1@kafka1:9093,2@kafka2:9093,3@kafka3:9093'

# Node 2
KAFKA_NODE_ID: 2
KAFKA_CONTROLLER_QUORUM_VOTERS: '1@kafka1:9093,2@kafka2:9093,3@kafka3:9093'  # Same!

# Node 3
KAFKA_NODE_ID: 3
KAFKA_CONTROLLER_QUORUM_VOTERS: '1@kafka1:9093,2@kafka2:9093,3@kafka3:9093'  # Same!
```

**Why?**
- Controllers use it for Raft consensus
- Brokers use it to find controllers for metadata

### Listener Names Are Arbitrary

**Names** = labels you choose  
**Protocols** = actual security mechanisms

```yaml
# Standard naming
KAFKA_LISTENERS: 'PLAINTEXT://0.0.0.0:9092,CONTROLLER://0.0.0.0:9093'
KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: 'PLAINTEXT:PLAINTEXT,CONTROLLER:PLAINTEXT'

# Custom naming (equally valid!)
KAFKA_LISTENERS: 'BANANAS://0.0.0.0:9092,IAMYOURDADDY://0.0.0.0:9093'
KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: 'BANANAS:SSL,IAMYOURDADDY:PLAINTEXT'
```

**Available security protocols:**
- `PLAINTEXT` - No encryption, no auth (dev only)
- `SSL` - TLS encryption, optional client certs
- `SASL_PLAINTEXT` - Auth but no encryption (rarely used)
- `SASL_SSL` - Auth + encryption (production standard)

### LISTENERS vs ADVERTISED_LISTENERS

**Key difference:**
- `KAFKA_LISTENERS` = What Kafka **binds to** (server-side)
- `KAFKA_ADVERTISED_LISTENERS` = What Kafka **tells clients** (client-side)

**Why both needed?**

**Flow:**
1. Client connects to bootstrap server (any broker)
2. Client asks: "Give me cluster metadata"
3. Kafka responds with **advertised addresses** of all brokers
4. Client connects to **advertised addresses** for actual data transfer

**Docker example:**
```yaml
KAFKA_LISTENERS: 'PLAINTEXT://0.0.0.0:9092'           # Bind inside container
KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://localhost:9092'  # Clients use localhost
```

**Common mistake:**
```yaml
KAFKA_ADVERTISED_LISTENERS: 'PLAINTEXT://0.0.0.0:9092'  # ❌ Clients can't reach 0.0.0.0!
```

**Multi-network example:**
```yaml
# Internal + External access
KAFKA_LISTENERS: 'INTERNAL://0.0.0.0:9092,EXTERNAL://0.0.0.0:9093'
KAFKA_ADVERTISED_LISTENERS: 'INTERNAL://10.0.1.5:9092,EXTERNAL://203.0.113.5:9093'
```

### Client Discovery & Connection

**Bootstrap process:**
1. Client configured with 1+ bootstrap servers (for redundancy)
2. Client connects to **any** working bootstrap server
3. Client fetches metadata: broker list + partition leaders
4. Client connects **only to brokers hosting needed partitions**

**Example:**
```
Topic has 3 partitions across 3 brokers:
- Partition 0 leader: broker1
- Partition 1 leader: broker2
- Partition 2 leader: broker3

Producer writing to partition 1 → connects to broker2 only
Consumer reading partitions 0,2 → connects to broker1 & broker3 only
```

**Not all connections active** - only to brokers with relevant partitions.

## Production vs Development Settings

| Config | Development | Production |
|--------|-------------|------------|
| Replication Factors | 1 | 3 |
| Min ISR | 1 | 2 |
| Security Protocol | PLAINTEXT | SASL_SSL |
| Advertised Listeners | localhost | actual hostname/IP |
| Process Roles | broker,controller | separate roles |
| Log Dirs | /tmp or local | dedicated disk/volume |
| Number of Nodes | 1 | 3 (or 5) |

