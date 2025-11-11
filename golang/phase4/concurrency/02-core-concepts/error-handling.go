package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Result struct {
	Value int
	Err   error
}

// TUTOR: Error handling in concurrent programs requires careful coordination.
// Goroutines cannot return errors directly to their callers.
// Errors must be communicated through channels or shared data structures.
// Proper error handling prevents silent failures in concurrent systems.
// Error propagation patterns ensure failures are detected and handled.
// TODO: Demonstrate basic error communication in goroutines
func demonstrateBasicErrorCommunication() {
	fmt.Println("=== Basic Error Communication ===")

	// TODO: Create goroutine that can produce errors
	// TODO: Use channels to send errors back to main goroutine
	// TODO: Show proper error checking and handling patterns
	// TODO: Demonstrate different error communication strategies

	resultCh := make(chan Result, 1)

	// this is a routine that produce either a result or an error
	go func() {
		defer close(resultCh)
		time.Sleep(1 * time.Second)
		result := rand.Intn(2)
		if result == 0 {
			resultCh <- Result{Err: errors.New("error")}
		} else {
			resultCh <- Result{Value: 42}
		}
	}()

	result := <-resultCh
	if result.Err != nil {
		fmt.Println("Error:", result.Err)
	} else {
		fmt.Println("Result:", result.Value)
	}
}

// TUTOR: Error channels provide dedicated paths for error communication.
// Separate error channels keep errors distinct from normal data flow.
// Error channels can be buffered to prevent blocking on error conditions.
// This pattern makes error handling explicit and systematic.
// Dedicated error channels improve code clarity and maintainability.
// TODO: Demonstrate dedicated error channel patterns
func demonstrateErrorChannels() {
	fmt.Println("\n=== Dedicated Error Channels ===")

	// EXAMPLE: Processing multiple files concurrently
	files := []string{"config.txt", "data.csv", "invalid.xml", "readme.md", "missing.json"}

	// Separate channels for results and errors
	dataCh := make(chan string, len(files))
	errCh := make(chan error, len(files))

	// Process files concurrently
	for _, filename := range files {
		go func(file string) {
			// Simulate file processing
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			// Simulate different outcomes
			switch {
			case strings.Contains(file, "missing"):
				errCh <- fmt.Errorf("file not found: %s", file)
			case strings.Contains(file, "invalid"):
				errCh <- fmt.Errorf("invalid format: %s", file)
			default:
				dataCh <- fmt.Sprintf("processed: %s", file)
			}
		}(filename)
	}

	// Collect results with proper handling
	processed := 0
	errors := []error{}
	results := []string{}

	for processed < len(files) {
		select {
		case result := <-dataCh:
			results = append(results, result)
			processed++
			fmt.Printf("✅ %s\n", result)

		case err := <-errCh:
			errors = append(errors, err)
			processed++
			fmt.Printf("❌ %v\n", err)

		case <-time.After(200 * time.Millisecond):
			fmt.Println("⏰ Timeout waiting for results")
			break
		}
	}

	// Summary
	fmt.Printf("\n📊 Summary: %d successful, %d errors\n", len(results), len(errors))

	// ALTERNATIVE: Error aggregation pattern
	fmt.Println("\n🔄 Alternative: Error aggregation:")

	allErrorsCh := make(chan error, 10)
	dataStreamCh := make(chan int, 10)

	// Multiple workers, some failing
	for i := 0; i < 5; i++ {
		go func(id int) {
			if id%2 == 0 {
				// Success case
				dataStreamCh <- id * 10
			} else {
				// Error case
				allErrorsCh <- fmt.Errorf("worker %d failed", id)
			}
		}(i)
	}

	// Collect everything
	time.Sleep(50 * time.Millisecond)
	close(dataStreamCh)
	close(allErrorsCh)

	fmt.Println("📥 Data received:")
	for data := range dataStreamCh {
		fmt.Printf("  Data: %d\n", data)
	}

	fmt.Println("🚨 Errors received:")
	for err := range allErrorsCh {
		fmt.Printf("  Error: %v\n", err)
	}
}

// TUTOR: Result types combine values and errors in single channel messages.
// Result structs contain both the value and any error that occurred.
// This pattern simplifies channel management for operations that can fail.
// Result types make error handling more systematic and less error-prone.
// Single-channel communication reduces coordination complexity.
// TODO: Demonstrate result type patterns for error handling
func demonstrateResultTypes() {
	fmt.Println("\n=== Result Type Patterns ===")
	fmt.Println("\nImplemented in demonstrateBasicErrorCommunication")
	// TODO: Define result type that contains value and error
	// TODO: Show goroutines sending result objects through channels
	// TODO: Demonstrate result unpacking and error checking
	// TODO: Compare with separate channel approaches
}

