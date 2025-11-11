# Kafka Versioning & ZooKeeper History

## Version Comparison: Confluent vs Apache

**Two distributions with different versioning:**

- **Apache Kafka** (vanilla): 4.0.x
  - Open-source project from Apache Foundation
  - Image: `apache/kafka:latest` or `apache/kafka:4.0.0`

- **Confluent Platform**: 8.1.0
  - Commercial distribution bundling Kafka + extras
  - Image: `confluentinc/cp-kafka:latest`
  - Includes: Schema Registry, ksqlDB, connectors, etc.

**Version mapping:**
- Confluent Platform 8.1.0 ≈ Apache Kafka 3.8-3.9.x internally
- Confluent uses its own version numbers for the entire platform
- Core Kafka functionality is the same

**License:**
- Confluent Community Edition: Free (Apache 2.0), no restrictions
- Confluent Enterprise: Paid (RBAC, tiered storage, support)

## ZooKeeper Deprecation Timeline

**KRaft Mode Evolution:**

| Version | Status | Notes |
|---------|--------|-------|
| Kafka 2.8.0 (2021) | Preview/Early Access | KRaft introduced, not production-ready |
| Kafka 3.3.0 (2022) | **Production-Ready** ✅ | Safe to use KRaft in production |
| Kafka 3.5.0+ | ZooKeeper Deprecated | Official recommendation to migrate |
| Kafka 4.0 (2024) | **ZooKeeper Removed** | KRaft only, no ZooKeeper support |

**Why KRaft?**
- Simpler architecture (one less dependency)
- Faster metadata operations
- Better scalability (no ZooKeeper bottleneck)
- Easier to deploy and maintain

**Current setup (CP 8.1 ≈ Kafka 3.8/3.9):**
- Using KRaft mode
- No ZooKeeper needed
- Production-ready and stable

