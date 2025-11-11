# Kafka Producers & Consumers

## Producers

### What is a Producer?

**Producer** = Application that publishes messages to Kafka topics

- Can be any application (web server, IoT device, ETL job, etc.)
- Sends messages to partition leaders
- Handles partitioning logic
- Can batch messages for efficiency

### Producer Message Flow

```
1. Producer creates message (key, value, headers)
2. Serializer converts to bytes
3. Partitioner determines which partition
4. Message added to batch for that partition
5. Sender thread sends batch to partition leader
6. Leader writes & replicates (based on acks config)
7. Leader sends ACK back to producer
```

### Key Producer Configurations

#### `acks` (Acknowledgment Mode)

**Most important config for durability!**

| acks  | Meaning          | Durability     | Latency | Use Case                         |
| ----- | ---------------- | -------------- | ------- | -------------------------------- |
| `0`   | Fire and forget  | Lowest ❌      | Fastest | Metrics, logs (data loss OK)     |
| `1`   | Leader only      | Medium         | Fast    | Balanced (some risk)             |
| `all` | All ISR replicas | **Highest** ✅ | Slowest | Critical data (payments, orders) |

**Examples:**

```python
# acks=0 - No confirmation (fire and forget)
producer = KafkaProducer(acks=0)
producer.send('topic', message)  # Don't wait for any ACK

# acks=1 - Leader only (default)
producer = KafkaProducer(acks=1)
producer.send('topic', message)  # Wait for leader to write

# acks=all - Wait for all ISR replicas
producer = KafkaProducer(acks='all')
producer.send('topic', message)  # Wait for leader + followers
```

**With acks=all + min.insync.replicas=2:**

- Producer waits for leader + at least 1 follower
- Strong durability guarantee
- Higher latency

#### `retries` & `max.in.flight.requests.per.connection`

```
retries=3  # Retry up to 3 times on failure
max.in.flight.requests.per.connection=5  # Max unACKed requests
```

**What these mean:**
- `retries` - How many times to retry a failed send
- `max.in.flight.requests.per.connection` - How many requests can be "in the air" waiting for ACK

**The Problem: Message Reordering**

```python
producer = KafkaProducer(
    retries=3,                              # Retry failed sends
    max.in.flight.requests.per.connection=5  # Allow 5 unACKed requests
)
```

**Reordering scenario:**

```
Time 1: Producer sends msg1 → Kafka (request in-flight)
Time 2: Producer sends msg2 → Kafka (request in-flight)
        # Doesn't wait for msg1 ACK because max.in.flight=5
Time 3: msg1 FAILS (network hiccup) ❌
Time 4: msg2 SUCCEEDS ✅ (written to partition)
Time 5: msg1 RETRIED (retry #1)
Time 6: msg1 SUCCEEDS ✅ (written to partition)

Result in Kafka partition:
  [msg2][msg1]  ← REORDERED! ❌
```

**Why?** msg2 succeeded while msg1 was being retried.

**Solution 1: Strict Ordering (Low Throughput)**

```python
producer = KafkaProducer(
    retries=3,
    max.in.flight.requests.per.connection=1  # Only 1 unACKed request
)
```

**Timeline:**

```
Time 1: Producer sends msg1 → Kafka
Time 2: Producer WAITS (can't send msg2 yet, max.in.flight=1)
Time 3: msg1 FAILS ❌
Time 4: msg1 RETRIED
Time 5: msg1 SUCCEEDS ✅
Time 6: Producer sends msg2 → Kafka
Time 7: msg2 SUCCEEDS ✅

Result: [msg1][msg2]  ← Correct order! ✅
```

**Trade-off:** Lower throughput (waits for each ACK)

**Solution 2: Idempotent Producer (Best!)**

```python
producer = KafkaProducer(
    enable_idempotence=True,  # Handles ordering + deduplication
    retries=3,
    max.in.flight.requests.per.connection=5  # Can still use 5!
)
```

