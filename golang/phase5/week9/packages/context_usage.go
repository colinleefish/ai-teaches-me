// Week 9: Context Usage for Cancellation and Timeouts
// This file demonstrates the context package for request cancellation and timeouts

package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// TODO: Demonstrate basic context creation and usage
func demonstrateBasicContext() {
	fmt.Println("=== Basic Context Usage ===")
	
	// TODO: Background context (root context)
	ctx := context.Background()
	fmt.Printf("Background context: %v\n", ctx)
	
	// TODO: TODO context (placeholder)
	todoCtx := context.TODO()
	fmt.Printf("TODO context: %v\n", todoCtx)
	
	// TODO: Context with value
	valueCtx := context.WithValue(ctx, "user", "alice")
	fmt.Printf("Context with value: %v\n", valueCtx.Value("user"))
	
	// TODO: Show context hierarchy
	userCtx := context.WithValue(ctx, "userID", 123)
	requestCtx := context.WithValue(userCtx, "requestID", "req-456")
	
	fmt.Printf("UserID: %v\n", requestCtx.Value("userID"))
	fmt.Printf("RequestID: %v\n", requestCtx.Value("requestID"))
	fmt.Printf("Missing key: %v\n", requestCtx.Value("missing"))
}

// TODO: Demonstrate context with timeout
func demonstrateContextTimeout() {
	fmt.Println("\n=== Context with Timeout ===")
	
	// TODO: Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Always call cancel to release resources
	
	// TODO: Simulate work that respects context
	simulateWork := func(ctx context.Context, workName string, duration time.Duration) {
		fmt.Printf("Starting %s (duration: %v)\n", workName, duration)
		
		select {
		case <-time.After(duration):
			fmt.Printf("%s completed successfully\n", workName)
		case <-ctx.Done():
			fmt.Printf("%s cancelled: %v\n", workName, ctx.Err())
		}
	}
	
	// TODO: Test with work that completes in time
	simulateWork(ctx, "fast work", 1*time.Second)
	
	// TODO: Test with work that times out
	simulateWork(ctx, "slow work", 3*time.Second)
	
	// TODO: Show context deadline
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Printf("Context deadline: %v\n", deadline)
		fmt.Printf("Time until deadline: %v\n", time.Until(deadline))
	}
}

// TODO: Demonstrate context with cancellation
func demonstrateContextCancellation() {
	fmt.Println("\n=== Context with Cancellation ===")
	
	// TODO: Create cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	
	// TODO: Start background work
	var wg sync.WaitGroup
	
	// Worker 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker 1 cancelled at iteration %d: %v\n", i, ctx.Err())
				return
			default:
				fmt.Printf("Worker 1: iteration %d\n", i)
				time.Sleep(500 * time.Millisecond)
			}
		}
		fmt.Println("Worker 1 completed all iterations")
	}()
	
	// Worker 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker 2 cancelled at iteration %d: %v\n", i, ctx.Err())
				return
			default:
				fmt.Printf("Worker 2: iteration %d\n", i)
				time.Sleep(300 * time.Millisecond)
			}
		}
		fmt.Println("Worker 2 completed all iterations")
	}()
	
	// TODO: Cancel after some time
	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling context...")
	cancel()
	
	// TODO: Wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers finished")
}

// TODO: Demonstrate context with deadline
func demonstrateContextDeadline() {
	fmt.Println("\n=== Context with Deadline ===")
	
	// TODO: Create context with specific deadline
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	// TODO: Long-running operation that checks context
	longOperation := func(ctx context.Context) error {
		for i := 0; i < 10; i++ {
			// Check if context is done
			if ctx.Err() != nil {
				return ctx.Err()
			}
			
			fmt.Printf("Operation step %d/10\n", i+1)
			time.Sleep(500 * time.Millisecond)
		}
		return nil
	}
	
	// TODO: Execute operation
	start := time.Now()
	err := longOperation(ctx)
	elapsed := time.Since(start)
	
	if err != nil {
		fmt.Printf("Operation failed after %v: %v\n", elapsed, err)
	} else {
		fmt.Printf("Operation completed successfully in %v\n", elapsed)
	}
}

// TODO: Demonstrate context values
func demonstrateContextValues() {
	fmt.Println("\n=== Context Values ===")
	
	// TODO: Context keys should be unique types
	type contextKey string
	
	const (
		userIDKey    contextKey = "userID"
		requestIDKey contextKey = "requestID"
		traceIDKey   contextKey = "traceID"
	)
	
	// TODO: Create context with multiple values
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDKey, 12345)
	ctx = context.WithValue(ctx, requestIDKey, "req-abc-123")
	ctx = context.WithValue(ctx, traceIDKey, "trace-xyz-789")
	
	// TODO: Function that uses context values
	processRequest := func(ctx context.Context) {
		// Extract values from context
		userID, ok := ctx.Value(userIDKey).(int)
		if !ok {
			fmt.Println("No user ID found in context")
			return
		}
		
		requestID := ctx.Value(requestIDKey).(string)
		traceID := ctx.Value(traceIDKey).(string)
		
		fmt.Printf("Processing request:\n")
		fmt.Printf("  User ID: %d\n", userID)
		fmt.Printf("  Request ID: %s\n", requestID)
		fmt.Printf("  Trace ID: %s\n", traceID)
		
		// Simulate processing
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Request processed successfully")
	}
	
	// TODO: Process request with context
	processRequest(ctx)
	
	// TODO: Best practices for context values
	fmt.Println("\nContext value best practices:")
	fmt.Println("- Use custom types for keys to avoid collisions")
	fmt.Println("- Store request-scoped data, not optional parameters")
	fmt.Println("- Don't store configuration or dependencies")
	fmt.Println("- Keep values immutable")
}

