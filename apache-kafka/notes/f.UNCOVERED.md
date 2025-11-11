# Kafka Topics Not Yet Covered

This document lists advanced Kafka topics not covered in the current learning materials. These are topics to explore next for reaching expert-level Kafka knowledge.

## Retention Policies

### Log Compaction
- **What**: Keep only the latest value for each key (not time-based)
- **Use cases**: Database change streams, user profile updates, configuration management
- **Config**: `cleanup.policy=compact`
- **How it works**: Background cleaner thread removes old values for duplicate keys
- **vs Delete**: `cleanup.policy=delete` (time/size-based retention)

### Time-based Retention
- `retention.ms` - How long to keep messages
- `retention.bytes` - Max size before deletion
- Segment-based deletion

### Size-based Retention
- `segment.bytes` - Size of each log segment
- `segment.ms` - Time before rolling new segment
- Interaction with retention policies

### Cleanup Policies
- `delete` - Delete old segments
- `compact` - Keep latest per key
- `compact,delete` - Compact then delete old segments

## Operations & Monitoring

### Performance Tuning
- Broker configs: `num.network.threads`, `num.io.threads`
- OS tuning: file descriptors, swap, filesystem (XFS vs ext4)
- JVM tuning: heap size, GC settings
- Disk I/O optimization

### Metrics to Monitor
- **Broker metrics**: UnderReplicatedPartitions, OfflinePartitionsCount, ActiveControllerCount
- **Producer metrics**: record-send-rate, record-error-rate, request-latency-avg
- **Consumer metrics**: records-lag-max, fetch-rate, commit-latency-avg
- **Disk metrics**: disk usage, I/O wait

### Common Issues & Troubleshooting
- Consumer lag
- Rebalancing storms
- Leader election failures
- Disk full scenarios
- Network partition handling
- Split brain prevention

### Cluster Management
- Adding/removing brokers
- Partition reassignment
- Leader rebalancing
- Preferred leader election
- Controlled shutdown

## Advanced Features

### Kafka Streams
- Stream processing library
- Stateful/stateless operations
- Windowing & aggregations
- Joins (stream-stream, stream-table, table-table)
- Interactive queries
- Exactly-once semantics

### ksqlDB
- SQL interface for Kafka
- Stream processing with SQL
- Materialized views
- Push/pull queries
- User-defined functions (UDFs)

### Kafka Connect
- **What**: Framework for connecting Kafka to external systems
- **Source connectors**: Database (CDC), files, APIs → Kafka
- **Sink connectors**: Kafka → databases, data warehouses, S3
- **Common connectors**: Debezium (CDC), JDBC, S3, Elasticsearch
- **Distributed mode**: Scalable connector deployment
- **Transforms**: Single Message Transforms (SMTs)

### Transactions API
- **What**: Exactly-once semantics across topics
- **Use case**: Read-process-write patterns
- Transactional producers
- Read committed isolation level
- Transaction coordinator
- More complex than idempotence

### Quotas
- Produce/consume rate limiting
- Client ID based quotas
- User-based quotas
- Preventing noisy neighbors

## Security

### Authentication (SASL)
- **SASL/PLAIN** - Username/password
- **SASL/SCRAM** - Salted Challenge Response
- **SASL/GSSAPI** - Kerberos
- **SASL/OAUTHBEARER** - OAuth tokens
- Delegation tokens

### Encryption (SSL/TLS)
- Broker-to-broker encryption
- Client-to-broker encryption
- Certificate management
- Performance impact

### Authorization (ACLs)
- Resource-based permissions (topics, consumer groups, clusters)
- Principal types (users, groups)
- Operations (read, write, create, delete, etc.)
- ACL management CLI
- Super users

### Network Security
- IP whitelisting
- VPC/network isolation
- Firewall rules
- DMZ deployments

## Multi-Datacenter & DR

### MirrorMaker 2.0
- **What**: Replication between Kafka clusters
- Active-passive setup
- Active-active setup
- Topic/group replication
- Offset translation
- Monitoring replication lag

### Cluster Linking (Confluent)
- Alternative to MirrorMaker
- Topic-level replication
- Zero-copy replication
- Offset preservation

### Disaster Recovery
- Backup strategies
- Cluster failover procedures
- Data loss scenarios
- RPO/RTO considerations

### Geo-replication Patterns
- Hub-and-spoke
- Active-active
- Aggregation patterns
- Conflict resolution

## Advanced Patterns

### Event Sourcing
- Storing all state changes as events
- Rebuilding state from events
- Snapshotting strategies
- Event schema evolution

### CQRS (Command Query Responsibility Segregation)
- Separate read/write models
- Using Kafka as event log
- Materialized views
- Eventual consistency

### Saga Pattern
- Distributed transactions across services
- Choreography vs orchestration
- Compensating transactions
- Using Kafka for coordination

### Change Data Capture (CDC)
- Capturing database changes
- Debezium connector
- Log-based vs query-based CDC
- Schema evolution

## Schema Management

### Schema Registry
- **What**: Centralized schema storage for Avro/Protobuf/JSON
- Schema evolution rules
- Compatibility modes (backward, forward, full)
- Schema versioning
- Integration with producers/consumers

### Avro
- Binary serialization format
- Schema evolution support
- Smaller message size
- Better than JSON for production

### Protobuf
- Google's serialization format
- Strongly typed
- Schema evolution
- Code generation

## Performance & Scalability

### Partition Strategy
- Determining optimal partition count
- Impact on throughput
- Impact on latency
- Partition key design

### Tiered Storage
- **What**: Offload old data to object storage (S3, etc.)
- Reduce local disk usage
- Cost optimization
- Query cold data

### Zero-Copy
- sendfile() system call
- Reduces CPU usage
- Faster data transfer
- OS-level optimization

### Page Cache Utilization
- How Kafka uses OS page cache
- Why Kafka is fast
- Memory vs disk trade-offs

## Testing & Development

### Embedded Kafka
- In-memory Kafka for unit tests
- Test containers
- MockProducer/MockConsumer

### Consumer Testing
- Idempotent processing
- Exactly-once testing
- Rebalance testing

### Performance Testing
- Load testing tools
- Benchmarking
- Capacity planning

## Ecosystem Tools

### Kafka Manager / CMAK
- Cluster management UI
- Topic/partition management
- Consumer group monitoring

### Cruise Control
- Automated cluster operations
- Anomaly detection
- Load balancing
- Capacity planning

### Kafka UI Tools
- Kafdrop
- AKHQ (previously KafkaHQ)
- Conduktor
- Confluent Control Center

### Monitoring Tools
- Prometheus + Grafana
- JMX metrics exporters
- Datadog/New Relic integrations
- Custom metrics

## Next Steps

**Recommended learning order:**

1. **Log Compaction** - Critical for many real-world use cases
2. **Kafka Connect** - Most common integration tool
3. **Monitoring & Operations** - Essential for production
4. **Schema Registry + Avro** - Production data serialization
5. **Kafka Streams** - Stream processing without external tools
6. **Security** - Authentication, encryption, authorization
7. **Multi-DC Replication** - Disaster recovery & geo-distribution

**You're currently at intermediate-advanced level (65-70%). These topics will take you to expert level (90%+)!**

