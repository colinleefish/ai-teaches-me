// Week 7: Goroutine Practice Exercises
// Complete these exercises to master goroutine fundamentals

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TODO: Exercise 1 - Basic Goroutine Creation
func exercise1_BasicGoroutines() {
	fmt.Println("=== Exercise 1: Basic Goroutines ===")

	// TODO: Create 5 goroutines that each print their ID and a message
	// Each goroutine should:
	// 1. Print "Goroutine X starting" (where X is the ID)
	// 2. Sleep for a random duration (100-500ms)
	// 3. Print "Goroutine X finished"
	//
	// Use proper synchronization to ensure all complete before main exits
	// Expected output: All 5 goroutines should start and finish

	fmt.Println("All goroutines completed!")
}

// TODO: Exercise 2 - WaitGroup Coordination
func exercise2_WaitGroupCoordination() {
	fmt.Println("\n=== Exercise 2: WaitGroup Coordination ===")

	// TODO: Create a function that downloads multiple URLs concurrently
	// Simulate downloads with different durations
	urls := []string{
		"https://api.github.com",
		"https://httpbin.org/delay/1",
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://reqres.in/api/users/1",
		"https://httpbin.org/delay/2",
	}

	// TODO: Implement concurrent downloader
	download := func(url string, wg *sync.WaitGroup) {
		// TODO: Simulate download with random delay (1-3 seconds)
		// TODO: Print download start and completion
		// TODO: Don't forget to call wg.Done()
	}

	// TODO: Download all URLs concurrently and wait for completion
	// Measure total time and compare with sequential approach

	fmt.Println("All downloads completed!")
}

// TODO: Exercise 3 - Safe Counter Implementation
func exercise3_SafeCounter() {
	fmt.Println("\n=== Exercise 3: Safe Counter Implementation ===")

	// TODO: Implement a thread-safe counter that can be incremented by multiple goroutines
	type SafeCounter struct {
		// TODO: Add necessary fields for thread safety
	}

	// TODO: Implement methods for SafeCounter
	// NewSafeCounter() *SafeCounter
	// Increment()
	// Value() int
	// Add(delta int)
	// Reset()

	// TODO: Test the counter with 100 goroutines, each incrementing 1000 times
	// Final value should be exactly 100,000

	// TODO: Also implement using atomic operations and compare performance

	fmt.Println("Counter implementation completed!")
}

// TODO: Exercise 4 - Worker Pool Implementation
func exercise4_WorkerPool() {
	fmt.Println("\n=== Exercise 4: Worker Pool Implementation ===")

	// TODO: Implement a generic worker pool that can process any type of job
	type Job interface {
		Process() Result
	}

	type Result interface {
		String() string
	}

	type WorkerPool struct {
		// TODO: Add fields for managing workers and jobs
	}

	// TODO: Implement WorkerPool methods:
	// NewWorkerPool(numWorkers int) *WorkerPool
	// Start()
	// Submit(job Job) <-chan Result
	// Stop()

	// TODO: Create a concrete job type for testing
	type MathJob struct {
		A, B int
		Op   string // "add", "multiply", "power"
	}

	// TODO: Test with 50 math jobs using 5 workers
	// Show that work is distributed among workers

	fmt.Println("Worker pool implementation completed!")
}

// TODO: Exercise 5 - Producer-Consumer with Buffer
func exercise5_ProducerConsumerBuffer() {
	fmt.Println("\n=== Exercise 5: Producer-Consumer with Buffer ===")

	// TODO: Implement a bounded buffer for producer-consumer pattern
	type BoundedBuffer struct {
		// TODO: Add fields for buffer management
		// Should support Put(item) and Get() operations
		// Should block when buffer is full (Put) or empty (Get)
	}

	// TODO: Implement BoundedBuffer methods:
	// NewBoundedBuffer(capacity int) *BoundedBuffer
	// Put(item interface{}) error
	// Get() (interface{}, error)
	// Close()
	// IsClosed() bool

	// TODO: Test with multiple producers and consumers
	// 3 producers generating items at different rates
	// 2 consumers processing items at different speeds
	// Buffer capacity of 10 items

	fmt.Println("Producer-consumer implementation completed!")
}

