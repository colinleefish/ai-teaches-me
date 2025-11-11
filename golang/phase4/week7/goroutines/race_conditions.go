// Week 7: Race Conditions and Concurrent Safety
// This file demonstrates race conditions and how to detect and prevent them

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Shared variables for demonstration
var (
	unsafeCounter  int
	safeCounter    int64
	counterMutex   sync.Mutex
	counterRWMutex sync.RWMutex
)

// TODO: Demonstrate what race conditions are
func demonstrateRaceCondition() {
	fmt.Println("=== Race Condition Example ===")

	// TODO: Create a race condition with shared variable
	// Reset counter
	unsafeCounter = 0

	// TODO: Launch multiple goroutines that increment the same variable
	// Each goroutine should increment the counter many times
	// Show how the final result is unpredictable

	// TODO: Use WaitGroup to wait for all goroutines
	var wg sync.WaitGroup
	numGoroutines := 10
	incrementsPerGoroutine := 1000

	// TODO: Implement the race condition
	// Each goroutine increments unsafeCounter
	// Show that final result is less than expected

	fmt.Printf("Expected: %d, Actual: %d\n", numGoroutines*incrementsPerGoroutine, unsafeCounter)
	fmt.Println("Run with 'go run -race race_conditions.go' to detect race!")
}

// TODO: Demonstrate race detection
func demonstrateRaceDetection() {
	fmt.Println("\n=== Race Detection ===")

	// TODO: Show how to use Go's race detector
	// Explain how to run: go run -race race_conditions.go
	// Show what the race detector output looks like

	// TODO: Create a deliberate race condition
	// Use a shared map that multiple goroutines modify

	sharedMap := make(map[int]int)
	var wg sync.WaitGroup

	// TODO: Launch goroutines that modify the map concurrently
	// This will trigger race detector warnings
}

// TODO: Fix race conditions with Mutex
func demonstrateMutexSolution() {
	fmt.Println("\n=== Mutex Solution ===")

	// TODO: Show how to fix race conditions using sync.Mutex
	// Reset safe counter
	safeCounter = 0

	// TODO: Create the same scenario as race condition demo
	// But use mutex to protect the shared variable

	safeIncrement := func() {
		// TODO: Implement safe increment using mutex
		// 1. Lock the mutex
		// 2. Increment the counter
		// 3. Unlock the mutex
	}

	// TODO: Launch multiple goroutines using safe increment
	// Show that the result is now predictable
}

// TODO: Demonstrate RWMutex for read-heavy scenarios
func demonstrateRWMutex() {
	fmt.Println("\n=== RWMutex for Read-Heavy Scenarios ===")

	// TODO: Show when to use RWMutex vs regular Mutex
	// Create a scenario with many readers and few writers

	data := make(map[string]int)
	data["counter"] = 0

	// TODO: Implement readers that only read the data
	reader := func(id int, wg *sync.WaitGroup) {
		// TODO: Use RLock/RUnlock for reading
		// Read the counter value multiple times
	}

	// TODO: Implement writers that modify the data
	writer := func(id int, wg *sync.WaitGroup) {
		// TODO: Use Lock/Unlock for writing
		// Increment the counter
	}

	// TODO: Launch many readers and few writers
	// Show the performance difference vs regular mutex
}

// TODO: Demonstrate atomic operations
func demonstrateAtomicOperations() {
	fmt.Println("\n=== Atomic Operations ===")

	// TODO: Show how to use sync/atomic for simple operations
	// Reset atomic counter
	atomic.StoreInt64(&safeCounter, 0)

	// TODO: Create goroutines that use atomic operations
	atomicIncrement := func(wg *sync.WaitGroup) {
		// TODO: Use atomic.AddInt64 to increment safely
		// Show how atomic operations avoid locks for simple cases
	}

	// TODO: Compare performance: atomic vs mutex vs unsafe
	// Measure time for each approach
}

// TODO: Demonstrate complex race conditions
func demonstrateComplexRaces() {
	fmt.Println("\n=== Complex Race Conditions ===")

	// TODO: Show race conditions in more complex scenarios

	// TODO: Scenario 1: Race in slice operations
	// Multiple goroutines appending to the same slice

	// TODO: Scenario 2: Race in channel operations
	// Closing channels while other goroutines are sending

	// TODO: Scenario 3: Race in interface values
	// Multiple goroutines modifying interface variables

	// TODO: Show how to detect and fix each type
}

// TODO: Demonstrate happens-before relationships
func demonstrateHappensBefore() {
	fmt.Println("\n=== Happens-Before Relationships ===")

	// TODO: Explain memory model and happens-before
	// Show examples of guaranteed ordering vs race conditions

	// TODO: Example 1: Channel send happens-before receive
	// TODO: Example 2: Goroutine creation happens-before execution
	// TODO: Example 3: WaitGroup.Done() happens-before Wait() returns

	// TODO: Show examples where ordering is NOT guaranteed
}

