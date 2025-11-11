# Level 2: Core Concepts (Safety & Types) 🛡️

Building on the 4 cornerstones, this level teaches you how to use them **safely** and **effectively**. You'll learn about concurrent safety, channel types, flow control, resource management, and error handling.

## Prerequisites ✅

Complete Level 1 (Cornerstones) first:
- [x] 1.1 Goroutines 
- [x] 1.2 WaitGroups
- [x] 1.3 Channels  
- [x] 1.4 Select

## Learning Objectives 🎯

By the end of Level 2, you will:

- **Identify and prevent** race conditions in concurrent programs
- **Design APIs** using directional channels for type safety
- **Control flow** with buffered channels and backpressure
- **Manage resources** properly with channel closing patterns
- **Handle errors** systematically in concurrent systems

## Topics Overview 📋

### 2.1 Race Conditions (`race-conditions.go`)
**Master concurrent safety fundamentals**

- Understand what causes race conditions and how to detect them
- Use Go's race detector to find concurrent access violations  
- Prevent races with mutexes, channels, and atomic operations
- Compare different synchronization approaches

### 2.2 Channel Types (`channel-types.go`)  
**Design type-safe APIs with directional channels**

- Restrict channel operations with send-only and receive-only types
- Create self-documenting APIs through channel directions
- Understand automatic type conversion from bidirectional channels
- Build producer-consumer systems with clear interfaces

### 2.3 Buffered Channels (`buffered-channels.go`)
**Control flow and performance with buffering**

- Manage channel capacity for asynchronous communication
- Use buffers for producer-consumer decoupling and flow control
- Implement backpressure and semaphore patterns
- Optimize performance through proper buffer sizing

### 2.4 Channel Closing (`channel-closing.go`)
**Manage resources with proper closing patterns**

- Signal completion and shutdown through channel closing
- Detect closure with comma-ok idiom and range loops
- Follow sender-closes principle for safe resource management
- Handle complex closure scenarios in pipelines and fan-out patterns

### 2.5 Error Handling (`error-handling.go`)
**Handle failures systematically in concurrent systems**

- Communicate errors between goroutines using channels
- Aggregate and process errors from multiple concurrent operations
- Implement fail-fast and graceful degradation strategies
- Recover from panics in goroutines safely

## Key Principles 💡

### Safety First
- **Race conditions** are bugs that can cause data corruption
- **Type safety** prevents API misuse at compile time
- **Resource management** prevents leaks and panics

### Flow Control
- **Buffering** enables asynchronous communication with bounded memory
- **Backpressure** prevents fast producers from overwhelming slow consumers
- **Closing** signals completion and enables clean shutdown

### Error Handling
- **Explicit communication** of errors through channels or result types
- **Graceful degradation** maintains system availability during failures
- **Panic recovery** prevents goroutine panics from crashing programs

## Practice Strategy 📝

1. **Start with templates** - Each file provides function signatures and TODOs
2. **Implement incrementally** - Build understanding step by step
3. **Run with race detector** - Use `go run -race` to catch concurrency bugs
4. **Test edge cases** - Try different buffer sizes, timing, and failure scenarios

## Common Patterns You'll Learn 🔧

- **Mutex protection** for shared data access
- **Producer-consumer** with buffered channels
- **Semaphore limiting** with channel capacity
- **Graceful shutdown** with channel closing
- **Error aggregation** from multiple goroutines

## Success Criteria ✅

You've mastered Level 2 when you can:

- [ ] Write race-free concurrent code consistently
- [ ] Design clear APIs using directional channels
- [ ] Choose appropriate buffer sizes for your workload
- [ ] Implement proper resource cleanup patterns
- [ ] Handle errors systematically in concurrent systems

## Next Steps ➡️

After mastering these core concepts, you'll be ready for **Level 3: Basic Patterns**, where you'll learn to compose the cornerstones into useful concurrent patterns like worker pools, pipelines, and fan-in/fan-out.

---

**Remember:** These concepts build directly on Level 1's cornerstones. Each safety mechanism and pattern uses combinations of goroutines, WaitGroups, channels, and select to achieve specific goals. Master these fundamentals before moving to more complex patterns!
