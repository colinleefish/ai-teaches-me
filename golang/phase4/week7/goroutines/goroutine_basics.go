// Week 7: Goroutine Basics
// This file demonstrates the fundamentals of goroutines in Go

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// Basic goroutine creation and execution
func demonstrateBasicGoroutines() {
	fmt.Println("=== Basic Goroutines ===")

	// Simple anonymous function goroutine
	go func() {
		fmt.Println("Hello from goroutine!")
	}()

	// Goroutines with parameters
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d: Hello from goroutine!\n", id)
		}(i)
	}

	fmt.Println("Main function continuing...")
	time.Sleep(1 * time.Second) // Wait for goroutines to complete
}

// Goroutine lifecycle demonstration
func demonstrateGoroutineLifecycle() {
	fmt.Println("\n=== Goroutine Lifecycle ===")

	// Create goroutines that show their lifecycle
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d started\n", id)
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	time.Sleep(4 * time.Second) // Wait for all to complete
}

func slowFunction(name string, duration time.Duration) {
	fmt.Printf("%s: Starting performance...\n", name)
	time.Sleep(duration)
	fmt.Printf("%s: Finished performance!\n", name)
}

// Compare synchronous vs asynchronous execution
func compareGoroutinesVsFunctions() {
	fmt.Println("\n=== Goroutines vs Regular Functions ===")

	// Synchronous execution
	fmt.Println("Calling functions synchronously:")
	start := time.Now()
	slowFunction("Task 1", 1*time.Second)
	slowFunction("Task 2", 1*time.Second)
	slowFunction("Task 3", 1*time.Second)
	fmt.Printf("Synchronous total time: %v\n", time.Since(start))

	// Asynchronous execution with goroutines
	fmt.Println("\nCalling functions as goroutines:")
	start = time.Now()
	go slowFunction("Async Task 1", 1*time.Second)
	go slowFunction("Async Task 2", 1*time.Second)
	go slowFunction("Async Task 3", 1*time.Second)
	time.Sleep(1500 * time.Millisecond) // Wait for all to complete
	fmt.Printf("Asynchronous total time: %v\n", time.Since(start))
}

// Demonstrate goroutine scheduling and concurrency
func demonstrateScheduling() {
	fmt.Println("\n=== Goroutine Scheduling ===")

	// Show interleaved execution
	fmt.Println("Goroutine Interleaving:")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			for j := 1; j <= 5; j++ {
				fmt.Printf("Goroutine %d: Step %d\n", id, j)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)

	// Show system information
	fmt.Printf("\nSystem Information:\n")
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}

// Anonymous function patterns and closure pitfalls
func demonstrateAnonymousFunctions() {
	fmt.Println("\n=== Anonymous Function Goroutines ===")

	// Simple anonymous function
	go func() {
		fmt.Println("Anonymous function goroutine finished")
	}()

	// Anonymous function with parameters
	for i := 1; i <= 2; i++ {
		go func(id int) {
			fmt.Printf("Anonymous function goroutine %d finished\n", id)
		}(i)
	}

	time.Sleep(500 * time.Millisecond)

	// RACE CONDITION DEMO: Counter with many goroutines
	fmt.Println("\n=== Race Condition Demonstration ===")
	counter := 0

	// Create many goroutines to increase race condition chances
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter++ // NOT atomic! Race condition here!
			}
		}()
	}

	time.Sleep(1 * time.Second) // Wait for all to finish

	fmt.Printf("Expected counter value: %d\n", 1000*100) // Should be 100,000
	fmt.Printf("Actual counter value: %d\n", counter)    // Will likely be less
	fmt.Println("Note: Actual value is usually less due to race conditions!")
}

// Error handling patterns in goroutines
func demonstrateErrorHandling() {
	fmt.Println("\n=== Error Handling in Goroutines ===")

	// Safe panic handling with recover
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		fmt.Println("About to panic in goroutine...")
		panic("Panic in goroutine")
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main function continuing after panic recovery...")

	// Simple error collection patterns
	fmt.Println("\n--- Simple Error Collection Patterns ---")

	// Pattern 1: Global error variable
	var hasError bool
	var errorMessage string

	go func() {
		time.Sleep(200 * time.Millisecond)
		hasError = true
		errorMessage = "Something went wrong in goroutine"
	}()

	time.Sleep(500 * time.Millisecond)

	if hasError {
		fmt.Printf("Error occurred: %s\n", errorMessage)
	} else {
		fmt.Println("No errors")
	}

	// Pattern 2: Error type with completion flag
	var capturedError error
	done := false

	go func() {
		time.Sleep(200 * time.Millisecond)
		capturedError = fmt.Errorf("file not found")
		done = true
	}()

	// Wait for completion
	for !done {
		time.Sleep(50 * time.Millisecond)
	}

	if capturedError != nil {
		fmt.Printf("Captured error: %v\n", capturedError)
	} else {
		fmt.Println("No error occurred")
	}
}