// TUTOR: Error aggregation collects multiple errors from concurrent operations.
// When multiple goroutines can fail, collect all errors for comprehensive reporting.
// Error slices or maps can accumulate errors from multiple sources.
// Partial failures require deciding whether to continue or abort operations.
// Error aggregation enables informed decision-making in concurrent systems.
// TODO: Demonstrate error aggregation from multiple goroutines
func demonstrateErrorAggregation() {
	fmt.Println("\n=== Error Aggregation ===")

	// TODO: Launch multiple goroutines that can fail
	// TODO: Collect errors from all goroutines
	// TODO: Show different aggregation strategies (first error, all errors)
	// TODO: Demonstrate decision-making based on error patterns

	wg := sync.WaitGroup{}
	workers := runtime.NumCPU()

	errCh := make(chan error, workers)
	resultCh := make(chan int, workers)

	critical := false
	results := []int{}
	mu := sync.Mutex{}

	// Start monitoring BEFORE workers
	monitorDone := make(chan bool)
	go func() {
		errCounter := 0
		completed := 0

		for completed < workers && !critical {
			select {
			case err := <-errCh:
				fmt.Println("❌ Error:", err)
				errCounter++
				completed++
				if errCounter > runtime.NumCPU()-5 { // Quit early if >2 errors
					mu.Lock()
					critical = true
					mu.Unlock()
					fmt.Println("🚨 TOO MANY ERRORS - ABORTING!")
				}
			case result := <-resultCh:
				fmt.Printf("✅ Success: worker %d\n", result)
				mu.Lock()
				results = append(results, result)
				mu.Unlock()
				completed++
			}
		}
		monitorDone <- true
	}()

	// Start workers AFTER monitoring is ready
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			duration := time.Duration(rand.Intn(500)) * time.Millisecond
			time.Sleep(duration)

			// Check if we should abort early
			mu.Lock()
			shouldAbort := critical
			mu.Unlock()

			if shouldAbort {
				fmt.Printf("⏹️ Worker %d: Aborting due to critical errors\n", id)
				return
			}

			seed := rand.Intn(100)
			if seed < 50 {
				errCh <- fmt.Errorf("worker %d failed", id)
			} else {
				resultCh <- id
			}
		}(i)
	}

	// Wait for monitor to finish (guaranteed to happen)
	<-monitorDone
	fmt.Println("🏁 Monitor finished")

	// Clean up
	wg.Wait()
	close(errCh)
	close(resultCh)

	mu.Lock()
	fmt.Printf("📊 Final: %d results, critical=%v\n", len(results), critical)
	fmt.Println("Results:", results)
	mu.Unlock()

}

// TUTOR: Fail-fast patterns stop processing when critical errors occur.
// Some errors indicate system state that makes continued processing futile.
// Context cancellation can signal other goroutines to stop work.
// Fail-fast prevents wasted computation and resource usage.
// Early termination on critical errors improves system responsiveness.
// TODO: Demonstrate fail-fast error handling patterns
func demonstrateFailFast() {
	fmt.Println("\n=== Fail-Fast Patterns ===")

	// TODO: Create scenario with critical and non-critical errors
	// TODO: Show early termination on critical failures
	// TODO: Demonstrate context-based cancellation signaling
	// TODO: Show resource cleanup on early termination
}

