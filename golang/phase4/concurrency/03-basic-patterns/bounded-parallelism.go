package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 🎚️ BOUNDED PARALLELISM PATTERN
// Controls the maximum number of concurrent operations
// Prevents resource exhaustion while maximizing throughput
// Essential for systems with limited resources (memory, connections, CPU)
// Provides predictable resource usage under varying load

// Task represents work that needs bounded parallelism
type Task struct {
	ID       int
	Workload int // Simulates different task complexities
	Data     string
}

// TaskResult represents the outcome of bounded parallel processing
type TaskResult struct {
	TaskID   int
	Output   string
	Duration time.Duration
	Error    error
}

// TUTOR: Semaphore pattern uses buffered channels to limit concurrency.
// Channel capacity determines maximum concurrent operations.
// Acquiring semaphore (sending to channel) starts operation.
// Releasing semaphore (receiving from channel) allows next operation.
// This pattern prevents unlimited resource consumption.
// TODO: Demonstrate basic semaphore-based concurrency limiting
func demonstrateBasicSemaphore() {
	fmt.Println("=== Basic Semaphore Pattern ===")

	// TODO: Create semaphore channel with capacity limit
	// TODO: Launch more tasks than semaphore capacity
	// TODO: Show tasks waiting for semaphore availability
	// TODO: Demonstrate controlled concurrent execution
	// TODO: Show proper semaphore release after task completion

	capacity := 3
	jobs := 100
	semaphore := make(chan struct{}, capacity)

	// wg := sync.WaitGroup{}
	jobStartWg := sync.WaitGroup{}

	for i := 0; i < jobs; i++ {
		semaphore <- struct{}{}
		jobStartWg.Add(1)
		go func(i int) {
			defer jobStartWg.Done()
			fmt.Println("job", i, "started")
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println("job", i, "completed")
			<-semaphore
		}(i)
	}

	jobStartWg.Wait()

}

// TUTOR: Worker pool semaphores limit active workers instead of total workers.
// Fixed worker count vs dynamic concurrency control.
// Semaphores provide finer-grained resource management.
// This pattern combines worker pools with bounded parallelism.
// Useful when workers have different resource requirements.
// TODO: Demonstrate semaphore-controlled worker activation
func demonstrateWorkerSemaphore() {
	fmt.Println("\n=== Worker Pool with Semaphore Control ===")

	// TODO: Create worker pool with more workers than allowed concurrent operations
	// TODO: Use semaphore to control how many workers are active
	// TODO: Show workers waiting for semaphore before processing
	// TODO: Demonstrate resource-aware worker management
	// TODO: Show dynamic concurrency adjustment
}

// TUTOR: Rate limiting controls operation frequency rather than concurrency.
// Prevents overwhelming external systems with too many requests.
// Token bucket or sliding window algorithms control request rates.
// Rate limiting complements concurrency limiting for external resource protection.
// Essential for API clients and external service interactions.
// TODO: Demonstrate rate limiting with bounded parallelism
func demonstrateRateLimiting() {
	fmt.Println("\n=== Rate Limiting with Bounded Parallelism ===")

	// Rate limiting setup
	requestsPerSecond := 5
	maxConcurrent := 2
	totalRequests := 20

	// TODO: Create rate limiter (time.NewTicker)
	rateLimiter := time.NewTicker(time.Second / time.Duration(requestsPerSecond))
	defer rateLimiter.Stop()

	// TODO: Create semaphore for concurrency limiting
	semaphore := make(chan struct{}, maxConcurrent)

	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < totalRequests; i++ {
		// TODO: Wait for rate limit token
		<-rateLimiter.C

		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// TODO: Acquire semaphore (concurrency limit)
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			// TODO: Simulate API request work
			fmt.Printf("[%v] Request %d started (concurrent slots used)\n",
				time.Since(start).Round(100*time.Millisecond), id)

			// TODO: Add work simulation
			time.Sleep(time.Duration(rand.Intn(800)+200) * time.Millisecond)

			fmt.Printf("[%v] Request %d completed\n",
				time.Since(start).Round(100*time.Millisecond), id)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Total time: %v\n", time.Since(start).Round(100*time.Millisecond))
}

// TUTOR: Resource pools manage limited shared resources efficiently.
// Database connections, file handles, or API clients are pooled.
// Pool size determines maximum concurrent resource usage.
// Resource acquisition and release must be carefully managed.
// Proper pooling prevents resource exhaustion and improves performance.
// TODO: Demonstrate resource pool management

type DBConn struct{ connID int }

func demonstrateResourcePooling() {
	fmt.Println("\n=== Resource Pool Management ===")

	// TODO: Create pool of limited resources (simulated database connections)
	// TODO: Show resource acquisition and release patterns
	// TODO: Demonstrate resource reuse across multiple operations
	// TODO: Show proper resource cleanup and pool shutdown
	// TODO: Illustrate pool utilization monitoring

	poolCount := 3
	jobCount := 50

	// build connection pool
	pool := make(chan *DBConn, poolCount)
	for i := 0; i < poolCount; i++ {
		pool <- &DBConn{connID: i}
	}

	jobs := make(chan int, jobCount)
	for j := 0; j < jobCount; j++ {
		jobs <- j
	}

	close(jobs)

	var wg sync.WaitGroup
	// start exactly poolCount workers
	wg.Add(poolCount)
	for w := 0; w < poolCount; w++ {
		go rpWorker(w, jobs, pool, &wg)
	}
	wg.Wait()

	close(pool)
}

func rpWorker(w int, jobs <-chan int, pool chan *DBConn, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		dbConn := <-pool
		fmt.Println("using db conn: ", dbConn.connID, "worker: ", w, "job: ", j)
		time.Sleep(time.Millisecond * 100)
		pool <- dbConn
	}
}

