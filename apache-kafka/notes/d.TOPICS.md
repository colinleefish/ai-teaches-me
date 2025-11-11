# Kafka Topics & Partitions

## What's a Topic?

**Topic** = Logical category/feed for messages

- Similar to a table in a database or folder for files
- Examples: "user-clicks", "orders", "payment-events", "application-logs"
- Messages are published to topics
- Consumers subscribe to topics to read messages

**Key characteristics:**

Every topic has:

1. **Name** - Unique identifier (e.g., "user-clicks")
2. **Partitions** - Physical subdivisions for parallelism (1 or more)
3. **Replication Factor** - Number of copies across brokers for durability
4. **Retention Policy** - How long to keep messages (time or size based)

**Example:**

```bash
Topic: "user-clicks"
Partitions: 3
Replication Factor: 3
Retention: 7 days
```

## Creating a Topic

```bash
docker exec kafka1 kafka-topics --create \
  --topic user-clicks \
  --bootstrap-server localhost:9092 \
  --partitions 3 \
  --replication-factor 3
```

## 5 Core Concepts to Understand Topics

### 1. Partitions (THE key concept)

- Topics are split into partitions for parallelism
- Each partition = ordered, immutable sequence of messages
- Partitions distributed across brokers
- More partitions = more throughput
- **Most important concept in Kafka**

### 2. Message Structure

Every message has:

- **Key** (optional) - Used to determine which partition (same key → same partition)
- **Value** - The actual data/payload
- **Headers** - Metadata (key-value pairs)
- **Timestamp** - When message was produced
- Messages with same key always go to same partition (ordering guarantee)

### 3. Ordering Guarantees

- ✅ **Guaranteed within a partition** - Messages in partition 0 are strictly ordered
- ❌ **NOT guaranteed across partitions** - Partition 0 and Partition 1 have independent ordering
- **Want strict order?** Use same key for related messages OR use single partition

### 4. Replication & Leadership

- Each partition has copies (replicas) across multiple brokers
- **Leader replica** - Handles all reads and writes for the partition
- **Follower replicas** - Replicate data for failover
- If leader dies, a follower automatically becomes the new leader
- Replication factor of 3 = 1 leader + 2 followers

### 5. Offsets (Message Position)

- Each message has a sequential ID per partition (0, 1, 2, 3, ...)
- **Offset** = position of message in partition
- Consumers track their position via offsets
- Can replay/rewind by changing offset
- Offsets are partition-specific, not topic-wide

## Partitions Deep Dive

### What is a Partition?

**Partition** = Ordered, immutable sequence of messages (like an append-only log)

```
Topic: "user-clicks" with 3 partitions

Partition 0: [msg0] → [msg1] → [msg2] → [msg3] → ...
Partition 1: [msg0] → [msg1] → [msg2] → ...
Partition 2: [msg0] → [msg1] → ...
```

**Key characteristics:**

- Each partition is **independent**
- Messages appended to end only (immutable)
- Each message gets sequential offset within its partition
- Partitions distributed across brokers

### Why Partitions?

**1. Parallelism**

- Multiple consumers can read from different partitions simultaneously
- 3 partitions = 3 consumers can work in parallel

**2. Scalability**

- Spread data across multiple brokers
- Each broker handles subset of partitions

**3. Ordering within streams**

- Messages with same key go to same partition
- Ordering guaranteed within partition

**4. Fault tolerance**

- Partitions replicated across brokers
- If one broker fails, other brokers have copies

### How Messages are Assigned to Partitions

**Producer chooses partition based on:**

**1. With Key** (most common):

```
hash(key) % num_partitions = partition_id

Examples:
- user_id=123 → always partition 1
- user_id=456 → always partition 2
- Same user's events stay in order!
```

**2. Round-robin** (no key):

```
Messages distributed evenly across partitions
- msg1 → partition 0
- msg2 → partition 1
- msg3 → partition 2
- msg4 → partition 0
```

**3. Custom partitioner**:

```
Write your own logic to determine partition
```

### Partition Distribution Example

