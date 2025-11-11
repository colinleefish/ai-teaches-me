// Retry with Backoff: Handling transient failures
// Key concepts: Exponential backoff, jitter, retry strategies

package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// RetryConfig defines retry behavior
type RetryConfig struct {
	MaxAttempts   int
	BaseDelay     time.Duration
	MaxDelay      time.Duration
	BackoffFactor float64
	Jitter        bool
}

// DefaultRetryConfig provides sensible defaults
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxAttempts:   3,
		BaseDelay:     100 * time.Millisecond,
		MaxDelay:      10 * time.Second,
		BackoffFactor: 2.0,
		Jitter:        true,
	}
}

// RetryableError indicates if an error should be retried
type RetryableError struct {
	Err       error
	Retryable bool
}

func (re *RetryableError) Error() string {
	return re.Err.Error()
}

// IsRetryable checks if an error should be retried
func IsRetryable(err error) bool {
	if re, ok := err.(*RetryableError); ok {
		return re.Retryable
	}
	// Default: retry on unknown errors
	return true
}

// Retry executes a function with retry logic
func Retry(ctx context.Context, config RetryConfig, fn func() error) error {
	var lastErr error

	for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
		fmt.Printf("🔄 Attempt %d/%d", attempt, config.MaxAttempts)

		err := fn()
		if err == nil {
			fmt.Println(" - ✅ SUCCESS")
			return nil
		}

		lastErr = err
		fmt.Printf(" - ❌ FAILED: %v", err)

		// Check if error is retryable
		if !IsRetryable(err) {
			fmt.Println(" (non-retryable)")
			return err
		}

		// Don't wait after last attempt
		if attempt == config.MaxAttempts {
			fmt.Println(" (max attempts reached)")
			break
		}

		// Calculate backoff delay
		delay := calculateBackoff(config, attempt-1)
		fmt.Printf(" - ⏳ Waiting %v before retry...\n", delay)

		// Wait with context cancellation support
		select {
		case <-time.After(delay):
			// Continue to next attempt
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return lastErr
}

func calculateBackoff(config RetryConfig, attempt int) time.Duration {
	// Exponential backoff: baseDelay * backoffFactor^attempt
	delay := float64(config.BaseDelay) * math.Pow(config.BackoffFactor, float64(attempt))

	// Apply maximum delay cap
	if delay > float64(config.MaxDelay) {
		delay = float64(config.MaxDelay)
	}

	// Add jitter to avoid thundering herd
	if config.Jitter {
		jitter := rand.Float64()*0.3 + 0.85 // Random between 0.85 and 1.15
		delay *= jitter
	}

	return time.Duration(delay)
}

// Advanced retry with different strategies
type RetryStrategy int

const (
	StrategyFixed RetryStrategy = iota
	StrategyLinear
	StrategyExponential
	StrategyExponentialJitter
)

func RetryWithStrategy(ctx context.Context, maxAttempts int, strategy RetryStrategy, baseDelay time.Duration, fn func() error) error {
	var lastErr error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("🎯 Strategy %v - Attempt %d/%d", strategy, attempt, maxAttempts)

		err := fn()
		if err == nil {
			fmt.Println(" - ✅ SUCCESS")
			return nil
		}

		lastErr = err
		fmt.Printf(" - ❌ FAILED: %v", err)

		if attempt == maxAttempts {
			fmt.Println(" (max attempts reached)")
			break
		}

		var delay time.Duration

		switch strategy {
		case StrategyFixed:
			delay = baseDelay
		case StrategyLinear:
			delay = baseDelay * time.Duration(attempt)
		case StrategyExponential:
			delay = time.Duration(float64(baseDelay) * math.Pow(2, float64(attempt-1)))
		case StrategyExponentialJitter:
			exponential := time.Duration(float64(baseDelay) * math.Pow(2, float64(attempt-1)))
			jitter := rand.Float64()*0.3 + 0.85
			delay = time.Duration(float64(exponential) * jitter)
		}

		fmt.Printf(" - ⏳ Waiting %v...\n", delay)

		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return lastErr
}

// Simulated unreliable service
type UnreliableService struct {
	failureRate    float64
	transientError bool
}

func NewUnreliableService(failureRate float64, transientError bool) *UnreliableService {
	return &UnreliableService{
		failureRate:    failureRate,
		transientError: transientError,
	}
}

func (us *UnreliableService) Call() error {
	// Simulate processing time
	time.Sleep(50 * time.Millisecond)

	if rand.Float64() < us.failureRate {
		if us.transientError {
			return &RetryableError{
				Err:       errors.New("temporary network timeout"),
				Retryable: true,
			}
		} else {
			return &RetryableError{
				Err:       errors.New("permanent authentication failure"),
				Retryable: false,
			}
		}
	}

	return nil
}

// Batch retry with concurrency
func batchRetryDemo() {
	fmt.Println("\n=== Concurrent Batch Retry ===")

	service := NewUnreliableService(0.6, true)
	config := RetryConfig{
		MaxAttempts:   4,
		BaseDelay:     200 * time.Millisecond,
		MaxDelay:      2 * time.Second,
		BackoffFactor: 1.5,
		Jitter:        true,
	}

	// Process multiple items concurrently
	items := []string{"item1", "item2", "item3"}
	results := make(chan string, len(items))

	for _, item := range items {
		go func(itemName string) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			err := Retry(ctx, config, func() error {
				fmt.Printf("   📦 Processing %s\n", itemName)
				return service.Call()
			})

			if err != nil {
				results <- fmt.Sprintf("%s: FAILED - %v", itemName, err)
			} else {
				results <- fmt.Sprintf("%s: SUCCESS", itemName)
			}
		}(item)
	}

	// Collect results
	for i := 0; i < len(items); i++ {
		fmt.Println("📋", <-results)
	}
}