// TODO: Demonstrate context in HTTP requests
func demonstrateHTTPWithContext() {
	fmt.Println("\n=== HTTP Requests with Context ===")
	
	// TODO: HTTP request with timeout
	makeRequestWithTimeout := func(url string, timeout time.Duration) error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return fmt.Errorf("creating request: %w", err)
		}
		
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			if ctx.Err() == context.DeadlineExceeded {
				return fmt.Errorf("request timed out after %v", timeout)
			}
			return fmt.Errorf("request failed: %w", err)
		}
		defer resp.Body.Close()
		
		fmt.Printf("HTTP request successful: %s\n", resp.Status)
		return nil
	}
	
	// TODO: Test with fast timeout (will timeout)
	fmt.Println("Making request with 100ms timeout...")
	err := makeRequestWithTimeout("https://httpbin.org/delay/1", 100*time.Millisecond)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	// TODO: Test with reasonable timeout (will succeed)
	fmt.Println("Making request with 5s timeout...")
	err = makeRequestWithTimeout("https://httpbin.org/get", 5*time.Second)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// TODO: Demonstrate context in database operations (simulated)
func demonstrateDBWithContext() {
	fmt.Println("\n=== Database Operations with Context ===")
	
	// TODO: Simulate database operations
	type DB struct {
		name string
	}
	
	func (db *DB) Query(ctx context.Context, query string) ([]map[string]interface{}, error) {
		// Simulate query execution time
		queryTime := time.Duration(rand.Intn(2000)+500) * time.Millisecond
		
		select {
		case <-time.After(queryTime):
			// Query completed
			result := []map[string]interface{}{
				{"id": 1, "name": "Alice"},
				{"id": 2, "name": "Bob"},
			}
			fmt.Printf("Query completed in %v\n", queryTime)
			return result, nil
		case <-ctx.Done():
			// Context cancelled or timed out
			return nil, fmt.Errorf("query cancelled: %w", ctx.Err())
		}
	}
	
	func (db *DB) Transaction(ctx context.Context, fn func(context.Context) error) error {
		// Simulate transaction start
		fmt.Println("Starting transaction...")
		
		// Check context before proceeding
		if ctx.Err() != nil {
			return ctx.Err()
		}
		
		// Execute transaction function
		err := fn(ctx)
		if err != nil {
			fmt.Println("Rolling back transaction...")
			return err
		}
		
		fmt.Println("Committing transaction...")
		return nil
	}
	
	// TODO: Test database operations with context
	db := &DB{name: "testdb"}
	
	// Query with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	result, err := db.Query(ctx, "SELECT * FROM users")
	if err != nil {
		fmt.Printf("Query error: %v\n", err)
	} else {
		fmt.Printf("Query result: %v\n", result)
	}
	
	// Transaction with cancellation
	transactionCtx, transactionCancel := context.WithCancel(context.Background())
	defer transactionCancel()
	
	err = db.Transaction(transactionCtx, func(ctx context.Context) error {
		// Simulate work in transaction
		_, err := db.Query(ctx, "UPDATE users SET active = true")
		return err
	})
	
	if err != nil {
		fmt.Printf("Transaction error: %v\n", err)
	}
}

// TODO: Demonstrate context patterns and best practices
func demonstrateContextPatterns() {
	fmt.Println("\n=== Context Patterns and Best Practices ===")
	
	// TODO: Pattern 1: Context propagation through call chain
	type Service struct {
		name string
	}
	
	func (s *Service) ProcessData(ctx context.Context, data string) error {
		// Add service info to context
		serviceCtx := context.WithValue(ctx, "service", s.name)
		
		// Call validation
		if err := s.validateData(serviceCtx, data); err != nil {
			return err
		}
		
		// Call processing
		return s.processInternal(serviceCtx, data)
	}
	
	func (s *Service) validateData(ctx context.Context, data string) error {
		fmt.Printf("Validating data in service: %v\n", ctx.Value("service"))
		
		// Check context cancellation
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		
		// Simulate validation
		time.Sleep(100 * time.Millisecond)
		return nil
	}
	
	func (s *Service) processInternal(ctx context.Context, data string) error {
		fmt.Printf("Processing data in service: %v\n", ctx.Value("service"))
		
		// Check context throughout processing
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				return fmt.Errorf("processing cancelled at step %d: %w", i, ctx.Err())
			default:
			}
			
			time.Sleep(200 * time.Millisecond)
		}
		
		return nil
	}
	
	// TODO: Test service with context
	service := &Service{name: "DataProcessor"}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	err := service.ProcessData(ctx, "test-data")
	if err != nil {
		fmt.Printf("Service error: %v\n", err)
	} else {
		fmt.Println("Service processing completed successfully")
	}
	
	// TODO: Pattern 2: Context for request tracing
	demonstrateRequestTracing := func() {
		fmt.Println("\nRequest tracing pattern:")
		
		type traceKey string
		const traceIDKey traceKey = "traceID"
		
		// Middleware to add trace ID
		addTraceID := func(ctx context.Context) context.Context {
			traceID := fmt.Sprintf("trace-%d", time.Now().UnixNano())
			return context.WithValue(ctx, traceIDKey, traceID)
		}
		
		// Function that logs with trace ID
		logWithTrace := func(ctx context.Context, message string) {
			traceID := ctx.Value(traceIDKey)
			fmt.Printf("[%v] %s\n", traceID, message)
		}
		
		// Test tracing
		ctx := addTraceID(context.Background())
		logWithTrace(ctx, "Request started")
		logWithTrace(ctx, "Processing data")
		logWithTrace(ctx, "Request completed")
	}
	
	demonstrateRequestTracing()
}

