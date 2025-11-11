# Kafka

## What is Kafka

Apache Kafka is a distributed event streaming platform for high-throughput, fault-tolerant data pipelines. It acts as a message broker where producers publish events to topics and consumers subscribe to those topics to process events in real-time or batch mode.

**Core Concepts:**

- **Topics**: Categories/feeds where records are stored
- **Producers**: Applications that publish events to topics
- **Consumers**: Applications that subscribe and process events from topics
- **Brokers**: Kafka servers that store and serve data
- **Partitions**: Topics split into partitions for parallelism and scalability

## Scenarios to Use It

1. **Real-time Data Pipelines**: Stream data between systems (e.g., CDC from databases to data warehouses)
2. **Event Sourcing**: Store all state changes as immutable events
3. **Log Aggregation**: Collect logs from multiple services for centralized processing
4. **Metrics & Monitoring**: Stream metrics from microservices to monitoring systems
5. **Message Queue**: Decouple microservices with async communication
6. **Stream Processing**: Real-time analytics, fraud detection, recommendation engines

## Setup Architecture Example

### Basic Single-Node Setup

```
┌─────────────┐      ┌──────────────────┐      ┌─────────────┐
│  Producer   │─────▶│   Kafka Broker   │─────▶│  Consumer   │
│ Application │      │                  │      │ Application │
└─────────────┘      │  - ZooKeeper/    │      └─────────────┘
                     │    KRaft Mode    │
                     │  - Topic Storage │
                     └──────────────────┘
```

### Production Cluster Setup

```
                     ┌─────────────────────────────────┐
                     │      Kafka Cluster (3 nodes)    │
┌──────────┐         │  ┌────────┐  ┌────────┐  ┌────────┐
│Producer 1│────┐    │  │Broker 1│  │Broker 2│  │Broker 3│
└──────────┘    │    │  └────────┘  └────────┘  └────────┘
┌──────────┐    ├───▶│       Partition Replication        │
│Producer 2│────┘    │                                     │
└──────────┘         └─────────────────────────────────────┘
                                    │
                                    ▼
                     ┌──────────────────────────────┐
                     │    Consumer Groups           │
                     │  ┌──────┐  ┌──────┐  ┌──────┐
                     │  │ C1   │  │ C2   │  │ C3   │
                     │  └──────┘  └──────┘  └──────┘
                     └──────────────────────────────┘
```

### Docker Compose Example

```yaml
version: "3"
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
    ports:
      - "2181:2181"

  kafka:
    image: confluentinc/cp-kafka:latest
    depends_on:
      - zookeeper
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
    ports:
      - "9092:9092"
```

### Key Ports

- **9092**: Kafka broker (default)
- **2181**: ZooKeeper (coordination)
- **8081**: Schema Registry (optional)
- **9021**: Control Center UI (optional)