**How it works:**
- Kafka assigns sequence numbers to messages
- Even if msg2 arrives before retried msg1, Kafka reorders based on sequence numbers
- Best of both worlds: **high throughput + ordering guarantee**

**Visual Comparison:**

```
max.in.flight=5 (without idempotence):
Producer → [msg1][msg2][msg3][msg4][msg5] → Kafka
              ↓     ✓     ✓     ✓     ✓
           (fails, retries later)
Kafka receives: msg2, msg3, msg4, msg5, msg1 ❌ REORDERED

max.in.flight=1:
Producer → [msg1] → wait for ACK → [msg2] → wait → [msg3]...
              ↓                       ↓
           (retry if fail)         (waits)
Kafka receives: msg1, msg2, msg3 ✅ ORDERED (but slow)

max.in.flight=5 + idempotence:
Producer → [msg1:seq0][msg2:seq1][msg3:seq2]... → Kafka
              ↓          ✓          ✓
           (fails, retries)
Kafka receives out of order but reorders by sequence number ✅
Final: msg1, msg2, msg3 (ORDERED + FAST)
```

**Bottom line:** Use `enable_idempotence=True` to get ordering + performance!

#### `compression.type`

```python
producer = KafkaProducer(
    compression_type='snappy'  # or 'gzip', 'lz4', 'zstd', 'none'
)
```

**Compression types:**

- `snappy` - Good balance (speed + ratio)
- `lz4` - Faster, lower compression
- `gzip` - Slower, best compression
- `zstd` - Best balance (Kafka 2.1+)

**Benefits:**

- Reduce network bandwidth
- Reduce disk usage
- Better throughput

#### `batch.size` & `linger.ms`

**Batching for efficiency:**

```python
batch.size=16384      # 16KB batch before sending
linger.ms=10          # Wait 10ms for more messages
```

**Trade-off:**

- Larger batch.size = better throughput, higher latency
- Higher linger.ms = more batching, higher latency
- Lower values = lower latency, worse throughput

#### `enable.idempotence`

```python
producer = KafkaProducer(
    enable_idempotence=True,
    acks='all',
    retries=5
)
```

**Idempotent producer:**

- Guarantees exactly-once to a partition
- Handles retries without duplicates
- Automatically sets: `acks=all`, `retries=MAX_INT`, `max.in.flight=5`
- **Recommended for production!**

### Producer Partitioning Strategies

#### 1. Key-based (default)

```python
producer.send('user-clicks',
    key=b'user123',      # Same key → same partition
    value=b'clicked button'
)

# hash(key) % num_partitions = partition_id
```

**Use case:** Maintain ordering for related events (same user, same order, etc.)

#### 2. Round-robin (no key)

```python
producer.send('logs',
    value=b'log entry'   # No key → round-robin distribution
)
```

**Use case:** Load balancing when order doesn't matter

#### 3. Custom Partitioner

```python
from kafka import KafkaProducer
from kafka.partitioner import Partitioner

class CustomPartitioner(Partitioner):
    def partition(self, key, all_partitions, available_partitions):
        # Custom logic
        if key.startswith(b'premium'):
            return 0  # Premium users to partition 0
        else:
            return hash(key) % len(all_partitions)

producer = KafkaProducer(partitioner=CustomPartitioner())
```

### Producer Example (Python)

```python
from kafka import KafkaProducer
import json

# Create producer
producer = KafkaProducer(
    bootstrap_servers=['localhost:9092', 'localhost:9093'],
    acks='all',                        # Wait for all ISR replicas
    retries=3,                         # Retry on failure
    enable_idempotence=True,           # Exactly-once semantics
    compression_type='snappy',         # Compress messages
    value_serializer=lambda v: json.dumps(v).encode('utf-8')
)

# Send message with key
future = producer.send(
    'user-clicks',
    key=b'user123',
    value={'event': 'click', 'button': 'signup', 'timestamp': 1234567890}
)

# Wait for ACK (synchronous)
try:
    record_metadata = future.get(timeout=10)
    print(f"Sent to partition {record_metadata.partition}, offset {record_metadata.offset}")
except Exception as e:
    print(f"Failed to send: {e}")

# Close producer (flush pending messages)
producer.close()
```