// TODO: Demonstrate context gotchas and anti-patterns
func demonstrateContextGotchas() {
	fmt.Println("\n=== Context Gotchas and Anti-patterns ===")
	
	// TODO: Anti-pattern 1: Storing context in structs
	fmt.Println("❌ DON'T store context in structs:")
	fmt.Println("type BadService struct {")
	fmt.Println("    ctx context.Context  // Anti-pattern!")
	fmt.Println("}")
	
	// TODO: Anti-pattern 2: Passing nil context
	fmt.Println("\n❌ DON'T pass nil context:")
	fmt.Println("someFunction(nil)  // Use context.Background() instead")
	
	// TODO: Anti-pattern 3: Using context for optional parameters
	fmt.Println("\n❌ DON'T use context for optional parameters:")
	fmt.Println("ctx = context.WithValue(ctx, \"timeout\", 5*time.Second)  // Wrong!")
	
	// TODO: Best practices
	fmt.Println("\n✅ Context Best Practices:")
	fmt.Println("1. Always pass context as first parameter")
	fmt.Println("2. Don't store contexts in structs")
	fmt.Println("3. Don't pass nil context - use context.Background()")
	fmt.Println("4. Use context.WithValue sparingly")
	fmt.Println("5. Always call cancel() to release resources")
	fmt.Println("6. Check ctx.Done() in long-running operations")
	fmt.Println("7. Context values should be request-scoped data")
	
	// TODO: Demonstrate proper context checking
	properContextUsage := func(ctx context.Context) error {
		for i := 0; i < 1000; i++ {
			// Check context periodically
			if i%100 == 0 {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
				}
			}
			
			// Simulate work
			time.Sleep(time.Millisecond)
		}
		return nil
	}
	
	// Test proper usage
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	
	err := properContextUsage(ctx)
	if err != nil {
		fmt.Printf("Operation properly cancelled: %v\n", err)
	}
}

func main() {
	fmt.Println("🔄 Welcome to Context Usage! 🔄")
	fmt.Println("This file teaches you Go's context package for cancellation and timeouts")
	
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	// TODO: Implement each demonstration function
	// Start with basic context and progress to advanced patterns
	
	demonstrateBasicContext()
	// demonstrateContextTimeout()
	// demonstrateContextCancellation()
	// demonstrateContextDeadline()
	// demonstrateContextValues()
	// demonstrateHTTPWithContext()
	// demonstrateDBWithContext()
	// demonstrateContextPatterns()
	// demonstrateContextGotchas()
	
	fmt.Println("\n🎉 Congratulations! You've mastered context usage in Go!")
	fmt.Println("Next: Practice with packages_practice.go")
}

/*
🔍 Key Concepts to Remember:

1. **Context Types**: Background(), TODO(), WithTimeout(), WithCancel()
2. **Context Propagation**: Pass context as first parameter through call chain
3. **Cancellation**: Check ctx.Done() in long-running operations
4. **Values**: Use sparingly for request-scoped data
5. **Timeouts**: Set reasonable timeouts for operations
6. **Resource Cleanup**: Always call cancel() to release resources
7. **Error Handling**: Check context.Err() for cancellation reasons

📋 Essential Context Operations:
```go
// Create contexts
ctx := context.Background()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()

// Check cancellation
select {
case <-ctx.Done():
    return ctx.Err()
default:
    // Continue work
}

// With values
ctx = context.WithValue(ctx, key, value)
value := ctx.Value(key)

// HTTP with context
req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
```

🚨 Common Mistakes:
- Storing context in structs (anti-pattern)
- Passing nil context instead of context.Background()
- Not calling cancel() (resource leak)
- Using context.WithValue for configuration
- Not checking ctx.Done() in loops
- Ignoring context.Err() return values

🎯 Next Steps:
- Practice with packages_practice.go exercises
- Learn file I/O and system programming
- Build context-aware applications
- Master timeout and cancellation patterns
*/
