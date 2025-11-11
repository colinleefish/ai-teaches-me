// Level 1.1: Goroutines - The Foundation of Go Concurrency
// This file teaches the basic building blocks of goroutines

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"time"
)

// Helper function for visual separation
func Divider(char string) {
	fmt.Println()
	for i := 0; i < 50; i++ {
		fmt.Print(char)
	}
	fmt.Println()
}

// TUTOR: Goroutines are lightweight threads managed by the Go runtime, not the OS.
// They're incredibly cheap - you can create thousands without worrying about memory.
// Think of them as "green threads" - the Go scheduler multiplexes them onto OS threads.
// Use goroutines when you want something to run concurrently without blocking the main flow.
// Key insight: goroutines make concurrent programming as easy as adding the `go` keyword.
// TODO: Create and run basic goroutines using the `go` keyword
func demonstrateBasicGoroutines() {
	fmt.Println("=== Basic Goroutine Creation ===")

	// TODO: Launch a simple anonymous function as a goroutine
	go func() {
		fmt.Println("Hello from goroutine")
	}()
	time.Sleep(1 * time.Second)

	Divider("*")

	// TODO: Launch a named function as a goroutine
	sayHello := func() {
		fmt.Println("I say Hello")
	}
	go sayHello()
	time.Sleep(1 * time.Second)

	Divider("*")

	// TODO: Show that main function doesn't wait for goroutines by default
	go func() {
		fmt.Println("I would never reach the terminal")
	}()
	fmt.Println("Main function continuing...")
	// TODO: Use time.Sleep to give goroutines a chance to run

	Divider("*")

	go func() {
		fmt.Println("I WILL reach the terminal")
	}()
	fmt.Println("Main function continuing...")
	time.Sleep(1 * time.Second)
}

// TUTOR: Goroutines have a lifecycle: created -> scheduled -> running -> finished.
// Unlike OS threads, creating a goroutine doesn't immediately run it - the scheduler decides when.
// Goroutines start with a small stack (2KB) that grows as needed, making them very memory efficient.
// Understanding lifecycle helps you reason about when goroutines actually execute.
// Key insight: goroutines are cooperative - they yield control at certain points.
// TODO: Demonstrate goroutine lifecycle and scheduling behavior
func demonstrateGoroutineLifecycle() {
	fmt.Println("\n=== Goroutine Lifecycle ===")

	// TODO: Create goroutines that announce their start and finish

	go func() {
		defer fmt.Println("Goroutine OUT!")
		fmt.Println("Goroutine IN!")
		time.Sleep(1 * time.Second)
	}()

	fmt.Println("Main function continuing...")
	time.Sleep(2 * time.Second)

	// TODO: Show how goroutines can run in different orders
	numberedHello := func(number int) {
		fmt.Printf("Hello %d\n", number)
	}
	go numberedHello(1)
	go numberedHello(2)
	go numberedHello(3)
	time.Sleep(1 * time.Second)

	Divider("*")

	// TODO: Demonstrate that goroutines don't run instantly when created

	go numberedHello(10086)
	time.Sleep(1 * time.Microsecond)
	fmt.Println("Main function continuing after numberedHello(10086) is triggered...")
	time.Sleep(1 * time.Second)

	Divider("*")
	// TODO: Use runtime.Gosched() to yield control manually

	runtime.GOMAXPROCS(2)
	fmt.Println("🎭 Demonstration: WITH vs WITHOUT Gosched()")

	// Helper function that does CPU-intensive work
	busyWork := func(iterations int) {
		for i := 0; i < iterations; i++ {
			_ = i * i // Simulate CPU work
		}
	}

	fmt.Println("\n--- WITHOUT Gosched() (might run to completion) ---")

	selfishWorker := func(id int) {
		fmt.Printf("🏃 Worker %d: Starting...\n", id)

		for step := 1; step <= 3; step++ {
			fmt.Printf("🏃 Worker %d: Step %d (working hard, not yielding)\n", id, step)
			busyWork(100000) // Do work without yielding
		}

		fmt.Printf("🏁 Worker %d: FINISHED!\n", id)
	}

	go selfishWorker(1)
	go selfishWorker(2)
	go selfishWorker(3)
	go selfishWorker(4)
	time.Sleep(100 * time.Millisecond) // Give them time to run

	fmt.Println("\n--- WITH Gosched() (cooperative yielding) ---")

	politeWorker := func(id int) {
		fmt.Printf("🤝 Worker %d: Starting...\n", id)

		for step := 1; step <= 3; step++ {
			fmt.Printf("🤝 Worker %d: Step %d (working, then yielding)\n", id, step)
			busyWork(100000)  // Do work
			runtime.Gosched() // ✨ YIELD CONTROL - "Others can go now!"
			fmt.Printf("🤝 Worker %d: Back from yield for step %d\n", id, step)
		}

		fmt.Printf("🏁 Worker %d: FINISHED!\n", id)
	}

	go politeWorker(1)
	go politeWorker(2)
	go politeWorker(3)
	go politeWorker(4)
	time.Sleep(100 * time.Millisecond) // Give them time to run

	fmt.Println("\n🎯 Notice the difference:")
	fmt.Println("   WITHOUT Gosched(): Workers might complete entirely before others start")
	fmt.Println("   WITH Gosched(): Workers take turns more fairly")
}

