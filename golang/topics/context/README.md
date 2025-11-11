# Go Context 🎯

## What is Context?

**Context** is Go's built-in system for handling **timeouts, cancellation, and request-scoped data** across your application. Think of it as a "control signal" that travels with your operations.

## 🎪 The Problem Context Solves

### Without Context (Problematic):
```go
// How do you cancel this if it takes too long?
result := slowOperation()  // Might hang forever

// How do you coordinate timeouts across multiple functions?
func parentFunc() {
    childFunc()  // No way to tell child to stop
}
```

### With Context (Solution):
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

select {
case result := <-slowOperation(ctx):
    // Success
case <-ctx.Done():
    // Timeout or cancellation
}
```

## 🔧 Context Types

### 1. **Background Context** (Root)
```go
ctx := context.Background()  // Empty context, never cancelled
// Use as root for other contexts
```

### 2. **Timeout Context**
```go
// Automatically cancels after duration
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()  // Always clean up!
```

### 3. **Deadline Context**
```go
// Cancels at specific time
deadline := time.Now().Add(30 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

### 4. **Cancellable Context**
```go
// Manually controllable cancellation
ctx, cancel := context.WithCancel(context.Background())

// Cancel from anywhere
go func() {
    time.Sleep(5 * time.Second)
    cancel()  // Manual cancellation
}()
```

## ⚡ Basic Usage Pattern

```go
func doWorkWithContext(ctx context.Context) error {
    // Create operation channel
    resultCh := make(chan string, 1)
    
    // Start work in goroutine
    go func() {
        // Simulate work
        time.Sleep(3 * time.Second)
        resultCh <- "Work completed"
    }()
    
    // Race: work vs context cancellation
    select {
    case result := <-resultCh:
        fmt.Println("Success:", result)
        return nil
    case <-ctx.Done():
        fmt.Println("Cancelled:", ctx.Err())
        return ctx.Err()
    }
}
```

## 🌐 Context Propagation

**Contexts automatically propagate down the call chain:**

```go
func parentFunction() {
    // Parent sets 10-second timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    childFunction(ctx)  // Pass context down
}

func childFunction(ctx context.Context) {
    // Child inherits the 10-second timeout
    grandchildFunction(ctx)  // Pass it further down
}

func grandchildFunction(ctx context.Context) {
    // Grandchild also has the same 10-second timeout
    select {
    case result := <-doWork():
        return result
    case <-ctx.Done():
        return ctx.Err()  // All levels cancelled together
    }
}
```

## 🔍 Context Interface

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)  // When will it timeout?
    Done() <-chan struct{}                    // Channel that closes when cancelled
    Err() error                               // Why was it cancelled?
    Value(key interface{}) interface{}        // Get request-scoped data
}
```

## 🎯 Real-World Examples

### HTTP Request with Timeout
```go
func fetchURLWithContext(url string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return err
    }
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err  // Includes timeout errors
    }
    defer resp.Body.Close()
    
    fmt.Printf("HTTP %d received\n", resp.StatusCode)
    return nil
}
```

### Database Query with Context
```go
func queryUserWithContext(ctx context.Context, userID int) (*User, error) {
    // Most database drivers support context
    query := "SELECT name, email FROM users WHERE id = ?"
    
    row := db.QueryRowContext(ctx, query, userID)
    
    var user User
    err := row.Scan(&user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    
    return &user, nil
}
```

### Worker Pool with Context
```go
func workerPool(ctx context.Context, jobs <-chan Job) {
    for {
        select {
        case job := <-jobs:
            processJob(ctx, job)
        case <-ctx.Done():
            fmt.Println("Worker stopping:", ctx.Err())
            return  // Clean shutdown
        }
    }
}
```

## ⚠️ Context Best Practices

### ✅ DO:
```go
// Always use defer cancel()
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()  // Prevents resource leaks

// Pass context as first parameter
func myFunction(ctx context.Context, name string) error

// Check context in long-running operations
for i := 0; i < 1000000; i++ {
    select {
    case <-ctx.Done():
        return ctx.Err()  // Stop early if cancelled
    default:
        // Continue work
    }
}
```

### ❌ DON'T:
```go
// Don't store context in structs
type Worker struct {
    ctx context.Context  // ❌ Bad practice
}

// Don't pass nil context
myFunction(nil, "data")  // ❌ Use context.Background() instead

// Don't ignore context in select
select {
case result := <-operation:
    // ❌ Missing context case - can't be cancelled
}
```

## 🏭 Context vs time.After

| Feature | time.After | Context |
|---------|------------|---------|
| **Propagation** | ❌ No | ✅ Automatic down call stack |
| **Manual cancellation** | ❌ No | ✅ cancel() function |
| **Resource cleanup** | ❌ Manual | ✅ Automatic with defer |
| **Error information** | ❌ No | ✅ ctx.Err() explains why |
| **Composability** | ❌ Limited | ✅ Excellent |
| **Professional use** | ❌ Simple cases only | ✅ Production applications |

## 🚀 Advanced Patterns

### Context with Values (Request Tracing)
```go
type ctxKey string

func handleRequest(w http.ResponseWriter, r *http.Request) {
    // Add request ID to context
    requestID := generateRequestID()
    ctx := context.WithValue(r.Context(), ctxKey("requestID"), requestID)
    
    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    requestID := ctx.Value(ctxKey("requestID")).(string)
    fmt.Printf("Processing request %s\n", requestID)
}
```

### Cascading Timeouts
```go
func cascadingTimeouts() {
    // Root: 30-second total timeout
    rootCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // Step 1: 10-second timeout for first operation
    step1Ctx, cancel1 := context.WithTimeout(rootCtx, 10*time.Second)
    defer cancel1()
    
    result1 := doStep1(step1Ctx)
    
    // Step 2: Use remaining time from root context
    result2 := doStep2(rootCtx, result1)
    
    return combineResults(result1, result2)
}
```

## 💡 When to Use Context

**Use Context for:**
- ✅ HTTP request handling
- ✅ Database operations  
- ✅ External API calls
- ✅ Long-running computations
- ✅ Worker pools and background jobs
- ✅ Any operation that might need cancellation

**Context is Go's professional way to handle timeouts and cancellation!** 

Essential for building robust, responsive applications that can gracefully handle failures and resource constraints. 🌟

## 🔗 Key Takeaways

1. **Always use `defer cancel()`** to prevent resource leaks
2. **Pass context as first parameter** to functions
3. **Check `ctx.Done()`** in long-running operations  
4. **Use `context.Background()`** as root context
5. **Context automatically propagates** down the call stack
6. **Professional Go code uses Context** for timeouts and cancellation
