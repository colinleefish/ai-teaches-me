package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// TUTOR: Race conditions occur when multiple goroutines access shared data concurrently.
// At least one goroutine modifies the data, creating unpredictable results.
// Race conditions are bugs that may not manifest consistently, making them dangerous.
// Go provides tools to detect and prevent race conditions in concurrent programs.
// Understanding race conditions is essential for writing safe concurrent code.
// TODO: Demonstrate a classic race condition scenario
func demonstrateRaceCondition() {
	fmt.Println("=== Race Condition Demonstration ===")

	// TODO: Create a shared variable that multiple goroutines will modify
	// TODO: Launch multiple goroutines that increment the shared variable
	// TODO: Show that the final result is unpredictable due to race conditions
	// TODO: Run multiple times to show inconsistent results

	sharedCounter := 0

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			sharedCounter++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			sharedCounter++
		}
	}()

	wg.Wait()
	fmt.Println("Shared counter:", sharedCounter)
}

// TUTOR: Go's race detector can automatically find race conditions in your code.
// Use 'go run -race' or 'go build -race' to enable race detection.
// The race detector uses runtime analysis to catch concurrent access violations.
// It reports the exact goroutines and source locations involved in races.
// Race detection is essential for validating concurrent program correctness.
// TODO: Demonstrate how to use Go's race detector
func demonstrateRaceDetector() {
	fmt.Println("\n=== Race Detector Usage ===")

	// TODO: Create code that will trigger the race detector
	// TODO: Show how to run with 'go run -race' to detect races
	// TODO: Explain the race detector output format
	// TODO: Demonstrate that the race detector catches the problem
}

// TUTOR: Mutexes provide mutual exclusion to prevent race conditions.
// sync.Mutex ensures only one goroutine can access protected data at a time.
// Lock() acquires exclusive access, Unlock() releases it.
// Mutexes turn concurrent access into sequential access for critical sections.
// Proper mutex usage eliminates race conditions but may reduce concurrency.
// TODO: Demonstrate race condition prevention with mutexes
func demonstrateMutexSolution() {
	fmt.Println("\n=== Mutex Solution ===")

	// TODO: Use sync.Mutex to protect shared data access
	// TODO: Show that mutex prevents race conditions
	// TODO: Compare results with and without mutex protection
	// TODO: Demonstrate proper Lock/Unlock patterns

	sharedCounter := 0

	wg := sync.WaitGroup{}
	wg.Add(2)

	mu := sync.Mutex{}

	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			mu.Lock()
			sharedCounter++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			mu.Lock()
			sharedCounter++
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("Shared counter:", sharedCounter)

}

// TUTOR: Channels provide race-free communication between goroutines.
// Channel operations are atomic and synchronized by the Go runtime.
// Using channels for data sharing eliminates the need for explicit locks.
// Channels follow the principle: "Don't communicate by sharing memory, share memory by communicating."
// Channel-based solutions are often cleaner than mutex-based approaches.
// TODO: Demonstrate race-free communication using channels
func demonstrateChannelSolution() {
	fmt.Println("\n=== Channel Solution ===")

	// TODO: Use channels to avoid shared mutable state
	// TODO: Show how channels eliminate race conditions naturally
	// TODO: Demonstrate the "share memory by communicating" principle
	// TODO: Compare channel approach with mutex approach

	sum := 0
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			ch1 <- 1
		}
		close(ch1)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			ch2 <- 1
		}
		close(ch2)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			ch3 <- 1
		}
		close(ch3)
	}()

	for {
		select {
		case i, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				sum += i
			}
		case i, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				sum += i
			}
		case i, ok := <-ch3:
			if !ok {
				ch3 = nil
			} else {
				sum += i
			}
		}
		if ch1 == nil && ch2 == nil && ch3 == nil {
			break
		}
	}

	fmt.Println("Sum:", sum)
}

// TUTOR: Read-write mutexes allow concurrent readers but exclusive writers.
// sync.RWMutex optimizes for read-heavy workloads with occasional writes.
// RLock()/RUnlock() for readers, Lock()/Unlock() for writers.
// Multiple readers can proceed simultaneously, but writers get exclusive access.
// RWMutex improves performance when reads vastly outnumber writes.
func demonstrateRWMutex() {
	fmt.Println("\n=== Read-Write Mutex Comparison ===")

	counter := 0

	// Test with regular Mutex - all operations are exclusive
	fmt.Println("Testing with regular Mutex...")
	testWithMutex(&counter)

	// Test with RWMutex - readers can run concurrently
	fmt.Println("\nTesting with RWMutex...")
	testWithRWMutex(&counter)
}