// TUTOR: Anonymous function goroutines are extremely common and powerful for closures.
// You can capture variables from the surrounding scope, but be careful with loop variables!
// Anonymous goroutines are perfect for quick concurrent tasks without defining separate functions.
// The closure behavior lets you pass data naturally without explicit parameters.
// Key insight: goroutines + closures = elegant concurrent programming.
// TODO: Demonstrate anonymous function goroutines and closure behavior
func demonstrateAnonymousGoroutines() {
	fmt.Println("\n=== Anonymous Function Goroutines ===")

	// CLASSIC MISTAKE: Loop variable capture problem
	fmt.Println("❌ Wrong way - captures loop variable:")
	for i := 0; i < 3; i++ {
		fmt.Println("--before go func:", i)
		go func() {
			fmt.Println("--in go func:", i)
			time.Sleep(1 * time.Second)
			fmt.Printf("Wrong: %d\n", i) // Usually prints 3, 3, 3
		}()
		fmt.Println("--after go func:", i)
	}
	time.Sleep(2 * time.Second)

	// CORRECT WAY 1: Pass as parameter
	fmt.Println("\n✅ Correct way 1 - pass as parameter:")
	for i := 0; i < 3; i++ {
		go func(val int) {
			fmt.Printf("Correct: %d\n", val)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)

	// CORRECT WAY 2: Capture in local variable
	fmt.Println("\n✅ Correct way 2 - capture in local variable:")
	for i := 0; i < 3; i++ {
		i := i // Create new variable in loop scope
		go func() {
			fmt.Printf("Correct: %d\n", i)
		}()
	}
	time.Sleep(100 * time.Millisecond)

	// REAL-WORLD EXAMPLE: Processing tasks concurrently (simple version)
	fmt.Println("\n🌍 Real-world example - processing tasks:")
	tasks := []string{"task-1", "task-2", "task-3"}

	for _, task := range tasks {
		go func(t string) {
			// Simulate work
			fmt.Printf("Processing %s...\n", t)
			time.Sleep(time.Duration(100+len(t)*10) * time.Millisecond)
			fmt.Printf("✅ Completed %s\n", t)
		}(task)
	}

	// Wait for all goroutines to finish (simple approach)
	time.Sleep(500 * time.Millisecond)

	// CLOSURE EXAMPLE: Shared variable
	fmt.Println("\n🔒 Closure sharing variable:")
	counter := 0

	increment := func() {
		counter++
		fmt.Printf("Counter: %d\n", counter)
	}

	// This is UNSAFE - race condition on shared counter!
	for i := 0; i < 3; i++ {
		go increment()
	}
	time.Sleep(100 * time.Millisecond)
}

// TUTOR: Passing parameters to goroutines avoids closure pitfalls and makes data flow explicit.
// Each goroutine gets its own copy of the parameters, preventing race conditions.
// Parameter passing is the safe way to give goroutines the data they need.
// This approach makes goroutines more like pure functions - input in, work done.
// Key insight: explicit parameters = safer concurrent programming.
// TODO: Demonstrate proper parameter passing to goroutines
func demonstrateGoroutineParameters() {
	fmt.Println("\n=== Goroutine Parameter Passing ===")

	// TODO: Create goroutines with different parameter types
	// TODO: Show how each goroutine gets its own copy of parameters
	// TODO: Demonstrate passing structs and complex data
	// TODO: Compare parameter passing vs closure capturing

	go func(a int, b int) {
		fmt.Println(a, b)
	}(1, 2) // 1, 2 are passed to the goroutine

	go func(a int, b int) {
		fmt.Println(a, b)
	}(5, 6)

	time.Sleep(1 * time.Second)
}

// TUTOR: The Go scheduler automatically manages goroutines across available CPU cores.
// You can see scheduler behavior with runtime functions and understand performance.
// The scheduler uses M:N threading - many goroutines on fewer OS threads.
// Understanding scheduling helps you write efficient concurrent programs.
// Key insight: trust the scheduler, but understand what it's doing.
// TODO: Demonstrate scheduler behavior and runtime inspection
func demonstrateScheduler() {
	fmt.Println("\n=== Goroutine Scheduler ===")

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// TODO: Show current number of goroutines with runtime.NumGoroutine()
	// TODO: Display CPU count and GOMAXPROCS settings
	// TODO: Create many goroutines and watch scheduler distribute them
	// TODO: Use runtime.Gosched() to manually yield control

	counter := 0

	myGoRoutine := func() {
		sleepTime := time.Duration(rand.Intn(10)) * time.Second
		time.Sleep(sleepTime)
		return
	}

	for {
		go myGoRoutine()
		counter++
		if counter > 100 {
			break
		}
	}

	// This is a way to keep the main goroutine alive

	start := time.Now()
	for {
		runtime.ReadMemStats(&m)
		debug.FreeOSMemory()
		runtime.GC()
		fmt.Println("🔍 Runtime.NumGoroutine():", runtime.NumGoroutine(), "Memory Usage:", m.Alloc)
		time.Sleep(1 * time.Second)
		if time.Since(start) > 10*time.Second {
			break
		}
	}
}

// TUTOR: Goroutines can run indefinitely if not properly managed - this is a "goroutine leak".
// Always ensure goroutines have a way to terminate, or they'll consume memory forever.
// Simple termination strategies: counters, time limits, or external stop signals.
// Proper goroutine lifecycle management is crucial for long-running programs.
// Key insight: every goroutine should have a clear way to stop.
// TODO: Demonstrate goroutine termination patterns
func demonstrateGoroutineTermination() {
	fmt.Println("\n=== Goroutine Termination ===")

	// 1. NATURAL TERMINATION - function ends, goroutine dies
	// fmt.Println("✅ Natural termination:")
	// go func() {
	// 	fmt.Println("I'm doing work...")
	// 	time.Sleep(100 * time.Millisecond)
	// 	fmt.Println("Work done! Goroutine exits naturally.")
	// }()

	// // 2. COUNTER-BASED TERMINATION
	// fmt.Println("\n🔢 Counter-based termination:")
	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Printf("Working... %d/3\n", i+1)
	// 		time.Sleep(50 * time.Millisecond)
	// 	}
	// 	fmt.Println("Counter reached limit, terminating!")
	// }()

	// // 3. TIME-BASED TERMINATION
	// fmt.Println("\n⏰ Time-based termination:")
	// go func() {
	// 	start := time.Now()
	// 	for time.Since(start) < 200*time.Millisecond {
	// 		fmt.Println("Still running...")
	// 		time.Sleep(50 * time.Millisecond)
	// 	}
	// 	fmt.Println("Time limit reached, terminating!")
	// }()

	// 4. SIMPLE SIGNALING (shared variable)
	fmt.Println("\n🚦 Simple signaling termination:")
	stop := false

	go func() {
		for !stop {
			fmt.Println("Monitoring...")
			time.Sleep(30 * time.Millisecond)
		}
		fmt.Println("Stop signal received, terminating!")
	}()

	// Signal to stop after some time
	time.Sleep(120 * time.Millisecond)
	stop = true

	// ❌ INFINITE GOROUTINE (avoid this!)
	// fmt.Println("\n❌ Don't do this - infinite goroutine:")
	// go func() {
	// 	for {
	// 		// This goroutine never terminates!
	// 		// In real code, this causes memory leaks
	// 		time.Sleep(1 * time.Hour) // Sleeps forever
	// 	}
	// }()
	// fmt.Println("Created an infinite goroutine (bad practice!)")

	// Wait for demonstrations to complete
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\nMain function ending - all finite goroutines should be done!")
}

// TUTOR: Monitoring goroutines helps you understand your program's concurrent behavior.
// Runtime package provides tools to inspect goroutine count and behavior.
// Monitoring is essential for debugging concurrency issues and performance.
// Simple monitoring can prevent complex problems in production.
// Key insight: visibility into goroutines = better concurrent programs.
// TODO: Demonstrate basic goroutine monitoring techniques
func demonstrateBasicMonitoring() {
	fmt.Println("\n=== Basic Goroutine Monitoring ===")

	// 1. BASELINE MONITORING
	fmt.Printf("🏁 Starting goroutines: %d\n", runtime.NumGoroutine())

	// 2. CREATE SOME WORKERS WITH LOGGING
	fmt.Println("\n📊 Creating worker goroutines...")
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("🔧 Worker %d: Starting\n", id)
			time.Sleep(time.Duration(100+id*50) * time.Millisecond)
			fmt.Printf("✅ Worker %d: Completed\n", id)
		}(i)
	}

	fmt.Printf("📈 After creating workers: %d goroutines\n", runtime.NumGoroutine())

	// 3. TRACK COMPLETION OVER TIME
	fmt.Println("\n🕐 Monitoring completion...")
	for i := 0; i < 8; i++ {
		count := runtime.NumGoroutine()
		fmt.Printf("⏱️  After %dms: %d active goroutines\n", i*50, count)
		time.Sleep(50 * time.Millisecond)
		if count <= 1 { // Only main goroutine left
			break
		}
	}

	// 4. MEMORY TRACKING WITH GOROUTINES
	fmt.Println("\n💾 Memory impact monitoring:")
	var m runtime.MemStats

	// Before burst
	runtime.ReadMemStats(&m)
	fmt.Printf("Before burst: %d KB\n", m.Alloc/1024)

	// Create burst of short-lived goroutines
	for i := 0; i < 20000; i++ {
		go func(id int) {
			// Quick work
			_ = make([]byte, 1024) // Allocate 1KB
			time.Sleep(10 * time.Millisecond)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	runtime.ReadMemStats(&m)
	fmt.Printf("After burst: %d KB\n", m.Alloc/1024)

	// 5. GOROUTINE LEAK DETECTION
	fmt.Println("\n🚨 Leak detection pattern:")
	baseline := runtime.NumGoroutine()

	// Simulate some work that should not leak
	for i := 0; i < 3; i++ {
		go func() {
			time.Sleep(50 * time.Millisecond)
		}()
	}

	time.Sleep(200 * time.Millisecond) // Wait for completion

	final := runtime.NumGoroutine()
	if final > baseline {
		fmt.Printf("🚨 POTENTIAL LEAK: Started with %d, ended with %d\n", baseline, final)
	} else {
		fmt.Printf("✅ No leaks detected: %d → %d goroutines\n", baseline, final)
	}
}

func main() {
	fmt.Println("🧵 Welcome to Goroutines - Foundation of Go Concurrency! 🧵")
	fmt.Println("Master these basics before moving to coordination and communication")

	// TODO: Implement each demonstration function
	// Start with the simplest concepts and build understanding

	// demonstrateBasicGoroutines()
	// demonstrateGoroutineLifecycle()
	// demonstrateAnonymousGoroutines()
	// demonstrateGoroutineParameters()
	// demonstrateScheduler()
	// demonstrateGoroutineTermination()
	demonstrateBasicMonitoring()

	fmt.Println()
	fmt.Println()
	fmt.Println("==============^=^=^=^=^===============")
	fmt.Println("END OF FILE: goroutines.go")
}

/*
🧵 Goroutine Foundation Concepts:

1. **Creation**: `go functionName()` or `go func(){...}()`
2. **Lifecycle**: Created → Scheduled → Running → Finished
3. **Scheduling**: Go runtime manages when goroutines actually run
4. **Parameters**: Pass data explicitly to avoid closure pitfalls
5. **Termination**: Always provide a way for goroutines to stop
6. **Monitoring**: Use runtime.NumGoroutine() to track active count

🎯 Key Principles:
- Goroutines are cheap - don't fear creating them
- Main function doesn't wait for goroutines automatically
- Use time.Sleep() temporarily to let goroutines run
- Pass parameters explicitly for safety
- Every goroutine should have a termination condition

🚨 Common Beginner Mistakes:
- Forgetting main function exits before goroutines finish
- Capturing loop variables in closures incorrectly
- Creating goroutines without termination conditions
- Not understanding that goroutines run concurrently, not instantly

🔗 What's Next:
After mastering basic goroutines, learn how to coordinate them with WaitGroups!
*/
