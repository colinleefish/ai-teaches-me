// Level 1.2: WaitGroups - Coordinating Goroutines
// This file teaches basic goroutine synchronization using sync.WaitGroup

package main

import (
	"fmt"
	"sync"
	"time"
)

// TUTOR: WaitGroups solve the fundamental problem: "How do I wait for goroutines to finish?"
// Think of WaitGroup as a counter that tracks running goroutines.
// Add(n) increases counter, Done() decreases it, Wait() blocks until counter reaches zero.
// This is the most basic and essential coordination primitive in Go.
// Use WaitGroups when you need to wait for a known number of goroutines to complete.
// TODO: Demonstrate basic WaitGroup usage pattern
func demonstrateBasicWaitGroup() {
	fmt.Println("=== Basic WaitGroup Usage ===")

	// TODO: Create a sync.WaitGroup variable
	// TODO: Launch several goroutines, calling wg.Add(1) for each
	// TODO: Inside each goroutine, use defer wg.Done()
	// TODO: Call wg.Wait() to block until all complete
	// TODO: Show the difference with and without WaitGroup

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine 2")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine 3")
	}()

	wg.Wait()
}

// TUTOR: The WaitGroup lifecycle follows a strict pattern: Add before goroutine, Done inside goroutine, Wait in coordinator.
// Add() must be called before launching the goroutine to avoid race conditions.
// Done() should be deferred immediately inside the goroutine to ensure cleanup.
// Wait() blocks the calling goroutine until the counter reaches zero.
// Understanding this lifecycle prevents common synchronization bugs.
// TODO: Demonstrate proper WaitGroup lifecycle and timing
func demonstrateWaitGroupLifecycle() {
	fmt.Println("\n=== WaitGroup Lifecycle ===")

	// TODO: Show the correct order: Add() → go func() → defer Done() → Wait()
	// TODO: Demonstrate what happens if you Add() after launching goroutine
	// TODO: Show proper cleanup with defer wg.Done()
	// TODO: Illustrate how Wait() blocks until counter reaches zero
	wg := sync.WaitGroup{}
	go func() {
		defer wg.Done()
		wg.Add(1)
		fmt.Println("Hello from goroutine")
	}()
	wg.Wait()
}

// TUTOR: WaitGroups can coordinate any number of goroutines dynamically.
// You can Add() multiple times, and the counter accumulates correctly.
// Add(n) is equivalent to calling Add(1) n times, which is more efficient.
// This flexibility makes WaitGroups perfect for batch processing scenarios.
// Dynamic coordination is key for real-world concurrent programming.
// TODO: Demonstrate dynamic WaitGroup coordination with varying goroutine counts
func demonstrateDynamicCoordination() {
	fmt.Println("\n=== Dynamic WaitGroup Coordination ===")

	// TODO: Create goroutines in a loop with dynamic count
	// TODO: Use wg.Add(1) for each goroutine OR wg.Add(n) once
	// TODO: Show how WaitGroup handles different numbers of goroutines
	// TODO: Demonstrate that Wait() works regardless of goroutine count
}

// TUTOR: WaitGroups coordinate goroutines that do real work, not just printing.
// Each goroutine should perform some meaningful task before calling Done().
// Work simulation helps you understand timing and coordination in realistic scenarios.
// This pattern is the foundation for worker pools and parallel processing.
// Coordinated work is where WaitGroups really shine in practice.
// TODO: Demonstrate WaitGroups coordinating goroutines doing actual work
func demonstrateWorkCoordination() {
	fmt.Println("\n=== Work Coordination with WaitGroups ===")

	// TODO: Create goroutines that simulate different types of work
	// TODO: Use time.Sleep() to simulate varying work durations
	// TODO: Show how WaitGroup waits for the slowest goroutine
	// TODO: Demonstrate that all work completes before proceeding

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		simulateWork(1, 10*time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		simulateWork(2, 5*time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		simulateWork(3, 1*time.Second)
	}()

	wg.Wait()
}

// TUTOR: Multiple WaitGroups can coordinate different groups of goroutines independently.
// This allows for complex coordination patterns: parallel groups, sequential phases, etc.
// Independent WaitGroups help structure complex concurrent programs clearly.
// Think of it as having multiple coordination points in your program.
// Multiple coordinators enable sophisticated concurrent architectures.
// TODO: Demonstrate using multiple WaitGroups for different coordination needs
func demonstrateMultipleWaitGroups() {
	fmt.Println("\n=== Multiple WaitGroups ===")

	// TODO: Create two separate WaitGroups for different tasks
	// TODO: Launch different groups of goroutines with their respective WaitGroups
	// TODO: Show how you can wait for groups independently
	// TODO: Demonstrate sequential coordination: group1, then group2

	wg1 := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}

	wg1.Add(1)
	go func() {
		defer wg1.Done()
		simulateWork(1, 10*time.Second)
	}()

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		simulateWork(2, 5*time.Second)
	}()

	wg1.Wait()
	wg2.Wait()
}