// TUTOR: Adaptive parallelism adjusts concurrency based on system conditions.
// Monitor system metrics to determine optimal concurrency levels.
// Increase parallelism when resources are available.
// Decrease parallelism when system is under stress.
// Adaptive systems maintain performance across varying conditions.
// TODO: Demonstrate adaptive concurrency control
func demonstrateAdaptiveParallelism() {
	fmt.Println("\n=== Adaptive Parallelism ===")

	// TODO: Monitor system metrics (CPU, memory, response times)
	// TODO: Adjust concurrency limits based on system health
	// TODO: Show automatic scaling up during low load
	// TODO: Demonstrate scaling down during high load
	// TODO: Show metrics-driven concurrency decisions

	// fake metrics that jumps from 0 to 100, 0 is idle and 100 is max load, update every 1 second
	// don't have to use channel, just assign value to a variable

	doTheJob := func(jobId int) {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		fmt.Println("job", jobId, "completed")
	}

	jobCount := 200

	metrics := int32(0)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			atomic.StoreInt32(&metrics, rand.Int31n(100))
		}
	}()

	jobCh := make(chan int, jobCount)
	for i := 0; i < jobCount; i++ {
		jobCh <- i
	}
	close(jobCh)

	// Worker management
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	activeWorkers := int32(0)
	var wg sync.WaitGroup

	// TODO: Scale workers based on metrics (11 - metrics/10)
	go func() {
		ticker := time.NewTicker(2 * time.Second) // Check every 2 seconds
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				currentMetrics := atomic.LoadInt32(&metrics)
				desiredWorkers := 11 - int(currentMetrics/10)
				if desiredWorkers < 1 {
					desiredWorkers = 1
				}

				current := int(atomic.LoadInt32(&activeWorkers))
				fmt.Printf("Metrics: %d, Current workers: %d, Desired: %d\n",
					currentMetrics, current, desiredWorkers)

				// TODO: Scale up - add workers
				if desiredWorkers > current {
					for i := current; i < desiredWorkers; i++ {
						atomic.AddInt32(&activeWorkers, 1)
						wg.Add(1)
						go worker(ctx, jobCh, doTheJob, &activeWorkers, &wg)
					}
				}

				// TODO: Scale down - workers will exit on context cancellation
				// (This is tricky - need coordination between workers)
				if desiredWorkers < current {
					fmt.Println("scaling down from", current, "to", desiredWorkers)
					fmt.Println("Do nothing at the moment")
				}

			case <-ctx.Done():
				return
			}
		}
	}()

	// TODO: Wait for all jobs completion
	wg.Wait()
}

func worker(ctx context.Context, jobs <-chan int, doJob func(int), activeWorkers *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	defer atomic.AddInt32(activeWorkers, -1)

	for {
		select {
		case jobId, ok := <-jobs:
			if !ok {
				return // Channel closed
			}
			doJob(jobId)
		case <-ctx.Done():
			return // Context cancelled (scaling down)
		}
	}
}

