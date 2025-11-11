// Timeout & Cancellation: Handling long-running operations
// Key concepts: time.After, context.Context, graceful cancellation
//
// DISCUSSION INSIGHTS:
// 1. Timeouts are RACES between channels - work vs time.After vs ctx.Done()
// 2. ctx.Done() is a BROADCAST SIGNAL - when context cancels, ALL listeners
//    across the application unblock simultaneously (channel closure behavior)
// 3. It's not about sending data, but about the STOP SIGNAL (channel closure)
// 4. 'default' in select means "if nothing else ready, do this immediately"
// 5. Coordinated shutdown: multiple goroutines share same context = fire alarm effect
// 6. Infinite loops + graceful drainage handle unpredictable timeout timing
// 7. Context prevents goroutine leaks - without it, workers run forever!

package main

import (
	"context"
	"fmt"
	"time"
)

// Basic timeout with time.After
// INSIGHT: This is a RACE between two channels - workDone vs time.After
// Whichever channel becomes ready first "wins" the select statement
func doWorkWithTimeout(workDuration time.Duration, timeout time.Duration) {
	fmt.Printf("Starting work (duration: %v, timeout: %v)\n", workDuration, timeout)

	workDone := make(chan string)

	// Start work in separate goroutine
	go func() {
		time.Sleep(workDuration) // Simulate work
		workDone <- "Work completed successfully"
	}()

	// RACE: work completion vs timeout timer
	// If both are ready simultaneously, Go picks randomly (50/50 chance)
	select {
	case result := <-workDone:
		fmt.Println("✅", result) // Work won the race
	case <-time.After(timeout):
		fmt.Println("❌ Work timed out") // Timeout won the race
	}
}

// Context-based cancellation
// INSIGHT: Similar race pattern, but using ctx.Done() instead of time.After
// ctx.Done() is a BROADCAST SIGNAL - when context cancels, ALL listeners
// across the entire application get unblocked simultaneously!
func doWorkWithContext(ctx context.Context, workDuration time.Duration) string {
	workDone := make(chan string)
	go func() {
		defer close(workDone)
		// RACE: context cancellation vs work completion
		select {
		case <-ctx.Done():
			// Context cancelled - this could be from timeout, manual cancel, etc.
			workDone <- ctx.Err().Error() // ctx.Err() tells us WHY it was cancelled
			break
		case <-time.After(workDuration):
			// Work completed before context was cancelled
			workDone <- "Work Completed"
			break
		}
	}()

	result := <-workDone
	return result
}

// Multiple workers with shared context - demonstrates COORDINATED SHUTDOWN
// INSIGHT: All workers share the SAME context - like a fire alarm!
// When context cancels, ALL workers stop simultaneously.
func coordinatedWork(ctx context.Context) {
	numWorkers := 3
	results := make(chan string, numWorkers)

	// Start workers - they all share the same ctx for coordination
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for { // Infinite loop until context cancellation
				select {
				case <-ctx.Done():
					// BROADCAST SIGNAL: When context closes, ALL workers get this simultaneously!
					// ctx.Done() channel closure unblocks ALL listeners at once
					results <- fmt.Sprintf("Worker %d stopped: %v", workerID, ctx.Err())
					return // Worker exits cleanly
				default:
					// INSIGHT: 'default' means "if no other case is ready, do this immediately"
					// - If context NOT cancelled → default runs → worker keeps working
					// - If context IS cancelled → ctx.Done() fires → worker stops
					time.Sleep(100 * time.Millisecond)
					results <- fmt.Sprintf("Worker %d produced result", workerID)
				}
			}
		}(i)
	}

	// Let workers run briefly
	time.Sleep(350 * time.Millisecond)

	// IMPROVED: Collect results until context is cancelled
	// INSIGHT: Changed from fixed loop (numWorkers*2) to infinite loop
	// This handles unpredictable timeout gracefully - no matter when context cancels!
	for {
		select {
		case result := <-results:
			fmt.Println(result) // Keep printing results as they come
		case <-ctx.Done():
			// SAME CONTEXT: Main loop also listens to ctx.Done() - coordinated shutdown!
			// This demonstrates that ctx.Done() is a BROADCAST - workers AND main get signal
			fmt.Println("Main: stopping due to context cancellation")

			// GRACEFUL DRAINAGE: Workers might have sent final results before stopping
			// Non-blocking drain ensures we don't lose any final messages
			fmt.Println("Draining remaining results...")
			for {
				select {
				case result := <-results:
					fmt.Println("Final:", result) // Collect any remaining results
				default:
					// INSIGHT: 'default' here means "no more results available"
					// This prevents blocking - we exit cleanly when channel is empty
					fmt.Println("All results collected. Shutdown complete.")
					return
				}
			}
		}
	}
}

// Timeout with cleanup - demonstrates MULTIPLE GOROUTINES listening to SAME context
func doWorkWithCleanup(workDuration, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cleanup := make(chan struct{})

	// INSIGHT: This cleanup goroutine ALSO listens to ctx.Done()
	// Same context, multiple listeners - when timeout hits, BOTH this
	// goroutine AND doWorkWithContext() get the signal simultaneously!
	go func() {
		defer close(cleanup)

		select {
		case <-ctx.Done():
			// This runs at the SAME TIME as the ctx.Done() in doWorkWithContext()
			// demonstrating the broadcast nature of context cancellation
			fmt.Println("🧹 Starting cleanup due to timeout...")
			time.Sleep(100 * time.Millisecond) // Cleanup work
			fmt.Println("🧹 Cleanup completed")
		}
	}()

	// Do work - this will ALSO listen to the same ctx.Done()
	err := doWorkWithContext(ctx, workDuration)
	if err != "" {
		fmt.Printf("❌ Work failed: %s\n", err)
		<-cleanup // Wait for cleanup to complete
	} else {
		fmt.Println("✅ Work completed successfully")
		cancel() // Clean shutdown
	}
}

func main() {
	fmt.Println("=== Timeout & Cancellation Patterns ===")

	// Example 1: Basic timeout
	fmt.Println("\n1. Basic Timeout:")
	doWorkWithTimeout(200*time.Millisecond, 300*time.Millisecond) // Success
	doWorkWithTimeout(400*time.Millisecond, 300*time.Millisecond) // Timeout

	// Example 2: Context cancellation
	fmt.Println("\n2. Context Cancellation:")
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(250 * time.Millisecond)
		fmt.Println("🛑 Cancelling context...")
		cancel()
	}()

	err := doWorkWithContext(ctx, 500*time.Millisecond)
	if err != "" {
		fmt.Printf("Work cancelled: %s\n", err)
	}

	// Example 3: Coordinated cancellation
	fmt.Println("\n3. Coordinated Worker Cancellation:")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel2()

	coordinatedWork(ctx2)

	// Example 4: Timeout with cleanup
	fmt.Println("\n4. Timeout with Cleanup:")
	doWorkWithCleanup(600*time.Millisecond, 400*time.Millisecond) // Will timeout and cleanup
}