### Producer Best Practices

1. **Use `acks=all` + `enable.idempotence=true` for critical data**
2. **Set appropriate timeouts** - `request.timeout.ms`, `delivery.timeout.ms`
3. **Handle failures** - implement retry logic at application level if needed
4. **Batch for throughput** - tune `batch.size` and `linger.ms`
5. **Compress data** - save network and disk
6. **Monitor metrics** - record-send-rate, record-error-rate, request-latency-avg
7. **Close producer gracefully** - calls `flush()` to send pending messages

## Consumers

### What is a Consumer?

**Consumer** = Application that reads messages from Kafka topics

- Subscribes to one or more topics
- Reads from partition leaders (or followers in newer versions)
- Tracks progress via offsets
- Can work in consumer groups for parallelism

### Consumer vs Consumer Group

**Single Consumer:**

```
Consumer1 reads from all partitions of a topic
- Partition 0 → Consumer1
- Partition 1 → Consumer1
- Partition 2 → Consumer1
```

**Consumer Group** (recommended):

```
Group "web-app-consumers" with 3 consumers:
- Partition 0 → Consumer1
- Partition 1 → Consumer2
- Partition 2 → Consumer3

Each partition consumed by exactly ONE consumer in the group
Parallel processing! 🚀
```

### Consumer Groups Explained

**Consumer Group** = Set of consumers working together to consume a topic

**Key properties:**

- Each partition assigned to **one consumer** in the group
- If consumer dies, partitions rebalanced to others
- Multiple groups can consume same topic independently

**Example:**

```
Topic: orders (3 partitions)

Group "order-processing":
  - Consumer A → Partition 0
  - Consumer B → Partition 1
  - Consumer C → Partition 2

Group "analytics" (different group):
  - Consumer X → Partition 0, 1, 2 (all partitions)

Both groups consume independently!
```

### Consumer Rebalancing

**Rebalance** = Reassigning partitions among consumers in a group

**Triggers:**

1. Consumer joins the group
2. Consumer leaves (crash or shutdown)
3. Consumer is considered dead (heartbeat timeout)
4. Partitions added to topic

**Process:**

```
Initial state:
- Consumer1: [P0, P1]
- Consumer2: [P2]

Consumer3 joins → Rebalance triggered

After rebalance:
- Consumer1: [P0]
- Consumer2: [P1]
- Consumer3: [P2]
```

**During rebalance:**

- Consumers stop consuming
- Partitions reassigned
- Consumers resume (latency spike!)

**Strategies:**

- `RangeAssignor` (default) - assigns contiguous ranges
- `RoundRobinAssignor` - round-robin across consumers
- `StickyAssignor` - minimizes movement during rebalance

### Offsets & Offset Management

**Offset** = Consumer's position in a partition

```
Partition 0:
[msg0][msg1][msg2][msg3][msg4][msg5]...
  ↑
Consumer offset = 3 (next message to read)
```

**Offset commit strategies:**

#### 1. Auto-commit (default)

```python
consumer = KafkaConsumer(
    'topic',
    enable_auto_commit=True,    # Auto-commit enabled
    auto_commit_interval_ms=5000  # Commit every 5 seconds
)

for message in consumer:
    process(message)  # Auto-commits in background
```

**Risk:** Can lose messages on crash (processed but not committed)

#### 2. Manual commit (synchronous)

```python
consumer = KafkaConsumer(
    'topic',
    enable_auto_commit=False
)

for message in consumer:
    process(message)
    consumer.commit()  # Commit after processing
```

**Benefit:** Process → commit → guaranteed at-least-once

#### 3. Manual commit (asynchronous)

