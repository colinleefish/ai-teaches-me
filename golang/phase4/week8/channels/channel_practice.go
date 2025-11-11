// Week 8: Channel Practice Exercises
// Complete these exercises to master channel programming

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// TODO: Exercise 1 - Basic Channel Communication
func exercise1_BasicChannelCommunication() {
	fmt.Println("=== Exercise 1: Basic Channel Communication ===")

	// TODO: Create a program that:
	// 1. Creates an unbuffered channel for string messages
	// 2. Starts a goroutine that sends 5 messages with delays
	// 3. Receives and prints all messages in main goroutine
	// 4. Uses proper channel closing
	// 5. Handles the case when channel is closed

	// Expected behavior:
	// - Messages should be received in order
	// - Program should not deadlock
	// - Should detect when channel is closed

	fmt.Println("Exercise 1 completed!")
}

// TODO: Exercise 2 - Buffered Channel Performance
func exercise2_BufferedChannelPerformance() {
	fmt.Println("\n=== Exercise 2: Buffered Channel Performance ===")

	// TODO: Compare performance of different buffer sizes:
	// 1. Test with unbuffered channel (0)
	// 2. Test with small buffer (10)
	// 3. Test with large buffer (1000)
	//
	// For each test:
	// - Send 10,000 integers through channel
	// - Use separate sender and receiver goroutines
	// - Measure total time taken
	// - Report throughput (messages/second)

	testBufferSize := func(bufferSize int) {
		// TODO: Implement performance test
		// TODO: Measure and report results
	}

	// TODO: Test different buffer sizes and compare results

	fmt.Println("Exercise 2 completed!")
}

// TODO: Exercise 3 - Select Statement Mastery
func exercise3_SelectStatementMastery() {
	fmt.Println("\n=== Exercise 3: Select Statement Mastery ===")

	// TODO: Implement a message router that:
	// 1. Receives messages from 3 input channels
	// 2. Routes messages based on message type to appropriate output channels
	// 3. Handles timeouts (1 second) when no messages arrive
	// 4. Implements graceful shutdown when signaled
	// 5. Uses default case for non-blocking operations when needed

	type Message struct {
		Type    string // "urgent", "normal", "low"
		Content string
		ID      int
	}

	// TODO: Create input channels and output channels
	// TODO: Implement message router with select
	// TODO: Test with different message types and timing
	// TODO: Test timeout behavior
	// TODO: Test graceful shutdown

	fmt.Println("Exercise 3 completed!")
}

// TODO: Exercise 4 - Pipeline Implementation
func exercise4_PipelineImplementation() {
	fmt.Println("\n=== Exercise 4: Pipeline Implementation ===")

	// TODO: Build a data processing pipeline:
	// Stage 1: Number generator (1-1000)
	// Stage 2: Prime number filter
	// Stage 3: Square the primes
	// Stage 4: Format as strings ("prime: square")
	// Stage 5: Collect results

	// Requirements:
	// - Each stage runs in separate goroutine
	// - Use channels for communication between stages
	// - Support context cancellation
	// - Handle errors gracefully
	// - Measure processing time for each stage

	// TODO: Implement each pipeline stage
	// TODO: Connect stages with channels
	// TODO: Add context support for cancellation
	// TODO: Add timing measurements
	// TODO: Test with early cancellation

	fmt.Println("Exercise 4 completed!")
}

// TODO: Exercise 5 - Fan-In/Fan-Out System
func exercise5_FanInFanOutSystem() {
	fmt.Println("\n=== Exercise 5: Fan-In/Fan-Out System ===")

	// TODO: Implement a job processing system:
	// 1. Job generator creates 100 jobs
	// 2. Fan-out distributes jobs to 5 workers
	// 3. Workers process jobs (simulate with random delay)
	// 4. Fan-in collects results from all workers
	// 5. Results aggregator computes statistics

	type Job struct {
		ID       int
		WorkType string
		Data     []int
	}

	type JobResult struct {
		JobID       int
		ProcessTime time.Duration
		Result      int
		WorkerID    int
	}

	// TODO: Implement job generator
	// TODO: Implement fan-out distributor
	// TODO: Implement worker goroutines
	// TODO: Implement fan-in collector
	// TODO: Implement results aggregator
	// TODO: Measure total system throughput

	fmt.Println("Exercise 5 completed!")
}

// TODO: Exercise 6 - Rate-Limited API Client
func exercise6_RateLimitedAPIClient() {
	fmt.Println("\n=== Exercise 6: Rate-Limited API Client ===")

	// TODO: Implement a rate-limited API client:
	// 1. Rate limit: 10 requests per second
	// 2. Burst capacity: 5 requests
	// 3. Request queue with timeout
	// 4. Retry logic for failed requests
	// 5. Circuit breaker for repeated failures

	type APIRequest struct {
		URL     string
		Method  string
		Payload interface{}
		Timeout time.Duration
	}

	type APIResponse struct {
		StatusCode int
		Body       string
		Error      error
		Duration   time.Duration
	}

	type APIClient struct {
		// TODO: Add fields for rate limiting and circuit breaking
	}

	// TODO: Implement rate limiter using channels
	// TODO: Implement request queue with timeout
	// TODO: Implement circuit breaker logic
	// TODO: Implement retry mechanism
	// TODO: Test with various request patterns

	fmt.Println("Exercise 6 completed!")
}

