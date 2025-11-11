package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 🏗️ WORKER POOL PATTERN
// Fixed number of workers processing jobs from a shared queue
// Controls resource usage while maximizing throughput
// Prevents goroutine explosion in high-load scenarios
// Essential pattern for production systems with limited resources

// Job represents work to be processed by workers
type Job struct {
	ID       int
	Data     string
	Priority int
}

// Result represents the outcome of job processing
type JobResult struct {
	JobID    int
	Output   string
	Duration time.Duration
	Error    error
}

// TUTOR: Basic worker pools use fixed number of goroutines to process jobs.
// Jobs are queued in a channel, workers pull jobs when available.
// This pattern prevents unlimited goroutine creation under load.
// Worker pools provide predictable resource usage and controlled concurrency.
// The job queue acts as a buffer between job creation and processing.
// TODO: Demonstrate basic worker pool with job queue
func demonstrateBasicWorkerPool() {
	fmt.Println("=== Basic Worker Pool ===")

	// TODO: Create job channel for work queue
	// TODO: Create result channel for processed jobs
	// TODO: Launch fixed number of worker goroutines
	// TODO: Workers pull jobs from queue and process them
	// TODO: Show job distribution among workers
	// TODO: Demonstrate proper pool shutdown

	jobCh := make(chan Job, 10)

	workerWg := sync.WaitGroup{}
	workerNum := runtime.NumCPU()

	for i := 0; i < workerNum; i++ {
		workerWg.Add(1)
		go func(workerId int) {
			rate := rand.Intn(300)
			defer workerWg.Done()
			for job := range jobCh {
				time.Sleep(time.Duration(rate) * time.Millisecond)
				fmt.Println("worker", workerId, "processing job", job.ID)
			}
		}(i)
	}

	producerWg := sync.WaitGroup{}
	producerWg.Add(1)
	go func() {
		defer producerWg.Done()
		for i := 0; i < 1000; i++ {
			jobCh <- Job{ID: i, Data: fmt.Sprintf("job %d", i), Priority: rand.Intn(10)}
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}
	}()

	go func() {
		producerWg.Wait()
		close(jobCh)
	}()

	workerWg.Wait()
}

// TUTOR: Worker pool sizing affects performance and resource usage.
// Too few workers underutilize CPU and create processing bottlenecks.
// Too many workers waste memory and increase context switching overhead.
// Optimal sizing depends on workload characteristics and system resources.
// Dynamic sizing can adapt to changing load conditions.
// TODO: Demonstrate worker pool sizing strategies
func demonstrateWorkerPoolSizing() {
	fmt.Println("\n=== Worker Pool Sizing ===")

	// TODO: Compare different worker pool sizes
	// TODO: Show performance characteristics of each size
	// TODO: Demonstrate CPU utilization vs worker count
	// TODO: Show memory usage implications
	// TODO: Illustrate optimal sizing for different workload types
}

// TUTOR: Buffered job queues provide elasticity in worker pool systems.
// Large job queues smooth out burst traffic but increase memory usage.
// Small job queues provide backpressure but may block job submission.
// Queue depth monitoring helps identify system bottlenecks.
// Bounded queues prevent memory exhaustion under extreme load.
// TODO: Demonstrate job queue buffer effects
func demonstrateJobQueueBuffering() {
	fmt.Println("\n=== Job Queue Buffering ===")

	// TODO: Compare unbuffered vs buffered job queues
	// TODO: Show how buffer size affects job submission blocking
	// TODO: Demonstrate queue depth monitoring
	// TODO: Show backpressure effects on job producers
	// TODO: Illustrate memory usage vs buffer size trade-offs
}

// TUTOR: Worker specialization assigns different job types to different workers.
// Specialized workers can be optimized for specific task characteristics.
// Job routing directs work to appropriate worker types.
// This pattern enables fine-tuned resource allocation.
// Mixed workloads benefit from worker specialization.
// TODO: Demonstrate worker specialization patterns
func demonstrateWorkerSpecialization() {
	fmt.Println("\n=== Worker Specialization ===")

	// TODO: Create different job types requiring different skills
	// TODO: Launch specialized worker pools for each job type
	// TODO: Show job routing to appropriate worker pools
	// TODO: Demonstrate resource optimization through specialization
	// TODO: Show monitoring of specialized worker utilization
}

// TUTOR: Dynamic worker pools adjust worker count based on load.
// Pool expansion increases capacity during high load periods.
// Pool contraction saves resources during low load periods.
// Load metrics drive scaling decisions.
// Dynamic pools balance performance and resource efficiency.
// TODO: Demonstrate dynamic worker pool scaling
func demonstrateDynamicWorkerPool() {
	fmt.Println("\n=== Dynamic Worker Pool Scaling ===")

	// TODO: Start with minimal worker count
	// TODO: Monitor job queue depth and processing rates
	// TODO: Add workers when queue grows beyond threshold
	// TODO: Remove workers when load decreases
	// TODO: Show scaling decision logic and metrics

	productionRate := atomic.Int32{}
	productionRate.Store(rand.Int31n(100))
	workerRate := 100
	numWorkers := atomic.Int32{}
	numWorkers.Store(3)

	jobChannel := make(chan Job, 1000)

	workerWg := sync.WaitGroup{}
	producerWg := sync.WaitGroup{}

	currWorkers := atomic.Int32{}
	currWorkers.Store(0)

	worker := func(workerId int) {
		defer workerWg.Done()
		defer currWorkers.Add(-1)
		for range jobChannel {
			time.Sleep(time.Duration(workerRate) * time.Millisecond)
			// fmt.Println("worker", workerId, "processing job", job.ID)
			// if workerId is no less than numWorkers, then we need to kill ourself
			if workerId >= int(numWorkers.Load()) {
				return
			}
		}
	}

	producerWg.Add(1)
	go func() {
		defer producerWg.Done()
		count := 0
		for {
			jobChannel <- Job{ID: count, Data: fmt.Sprintf("job %d", count), Priority: rand.Intn(10)}
			count++
			time.Sleep(time.Duration(productionRate.Load()) * time.Millisecond)
		}
	}()

	tickerWg := sync.WaitGroup{}
	tickerWg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				productionRate.Store(rand.Int31n(100))
				waterLevel := float64(len(jobChannel)) / float64(cap(jobChannel))
				if waterLevel > 0.75 {
					numWorkers.Store(int32(8))
				} else if waterLevel < 0.25 {
					numWorkers.Store(int32(2))
				}
				for currWorkers.Load() < numWorkers.Load() {
					workerWg.Add(1)
					currWorkers.Add(1)
					go worker(int(currWorkers.Load()))
				}
				fmt.Println("PROD RATE:", productionRate.Load(), "WORKERS:", currWorkers.Load(), "QUEUE:", len(jobChannel))
			}
		}
	}()

	tickerWg.Wait()
}

