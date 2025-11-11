// Week 7: Common Goroutine Patterns
// This file demonstrates proven patterns for structuring concurrent Go programs

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TODO: Implement Worker Pool pattern
func demonstrateWorkerPool() {
	fmt.Println("=== Worker Pool Pattern ===")

	// TODO: Create a worker pool that processes jobs concurrently
	// but limits the number of workers to avoid overwhelming the system

	type Job struct {
		ID   int
		Data string
	}

	type Result struct {
		Job    Job
		Output string
		Error  error
	}

	// TODO: Implement worker function
	worker := func(id int, jobs <-chan Job, results chan<- Result) {
		// TODO: Process jobs from the channel
		// 1. Read jobs from jobs channel
		// 2. Process each job (simulate work)
		// 3. Send result to results channel
		// 4. Continue until jobs channel is closed
	}

	// TODO: Set up the worker pool
	numWorkers := 3
	numJobs := 10

	// TODO: Create channels for jobs and results
	// TODO: Start workers
	// TODO: Send jobs to workers
	// TODO: Collect results
	// TODO: Show how work is distributed among workers
}

// TODO: Implement Producer-Consumer pattern
func demonstrateProducerConsumer() {
	fmt.Println("\n=== Producer-Consumer Pattern ===")

	// TODO: Implement multiple producers and consumers
	// Producers generate data, consumers process it

	// TODO: Create producer function
	producer := func(id int, output chan<- string, wg *sync.WaitGroup) {
		// TODO: Generate data and send to channel
		// Create multiple items with producer ID
	}

	// TODO: Create consumer function
	consumer := func(id int, input <-chan string, wg *sync.WaitGroup) {
		// TODO: Consume data from channel
		// Process until channel is closed
	}

	// TODO: Set up multiple producers and consumers
	// TODO: Show how to coordinate shutdown
}

// TODO: Implement Pipeline pattern
func demonstratePipeline() {
	fmt.Println("\n=== Pipeline Pattern ===")

	// TODO: Create a processing pipeline with multiple stages
	// Stage 1: Generate numbers
	// Stage 2: Square the numbers
	// Stage 3: Filter even numbers
	// Stage 4: Format as strings

	// TODO: Stage 1 - Number generator
	generateNumbers := func(ctx context.Context, max int) <-chan int {
		// TODO: Generate numbers from 1 to max
		// Return receive-only channel
		return nil
	}

	// TODO: Stage 2 - Square numbers
	squareNumbers := func(ctx context.Context, input <-chan int) <-chan int {
		// TODO: Read from input, square each number, send to output
		return nil
	}

	// TODO: Stage 3 - Filter even numbers
	filterEven := func(ctx context.Context, input <-chan int) <-chan int {
		// TODO: Only pass through even numbers
		return nil
	}

	// TODO: Stage 4 - Format as strings
	formatStrings := func(ctx context.Context, input <-chan int) <-chan string {
		// TODO: Convert numbers to formatted strings
		return nil
	}

	// TODO: Connect the pipeline stages
	// TODO: Show how data flows through the pipeline
}

// TODO: Implement Fan-out/Fan-in pattern
func demonstrateFanOutFanIn() {
	fmt.Println("\n=== Fan-Out/Fan-In Pattern ===")

	// TODO: Fan-out: Distribute work from one source to multiple workers
	// Fan-in: Collect results from multiple workers into one stream

	// TODO: Create work generator
	generateWork := func() <-chan int {
		// TODO: Generate work items
		return nil
	}

	// TODO: Create worker that processes items
	processWork := func(input <-chan int, output chan<- int) {
		// TODO: Process items and send results
	}

	// TODO: Fan-out: Distribute work to multiple processors
	// TODO: Fan-in: Merge results from all processors

	// TODO: Show how this pattern scales with more workers
}