// TODO: Exercise 7 - Pub-Sub Message System
func exercise7_PubSubMessageSystem() {
	fmt.Println("\n=== Exercise 7: Pub-Sub Message System ===")

	// TODO: Build a publish-subscribe message system:
	// 1. Support multiple topics
	// 2. Multiple publishers can publish to same topic
	// 3. Multiple subscribers can subscribe to same topic
	// 4. Message persistence (keep last N messages per topic)
	// 5. Subscriber filtering based on message content
	// 6. Graceful handling of slow subscribers

	type Publisher struct {
		ID string
	}

	type Subscriber struct {
		ID     string
		Filter func(Message) bool
	}

	type Message struct {
		Topic     string
		Content   string
		Publisher string
		Timestamp time.Time
	}

	type MessageBroker struct {
		// TODO: Add fields for managing topics, subscribers, and messages
	}

	// TODO: Implement message broker
	// TODO: Implement publisher registration
	// TODO: Implement subscriber registration with filtering
	// TODO: Implement message routing and delivery
	// TODO: Handle slow subscribers (drop or buffer)
	// TODO: Test with multiple publishers and subscribers

	fmt.Println("Exercise 7 completed!")
}

// TODO: Exercise 8 - Distributed Task Scheduler
func exercise8_DistributedTaskScheduler() {
	fmt.Println("\n=== Exercise 8: Distributed Task Scheduler ===")

	// TODO: Implement a distributed task scheduler:
	// 1. Task submission with priority levels
	// 2. Multiple worker nodes with different capabilities
	// 3. Task assignment based on worker availability and capability
	// 4. Task retry on failure
	// 5. Worker health monitoring
	// 6. Load balancing across workers

	type Task struct {
		ID           string
		Priority     int // 1-5, 5 being highest
		RequiredCaps []string
		Payload      interface{}
		MaxRetries   int
		Timeout      time.Duration
	}

	type Worker struct {
		ID            string
		Capabilities  []string
		MaxConcurrent int
		Status        string // "healthy", "busy", "failed"
	}

	type TaskScheduler struct {
		// TODO: Add fields for task queue, workers, assignments
	}

	// TODO: Implement task scheduler
	// TODO: Implement worker registration and health monitoring
	// TODO: Implement task assignment algorithm
	// TODO: Implement retry logic
	// TODO: Implement load balancing
	// TODO: Test with various scenarios

	fmt.Println("Exercise 8 completed!")
}

// TODO: Exercise 9 - Real-time Data Aggregator
func exercise9_RealTimeDataAggregator() {
	fmt.Println("\n=== Exercise 9: Real-time Data Aggregator ===")

	// TODO: Build a real-time data aggregation system:
	// 1. Multiple data sources sending metrics
	// 2. Different aggregation windows (1min, 5min, 1hour)
	// 3. Multiple aggregation types (sum, avg, max, min, count)
	// 4. Real-time alerting when thresholds exceeded
	// 5. Data persistence for historical analysis
	// 6. Query interface for historical data

	type Metric struct {
		Name      string
		Value     float64
		Tags      map[string]string
		Timestamp time.Time
	}

	type AggregationRule struct {
		MetricName string
		Window     time.Duration
		Function   string // "sum", "avg", "max", "min", "count"
		GroupBy    []string
	}

	type Alert struct {
		MetricName string
		Threshold  float64
		Condition  string // "gt", "lt", "eq"
		Window     time.Duration
	}

	type DataAggregator struct {
		// TODO: Add fields for managing metrics, rules, and alerts
	}

	// TODO: Implement data ingestion
	// TODO: Implement windowed aggregation
	// TODO: Implement alerting system
	// TODO: Implement data persistence
	// TODO: Implement query interface
	// TODO: Test with high-volume data streams

	fmt.Println("Exercise 9 completed!")
}

// TODO: Exercise 10 - Chaos Engineering Framework
func exercise10_ChaosEngineeringFramework() {
	fmt.Println("\n=== Exercise 10: Chaos Engineering Framework ===")

	// TODO: Build a chaos engineering framework:
	// 1. System under test with multiple components
	// 2. Chaos experiments (network delays, failures, resource exhaustion)
	// 3. Health monitoring and metrics collection
	// 4. Experiment scheduling and coordination
	// 5. Recovery verification after chaos
	// 6. Report generation

	type Component struct {
		Name   string
		Status string
		Health func() bool
	}

	type ChaosExperiment struct {
		Name       string
		Target     string
		Type       string // "delay", "failure", "resource"
		Duration   time.Duration
		Intensity  float64
		Hypothesis string
	}

	type ChaosFramework struct {
		// TODO: Add fields for managing components and experiments
	}

	// TODO: Implement system monitoring
	// TODO: Implement chaos injection
	// TODO: Implement experiment scheduling
	// TODO: Implement health verification
	// TODO: Implement recovery testing
	// TODO: Generate experiment reports

	fmt.Println("Exercise 10 completed!")
}