// TODO: Demonstrate data race vs race condition
func demonstrateDataRaceVsRaceCondition() {
	fmt.Println("\n=== Data Race vs Race Condition ===")

	// TODO: Explain the difference:
	// Data race: Concurrent access to same memory location, at least one write
	// Race condition: Behavior depends on timing of events

	// TODO: Example of data race (detected by race detector)

	// TODO: Example of race condition without data race
	// Use channels to show race condition that race detector won't catch
}

// TODO: Implement lock-free programming examples
func demonstrateLockFreeProgramming() {
	fmt.Println("\n=== Lock-Free Programming ===")

	// TODO: Show lock-free data structures using atomic operations

	// TODO: Implement a simple lock-free stack
	type LockFreeStack struct {
		// TODO: Use atomic pointer operations
	}

	// TODO: Implement a simple lock-free queue
	// Show compare-and-swap operations

	// TODO: Discuss when lock-free is appropriate vs complex
}

// TODO: Demonstrate debugging race conditions
func demonstrateDebuggingRaces() {
	fmt.Println("\n=== Debugging Race Conditions ===")

	// TODO: Show debugging techniques
	// 1. Race detector usage
	// 2. Adding logging to identify race patterns
	// 3. Using runtime.Gosched() to increase race probability
	// 4. Stress testing with higher goroutine counts

	// TODO: Show how to write race-free code from the start
	// Design patterns that avoid races
}

// TODO: Performance comparison of different approaches
func performanceComparison() {
	fmt.Println("\n=== Performance Comparison ===")

	// TODO: Benchmark different approaches:
	// 1. Unsafe (with race conditions)
	// 2. Mutex protection
	// 3. RWMutex protection
	// 4. Atomic operations
	// 5. Channel-based coordination

	// TODO: Show when each approach is appropriate
	// Consider contention level, read/write ratio, complexity
}

// TODO: Best practices for concurrent safety
func demonstrateBestPractices() {
	fmt.Println("\n=== Best Practices for Concurrent Safety ===")

	// TODO: Show best practices:

	// TODO: 1. Design for concurrency from the start
	// TODO: 2. Minimize shared mutable state
	// TODO: 3. Use channels for communication, not sharing
	// TODO: 4. Keep critical sections small
	// TODO: 5. Always use defer for unlocking
	// TODO: 6. Avoid nested locks
	// TODO: 7. Test with race detector regularly

	// TODO: Show examples of each practice
}

// Helper function to measure execution time
func measureTime(name string, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	fmt.Printf("%s took: %v\n", name, duration)
}

// Helper function to run with race detection info
func runWithRaceInfo(name string, fn func()) {
	fmt.Printf("\n--- %s ---\n", name)
	fmt.Println("💡 Run with: go run -race race_conditions.go")
	fn()
}

func main() {
	fmt.Println("⚠️  Welcome to Race Conditions! ⚠️")
	fmt.Println("This file teaches you about concurrent safety in Go")
	fmt.Printf("Running on %d CPUs with GOMAXPROCS=%d\n", runtime.NumCPU(), runtime.GOMAXPROCS(0))

	// TODO: Implement each demonstration function
	// Start with basic race conditions and work up to complex scenarios

	demonstrateRaceCondition()
	// demonstrateRaceDetection()
	// demonstrateMutexSolution()
	// demonstrateRWMutex()
	// demonstrateAtomicOperations()
	// demonstrateComplexRaces()
	// demonstrateHappensBefore()
	// demonstrateDataRaceVsRaceCondition()
	// demonstrateLockFreeProgramming()
	// demonstrateDebuggingRaces()
	// performanceComparison()
	// demonstrateBestPractices()

	fmt.Println("\n🎉 Congratulations! You understand concurrent safety!")
	fmt.Println("Next: Learn goroutine patterns in goroutine_patterns.go")
}

/*
🔍 Key Concepts to Remember:

1. **Race Condition**: Outcome depends on timing of concurrent operations
2. **Data Race**: Concurrent access to memory, at least one write, no synchronization
3. **Race Detector**: Use 'go run -race' to detect data races
4. **Mutex**: Mutual exclusion for protecting critical sections
5. **RWMutex**: Allows multiple readers OR single writer
6. **Atomic**: Lock-free operations for simple data types
7. **Happens-Before**: Memory model guarantees about operation ordering

🔒 Synchronization Tools:
- **sync.Mutex**: Exclusive access
- **sync.RWMutex**: Reader-writer locks
- **sync/atomic**: Lock-free atomic operations
- **Channels**: Communication and synchronization
- **sync.WaitGroup**: Wait for goroutine completion

🚨 Common Race Conditions:
- Incrementing shared counters
- Appending to shared slices
- Writing to shared maps
- Closing channels while sending
- Modifying interface values

💡 Prevention Strategies:
- Use race detector during development
- Minimize shared mutable state
- Prefer channels over shared memory
- Keep critical sections small
- Always use defer for unlocking
- Design for concurrency from start

🎯 Next Steps:
- Learn channel-based coordination
- Master goroutine patterns
- Build race-free concurrent applications
*/