// TUTOR: Priority-based bounded parallelism reserves capacity for important work.
// High-priority tasks get preferential access to limited resources.
// Multiple semaphores can implement priority levels.
// This pattern ensures critical work isn't blocked by bulk processing.
// Priority management prevents starvation of important operations.
// TODO: Demonstrate priority-aware concurrency limiting
func demonstratePriorityBounding() {
	fmt.Println("\n=== Priority-Based Bounded Parallelism ===")

	// TODO: Create separate semaphores for different priority levels
	// TODO: Show high-priority tasks getting preferential resource access
	// TODO: Demonstrate low-priority task deferral during resource contention
	// TODO: Show fair scheduling between priority levels
	// TODO: Illustrate starvation prevention mechanisms
}

// TUTOR: Timeout-based parallelism prevents long-running operations from blocking resources.
// Operations that exceed time limits are cancelled or terminated.
// Timeouts ensure predictable resource release timing.
// This pattern prevents resource starvation from slow operations.
// Timeout handling requires proper cleanup and error reporting.
// TODO: Demonstrate timeout-controlled bounded parallelism
func demonstrateTimeoutBounding() {
	fmt.Println("\n=== Timeout-Based Parallelism ===")

	// TODO: Implement operation timeouts with concurrency limits
	// TODO: Show automatic cancellation of long-running tasks
	// TODO: Demonstrate resource release on timeout
	// TODO: Show timeout monitoring and adjustment
	// TODO: Illustrate proper cleanup of timed-out operations
}

// TUTOR: Circuit breaker patterns protect against cascading failures.
// Failed operations temporarily reduce concurrency limits.
// Circuit breakers prevent overwhelming failing systems.
// Automatic recovery restores normal concurrency when systems recover.
// This pattern improves system resilience under failure conditions.
// TODO: Demonstrate circuit breaker with bounded parallelism
func demonstrateCircuitBreakerBounding() {
	fmt.Println("\n=== Circuit Breaker with Bounded Parallelism ===")

	// TODO: Implement circuit breaker state management
	// TODO: Show concurrency reduction during failure periods
	// TODO: Demonstrate automatic recovery and limit restoration
	// TODO: Show failure rate monitoring and thresholds
	// TODO: Illustrate system protection under adverse conditions
}

// TUTOR: Bounded parallelism monitoring tracks resource utilization.
// Monitor semaphore utilization and queue depths.
// Track operation success rates and processing times.
// Resource utilization metrics guide capacity planning.
// Proper monitoring enables proactive performance management.
// TODO: Demonstrate monitoring of bounded parallel systems
func demonstrateBoundedParallelismMonitoring() {
	fmt.Println("\n=== Bounded Parallelism Monitoring ===")

	// TODO: Track semaphore utilization and wait times
	// TODO: Monitor resource pool efficiency
	// TODO: Show queue depth and processing rate metrics
	// TODO: Demonstrate bottleneck identification
	// TODO: Show capacity planning using utilization data
}

// TUTOR: Real-world bounded parallelism appears in many production systems.
// Web servers limit concurrent request processing.
// Database clients limit concurrent connection usage.
// File processors limit concurrent file operations.
// Understanding bounded parallelism is essential for scalable systems.
// TODO: Show practical applications of bounded parallelism
func demonstrateRealWorldBoundedParallelism() {
	fmt.Println("\n=== Real-World Bounded Parallelism ===")

	// TODO: Show web server with connection limits
	// TODO: Demonstrate database client with connection pooling
	// TODO: Show file processor with concurrent file limits
	// TODO: Illustrate API client with rate and concurrency limits
	// TODO: Show memory-conscious batch processing
}

func main() {
	fmt.Println("🎚️ Bounded Parallelism - Controlled Concurrency for Stable Systems 🎚️")
	fmt.Println("Learn to manage resources while maximizing performance")

	// TODO: Implement each demonstration function
	// Build understanding of resource management in concurrent systems

	// demonstrateBasicSemaphore()
	// demonstrateWorkerSemaphore()
	// demonstrateRateLimiting()
	// demonstrateResourcePooling()
	demonstrateAdaptiveParallelism()
	// demonstratePriorityBounding()
	// demonstrateTimeoutBounding()
	// demonstrateCircuitBreakerBounding()
	// demonstrateBoundedParallelismMonitoring()
	// demonstrateRealWorldBoundedParallelism()
}