// Helper functions for exercises
func generateRandomJobs(count int) []Job {
	jobs := make([]Job, count)
	workTypes := []string{"compute", "io", "network", "memory"}

	for i := 0; i < count; i++ {
		jobs[i] = Job{
			ID:       i,
			WorkType: workTypes[rand.Intn(len(workTypes))],
			Data:     generateRandomData(rand.Intn(100) + 1),
		}
	}
	return jobs
}

func generateRandomData(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	return data
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func simulateAPICall(req APIRequest) APIResponse {
	// Simulate API call with random delay and possible failure
	time.Sleep(time.Duration(rand.Intn(200)+50) * time.Millisecond)

	// 10% chance of failure
	if rand.Float32() < 0.1 {
		return APIResponse{
			StatusCode: 500,
			Error:      fmt.Errorf("API call failed"),
		}
	}

	return APIResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Response for %s", req.URL),
	}
}

func measureExecutionTime(name string, fn func()) time.Duration {
	start := time.Now()
	fn()
	duration := time.Since(start)
	fmt.Printf("%s execution time: %v\n", name, duration)
	return duration
}

func createTestContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

func main() {
	fmt.Println("🏃‍♂️ Welcome to Channel Practice! 🏃‍♂️")
	fmt.Println("Complete these exercises to master channel programming")

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// TODO: Implement each exercise one by one
	// Start with basic exercises and progress to advanced systems
	// Uncomment each exercise as you complete the previous one

	exercise1_BasicChannelCommunication()
	// exercise2_BufferedChannelPerformance()
	// exercise3_SelectStatementMastery()
	// exercise4_PipelineImplementation()
	// exercise5_FanInFanOutSystem()
	// exercise6_RateLimitedAPIClient()
	// exercise7_PubSubMessageSystem()
	// exercise8_DistributedTaskScheduler()
	// exercise9_RealTimeDataAggregator()
	// exercise10_ChaosEngineeringFramework()

	fmt.Println("\n🎉 Congratulations! You've mastered channel programming!")
	fmt.Println("🚀 Ready for Phase 5: Standard Library & Packages!")
}

/*
🎯 Exercise Guidelines:

1. **Start Simple**: Begin with basic channel operations
2. **Add Complexity**: Gradually build more sophisticated systems
3. **Test Thoroughly**: Use race detector and stress testing
4. **Handle Errors**: Don't ignore error conditions
5. **Use Context**: Support cancellation in long-running operations
6. **Monitor Performance**: Measure throughput and latency
7. **Document Patterns**: Understand when to use each approach

📝 Completion Checklist:
□ Exercise 1: Basic channel communication and closing
□ Exercise 2: Buffer size performance comparison
□ Exercise 3: Advanced select statement usage
□ Exercise 4: Multi-stage pipeline with cancellation
□ Exercise 5: Fan-in/fan-out job processing system
□ Exercise 6: Rate-limited client with circuit breaker
□ Exercise 7: Pub-sub system with filtering
□ Exercise 8: Distributed task scheduler
□ Exercise 9: Real-time data aggregation
□ Exercise 10: Chaos engineering framework

🔧 Testing Commands:
```bash
# Run with race detector
go run -race channel_practice.go

# Run with CPU profiling
go run channel_practice.go -cpuprofile=cpu.prof

# Run with memory profiling
go run channel_practice.go -memprofile=mem.prof

# Test specific exercise
go run channel_practice.go -exercise=1
```

🚨 Common Mistakes to Avoid:
- Deadlocks from unbuffered channels
- Memory leaks from unclosed channels
- Not handling context cancellation
- Poor error propagation
- Missing graceful shutdown
- Ignoring backpressure in high-throughput systems

🎯 Success Criteria:
- All exercises pass with race detector
- Clean channel lifecycle management
- Proper error handling and resource cleanup
- Understanding of when to use each pattern
- Ability to debug complex channel interactions
- Performance optimization awareness

💡 Real-World Applications:
- Microservice communication
- Event-driven architectures
- Stream processing systems
- Distributed task processing
- Real-time monitoring systems
- High-throughput data pipelines
*/

// Additional types for exercises
type Job struct {
	ID       int
	WorkType string
	Data     []int
}

type APIRequest struct {
	URL     string
	Method  string
	Payload interface{}
	Timeout time.Duration
}

type APIResponse struct {
	StatusCode int
	Body       string
	Error      error
	Duration   time.Duration
}