// Resource management and goroutine leak prevention
func demonstrateResourceManagement() {
	fmt.Println("\n=== Resource Management and Leak Prevention ===")

	// Show goroutine leak examples
	fmt.Println("\n--- Goroutine Leak Examples ---")

	fmt.Println("Creating leaked goroutines (they run forever):")
	for i := 0; i < 3; i++ {
		go func(id int) {
			for {
				fmt.Printf("Leaked goroutine %d is running...\n", id)
				time.Sleep(2 * time.Second)
				// This never stops - goroutine leak!
			}
		}(i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Main continues, but leaked goroutines run forever...")

	// Show proper termination patterns
	fmt.Println("\n--- How to Avoid Leaks ---")

	// Pattern 1: Counter/limit
	go func() {
		counter := 0
		for counter < 3 { // EXIT CONDITION!
			fmt.Printf("Working... %d/3\n", counter+1)
			counter++
			time.Sleep(300 * time.Millisecond)
		}
		fmt.Println("Goroutine finished cleanly!")
	}()

	time.Sleep(1 * time.Second)

	// Pattern 2: Stop flag
	stopFlag := false
	go func() {
		count := 0
		for !stopFlag && count < 5 { // Check flag AND have backup limit
			fmt.Printf("Working with stop flag... %d\n", count+1)
			count++
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("Goroutine stopped!")
	}()

	time.Sleep(500 * time.Millisecond)
	stopFlag = true // Signal to stop
	time.Sleep(300 * time.Millisecond)

	// Pattern 3: Time-based limits
	startTime := time.Now()
	go func() {
		count := 0
		for time.Since(startTime) < 1*time.Second {
			fmt.Printf("Working with time limit... %d\n", count+1)
			count++
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("Goroutine stopped due to time limit!")
	}()

	time.Sleep(1200 * time.Millisecond)
}

// Performance comparison between sequential and concurrent execution
func performanceComparison() {
	fmt.Println("\n=== Performance Comparison ===")

	// Sequential vs Concurrent with meaningful work
	fmt.Println("\n--- Sequential vs Concurrent Execution ---")

	// Sequential execution
	fmt.Println("Running tasks sequentially...")
	start := time.Now()

	for i := 0; i < 4; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("Sequential task %d completed\n", i+1)
	}

	sequentialTime := time.Since(start)
	fmt.Printf("Sequential execution took: %v\n", sequentialTime)

	// Concurrent execution with atomic counter (no race condition)
	fmt.Println("\nRunning tasks concurrently...")
	start = time.Now()
	finished := int32(0)

	for i := 0; i < 4; i++ {
		go func(taskId int) {
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("Concurrent task %d completed\n", taskId+1)
			atomic.AddInt32(&finished, 1) // Thread-safe increment
		}(i)
	}

	// Wait for all to finish using atomic operations
	for atomic.LoadInt32(&finished) < 4 {
		time.Sleep(10 * time.Millisecond)
	}

	concurrentTime := time.Since(start)
	fmt.Printf("Concurrent execution took: %v\n", concurrentTime)
	fmt.Printf("Speedup: %.2fx faster\n", float64(sequentialTime)/float64(concurrentTime))

	// Goroutine overhead demonstration
	fmt.Println("\n--- Goroutine Overhead Analysis ---")

	numTasks := 1000
	fmt.Printf("Testing with %d small tasks:\n", numTasks)

	// Sequential small tasks
	start = time.Now()
	for i := 0; i < numTasks; i++ {
		_ = i * 2 // Tiny work
	}
	sequentialSmall := time.Since(start)
	fmt.Printf("Sequential small tasks: %v\n", sequentialSmall)

	// Concurrent small tasks
	start = time.Now()
	finished = 0

	for i := 0; i < numTasks; i++ {
		go func(n int) {
			_ = n * 2 // Same tiny work
			atomic.AddInt32(&finished, 1)
		}(i)
	}

	for atomic.LoadInt32(&finished) < int32(numTasks) {
		time.Sleep(1 * time.Millisecond)
	}

	concurrentSmall := time.Since(start)
	fmt.Printf("Concurrent small tasks: %v\n", concurrentSmall)

	if sequentialSmall < concurrentSmall {
		fmt.Printf("Sequential was %.2fx faster for small tasks!\n",
			float64(concurrentSmall)/float64(sequentialSmall))
		fmt.Println("Lesson: Don't use goroutines for tiny tasks - overhead isn't worth it")
	}
}

// Best practices for goroutine usage
func demonstrateBestPractices() {
	fmt.Println("\n=== Goroutine Best Practices ===")

	// Proper goroutine naming and identification
	fmt.Println("\n--- Tip 1: Use Clear Goroutine Identification ---")

	for i := 0; i < 3; i++ {
		go func(workerID int) {
			fmt.Printf("[WORKER-%d] Starting work...\n", workerID)

			for task := 0; task < 2; task++ {
				fmt.Printf("[WORKER-%d] Processing task %d\n", workerID, task+1)
				time.Sleep(200 * time.Millisecond)
			}

			fmt.Printf("[WORKER-%d] Completed all tasks\n", workerID)
		}(i)
	}

	time.Sleep(1 * time.Second)

	// CPU-based worker count
	fmt.Println("\n--- Tip 2: Use CPU-Based Worker Limits ---")

	maxWorkers := runtime.NumCPU() // Optimal for CPU-intensive work
	fmt.Printf("Using %d workers (detected %d CPU cores)\n", maxWorkers, runtime.NumCPU())

	workFinished := int32(0)
	totalWork := 20

	for i := 0; i < maxWorkers; i++ {
		go func(workerID int) {
			fmt.Printf("[CPU-WORKER-%d] Starting...\n", workerID)

			// Simulate work distribution
			for atomic.LoadInt32(&workFinished) < int32(totalWork) {
				currentWork := atomic.AddInt32(&workFinished, 1)
				if currentWork <= int32(totalWork) {
					fmt.Printf("[CPU-WORKER-%d] Processing work item %d\n", workerID, currentWork)
					time.Sleep(100 * time.Millisecond)
				}
			}

			fmt.Printf("[CPU-WORKER-%d] Finished\n", workerID)
		}(i)
	}

	// Wait for all work to complete
	for atomic.LoadInt32(&workFinished) < int32(totalWork) {
		time.Sleep(50 * time.Millisecond)
	}

	time.Sleep(200 * time.Millisecond) // Let final messages print
}

// Monitoring and debugging goroutines
func demonstrateMonitoring() {
	fmt.Println("\n=== Monitoring and Debugging ===")

	// Monitor goroutine count
	fmt.Printf("Starting goroutines: %d\n", runtime.NumGoroutine())

	// Create some goroutines to monitor
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("🐾 Worker %d starting...\n", id)
			time.Sleep(1 * time.Second)
			fmt.Printf("🐾 Worker %d finished\n", id)
		}(i)
	}

	fmt.Printf("After creating 5 goroutines: %d\n", runtime.NumGoroutine())
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Mid-execution: %d\n", runtime.NumGoroutine())
	time.Sleep(1 * time.Second)
	fmt.Printf("After completion: %d\n", runtime.NumGoroutine())

	// Memory monitoring
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Current memory allocated: %.2f KB\n", float64(memStats.Alloc)/1024)

	// Timeout detection pattern
	fmt.Println("\n--- Timeout Detection Example ---")

	finished := false

	go func() {
		fmt.Println("[TASK] Starting work that might hang...")
		time.Sleep(500 * time.Millisecond) // Simulate work
		fmt.Println("[TASK] Work completed")
		finished = true
	}()

	// Monitor with timeout
	maxWait := 1 * time.Second
	start := time.Now()

	for !finished && time.Since(start) < maxWait {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[MONITOR] Waiting... (%.1fs elapsed)\n", time.Since(start).Seconds())
	}

	if !finished {
		fmt.Println("⚠️ WARNING: Task may be hanging (exceeded timeout)")
	} else {
		fmt.Println("✅ Task completed within timeout")
	}

	fmt.Println("\n🔧 Key Monitoring Tips:")
	fmt.Println("- Use runtime.NumGoroutine() to track goroutine count")
	fmt.Println("- Add unique IDs to goroutines for easier debugging")
	fmt.Println("- Monitor memory usage with runtime.MemStats")
	fmt.Println("- Set timeouts to detect hanging goroutines")
	fmt.Printf("- Current GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}

func main() {
	fmt.Println("🧵 Welcome to Goroutines! 🧵")
	fmt.Println("This file teaches you the fundamentals of Go's lightweight threads")

	// Run all demonstrations
	demonstrateBasicGoroutines()
	demonstrateGoroutineLifecycle()
	compareGoroutinesVsFunctions()
	demonstrateScheduling()
	demonstrateAnonymousFunctions()
	demonstrateErrorHandling()
	demonstrateResourceManagement()
	performanceComparison()
	demonstrateBestPractices()
	demonstrateMonitoring()

	fmt.Println("\n🎉 Congratulations! You've learned goroutine basics!")
	fmt.Println("Next: Learn synchronization with channels and WaitGroups")
}

/*
🔍 Key Concepts Covered:

1. **Goroutine Creation**: Use 'go' keyword before any function call
2. **Lightweight Threads**: Start with ~2KB stack, grow as needed
3. **Concurrency vs Parallelism**: Concurrent execution on multiple CPU cores
4. **Race Conditions**: Multiple goroutines accessing shared data unsafely
5. **Error Handling**: Use defer/recover within goroutines
6. **Resource Management**: Prevent goroutine leaks with proper termination
7. **Performance**: When to use goroutines vs sequential execution
8. **Best Practices**: Naming, CPU-based limits, monitoring
9. **Debugging**: Track goroutine count, use timeouts, unique IDs

🚨 Common Pitfalls:
- Race conditions when sharing data between goroutines
- Goroutine leaks from infinite loops without exit conditions
- Using goroutines for tiny tasks (overhead > benefit)
- Not handling panics in goroutines (crashes entire program)
- Creating too many goroutines without bounds

🎯 Next Steps:
- Learn channels for safe communication between goroutines
- Master sync.WaitGroup for proper synchronization
- Understand context.Context for cancellation and timeouts
- Explore goroutine patterns: worker pools, pipelines, fan-out/fan-in
*/