// TODO: Implement Bounded Parallelism pattern
func demonstrateBoundedParallelism() {
	fmt.Println("\n=== Bounded Parallelism Pattern ===")

	// TODO: Process many items concurrently but limit concurrent operations
	// Use semaphore pattern with buffered channel

	items := make([]int, 20)
	for i := range items {
		items[i] = i + 1
	}

	// TODO: Create semaphore channel to limit concurrency
	maxConcurrent := 3
	// TODO: Use buffered channel as semaphore

	// TODO: Process each item with bounded parallelism
	processItem := func(item int, wg *sync.WaitGroup) {
		// TODO: Acquire semaphore
		// TODO: Process item
		// TODO: Release semaphore
	}

	// TODO: Launch goroutines for all items
	// TODO: Show how only 'maxConcurrent' run at once
}

// TODO: Implement Timeout and Cancellation patterns
func demonstrateTimeoutCancellation() {
	fmt.Println("\n=== Timeout and Cancellation Patterns ===")

	// TODO: Show different ways to handle timeouts and cancellation

	// TODO: Pattern 1: Simple timeout with time.After
	simpleTimeout := func() {
		// TODO: Create a goroutine that might take too long
		// TODO: Use select with time.After for timeout
	}

	// TODO: Pattern 2: Context-based cancellation
	contextCancellation := func() {
		// TODO: Use context.WithTimeout
		// TODO: Pass context to goroutines
		// TODO: Check ctx.Done() in goroutines
	}

	// TODO: Pattern 3: Manual cancellation with channels
	manualCancellation := func() {
		// TODO: Create done channel for cancellation signal
		// TODO: Use select to check for cancellation
	}

	// TODO: Demonstrate each pattern
}

// TODO: Implement Graceful Shutdown pattern
func demonstrateGracefulShutdown() {
	fmt.Println("\n=== Graceful Shutdown Pattern ===")

	// TODO: Show how to cleanly shut down a concurrent application

	// TODO: Create a service that runs multiple workers
	type Service struct {
		workers []chan bool
		wg      sync.WaitGroup
	}

	// TODO: Implement service start
	startService := func() *Service {
		// TODO: Start background workers
		// TODO: Return service handle for shutdown
		return nil
	}

	// TODO: Implement graceful shutdown
	shutdown := func(service *Service, timeout time.Duration) {
		// TODO: Signal all workers to stop
		// TODO: Wait for workers to finish (with timeout)
		// TODO: Force stop if timeout exceeded
	}

	// TODO: Demonstrate service lifecycle
}

// TODO: Implement Rate Limiting pattern
func demonstrateRateLimiting() {
	fmt.Println("\n=== Rate Limiting Pattern ===")

	// TODO: Implement different rate limiting strategies

	// TODO: Strategy 1: Token bucket using time.Ticker
	tokenBucket := func(requestsPerSecond int) {
		// TODO: Create ticker for rate limiting
		// TODO: Process requests at limited rate
	}

	// TODO: Strategy 2: Leaky bucket with buffered channel
	leakyBucket := func(capacity int, leakRate time.Duration) {
		// TODO: Use buffered channel as bucket
		// TODO: Leak tokens at steady rate
	}

	// TODO: Strategy 3: Sliding window rate limiter
	slidingWindow := func(windowSize time.Duration, maxRequests int) {
		// TODO: Track request timestamps
		// TODO: Allow requests within sliding window
	}

	// TODO: Test each strategy with burst traffic
}

// TODO: Implement Circuit Breaker pattern
func demonstrateCircuitBreaker() {
	fmt.Println("\n=== Circuit Breaker Pattern ===")

	// TODO: Implement circuit breaker for handling failures

	type CircuitBreaker struct {
		// TODO: Add fields for state management
		// States: Closed, Open, Half-Open
	}

	// TODO: Implement circuit breaker logic
	call := func(cb *CircuitBreaker, operation func() error) error {
		// TODO: Check circuit state
		// TODO: Execute operation or return error
		// TODO: Update state based on result
		return nil
	}

	// TODO: Simulate a service that sometimes fails
	// TODO: Show how circuit breaker prevents cascading failures
}