func main() {
	fmt.Println("=== Retry with Backoff Pattern ===")

	// Example 1: Basic retry with configuration
	fmt.Println("\n1. Basic Retry with Exponential Backoff:")

	unreliableService := NewUnreliableService(0.7, true) // 70% failure, retryable
	config := DefaultRetryConfig()
	config.MaxAttempts = 4

	ctx := context.Background()
	err := Retry(ctx, config, func() error {
		return unreliableService.Call()
	})

	if err != nil {
		fmt.Printf("🚨 Final result: FAILED - %v\n", err)
	} else {
		fmt.Println("🎉 Final result: SUCCESS")
	}

	// Example 2: Non-retryable errors
	fmt.Println("\n2. Non-Retryable Error Handling:")

	permanentFailService := NewUnreliableService(0.8, false) // Non-retryable errors

	err = Retry(ctx, config, func() error {
		return permanentFailService.Call()
	})

	if err != nil {
		fmt.Printf("🚨 Final result: FAILED - %v\n", err)
	}

	// Example 3: Different retry strategies
	fmt.Println("\n3. Different Retry Strategies:")

	moderateFailService := NewUnreliableService(0.5, true)
	baseDelay := 200 * time.Millisecond

	strategies := []RetryStrategy{
		StrategyFixed,
		StrategyLinear,
		StrategyExponential,
		StrategyExponentialJitter,
	}

	strategyNames := map[RetryStrategy]string{
		StrategyFixed:             "Fixed",
		StrategyLinear:            "Linear",
		StrategyExponential:       "Exponential",
		StrategyExponentialJitter: "Exponential+Jitter",
	}

	for _, strategy := range strategies {
		fmt.Printf("\n📈 Testing %s Strategy:\n", strategyNames[strategy])

		err := RetryWithStrategy(ctx, 4, strategy, baseDelay, func() error {
			return moderateFailService.Call()
		})

		if err != nil {
			fmt.Printf("   Result: FAILED\n")
		} else {
			fmt.Printf("   Result: SUCCESS\n")
		}
	}

	// Example 4: Context cancellation
	fmt.Println("\n4. Context Cancellation:")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	slowFailService := NewUnreliableService(0.9, true)
	longRetryConfig := RetryConfig{
		MaxAttempts:   10,
		BaseDelay:     500 * time.Millisecond,
		MaxDelay:      5 * time.Second,
		BackoffFactor: 2.0,
		Jitter:        false,
	}

	err = Retry(ctx, longRetryConfig, func() error {
		return slowFailService.Call()
	})

	if err != nil {
		fmt.Printf("🚨 Result: %v\n", err)
	}

	// Example 5: Concurrent batch processing
	batchRetryDemo()

	fmt.Println("\n✅ Retry with backoff demo complete")
}