// TUTOR: Graceful degradation continues with reduced functionality during errors.
// Non-critical failures shouldn't stop entire system operation.
// Error handling can enable fallback mechanisms and alternative paths.
// Resilient systems adapt to partial failures gracefully.
// Graceful degradation improves system availability and user experience.
// TODO: Demonstrate graceful degradation patterns
func demonstrateGracefulDegradation() {
	fmt.Println("\n=== Graceful Degradation ===")

	// EXAMPLE: User profile page with multiple data sources
	// userID := 123

	// Different services with different importance
	type ServiceResult struct {
		Name     string
		Data     string
		Error    error
		Critical bool // Some services are critical, others optional
	}

	results := make(chan ServiceResult, 4)

	// Critical service: User basic info (MUST work)
	go func() {
		time.Sleep(50 * time.Millisecond)
		// Simulate: this one always works
		results <- ServiceResult{
			Name:     "UserService",
			Data:     "John Doe, john@example.com",
			Critical: true,
		}
	}()

	// Optional service: User avatar (nice to have)
	go func() {
		time.Sleep(100 * time.Millisecond)
		// Simulate: 50% chance of failure
		if rand.Intn(2) == 0 {
			results <- ServiceResult{
				Name:     "AvatarService",
				Error:    errors.New("avatar service down"),
				Critical: false,
			}
		} else {
			results <- ServiceResult{
				Name:     "AvatarService",
				Data:     "avatar_url_123.jpg",
				Critical: false,
			}
		}
	}()

	// Optional service: Recommendations (nice to have)
	go func() {
		time.Sleep(80 * time.Millisecond)
		// Simulate: this one often fails
		results <- ServiceResult{
			Name:     "RecommendationService",
			Error:    errors.New("recommendations unavailable"),
			Critical: false,
		}
	}()

	// Optional service: Activity feed (nice to have)
	go func() {
		time.Sleep(30 * time.Millisecond)
		results <- ServiceResult{
			Name:     "ActivityService",
			Data:     "3 new notifications",
			Critical: false,
		}
	}()

	// Collect results and decide what to show user
	profile := make(map[string]string)
	criticalFailed := false
	optionalFailures := 0
	servicesReceived := 0

	for servicesReceived < 4 {
		result := <-results
		servicesReceived++

		if result.Error != nil {
			if result.Critical {
				criticalFailed = true
				fmt.Printf("🚨 CRITICAL FAILURE: %s - %v\n", result.Name, result.Error)
			} else {
				optionalFailures++
				fmt.Printf("⚠️  Optional failure: %s - %v\n", result.Name, result.Error)
			}
		} else {
			profile[result.Name] = result.Data
			fmt.Printf("✅ %s: %s\n", result.Name, result.Data)
		}
	}

	// Decision making: Can we still serve the user?
	fmt.Println("\n🎯 System Decision:")
	if criticalFailed {
		fmt.Println("❌ Cannot serve user - critical service failed")
	} else {
		fmt.Println("✅ Serving user with available data:")
		for service, data := range profile {
			fmt.Printf("  - %s: %s\n", service, data)
		}

		if optionalFailures > 0 {
			fmt.Printf("📉 Degraded experience: %d optional services failed\n", optionalFailures)
		} else {
			fmt.Println("🌟 Full experience: all services working")
		}
	}
}

// TUTOR: Error recovery patterns attempt to recover from transient failures.
// Retry mechanisms can overcome temporary network or resource issues.
// Circuit breakers prevent cascade failures in distributed systems.
// Recovery strategies depend on error types and system requirements.
// Proper recovery improves system resilience and availability.
// TODO: Demonstrate error recovery and retry patterns
func demonstrateErrorRecovery() {
	fmt.Println("\n=== Error Recovery Patterns ===")

	// TODO: Show retry mechanisms for transient failures
	// TODO: Demonstrate exponential backoff strategies
	// TODO: Show circuit breaker patterns for failure isolation
	// TODO: Illustrate different recovery strategies for different errors
}