// TODO: Implement Retry pattern with backoff
func demonstrateRetryPattern() {
	fmt.Println("\n=== Retry Pattern with Backoff ===")

	// TODO: Implement different retry strategies

	// TODO: Strategy 1: Fixed delay retry
	fixedRetry := func(operation func() error, maxRetries int, delay time.Duration) error {
		// TODO: Retry with fixed delay between attempts
		return nil
	}

	// TODO: Strategy 2: Exponential backoff
	exponentialBackoff := func(operation func() error, maxRetries int, baseDelay time.Duration) error {
		// TODO: Retry with exponentially increasing delays
		return nil
	}

	// TODO: Strategy 3: Jittered backoff
	jitteredBackoff := func(operation func() error, maxRetries int, baseDelay time.Duration) error {
		// TODO: Add randomness to prevent thundering herd
		return nil
	}

	// TODO: Test with a flaky operation
}

// TODO: Performance comparison of patterns
func performanceComparison() {
	fmt.Println("\n=== Performance Comparison ===")

	// TODO: Compare different patterns for the same problem
	// 1. Sequential processing
	// 2. Unlimited goroutines
	// 3. Worker pool
	// 4. Pipeline
	// 5. Bounded parallelism

	// TODO: Measure time, memory, and CPU usage
	// TODO: Show optimal patterns for different scenarios
}

// Helper function to simulate work
// func simulateWork(id int, duration time.Duration) string {
// 	time.Sleep(duration)
// 	return fmt.Sprintf("Work %d completed", id)
// }

// Helper function to simulate flaky operation
func flakyOperation() error {
	// 30% chance of failure
	if rand.Float32() < 0.3 {
		return fmt.Errorf("operation failed")
	}
	time.Sleep(100 * time.Millisecond)
	return nil
}

func main() {
	fmt.Println("🎨 Welcome to Goroutine Patterns! 🎨")
	fmt.Println("This file teaches you proven patterns for concurrent programming")

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// TODO: Implement each pattern demonstration
	// Start with simple patterns and progress to complex ones

	demonstrateWorkerPool()
	// demonstrateProducerConsumer()
	// demonstratePipeline()
	// demonstrateFanOutFanIn()
	// demonstrateBoundedParallelism()
	// demonstrateTimeoutCancellation()
	// demonstrateGracefulShutdown()
	// demonstrateRateLimiting()
	// demonstrateCircuitBreaker()
	// demonstrateRetryPattern()
	// performanceComparison()

	fmt.Println("\n🎉 Congratulations! You've learned essential goroutine patterns!")
	fmt.Println("Next: Practice with goroutine_practice.go")
}

/*
🔍 Key Patterns to Remember:

1. **Worker Pool**: Fixed number of workers processing from a queue
2. **Producer-Consumer**: Decoupled data generation and processing
3. **Pipeline**: Sequential stages of processing with channels
4. **Fan-Out/Fan-In**: Distribute work, then merge results
5. **Bounded Parallelism**: Limit concurrent operations with semaphore
6. **Timeout/Cancellation**: Handle long-running operations gracefully
7. **Graceful Shutdown**: Clean service termination
8. **Rate Limiting**: Control request throughput
9. **Circuit Breaker**: Fail fast when service is unhealthy
10. **Retry with Backoff**: Handle transient failures

🏗️ Pattern Selection Guide:
- **CPU-bound work**: Worker pool with GOMAXPROCS workers
- **I/O-bound work**: Higher number of workers
- **Stream processing**: Pipeline pattern
- **Batch processing**: Fan-out/fan-in
- **Resource protection**: Rate limiting + circuit breaker
- **Distributed work**: Producer-consumer

🚨 Anti-Patterns to Avoid:
- Creating unlimited goroutines
- Not handling goroutine lifecycle
- Ignoring errors in concurrent code
- Missing cancellation mechanisms
- Poor error propagation
- Resource leaks

🎯 Next Steps:
- Practice implementing these patterns
- Learn when to use each pattern
- Combine patterns for complex scenarios
- Move on to channels for communication
*/