```python
consumer = KafkaConsumer(
    'topic',
    enable_auto_commit=False
)

for message in consumer:
    process(message)
    consumer.commit_async()  # Non-blocking commit
```

**Benefit:** Better performance, no waiting for commit

### Offset Storage

**Kafka stores offsets in internal topic:**

```
Topic: __consumer_offsets (50 partitions by default)

Stores: (group_id, topic, partition) → offset
```

**Key structure:**
```
Key: (group_id, topic, partition)
Value: offset

Examples:
  ("fraud_detection", "user-activities", partition=0) → offset=5
  ("activity_analytics", "user-activities", partition=0) → offset=3
```

**Why consumer ID is NOT needed:**
- One partition assigned to ONE consumer in a group at a time
- Group tracks position, not individual consumers
- Rebalancing doesn't affect offsets - new consumer picks up from group's last offset

**Example scenario:**

```
Messages: [A][B][C][D][E]
          0  1  2  3  4

Group "fraud_detection":
  - Consumed: A, B
  - Committed offset: 2
  - Next to read: C (offset 2)

Group "activity_analytics":
  - Consumed: A, B, C
  - Committed offset: 3
  - Next to read: D (offset 3)

Stored in __consumer_offsets:
  ("fraud_detection", "user-activities", 0) → 2
  ("activity_analytics", "user-activities", 0) → 3
```

**How consumers fetch:**

```python
# fraud_detection consumer asks:
"Give me messages from partition 0, starting at offset 2"
Kafka reads __consumer_offsets: offset=2
Returns: [C][D][E]...

# activity_analytics consumer asks:
"Give me messages from partition 0, starting at offset 3"
Kafka reads __consumer_offsets: offset=3
Returns: [D][E]...
```

**On rebalance:**

```
Initial:
  Consumer A → P0 (group offset: 5)
  Consumer B → P1 (group offset: 8)

Consumer A crashes! Rebalance:
  Consumer B → P0, P1

Consumer B reads P0:
  Kafka: "Group last read P0 at offset 5"
  Consumer B picks up at offset 5 ✅
  
Doesn't matter that it was Consumer A before!
```

**Key insight:** The group is the unit of consumption, not individual consumers. This enables:
- Flexible consumer instances (can come and go)
- Seamless rebalancing
- Simpler offset management

**Allows:**

- Consumer restarts from last committed offset
- Multiple consumer groups track independently
- Offset reset/rewind
- Multiple groups consume same topic independently (Fan-Out pattern)

### Key Consumer Configurations

#### `group.id`

```python
consumer = KafkaConsumer(
    'topic',
    group_id='my-consumer-group'  # Join this group
)
```

**Critical:** Consumers with same `group.id` share partition assignments

#### `auto.offset.reset`

**What to do when no offset exists (new consumer group):**

```python
auto_offset_reset='earliest'  # Start from beginning
auto_offset_reset='latest'    # Start from end (default)
auto_offset_reset='none'      # Throw exception
```

#### `max.poll.records`

```python
max_poll_records=500  # Fetch up to 500 messages per poll()
```

**Tune based on:**

- Processing time per message
- Available memory

#### `session.timeout.ms` & `heartbeat.interval.ms`

```python
session_timeout_ms=10000     # 10s - consumer dead if no heartbeat
heartbeat_interval_ms=3000   # 3s - send heartbeat every 3s
```

**Relationship:**

- `heartbeat.interval.ms` < `session.timeout.ms`
- Typically: heartbeat = 1/3 of session timeout

**Too low:** False positives (unnecessary rebalances)
**Too high:** Slow failure detection

#### `fetch.min.bytes` & `fetch.max.wait.ms`

```python
fetch_min_bytes=1024       # Wait for at least 1KB
fetch_max_wait_ms=500      # But wait max 500ms
```

**Trade-off:**

- Higher min bytes = better throughput, higher latency
- Lower values = lower latency, more requests

### Consumer Example (Python)