// TUTOR: WaitGroups must be passed by pointer to goroutines to work correctly.
// Passing by value would create copies, and Done() calls wouldn't affect the original.
// Function parameters should use *sync.WaitGroup type for proper sharing.
// Understanding Go's value vs pointer semantics is crucial for WaitGroups.
// Proper parameter passing ensures your coordination actually works.
// TODO: Demonstrate correct WaitGroup parameter passing
func demonstrateWaitGroupParameters() {
	fmt.Println("\n=== WaitGroup Parameter Passing ===")

	// TODO: Create a helper function that takes *sync.WaitGroup parameter
	// TODO: Show the correct way to pass WaitGroup to functions
	// TODO: Demonstrate what happens if you pass by value (spoiler: it breaks)
	// TODO: Use the helper function with proper pointer semantics
}

// TUTOR: WaitGroup counter must never go negative, or your program will panic.
// This happens when Done() is called more times than Add() was called.
// Always match your Add() and Done() calls carefully to avoid panics.
// Understanding counter balance prevents runtime crashes in concurrent programs.
// Balanced Add/Done calls are essential for stable concurrent code.
// TODO: Demonstrate WaitGroup counter balance and potential issues
func demonstrateWaitGroupBalance() {
	fmt.Println("\n=== WaitGroup Counter Balance ===")

	// TODO: Show correct balanced usage: Add(3) with 3 Done() calls
	// TODO: Demonstrate what happens with mismatched Add/Done (in comments)
	// TODO: Show how to safely check WaitGroup state conceptually
	// TODO: Illustrate defensive programming with WaitGroups

	wg := sync.WaitGroup{}
	// Scenario 1: More Add than Done
	wg.Add(5)                 // counter = 5
	go func() { wg.Done() }() // counter = 4
	go func() { wg.Done() }() // counter = 3
	go func() { wg.Done() }() // counter = 2
	// Missing 2 Done() calls
	wg.Wait() // HANGS FOREVER (counter never reaches 0)
}

// Helper function to simulate work
func simulateWork(workerID int, duration time.Duration) {
	fmt.Printf("Worker %d starting work...\n", workerID)
	time.Sleep(duration)
	fmt.Printf("Worker %d finished work\n", workerID)
}

// Helper function demonstrating proper WaitGroup parameter usage
func workerWithWaitGroup(id int, wg *sync.WaitGroup) {
	// TODO: Implement proper WaitGroup usage in a helper function
	// TODO: Use defer wg.Done() for guaranteed cleanup
	// TODO: Do some work here
}

func main() {
	fmt.Println("🤝 Welcome to WaitGroups - Goroutine Coordination! 🤝")
	fmt.Println("Learn to coordinate goroutines before moving to communication")

	// TODO: Implement each demonstration function
	// Build understanding of coordination before communication

	// demonstrateBasicWaitGroup()
	// demonstrateWaitGroupLifecycle()
	// demonstrateDynamicCoordination()
	// demonstrateWorkCoordination()
	// demonstrateMultipleWaitGroups()
	// demonstrateWaitGroupParameters()
	demonstrateWaitGroupBalance()

	fmt.Println("\n🎉 Congratulations! You can coordinate goroutines!")
	fmt.Println("Next: Learn communication with channels.go")
}

/*
🤝 WaitGroup Foundation Concepts:

1. **Purpose**: Wait for a known number of goroutines to complete
2. **Methods**: Add(n) increases counter, Done() decreases, Wait() blocks
3. **Lifecycle**: Add() → go func() → defer Done() → Wait()
4. **Parameter Passing**: Always use *sync.WaitGroup (pointer)
5. **Balance**: Add() and Done() calls must match exactly

🎯 Essential Patterns:
```go
var wg sync.WaitGroup
wg.Add(1)          // Before launching goroutine
go func() {
    defer wg.Done() // First line inside goroutine
    // Do work here
}()
wg.Wait()          // Block until all complete
```

🚨 Common Beginner Mistakes:
- Calling Add() after launching goroutine (race condition)
- Forgetting Done() call (program hangs forever)
- Passing WaitGroup by value instead of pointer
- Mismatched Add/Done calls (panic or hang)
- Not using defer for Done() (skipped on early returns)

🔗 What's Next:
WaitGroups coordinate goroutines, but they can't communicate data.
Next, learn channels for safe data passing between goroutines!
*/