**3 brokers, 1 topic with 6 partitions, replication factor 3:**

```
Broker 1:
  - Partition 0 (Leader)
  - Partition 1 (Follower)
  - Partition 3 (Follower)
  - Partition 4 (Leader)

Broker 2:
  - Partition 1 (Leader)
  - Partition 2 (Follower)
  - Partition 4 (Follower)
  - Partition 5 (Leader)

Broker 3:
  - Partition 0 (Follower)
  - Partition 2 (Leader)
  - Partition 3 (Leader)
  - Partition 5 (Follower)
```

Kafka automatically balances leader partitions across brokers.

### How Many Partitions Should You Use?

**Factors to consider:**

**Too few partitions:**

- Limited parallelism
- Lower throughput
- Harder to scale

**Too many partitions:**

- More overhead (file handles, memory)
- Slower leader election on failures
- More end-to-end latency

**Guidelines:**

- **Start small**: 3-10 partitions per topic
- **Throughput-based**: Want 100 MB/s? Each partition ≈ 10 MB/s → need 10 partitions
- **Consumer parallelism**: Max concurrent consumers = num partitions (more partitions = more consumers)
- **Production**: 10-30 partitions common
- **Max**: ~2000 partitions per broker (Kafka limit)

**Example calculation:**

```
Target throughput: 300 MB/s
Per-partition throughput: 10 MB/s
Partitions needed: 300 / 10 = 30 partitions
```

### Partition Files on Disk

Each partition stored as directory with log segments:

```
logs/kafka1/
  user-clicks-0/                        ← Partition 0
    00000000000000000000.log            ← Segment 0 (actual messages)
    00000000000000000000.index          ← Index for fast lookup
    00000000000000000000.timeindex      ← Time-based index
    00000000000000001000.log            ← Segment 1 (after segment size reached)
    leader-epoch-checkpoint
    partition.metadata
```

**Segments**:

- Partition split into segments (default 1GB)
- Older segments can be deleted based on retention
- Only latest segment is "active" for writes

### Important Notes

**Cannot reduce partitions**:

- Can increase: 3 → 6 partitions ✅
- Cannot decrease: 6 → 3 partitions ❌
- Choose wisely at creation!

**Partition = unit of parallelism**:

- 3 partitions = max 3 consumers in same consumer group
- 4th consumer would be idle

**Cross-partition ordering impossible**:

- If you need strict ordering across all messages, use 1 partition
- Trade-off: ordering vs throughput

## Replication Deep Dive

### How Replication Works

**Followers continuously pull from the leader** (like consumers):

```
Leader (Broker1)
  ↓ (follower fetches via Fetch API)
Follower (Broker2) - pulls new messages, writes to disk
  ↓ (follower fetches via Fetch API)
Follower (Broker3) - pulls new messages, writes to disk
```

**Process:**

1. Producer sends message to partition leader
2. Leader writes to its local log
3. Followers continuously fetch from leader (using same Fetch API as consumers)
4. Followers write messages to their local logs
5. Leader tracks which followers are caught up
6. Once Min ISR replicas have the data, leader ACKs producer

**Key difference from consumers:**

- Followers replicate by writing to disk
- Regular consumers just process data

### ISR (In-Sync Replicas)

**ISR** = The set of replicas currently caught up with the leader

**Example:**

```
Partition 3: RF=3, replicas on [Broker1, Broker2, Broker3]
Leader: Broker2

Normal state (all caught up):
ISR = [Broker2, Broker1, Broker3]

Broker3 falls behind:
ISR = [Broker2, Broker1]  ← Kafka removes Broker3 from ISR

Broker3 catches up:
ISR = [Broker2, Broker1, Broker3]  ← Kafka adds Broker3 back
```

**Kafka dynamically tracks ISR:**

- Adds replicas when they catch up
- Removes replicas when they fall behind (configurable threshold)

### Min ISR (Minimum In-Sync Replicas)

**Configuration:** `min.insync.replicas=2`

**Meaning:** At least 2 replicas must acknowledge before write succeeds

**With RF=3, Min ISR=2:**

