# Phase 4: Concurrency 🚀

Welcome to Go's superpower! This phase covers Go's revolutionary approach to concurrent programming through goroutines and channels. You'll learn to write programs that can handle thousands of concurrent operations efficiently.

## 📚 What You'll Learn

### Week 7: Goroutines
- **Goroutine basics**: Lightweight threads vs OS threads
- **Creating goroutines**: The `go` keyword and function calls
- **Goroutine lifecycle**: Creation, execution, and termination
- **Synchronization**: WaitGroups and coordination patterns
- **Common pitfalls**: Race conditions and goroutine leaks

### Week 8: Channels
- **Channel fundamentals**: Communication between goroutines
- **Channel types**: Buffered vs unbuffered channels
- **Channel directions**: Send-only and receive-only channels
- **Select statements**: Non-blocking channel operations
- **Channel patterns**: Fan-in, fan-out, pipeline, worker pools

## 🎯 Learning Objectives

After completing this phase, you'll be able to:
- [ ] Create and manage goroutines effectively
- [ ] Synchronize goroutines using WaitGroups
- [ ] Design concurrent programs without race conditions
- [ ] Use channels for safe communication between goroutines
- [ ] Implement common concurrency patterns
- [ ] Build scalable concurrent applications
- [ ] Debug and profile concurrent Go programs

## 📁 Phase Structure

```
phase4/
├── week7/          # Goroutines
│   └── goroutines/
│       ├── README.md
│       ├── goroutine_basics.go     # Creating and running goroutines
│       ├── waitgroups.go           # Synchronization with WaitGroups
│       ├── race_conditions.go      # Common concurrency pitfalls
│       ├── goroutine_patterns.go   # Worker pools and coordination
│       └── goroutine_practice.go   # Practice exercises
│
└── week8/          # Channels
    └── channels/
        ├── README.md
        ├── channel_basics.go       # Channel creation and operations
        ├── buffered_channels.go    # Buffered vs unbuffered channels
        ├── select_statements.go    # Non-blocking operations
        ├── channel_patterns.go     # Advanced channel patterns
        └── channel_practice.go     # Practice exercises
```

## ⚡ Key Differences from Python

### Concurrency Models
```python
# Python - Threading with GIL limitations
import threading
import time

def worker(name):
    for i in range(5):
        print(f"Worker {name}: {i}")
        time.sleep(1)

# Limited by GIL for CPU-bound tasks
thread1 = threading.Thread(target=worker, args=("A",))
thread2 = threading.Thread(target=worker, args=("B",))
thread1.start()
thread2.start()
thread1.join()
thread2.join()
```

```go
// Go - True parallelism with goroutines
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(name string, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
        fmt.Printf("Worker %s: %d\n", name, i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    
    go worker("A", &wg)  // Goroutine - true parallelism
    go worker("B", &wg)
    
    wg.Wait()
}
```

### Communication Patterns
```python
# Python - Shared memory with locks
import threading
import queue

# Shared queue for communication
q = queue.Queue()
lock = threading.Lock()

def producer():
    for i in range(5):
        q.put(f"item-{i}")

def consumer():
    while True:
        try:
            item = q.get(timeout=1)
            print(f"Consumed: {item}")
            q.task_done()
        except queue.Empty:
            break
```

```go
// Go - Channels for communication
package main

import "fmt"

func producer(ch chan<- string) {
    for i := 0; i < 5; i++ {
        ch <- fmt.Sprintf("item-%d", i)
    }
    close(ch)
}

func consumer(ch <-chan string) {
    for item := range ch {
        fmt.Printf("Consumed: %s\n", item)
    }
}

func main() {
    ch := make(chan string, 2)  // Buffered channel
    
    go producer(ch)
    consumer(ch)  // Blocks until channel is closed
}
```

## 🚀 Getting Started

1. **Week 7**: Start with `week7/goroutines/README.md`
2. **Understand the model**: Learn how goroutines differ from threads
3. **Practice coordination**: Master WaitGroups and synchronization
4. **Week 8**: Move to channels and communication patterns
5. **Build concurrent apps**: Apply patterns in real projects

## 💡 Go Concurrency Philosophy

**"Don't communicate by sharing memory; share memory by communicating"**

### Goroutine Principles:
- Goroutines are cheap - create thousands without worry
- Use channels to coordinate, not shared variables
- Avoid locks when possible - use channels instead
- Design for concurrent execution from the start

### Channel Design Patterns:
- **Pipeline**: Chain processing stages with channels
- **Fan-out**: Distribute work across multiple goroutines
- **Fan-in**: Collect results from multiple sources
- **Worker Pool**: Limited workers processing from a queue

## ⚠️ Common Pitfalls

### Race Conditions
```go
// ❌ WRONG - Race condition
var counter int
for i := 0; i < 10; i++ {
    go func() {
        counter++  // Data race!
    }()
}

// ✅ CORRECT - Channel communication
ch := make(chan int, 10)
for i := 0; i < 10; i++ {
    go func(val int) {
        ch <- val
    }(i)
}
```

### Goroutine Leaks
```go
// ❌ WRONG - Goroutine leak
func leakyFunction() {
    ch := make(chan int)
    go func() {
        ch <- 42  // Blocks forever if no receiver
    }()
    // Function returns, goroutine stays alive
}

// ✅ CORRECT - Proper cleanup
func cleanFunction() {
    ch := make(chan int, 1)  // Buffered
    go func() {
        ch <- 42  // Won't block
    }()
    select {
    case result := <-ch:
        fmt.Println(result)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout")
    }
}
```

## 🎯 Success Criteria

You'll know you've mastered this phase when you can:
- Launch goroutines confidently and avoid leaks
- Coordinate multiple goroutines with WaitGroups
- Choose between buffered and unbuffered channels appropriately
- Use select statements for complex channel operations
- Implement worker pools and pipeline patterns
- Debug race conditions and deadlocks
- Design scalable concurrent architectures

## 🔗 What's Next

After mastering concurrency, you'll advance to **Phase 5: Standard Library & Packages** where you'll learn to build complete applications using Go's rich ecosystem!

## 📊 Performance Comparison

| Feature | Python Threading | Go Goroutines |
|---------|------------------|---------------|
| Creation Cost | ~50KB per thread | ~2KB per goroutine |
| OS Threads | 1:1 mapping | M:N multiplexing |
| GIL Impact | Limits parallelism | True parallelism |
| Max Recommended | ~100 threads | 100,000+ goroutines |
| Communication | Shared memory + locks | Channels (CSP) |

Let's build lightning-fast concurrent applications! ⚡🐹
