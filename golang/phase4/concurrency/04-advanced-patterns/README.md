# Advanced Patterns (Level 4)

Complex compositions using all 4 cornerstones: **Goroutines + WaitGroups + Channels + Select**

## 🎯 Key Concepts

| Pattern | Purpose | Key Techniques |
|---------|---------|----------------|
| **Fan-In** | Merge multiple streams | Select multiplexing, fair merging |
| **Fan-Out** | Distribute work to workers | Load distribution, result collection |
| **Timeout & Cancellation** | Handle long operations | `time.After`, `context.Context` |
| **Request-Response** | Bidirectional communication | Reply channels, correlation |

## 📁 Files

### 🔀 fan-in.go
**Merging multiple input streams**

```go
// Basic: Multiple channels → Single output
merged := fanIn(channel1, channel2, channel3)

// Advanced: Fair multiplexing with select
output := fanInSelect(input1, input2)
```

**Key concepts:**
- Stream merging
- Select-based fairness
- Goroutine per input

### 📤 fan-out.go
**Distributing work to multiple workers**

```go
// Distribute jobs to workers
results := fanOut(jobs, numWorkers)

// Collect results in order
orderedResults := fanOutOrdered(jobList, numWorkers)
```

**Key concepts:**
- Load distribution
- Result collection
- Worker coordination

### ⏰ timeout-cancellation.go
**Handling long-running operations**

```go
// Basic timeout
select {
case result := <-work:
    // Success
case <-time.After(timeout):
    // Timeout
}

// Context cancellation
ctx, cancel := context.WithTimeout(ctx, timeout)
err := doWork(ctx)
```

**Key concepts:**
- `time.After` for timeouts
- `context.Context` for cancellation
- Graceful cleanup

### 🔄 request-response.go
**Bidirectional communication patterns**

```go
// Request with reply channel
type Request struct {
    ID      int
    Data    string
    ReplyTo chan<- Response
}

// Async client-server
responses := asyncClient(requests, numRequests)
```

**Key concepts:**
- Reply channels
- Correlation IDs
- Async responses

## 🚀 Usage Examples

### Run Individual Patterns

```bash
# Fan-In: Merge streams from Alice, Bob, Charlie
go run fan-in.go

# Fan-Out: Distribute 10 jobs to 3 workers  
go run fan-out.go

# Timeout: Handle work with timeouts and context
go run timeout-cancellation.go

# Request-Response: Bidirectional communication
go run request-response.go
```

### Expected Output Examples

**Fan-In:**
```
Alice: message 0
Bob: message 0  
Alice: message 1
Charlie: message 0
Bob: message 1
...
```

**Fan-Out:**
```
Job 0 completed by worker 1: Processed: task-0
Job 2 completed by worker 0: Processed: task-2
Job 1 completed by worker 2: Processed: task-1
...
```

## 🎓 Learning Progression

1. **Start with Fan-In** - Learn stream merging with select
2. **Move to Fan-Out** - Understand work distribution  
3. **Add Timeouts** - Handle failure scenarios
4. **Master Request-Response** - Build interactive systems

## 💡 Design Patterns

### When to Use Each

- **Fan-In**: Multiple data sources → Single processor
- **Fan-Out**: Single work queue → Multiple workers  
- **Timeout**: External dependencies, user operations
- **Request-Response**: Interactive systems, APIs

### Composition

These patterns often combine:

```go
// Fan-Out + Fan-In + Timeout
requests := fanOut(jobs, workers)      // Distribute
results := fanIn(responses...)         // Merge results  
select {
case final := <-results:
    // Success
case <-time.After(deadline):
    // Timeout
}
```

## ⚡ Performance Tips

- **Fan-In**: Use select for fairness, avoid goroutine-per-channel for many inputs
- **Fan-Out**: Match worker count to CPU cores or I/O capacity
- **Timeout**: Use context cancellation for cleanup
- **Request-Response**: Pool reply channels to reduce allocations

Master these patterns and you can build sophisticated concurrent systems! 🐹⚡