```
Normal: 3 replicas in-sync
- Producer writes → Leader + 1 follower ACK → Success ✅

1 broker down: 2 replicas in-sync
- Producer writes → Leader + 1 follower ACK → Success ✅

2 brokers down: Only 1 replica (leader)
- Producer writes → Only leader ACK → FAIL ❌
- Writes BLOCKED until another replica comes back
```

**Why block writes?**

- **Safety over availability**
- Prevents data loss if leader crashes after accepting write with only 1 copy
- Forces cluster repair before accepting more data

**Trade-offs:**

| Min ISR | Durability          | Availability                         | Production Use |
| ------- | ------------------- | ------------------------------------ | -------------- |
| 1       | Lower (only leader) | Higher (tolerates 2 broker failures) | ❌ Risky       |
| 2       | **Higher** ✅       | Lower (tolerates 1 broker failure)   | ✅ Standard    |
| 3       | Highest             | Lowest (no failures tolerated)       | Rarely used    |

**Production standard:** RF=3, Min ISR=2 (balances safety + availability)

### Complete Message Flow Example

**Cluster:** 3 nodes, Topic: user-clicks, 6 partitions, RF=3

```
Partition 0: Leader=Node1, Followers=[Node2, Node3]
Partition 1: Leader=Node2, Followers=[Node1, Node3]
Partition 2: Leader=Node3, Followers=[Node1, Node2]
Partition 3: Leader=Node2, Followers=[Node1, Node3]
Partition 4: Leader=Node1, Followers=[Node2, Node3]
Partition 5: Leader=Node3, Followers=[Node1, Node2]
```

**When message arrives for Partition 3:**

1. **Producer discovery:**

   - Producer asks any broker: "Who leads partition 3?"
   - Broker responds: "Node2"
   - Producer caches this metadata

2. **Producer sends message:**

   - Producer connects directly to Node2
   - Node2 (leader) writes to local log

3. **Replication:**

   - Node1 (follower) pulls from Node2, writes to disk
   - Node3 (follower) pulls from Node2, writes to disk

4. **Acknowledgment:**

   - Node2 waits for Min ISR=2 replicas to ACK
   - Node2 + Node1 have written → ACK producer ✅

5. **Eventually:**
   - Node3 catches up (if it was slow)
   - All 3 replicas in-sync

**Producer always sends to partition leader** - never to followers!

### Network Traffic Considerations

**Inter-broker traffic is significant:**

```
Every message written = replicated to followers
6 partitions × RF3 = each message stored 3 times
2 network copies per message (leader → 2 followers)
```

**Traffic types:**

1. **Replication** (biggest) - followers continuously pulling from leaders
2. **Metadata/Coordination** - ISR updates, leader elections
3. **Client traffic** - producers/consumers

**Kafka optimizations:**

- ✅ **Batching** - followers fetch in batches, not per-message
- ✅ **Compression** - data compressed before network transfer
- ✅ **Zero-copy** - efficient kernel-level transfers
- ✅ **Sequential I/O** - fast disk writes/reads
- ✅ **Page cache** - OS caches hot data in memory

**Throughput impact:**

- Low (1K msg/s): minimal inter-broker traffic
- High (100K msg/s): **significant** network usage

**Production best practices:**

- Put brokers on **same network segment** (1-10Gbps links)
- Same datacenter or availability zone (low latency)
- Avoid cross-datacenter replication (use MirrorMaker for that)
- Monitor network saturation

**Trade-off:** High network traffic is the price of durability and fault tolerance!

### Replication Recovery

**What happens when a follower falls behind?**

```
1. Follower crashes or network issue
2. Leader removes follower from ISR
3. Leader continues accepting writes (if Min ISR still met)
4. Follower comes back online
5. Follower starts pulling missing messages from leader
6. Follower catches up to leader's latest offset
7. Leader adds follower back to ISR
```

**Not lost forever** - continuous catch-up process ensures eventual consistency!

**Throttling:**

- Kafka can throttle replication to avoid overwhelming recovering brokers
- Configurable via `leader.replication.throttled.rate`
