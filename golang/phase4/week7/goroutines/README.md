# Week 7: Goroutines 🧵

Welcome to Go's lightweight threading model! Goroutines are the foundation of Go's concurrency system - they're incredibly cheap and allow you to create thousands of concurrent operations with ease.

## 🎯 Learning Objectives

By the end of this week, you'll understand:
- What goroutines are and how they differ from OS threads
- How to create and manage goroutines with the `go` keyword
- Synchronization patterns using `sync.WaitGroup`
- Common concurrency pitfalls and how to avoid them
- Building concurrent programs that scale

## 📚 Topics Covered

### 1. Goroutine Basics (`goroutine_basics.go`)
- Creating goroutines with `go` keyword
- Goroutine lifecycle and scheduling
- Anonymous functions as goroutines
- Goroutine vs thread comparison

### 2. Synchronization with WaitGroups (`waitgroups.go`)
- `sync.WaitGroup` fundamentals
- Coordinating multiple goroutines
- Proper cleanup and error handling
- Dynamic goroutine management

### 3. Race Conditions & Safety (`race_conditions.go`)
- Understanding data races
- Using the race detector (`go run -race`)
- Safe vs unsafe concurrent access
- When to use mutexes vs channels

### 4. Goroutine Patterns (`goroutine_patterns.go`)
- Worker pool pattern
- Producer-consumer patterns
- Bounded parallelism
- Graceful shutdown patterns

### 5. Practice Exercises (`goroutine_practice.go`)
- Hands-on implementations
- Real-world scenarios
- Performance comparisons
- Debugging exercises

## ⚡ Goroutines vs Threads

| Aspect | OS Threads | Goroutines |
|--------|------------|------------|
| **Memory** | ~2MB stack | ~2KB initial stack |
| **Creation** | Expensive syscall | Cheap function call |
| **Scheduling** | OS kernel | Go runtime |
| **Context Switch** | ~1-2μs | ~10-50ns |
| **Maximum** | ~1000s | ~1,000,000s |

## 🚀 Quick Start Example

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    
    // Launch 3 goroutines
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d starting\n", id)
            time.Sleep(time.Second)
            fmt.Printf("Goroutine %d done\n", id)
        }(i)
    }
    
    wg.Wait() // Wait for all goroutines to complete
    fmt.Println("All goroutines finished!")
}
```

## 🛠️ How to Practice

1. **Start with basics**: Read `goroutine_basics.go` to understand fundamentals
2. **Learn coordination**: Study `waitgroups.go` for synchronization
3. **Understand safety**: Review `race_conditions.go` to avoid common mistakes  
4. **Apply patterns**: Implement examples in `goroutine_patterns.go`
5. **Practice**: Complete exercises in `goroutine_practice.go`

## 🧪 Testing Your Code

```bash
# Run with race detector
go run -race goroutine_basics.go

# Build and run
go build goroutine_basics.go
./goroutine_basics

# Run specific function
go run goroutine_practice.go
```

## ⚠️ Common Pitfalls

### 1. Forgetting to Wait
```go
// ❌ WRONG - Program exits before goroutines finish
for i := 0; i < 5; i++ {
    go fmt.Println(i)
}
// Program ends immediately

// ✅ CORRECT - Wait for completion
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(val int) {
        defer wg.Done()
        fmt.Println(val)
    }(i)
}
wg.Wait()
```

### 2. Closing Over Loop Variables
```go
// ❌ WRONG - All goroutines print the same value
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i) // Always prints 5!
    }()
}

// ✅ CORRECT - Pass value as parameter
for i := 0; i < 5; i++ {
    go func(val int) {
        fmt.Println(val) // Prints 0, 1, 2, 3, 4
    }(i)
}
```

### 3. Goroutine Leaks
```go
// ❌ WRONG - Goroutine never terminates
func leakyFunction() {
    go func() {
        for {
            time.Sleep(1 * time.Second) // Runs forever
        }
    }()
} // Function returns, but goroutine keeps running

// ✅ CORRECT - Provide exit condition
func cleanFunction(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return // Exit when context is cancelled
            default:
                time.Sleep(1 * time.Second)
            }
        }
    }()
}
```

## 🎯 Key Concepts to Master

### WaitGroup Patterns
```go
// Dynamic work distribution
var wg sync.WaitGroup
jobs := []string{"job1", "job2", "job3"}

for _, job := range jobs {
    wg.Add(1)
    go func(j string) {
        defer wg.Done()
        processJob(j)
    }(job)
}
wg.Wait()
```

### Error Handling in Goroutines
```go
// Using channels to collect errors
func processWithErrors(items []string) error {
    var wg sync.WaitGroup
    errCh := make(chan error, len(items))
    
    for _, item := range items {
        wg.Add(1)
        go func(i string) {
            defer wg.Done()
            if err := processItem(i); err != nil {
                errCh <- err
            }
        }(item)
    }
    
    wg.Wait()
    close(errCh)
    
    // Check for any errors
    for err := range errCh {
        if err != nil {
            return err
        }
    }
    return nil
}
```

## 🔍 Debugging Tools

```bash
# Race detector
go run -race program.go

# CPU profiling
go tool pprof http://localhost:6060/debug/pprof/profile

# Goroutine debugging
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

## 📈 Performance Tips

1. **Goroutines are cheap** - Don't be afraid to create many
2. **Avoid premature optimization** - Measure before optimizing
3. **Use bounded parallelism** - Don't overwhelm system resources
4. **Profile your code** - Use `go tool pprof` to find bottlenecks
5. **Test with `-race`** - Always check for race conditions

## 🔗 What's Next

After mastering goroutines, you'll learn **channels** in Week 8 - Go's way of enabling safe communication between goroutines without shared memory!

Ready to unleash the power of concurrent programming? Let's go! 🚀🐹
