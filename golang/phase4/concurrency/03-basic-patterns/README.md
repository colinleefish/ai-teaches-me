# Level 3: Basic Patterns 🔄

**Simple Compositions of Goroutines and Channels**

Now that you understand goroutines and channels individually, it's time to combine them into powerful patterns that solve real-world problems!

## 🎯 Learning Progression

### 3.1 Producer-Consumer 🏭
**File**: `producer-consumer.go`
**Core Concept**: Separate data generation from data processing

**Key Patterns:**
- Single producer → Single consumer
- Multiple producers → Single consumer  
- Single producer → Multiple consumers
- Fan-out/Fan-in combinations
- Buffered queues and backpressure

### 3.2 Worker Pools 🏗️  
**File**: `worker-pools.go`
**Core Concept**: Fixed workers processing unlimited jobs

**Key Patterns:**
- Job queue with worker pool
- Worker specialization
- Dynamic pool sizing
- Graceful shutdown
- Error handling and retries

### 3.3 Pipeline 🔄
**File**: `pipeline.go`  
**Core Concept**: Chain processing stages for data transformation

**Key Patterns:**
- Multi-stage data transformation
- Pipeline fan-out/fan-in
- Stage buffering and rate matching
- Error propagation
- Pipeline cancellation

### 3.4 Bounded Parallelism 🎚️
**File**: `bounded-parallelism.go`
**Core Concept**: Limit concurrency to protect resources

**Key Patterns:**
- Semaphore pattern for concurrency limits
- Resource pooling
- Rate limiting
- Adaptive concurrency
- Circuit breaker protection

## 🚀 Why These Patterns Matter

**🏭 Producer-Consumer**: Handles rate mismatches between data generation and processing  
**🏗️ Worker Pools**: Provides controlled concurrency for predictable resource usage  
**🔄 Pipeline**: Enables streaming data transformation with parallel stages  
**🎚️ Bounded Parallelism**: Prevents resource exhaustion under high load

## 💡 Implementation Strategy

1. **Start with templates** - Each file has function signatures and TODOs
2. **Follow TUTOR hints** - Understand the why before implementing
3. **Build incrementally** - Complete basic patterns before advanced features
4. **Test edge cases** - Shutdown, errors, and resource limits
5. **Monitor behavior** - Use `runtime.NumGoroutine()` and metrics

## ⚠️ Common Pitfalls in Patterns

**Producer-Consumer:**
- Forgetting to close channels → consumer hangs
- Mismatched production/consumption rates → memory bloat

**Worker Pools:**  
- Creating unlimited workers → resource exhaustion
- No worker termination → goroutine leaks

**Pipeline:**
- Unbalanced stage performance → bottlenecks
- No error handling → silent failures

**Bounded Parallelism:**
- Incorrect semaphore management → deadlocks
- Resource leaks in error scenarios → instability

## 🎯 Success Criteria

You've mastered Level 3 when you can:
- [ ] Build producer-consumer systems with proper shutdown
- [ ] Implement worker pools with controlled resource usage
- [ ] Design multi-stage pipelines with error handling
- [ ] Apply bounded parallelism to prevent resource exhaustion
- [ ] Monitor and tune concurrent pattern performance
- [ ] Choose appropriate patterns for different use cases

## 🔗 What's Next

After mastering these basic patterns, you'll advance to **Level 4: Advanced Patterns** where you'll learn context propagation, distributed patterns, and complex synchronization!

---

**"Patterns are the vocabulary of concurrent programming"** 🗣️

These four patterns form the foundation of almost every concurrent Go application. Master them, and you'll be able to build scalable, efficient systems! ⚡