// TUTOR: Worker pool shutdown requires coordination between all components.
// New job submission must be stopped first.
// Existing jobs in queue should be processed or drained.
// Workers must complete current jobs before terminating.
// Proper shutdown prevents job loss and resource leaks.
// TODO: Demonstrate proper worker pool shutdown
func demonstrateWorkerPoolShutdown() {
	fmt.Println("\n=== Worker Pool Shutdown ===")
	fmt.Println("\nhave implemented this in other examples")
	// TODO: Show graceful shutdown signal handling
	// TODO: Stop accepting new jobs while preserving existing queue
	// TODO: Allow workers to finish current jobs
	// TODO: Implement timeout-based forced shutdown
	// TODO: Show proper resource cleanup and job accounting

}

// TUTOR: Worker pool error handling requires error aggregation patterns.
// Individual worker errors shouldn't crash the entire pool.
// Failed jobs may need retry mechanisms or dead letter queues.
// Error rates can indicate system health and scaling needs.
// Proper error handling maintains pool stability under adverse conditions.
// TODO: Demonstrate error handling in worker pools
func demonstrateWorkerPoolErrorHandling() {
	fmt.Println("\n=== Worker Pool Error Handling ===")

	// TODO: Show worker error recovery without pool shutdown
	// TODO: Implement job retry mechanisms
	// TODO: Create dead letter queue for permanently failed jobs
	// TODO: Demonstrate error rate monitoring and alerting
	// TODO: Show error-based worker pool health assessment
}

// TUTOR: Worker pool monitoring provides visibility into system performance.
// Track job throughput, worker utilization, and queue depths.
// Monitor processing latencies and error rates.
// Use metrics to identify bottlenecks and optimization opportunities.
// Good monitoring enables proactive performance management.
// TODO: Demonstrate comprehensive worker pool monitoring
func demonstrateWorkerPoolMonitoring() {
	fmt.Println("\n=== Worker Pool Monitoring ===")

	// TODO: Implement job throughput measurement
	// TODO: Track individual worker utilization
	// TODO: Monitor job queue depth and wait times
	// TODO: Measure processing latencies and success rates
	// TODO: Show bottleneck identification and performance tuning
}

// TUTOR: Worker pool patterns form the foundation of scalable concurrent systems.
// These patterns appear in web servers, data processors, and distributed systems.
// Understanding worker pools enables building high-performance applications.
// Proper implementation prevents common concurrency pitfalls.
// Worker pools are essential for production Go programming.
func demonstrateRealWorldWorkerPools() {
	fmt.Println("\n=== Web Server Worker Pool ===")

	// Request represents an HTTP request to process
	type Request struct {
		ID   int
		Path string
		Done chan string
	}

	// Create request channel (buffer prevents blocking)
	requests := make(chan Request, 100)

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		go func(workerID int) {
			for req := range requests {
				// Simulate request processing (database query, etc.)
				fmt.Printf("Worker %d processing request %d: %s\n",
					workerID, req.ID, req.Path)

				time.Sleep(time.Millisecond * 100) // Simulate work

				// Send response back
				req.Done <- fmt.Sprintf("Response from worker %d for %s",
					workerID, req.Path)
			}
		}(i)
	}

	// Simulate incoming HTTP requests
	var wg sync.WaitGroup

	paths := []string{"/users", "/orders", "/products", "/health", "/metrics"}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()

			// Create request
			req := Request{
				ID:   requestID,
				Path: paths[requestID%len(paths)],
				Done: make(chan string),
			}

			// Send to worker pool
			requests <- req

			// Wait for response
			response := <-req.Done
			fmt.Printf("Request %d completed: %s\n", requestID, response)
		}(i)
	}

	wg.Wait()
	close(requests)

	fmt.Println("All requests processed!")
}

func main() {
	fmt.Println("🏗️ Worker Pools - Controlled Concurrency for Maximum Throughput 🏗️")
	fmt.Println("Master resource-efficient concurrent processing")

	// TODO: Implement each demonstration function
	// Build understanding from basic to advanced patterns

	// demonstrateBasicWorkerPool()
	// demonstrateWorkerPoolSizing()
	// demonstrateJobQueueBuffering()
	// demonstrateWorkerSpecialization()
	// demonstrateDynamicWorkerPool()
	// demonstrateWorkerPoolShutdown()
	// demonstrateWorkerPoolErrorHandling()
	// demonstrateWorkerPoolMonitoring()
	demonstrateRealWorldWorkerPools()
}