```python
from kafka import KafkaConsumer
import json

# Create consumer
consumer = KafkaConsumer(
    'user-clicks',                    # Topic to consume
    bootstrap_servers=['localhost:9092'],
    group_id='analytics-group',       # Consumer group
    auto_offset_reset='earliest',     # Start from beginning if no offset
    enable_auto_commit=False,         # Manual commit for safety
    value_deserializer=lambda m: json.loads(m.decode('utf-8'))
)

# Consume messages
try:
    for message in consumer:
        print(f"Partition {message.partition}, Offset {message.offset}")
        print(f"Key: {message.key}, Value: {message.value}")

        # Process message
        process(message.value)

        # Manually commit offset
        consumer.commit()

except KeyboardInterrupt:
    print("Stopping consumer...")
finally:
    consumer.close()
```

### Consumer Best Practices

1. **Use consumer groups** for parallelism (not single consumers)
2. **Match consumer count to partition count** (or less)
   - 3 partitions → max 3 consumers in group (4th would be idle)
3. **Manual commit** for critical data (at-least-once processing)
4. **Idempotent processing** - handle duplicate messages
5. **Monitor lag** - how far behind consumers are
6. **Handle rebalances gracefully** - pause/resume processing
7. **Set appropriate timeouts** - session.timeout, request.timeout
8. **Process in batches** when possible for efficiency

## Delivery Semantics

### At-Most-Once

**Config:**

```python
producer: acks=0, retries=0
consumer: enable_auto_commit=True
```

**Behavior:** Message may be lost, never duplicated
**Use case:** Metrics, logs where loss is acceptable

### At-Least-Once (Most Common)

**Config:**

```python
producer: acks=all, retries>0, enable.idempotence=False
consumer: enable_auto_commit=False, manual commit after processing
```

**Behavior:** Message never lost, may be duplicated
**Use case:** Most production workloads + idempotent processing

### Exactly-Once

**Config:**

```python
producer: enable_idempotence=True, transactional.id='unique-id'
consumer: isolation.level='read_committed'
```

**Behavior:** Message delivered exactly once (end-to-end)
**Use case:** Financial transactions, critical state changes
**Note:** Complex, requires transactional API

## Testing Producer & Consumer

### Test Producer

```bash
# Console producer (interactive)
docker exec -it kafka1 kafka-console-producer \
  --bootstrap-server localhost:9092 \
  --topic test-topic \
  --property "parse.key=true" \
  --property "key.separator=:"

# Then type: key1:value1
```

### Test Consumer

```bash
# Console consumer
docker exec kafka1 kafka-console-consumer \
  --bootstrap-server localhost:9092 \
  --topic test-topic \
  --from-beginning \
  --property print.key=true \
  --property key.separator=":"

# With consumer group
docker exec kafka1 kafka-console-consumer \
  --bootstrap-server localhost:9092 \
  --topic test-topic \
  --group my-test-group \
  --from-beginning
```

### Check Consumer Group Status

```bash
# List consumer groups
docker exec kafka1 kafka-consumer-groups \
  --bootstrap-server localhost:9092 \
  --list

# Describe group (shows lag!)
docker exec kafka1 kafka-consumer-groups \
  --bootstrap-server localhost:9092 \
  --group my-test-group \
  --describe
```

## Common Patterns

### Fan-Out Pattern

**One topic, multiple consumer groups:**

```
Topic: events

Group "email-service" → sends emails
Group "analytics" → tracks metrics
Group "audit-log" → stores for compliance

All groups consume same events independently
```

### Stream Processing

**Consumer reads, transforms, produces:**

```
Consumer (group "processor")
  ← reads from topic "raw-events"
  → processes/enriches
  → produces to topic "enriched-events"
```

### Dead Letter Queue (DLQ)

**Failed messages sent to separate topic:**

```python
for message in consumer:
    try:
        process(message)
        consumer.commit()
    except ProcessingError:
        producer.send('dlq-topic', message)  # Send to DLQ
        consumer.commit()  # Commit to avoid retry
```