func testWithMutex(counter *int) {
	var mu sync.Mutex
	var wg sync.WaitGroup

	readers := 5
	writers := 1
	operations := 100000

	start := time.Now()

	// Start multiple reader goroutines
	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				mu.Lock()                        // Exclusive lock for read
				_ = *counter                     // Simulate read work
				time.Sleep(time.Nanosecond * 10) // Small delay to show contention
				mu.Unlock()
			}
		}(i)
	}

	// Start writer goroutine
	for i := 0; i < writers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operations/10; j++ { // Fewer writes
				mu.Lock()
				(*counter)++
				time.Sleep(time.Nanosecond * 50) // Writes take longer
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("Mutex: %d readers, %d writers completed in %v\n", readers, writers, duration)
}

func testWithRWMutex(counter *int) {
	var rwmu sync.RWMutex
	var wg sync.WaitGroup

	readers := 5
	writers := 1
	operations := 100000

	start := time.Now()

	// Start multiple reader goroutines
	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < operations; j++ {
				rwmu.RLock()                     // Shared lock for read - can run concurrently!
				_ = *counter                     // Simulate read work
				time.Sleep(time.Nanosecond * 10) // Small delay to show benefit
				rwmu.RUnlock()
			}
		}(i)
	}

	// Start writer goroutine
	for i := 0; i < writers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operations/10; j++ { // Fewer writes
				rwmu.Lock() // Exclusive lock for write
				(*counter)++
				time.Sleep(time.Nanosecond * 50) // Writes take longer
				rwmu.Unlock()
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("RWMutex: %d readers, %d writers completed in %v\n", readers, writers, duration)
}

// TUTOR: Atomic operations provide lock-free synchronization for simple values.
// sync/atomic package offers atomic read/write operations for basic types.
// Atomic operations are faster than mutexes for simple counters and flags.
// They guarantee atomicity without the overhead of lock acquisition.
// Atomic operations are building blocks for more complex lock-free algorithms.
// TODO: Demonstrate atomic operations for race-free counters
func demonstrateAtomicOperations() {
	fmt.Println("\n=== Atomic Operations ===")

	// TODO: Use sync/atomic for lock-free operations
	// TODO: Show atomic increment operations
	// TODO: Compare performance with mutex-based approaches
	// TODO: Demonstrate atomic load/store operations

	atomicCounter := atomic.Int32{}
	counter := 0

	start := time.Now()
	testWithMutexVsAtomic(&counter)
	duration := time.Since(start)
	fmt.Println("Mutex vs Atomic duration:", duration)

	start = time.Now()
	testWithAtomic(&atomicCounter)
	duration = time.Since(start)
	fmt.Println("Atomic duration:", duration)
}

func testWithAtomic(atomicCounter *atomic.Int32) {
	var wg sync.WaitGroup

	operations := 100000000
	numGoroutines := runtime.NumCPU()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operations/numGoroutines; j++ {
				atomicCounter.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Atomic counter:", atomicCounter.Load())
}

func testWithMutexVsAtomic(counter *int) {
	var mu sync.Mutex
	var wg sync.WaitGroup

	operations := 100000000
	numGoroutines := runtime.NumCPU()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < operations/numGoroutines; j++ {
				mu.Lock()
				(*counter)++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Mutex counter:", *counter)
}

func main() {
	fmt.Println("🔒 Go Concurrency: Race Conditions & Safety")
	fmt.Printf("Using %d CPU cores\n", runtime.NumCPU())

	// Build understanding of concurrent safety
	// demonstrateRaceCondition()
	// demonstrateRaceDetector()
	// demonstrateMutexSolution()
	// demonstrateChannelSolution()
	// demonstrateRWMutex()
	demonstrateAtomicOperations()

	fmt.Println("\n✅ Race condition fundamentals complete!")
	fmt.Println("Next: Learn about directional channels for API safety")
}