// TUTOR: Error logging and monitoring are essential for concurrent systems.
// Errors from goroutines must be explicitly logged or they disappear.
// Structured logging helps correlate errors across concurrent operations.
// Error metrics enable monitoring and alerting for system health.
// Proper observability is crucial for debugging concurrent systems.
// TODO: Demonstrate error logging and monitoring patterns
func demonstrateErrorLogging() {
	fmt.Println("\n=== Error Logging and Monitoring ===")

	// Log entry structure
	type LogEntry struct {
		Timestamp time.Time
		Level     string // INFO, WARN, ERROR
		WorkerID  int
		Message   string
		Error     error
	}

	logCh := make(chan LogEntry, 20)

	// Centralized logger goroutine
	loggerDone := make(chan bool)
	go func() {
		fmt.Println("📝 Logger started...")
		for log := range logCh {
			timeStr := log.Timestamp.Format("15:04:05.000")
			if log.Error != nil {
				fmt.Printf("[%s] %s Worker-%d: %s | Error: %v\n",
					timeStr, log.Level, log.WorkerID, log.Message, log.Error)
			} else {
				fmt.Printf("[%s] %s Worker-%d: %s\n",
					timeStr, log.Level, log.WorkerID, log.Message)
			}
		}
		loggerDone <- true
	}()

	// PATTERN 1: Workers send logs to central logger
	fmt.Println("\n🔧 Pattern 1: Central logging channel")
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			// Worker logs its lifecycle
			logCh <- LogEntry{
				Timestamp: time.Now(),
				Level:     "INFO",
				WorkerID:  workerID,
				Message:   "Starting work",
			}

			// Simulate work with random outcome
			workTime := time.Duration(rand.Intn(200)) * time.Millisecond
			time.Sleep(workTime)

			if rand.Intn(3) == 0 {
				// Failure case
				logCh <- LogEntry{
					Timestamp: time.Now(),
					Level:     "ERROR",
					WorkerID:  workerID,
					Message:   "Task failed",
					Error:     fmt.Errorf("network timeout after %v", workTime),
				}
			} else {
				// Success case
				logCh <- LogEntry{
					Timestamp: time.Now(),
					Level:     "INFO",
					WorkerID:  workerID,
					Message:   fmt.Sprintf("Task completed in %v", workTime),
				}
			}

		}(i)
	}

	// PATTERN 2: Batch processing with error counting
	fmt.Println("\n📊 Pattern 2: Error metrics collection")
	errorCounts := make(map[string]int)
	mu := sync.Mutex{}

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			operation := []string{"database", "api", "file"}[rand.Intn(3)]

			if rand.Intn(4) == 0 { // 25% failure rate
				errType := fmt.Sprintf("%s_error", operation)

				logCh <- LogEntry{
					Timestamp: time.Now(),
					Level:     "ERROR",
					WorkerID:  workerID,
					Message:   fmt.Sprintf("%s operation failed", operation),
					Error:     fmt.Errorf("%s connection failed", operation),
				}

				// Thread-safe error counting
				mu.Lock()
				errorCounts[errType]++
				mu.Unlock()
			} else {
				logCh <- LogEntry{
					Timestamp: time.Now(),
					Level:     "INFO",
					WorkerID:  workerID,
					Message:   fmt.Sprintf("%s operation succeeded", operation),
				}
			}
		}(i)
	}

	// Wait for all workers, then close log channel
	wg.Wait()
	close(logCh)

	// Wait for logger to finish
	<-loggerDone

	// Print error summary
	fmt.Println("\n📈 Error Summary:")
	mu.Lock()
	if len(errorCounts) == 0 {
		fmt.Println("✅ No errors detected")
	} else {
		for errType, count := range errorCounts {
			fmt.Printf("  %s: %d occurrences\n", errType, count)
		}
	}
	mu.Unlock()
}

// TUTOR: Panic recovery in goroutines requires explicit defer/recover patterns.
// Unhandled panics in goroutines crash the entire program.
// defer/recover must be used within each goroutine that might panic.
// Recovered panics should be converted to errors for proper handling.
// Panic recovery is a last resort for unexpected failures.
// TODO: Demonstrate panic recovery in concurrent scenarios
func demonstratePanicRecovery() {
	fmt.Println("\n=== Panic Recovery ===")

	// TODO: Show goroutines that might panic
	// TODO: Demonstrate defer/recover patterns within goroutines
	// TODO: Show conversion of recovered panics to errors
	// TODO: Illustrate when panic recovery is appropriate
}

// TUTOR: Testing error scenarios in concurrent code requires special techniques.
// Error injection helps verify error handling paths are correct.
// Race conditions in error handling can cause subtle bugs.
// Timeout patterns prevent tests from hanging on error conditions.
// Comprehensive error testing improves system reliability.
// TODO: Demonstrate testing patterns for concurrent error handling
func demonstrateErrorTesting() {
	fmt.Println("\n=== Testing Error Scenarios ===")

	// TODO: Show error injection techniques for testing
	// TODO: Demonstrate testing of different error handling paths
	// TODO: Show timeout patterns for error scenario tests
	// TODO: Illustrate comprehensive error handling validation
}

func main() {
	fmt.Println("🚨 Go Concurrency: Error Handling")

	// Build understanding of concurrent error management
	// demonstrateBasicErrorCommunication()
	// demonstrateErrorChannels()
	// demonstrateResultTypes()
	// demonstrateErrorAggregation()
	// demonstrateFailFast()
	// demonstrateGracefulDegradation()
	// demonstrateErrorRecovery()
	// demonstrateErrorLogging()
	// demonstratePanicRecovery()
	// demonstrateErrorTesting()

	fmt.Println("\n✅ Concurrent error handling fundamentals complete!")
	fmt.Println("Next: Ready for Level 3 - Basic Patterns!")
}