// TODO: Exercise 6 - Pipeline Processing
func exercise6_PipelineProcessing() {
	fmt.Println("\n=== Exercise 6: Pipeline Processing ===")

	// TODO: Implement a data processing pipeline
	// Stage 1: Read numbers from input (1-100)
	// Stage 2: Filter prime numbers
	// Stage 3: Square the prime numbers
	// Stage 4: Format as "number: square" strings
	// Stage 5: Write to output

	// TODO: Each stage should be a separate goroutine communicating via channels
	// TODO: Handle proper pipeline shutdown when input is exhausted

	// Helper function to check if number is prime
	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	// TODO: Implement pipeline stages and connect them
	// TODO: Measure processing time and show results

	fmt.Println("Pipeline processing completed!")
}

// TODO: Exercise 7 - Rate Limiter Implementation
func exercise7_RateLimiter() {
	fmt.Println("\n=== Exercise 7: Rate Limiter Implementation ===")

	// TODO: Implement a token bucket rate limiter
	type RateLimiter struct {
		// TODO: Add fields for rate limiting logic
	}

	// TODO: Implement RateLimiter methods:
	// NewRateLimiter(rate int, burst int) *RateLimiter
	// Allow() bool (non-blocking)
	// Wait(ctx context.Context) error (blocking)
	// Stop()

	// TODO: Test the rate limiter
	// Configure for 10 requests per second with burst of 5
	// Send 50 requests as fast as possible
	// Show that requests are properly rate limited

	// TODO: Demonstrate both Allow() and Wait() methods

	fmt.Println("Rate limiter implementation completed!")
}

// TODO: Exercise 8 - Graceful Shutdown Service
func exercise8_GracefulShutdown() {
	fmt.Println("\n=== Exercise 8: Graceful Shutdown Service ===")

	// TODO: Implement a service that handles graceful shutdown
	type Service struct {
		// TODO: Add fields for service management
	}

	// TODO: Implement Service methods:
	// NewService() *Service
	// Start() error
	// Stop(timeout time.Duration) error
	// IsRunning() bool

	// TODO: The service should:
	// 1. Run multiple background workers
	// 2. Handle shutdown signal gracefully
	// 3. Allow workers to complete current tasks
	// 4. Force shutdown if timeout exceeded
	// 5. Clean up resources properly

	// TODO: Test the service:
	// 1. Start service with 3 workers
	// 2. Let it run for 2 seconds
	// 3. Initiate graceful shutdown with 3-second timeout
	// 4. Verify all workers stop cleanly

	fmt.Println("Graceful shutdown service completed!")
}

// TODO: Exercise 9 - Concurrent Map Implementation
func exercise9_ConcurrentMap() {
	fmt.Println("\n=== Exercise 9: Concurrent Map Implementation ===")

	// TODO: Implement a thread-safe map with reader-writer semantics
	type ConcurrentMap struct {
		// TODO: Add fields for concurrent access
	}

	// TODO: Implement ConcurrentMap methods:
	// NewConcurrentMap() *ConcurrentMap
	// Set(key string, value interface{})
	// Get(key string) (interface{}, bool)
	// Delete(key string) bool
	// Keys() []string
	// Len() int
	// Clear()

	// TODO: Test with concurrent readers and writers
	// 10 goroutines writing different keys
	// 20 goroutines reading random keys
	// 5 goroutines deleting random keys
	// Run for 5 seconds and show final state

	// TODO: Implement both mutex-based and channel-based versions
	// Compare performance

	fmt.Println("Concurrent map implementation completed!")
}

