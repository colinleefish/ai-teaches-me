package main

import (
	"fmt"
	"runtime"
	"time"
)

func busyWork(name string, duration time.Duration) {
	fmt.Printf("%s: Starting work...\n", name)
	start := time.Now()

	// Do some actual work (not just sleep)
	for time.Since(start) < duration {
		// Busy work - using CPU
		for i := 0; i < 1000000; i++ {
			_ = i * i
		}
	}

	fmt.Printf("%s: Finished after %v\n", name, time.Since(start))
}

func main() {
	fmt.Printf("Your computer has %d CPU cores\n", runtime.NumCPU())
	fmt.Printf("Go can use %d cores for goroutines\n", runtime.GOMAXPROCS(0))

	fmt.Println("\n=== Running 3 tasks that each take ~1 second ===")

	start := time.Now()

	// Start 3 goroutines
	go busyWork("Worker 1", 1*time.Second)
	go busyWork("Worker 2", 1*time.Second)
	go busyWork("Worker 3", 1*time.Second)

	// Wait for them to finish
	time.Sleep(2 * time.Second)

	fmt.Printf("\nTotal time for all 3 tasks: %v\n", time.Since(start))
	fmt.Println("Notice: If you have multiple cores, they might finish around the same time!")
}