// TODO: Exercise 10 - Goroutine Pool with Context
func exercise10_GoroutinePoolWithContext() {
	fmt.Println("\n=== Exercise 10: Goroutine Pool with Context ===")

	// TODO: Implement a goroutine pool that supports context cancellation
	type GoroutinePool struct {
		// TODO: Add fields for pool management
	}

	type Task func(ctx context.Context) error

	// TODO: Implement GoroutinePool methods:
	// NewGoroutinePool(size int) *GoroutinePool
	// Submit(ctx context.Context, task Task) error
	// SubmitWithTimeout(task Task, timeout time.Duration) error
	// Shutdown(ctx context.Context) error
	// Stats() (active, queued, completed int)

	// TODO: Test the pool:
	// 1. Create pool with 5 goroutines
	// 2. Submit 20 tasks with varying durations
	// 3. Cancel some tasks using context
	// 4. Show pool statistics
	// 5. Shutdown gracefully

	fmt.Println("Goroutine pool with context completed!")
}

// Helper functions for exercises
func randomDuration(min, max time.Duration) time.Duration {
	return min + time.Duration(rand.Int63n(int64(max-min)))
}

func simulateWork(name string, duration time.Duration) {
	fmt.Printf("%s: Starting work\n", name)
	time.Sleep(duration)
	fmt.Printf("%s: Work completed\n", name)
}

func measureExecutionTime(name string, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	fmt.Printf("%s execution time: %v\n", name, duration)
}

func main() {
	fmt.Println("🏃‍♂️ Welcome to Goroutine Practice! 🏃‍♂️")
	fmt.Println("Complete these exercises to master goroutine programming")

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// TODO: Implement each exercise one by one
	// Start with basic exercises and progress to advanced ones
	// Uncomment each exercise as you complete the previous one

	exercise1_BasicGoroutines()
	// exercise2_WaitGroupCoordination()
	// exercise3_SafeCounter()
	// exercise4_WorkerPool()
	// exercise5_ProducerConsumerBuffer()
	// exercise6_PipelineProcessing()
	// exercise7_RateLimiter()
	// exercise8_GracefulShutdown()
	// exercise9_ConcurrentMap()
	// exercise10_GoroutinePoolWithContext()

	fmt.Println("\n🎉 Congratulations! You've mastered goroutine programming!")
	fmt.Println("🚀 Ready for Week 8: Channels!")
}

/*
🎯 Exercise Guidelines:

1. **Start Simple**: Begin with basic goroutine creation and synchronization
2. **Add Complexity**: Gradually increase complexity with each exercise
3. **Test Thoroughly**: Test your implementations with race detector
4. **Measure Performance**: Compare different approaches when applicable
5. **Handle Errors**: Don't ignore error handling in concurrent code
6. **Clean Resources**: Always clean up goroutines and channels properly

📝 Completion Checklist:
□ Exercise 1: Basic goroutine creation and coordination
□ Exercise 2: WaitGroup for multiple goroutines
□ Exercise 3: Thread-safe data structures
□ Exercise 4: Worker pool pattern
□ Exercise 5: Producer-consumer with buffering
□ Exercise 6: Pipeline processing pattern
□ Exercise 7: Rate limiting implementation
□ Exercise 8: Graceful shutdown handling
□ Exercise 9: Concurrent data structures
□ Exercise 10: Context-based cancellation

🔧 Testing Commands:
```bash
# Run with race detector
go run -race goroutine_practice.go

# Run with CPU profiling
go run goroutine_practice.go -cpuprofile=cpu.prof

# Run specific exercise
go run goroutine_practice.go -exercise=1
```

🚨 Common Mistakes to Avoid:
- Not using defer for WaitGroup.Done()
- Capturing loop variables in closures incorrectly
- Creating goroutines without bounds
- Not handling context cancellation
- Ignoring goroutine leaks
- Poor error propagation in concurrent code

🎯 Success Criteria:
- All exercises pass with race detector
- Clean goroutine lifecycle management
- Proper error handling and resource cleanup
- Understanding of when to use each pattern
- Ability to debug concurrent issues
*/
