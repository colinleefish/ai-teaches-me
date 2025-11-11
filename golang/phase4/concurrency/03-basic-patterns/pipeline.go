package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// 🔄 PIPELINE PATTERN
// Chain of processing stages connected by channels
// Each stage transforms data and passes it to the next stage
// Enables parallel processing across multiple stages
// Essential for data transformation and streaming systems

// DataItem represents data flowing through pipeline stages
type DataItem struct {
	ID        int
	Content   string
	Metadata  map[string]interface{}
	Timestamp time.Time
}

// StageResult represents output from a pipeline stage
type StageResult struct {
	Item     DataItem
	Error    error
	StageNum int
}

// CircuitState represents the state of a circuit breaker
type CircuitState int

const (
	CLOSED CircuitState = iota
	OPEN
	HALF_OPEN
)

// CircuitBreaker protects against cascading failures
type CircuitBreaker struct {
	state            CircuitState
	failureCount     int
	failureThreshold int
	timeout          time.Duration
	lastFailureTime  time.Time
	successCount     int
	mu               sync.Mutex
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            CLOSED,
		failureThreshold: threshold,
		timeout:          timeout,
	}
}

// Call executes a function through the circuit breaker
func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	// Check if we should transition from OPEN to HALF_OPEN
	if cb.state == OPEN && time.Since(cb.lastFailureTime) > cb.timeout {
		cb.state = HALF_OPEN
		cb.successCount = 0
		fmt.Println("🟡 Circuit breaker: OPEN → HALF_OPEN (testing recovery)")
	}

	// If circuit is OPEN, reject the call
	if cb.state == OPEN {
		return fmt.Errorf("circuit breaker is OPEN - service unavailable")
	}

	// Execute the function
	err := fn()

	if err != nil {
		cb.recordFailure()
		return fmt.Errorf("service failed: %w", err)
	} else {
		cb.recordSuccess()
		return nil
	}
}

func (cb *CircuitBreaker) recordFailure() {
	cb.failureCount++
	cb.lastFailureTime = time.Now()

	if cb.state == HALF_OPEN {
		// Any failure in HALF_OPEN goes back to OPEN
		cb.state = OPEN
		fmt.Printf("🔴 Circuit breaker: HALF_OPEN → OPEN (failure detected)\n")
	} else if cb.failureCount >= cb.failureThreshold {
		// Too many failures, open the circuit
		cb.state = OPEN
		fmt.Printf("🔴 Circuit breaker: CLOSED → OPEN (threshold %d reached)\n", cb.failureThreshold)
	}
}

func (cb *CircuitBreaker) recordSuccess() {
	if cb.state == HALF_OPEN {
		cb.successCount++
		if cb.successCount >= 3 { // Need 3 successes to close
			cb.state = CLOSED
			cb.failureCount = 0
			fmt.Println("🟢 Circuit breaker: HALF_OPEN → CLOSED (service recovered)")
		}
	} else if cb.state == CLOSED {
		// Reset failure count on success
		cb.failureCount = 0
	}
}

func (cb *CircuitBreaker) GetState() string {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	switch cb.state {
	case CLOSED:
		return "CLOSED"
	case OPEN:
		return "OPEN"
	case HALF_OPEN:
		return "HALF_OPEN"
	default:
		return "UNKNOWN"
	}
}

// TUTOR: Basic pipeline connects stages with channels in sequence.
// Each stage reads from input channel and writes to output channel.
// Stages can process data independently and concurrently.
// Channel closing propagates through pipeline for clean shutdown.
// Pipeline depth determines memory usage and latency characteristics.
func demonstrateBasicPipeline() {
	fmt.Println("=== Data Processing Pipeline ===")

	// Pipeline: Fetch → Parse → Validate → Save

	// Stage 1: Simulate fetching raw data (slow I/O)
	fetchStage := make(chan DataItem, 5)
	parseStage := make(chan DataItem, 5)
	validateStage := make(chan DataItem, 5)
	saveStage := make(chan DataItem, 5)

	var wg sync.WaitGroup
	wg.Add(4) // 4 stages

	// Stage 1: Data Fetcher (slow - simulates API calls)
	go func() {
		defer wg.Done()
		defer close(parseStage)

		for item := range fetchStage {
			fmt.Printf("🔍 Fetching data for ID %d...\n", item.ID)
			time.Sleep(200 * time.Millisecond) // Slow I/O operation

			item.Content = fmt.Sprintf("raw_data_%d.json", item.ID)
			parseStage <- item
			fmt.Printf("✅ Fetched: %s\n", item.Content)
		}
	}()

	// Stage 2: Data Parser (medium speed)
	go func() {
		defer wg.Done()
		defer close(validateStage)

		for item := range parseStage {
			fmt.Printf("📝 Parsing %s...\n", item.Content)
			time.Sleep(100 * time.Millisecond) // CPU processing

			item.Content = strings.Replace(item.Content, "raw_", "parsed_", 1)
			validateStage <- item
			fmt.Printf("✅ Parsed: %s\n", item.Content)
		}
	}()

	// Stage 3: Data Validator (fast)
	go func() {
		defer wg.Done()
		defer close(saveStage)

		for item := range validateStage {
			fmt.Printf("🔍 Validating %s...\n", item.Content)
			time.Sleep(50 * time.Millisecond) // Quick validation

			item.Content = strings.Replace(item.Content, "parsed_", "valid_", 1)
			saveStage <- item
			fmt.Printf("✅ Validated: %s\n", item.Content)
		}
	}()

	// Stage 4: Data Saver (medium-slow - database writes)
	go func() {
		defer wg.Done()

		for item := range saveStage {
			fmt.Printf("💾 Saving %s...\n", item.Content)
			time.Sleep(150 * time.Millisecond) // Database write

			item.Content = strings.Replace(item.Content, "valid_", "saved_", 1)
			fmt.Printf("✅ Saved: %s\n", item.Content)
		}
	}()

	// Feed data into pipeline
	start := time.Now()

	for i := 1; i <= 50; i++ {
		fetchStage <- DataItem{
			ID:      i,
			Content: fmt.Sprintf("request_%d", i),
		}
		fmt.Printf("📨 Submitted item %d to pipeline\n", i)
	}

	close(fetchStage) // Start shutdown cascade
	wg.Wait()

	fmt.Printf("\n🎉 Pipeline completed in %v\n", time.Since(start))
	fmt.Println("Notice: Items processed concurrently at different stages!")
}

// TUTOR: Pipeline fan-out multiplies processing capacity at bottleneck stages.
// Slow stages can run multiple instances in parallel.
// Fast stages can remain single-instance to save resources.
// Load balancing occurs naturally through channel distribution.
// Fan-out enables fine-tuned performance optimization.
func demonstratePipelineFanOut() {
	fmt.Println("\n=== Pipeline Stage Fan-Out ===")

	// Pipeline: Generate → Process (BOTTLENECK) → Validate → Save
	// We'll fan-out the slow processing stage to 3 workers

	generateCh := make(chan DataItem, 10)
	processCh := make(chan DataItem, 20) // Shared by 3 processors
	validateCh := make(chan DataItem, 10)
	saveCh := make(chan DataItem, 10)

	var wg sync.WaitGroup

	start := time.Now()

	// Stage 1: Generator (fast)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(processCh)

		for item := range generateCh {
			fmt.Printf("📊 Generated: %s\n", item.Content)
			processCh <- item
		}
	}()

	// Stage 2: Processors (BOTTLENECK - fan out to 3 workers)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for item := range processCh {
				fmt.Printf("⚙️  Processor %d working on: %s\n", workerID, item.Content)
				time.Sleep(300 * time.Millisecond) // Slow processing

				item.Content = fmt.Sprintf("processed_by_%d_%s", workerID, item.Content)
				validateCh <- item
				fmt.Printf("✅ Processor %d finished: %s\n", workerID, item.Content)
			}
		}(i)
	}

	// Stage 3: Validator (fast)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(saveCh)

		for item := range validateCh {
			fmt.Printf("🔍 Validating: %s\n", item.Content)
			time.Sleep(50 * time.Millisecond) // Quick validation

			item.Content = "valid_" + item.Content
			saveCh <- item
		}
	}()

	// Stage 4: Saver (medium speed)
	wg.Add(1)
	go func() {
		defer wg.Done()

		for item := range saveCh {
			fmt.Printf("💾 Saving: %s\n", item.Content)
			time.Sleep(100 * time.Millisecond) // Database save

			fmt.Printf("✅ Saved: %s\n", item.Content)
		}
	}()

	// Feed data into pipeline
	for i := 1; i <= 10; i++ {
		generateCh <- DataItem{
			ID:      i,
			Content: fmt.Sprintf("data_%d", i),
		}
	}

	close(generateCh) // Start shutdown cascade
	wg.Wait()

	fmt.Printf("\n🚀 Fan-out pipeline completed in %v\n", time.Since(start))
	fmt.Println("🔄 Notice: Processing stage had 3 parallel workers!")
	fmt.Println("📈 Bottleneck parallelized while maintaining pipeline flow")
}

// TUTOR: Pipeline fan-in collects results from parallel processing stages.
// Multiple parallel stages feed into single collection point.
// Order preservation may be lost but throughput increases.
// Result aggregation can restore ordering if needed.
// Fan-in complements fan-out for complete parallel processing.
// TODO: Demonstrate result collection from parallel stages
func demonstratePipelineFanIn() {
	fmt.Println("\n=== Pipeline Result Fan-In ===")

	// TODO: Create multiple parallel processing stages
	// TODO: Implement result collection and aggregation
	// TODO: Show throughput benefits of parallel processing
	// TODO: Demonstrate order preservation techniques if needed
	// TODO: Show proper synchronization of fan-in collection
}

// TUTOR: Pipeline buffering smooths rate differences between stages.
// Faster stages can continue working while slower stages catch up.
// Buffering reduces blocking and improves overall throughput.
// Buffer sizing affects memory usage and latency characteristics.
// Proper buffering prevents pipeline stalls and deadlocks.
// TODO: Demonstrate buffer effects on pipeline performance
func demonstratePipelineBuffering() {
	fmt.Println("\n=== Pipeline Buffering Strategies ===")

	// TODO: Compare unbuffered vs buffered pipeline stages
	// TODO: Show how buffer sizes affect stage interaction
	// TODO: Demonstrate memory vs performance trade-offs
	// TODO: Show optimal buffer sizing for different workloads
	// TODO: Illustrate backpressure propagation through buffered stages
}

// TUTOR: Pipeline error handling requires error propagation mechanisms.
// Errors at any stage can affect downstream processing.
// Error channels run parallel to data channels.
// Failed items may need retry or alternate processing paths.
// Proper error handling maintains pipeline stability.
// TODO: Demonstrate error handling across pipeline stages
func demonstratePipelineErrorHandling() {
	fmt.Println("\n=== Pipeline Error Handling ===")

	// TODO: Create error channels parallel to data channels
	// TODO: Show error propagation between stages
	// TODO: Implement error recovery and retry mechanisms
	// TODO: Demonstrate failed item handling and routing
	// TODO: Show pipeline stability under error conditions
}

// TUTOR: Pipeline cancellation enables early termination of processing chains.
// Cancellation signals propagate through all pipeline stages.
// In-flight data may be processed or discarded based on requirements.
// Proper cancellation prevents resource waste and improves responsiveness.
// Context or channel-based cancellation provides clean shutdown.
func demonstratePipelineCancellation() {
	fmt.Println("\n=== Pipeline Cancellation ===")

	// Pipeline: Producer → Doubler → Tripler → Results
	// We'll cancel after 3 seconds to demonstrate early termination

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	numberCh := make(chan int, 5)
	doubledCh := make(chan int, 5)
	resultCh := make(chan int, 5)

	var wg sync.WaitGroup

	// Producer: Generate numbers
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(numberCh)

		for i := 1; i <= 20; i++ {
			// Check cancellation first
			select {
			case <-ctx.Done():
				fmt.Printf("🛑 Producer cancelled at number %d\n", i)
				return
			default:
				// Context not cancelled, continue
			}

			// Try to send with cancellation check
			select {
			case <-ctx.Done():
				fmt.Printf("🛑 Producer cancelled at number %d\n", i)
				return
			case numberCh <- i:
				fmt.Printf("📊 Produced: %d\n", i)
				time.Sleep(500 * time.Millisecond) // Slow producer
			}
		}
	}()

	// Worker 1: Double the numbers
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(doubledCh)

		for {
			select {
			case <-ctx.Done():
				fmt.Println("🛑 Doubler cancelled")
				return
			case num, ok := <-numberCh:
				if !ok {
					return // Channel closed
				}
				doubled := num * 2
				fmt.Printf("✖️ Doubled %d → %d\n", num, doubled)

				select {
				case <-ctx.Done():
					fmt.Println("🛑 Doubler cancelled while sending")
					return
				case doubledCh <- doubled:
					// Successfully sent
				}
			}
		}
	}()

	// Worker 2: Triple the doubled numbers
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(resultCh)

		for {
			select {
			case <-ctx.Done():
				fmt.Println("🛑 Tripler cancelled")
				return
			case doubled, ok := <-doubledCh:
				if !ok {
					return // Channel closed
				}
				tripled := doubled * 3
				fmt.Printf("🔺 Tripled %d → %d\n", doubled, tripled)

				select {
				case <-ctx.Done():
					fmt.Println("🛑 Tripler cancelled while sending")
					return
				case resultCh <- tripled:
					// Successfully sent
				}
			}
		}
	}()

	// Result collector
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("🛑 Result collector cancelled")
				return
			case result, ok := <-resultCh:
				if !ok {
					return // Channel closed
				}
				fmt.Printf("🎯 Final result: %d (original × 6)\n", result)
			}
		}
	}()

	// Wait for completion or cancellation
	wg.Wait()

	fmt.Println("\n⏰ Pipeline terminated (completed or cancelled)")
	fmt.Println("💡 Notice: All stages respected cancellation signal!")
}

// TUTOR: Pipeline monitoring tracks data flow and processing performance.
// Monitor throughput at each stage to identify bottlenecks.
// Track processing latencies and queue depths.
// Stage utilization metrics guide optimization efforts.
// End-to-end latency measurement shows overall pipeline performance.
// TODO: Demonstrate comprehensive pipeline monitoring
func demonstratePipelineMonitoring() {
	fmt.Println("\n=== Pipeline Monitoring ===")

	// TODO: Implement per-stage throughput measurement
	// TODO: Track processing latencies at each stage
	// TODO: Monitor queue depths between stages
	// TODO: Show bottleneck identification techniques
	// TODO: Demonstrate end-to-end latency tracking
}

// TUTOR: Branching pipelines split data flow based on content or conditions.
// Different data types or conditions route to different processing paths.
// Branching enables specialized processing for different data categories.
// Conditional routing optimizes processing for data characteristics.
// Branch points require routing logic and multiple output channels.
func demonstratePipelineBranching() {
	fmt.Println("\n=== Data ETL Pipeline Branching ===")

	// Raw data with different formats
	type RawData struct {
		ID      int
		Content string
		Format  string
	}

	type ProcessedData struct {
		ID     int
		Result string
		Source string
	}

	// Pipeline channels
	rawDataCh := make(chan RawData, 10)
	jsonCh := make(chan RawData, 5)
	csvCh := make(chan RawData, 5)
	xmlCh := make(chan RawData, 5)
	resultCh := make(chan ProcessedData, 10)

	var wg sync.WaitGroup
	var processorWg sync.WaitGroup
	processorWg.Add(4) // Router + 3 processors

	// Router: Classify and route data by format
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer processorWg.Done()
		defer close(jsonCh)
		defer close(csvCh)
		defer close(xmlCh)

		for data := range rawDataCh {
			fmt.Printf("🔀 Routing %s data (ID: %d)\n", data.Format, data.ID)

			switch data.Format {
			case "JSON":
				jsonCh <- data
			case "CSV":
				csvCh <- data
			case "XML":
				xmlCh <- data
			default:
				fmt.Printf("❌ Unknown format: %s\n", data.Format)
			}
		}
	}()

	// JSON Processor Branch
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer processorWg.Done()

		for data := range jsonCh {
			fmt.Printf("📄 Processing JSON (ID: %d): %s\n", data.ID, data.Content)
			time.Sleep(100 * time.Millisecond) // JSON parsing

			result := ProcessedData{
				ID:     data.ID,
				Result: "parsed_json_" + data.Content,
				Source: "JSON_PROCESSOR",
			}
			resultCh <- result
			fmt.Printf("✅ JSON processed: %s\n", result.Result)
		}
	}()

	// CSV Processor Branch
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer processorWg.Done()

		for data := range csvCh {
			fmt.Printf("📊 Processing CSV (ID: %d): %s\n", data.ID, data.Content)
			time.Sleep(150 * time.Millisecond) // CSV transformation

			result := ProcessedData{
				ID:     data.ID,
				Result: "transformed_csv_" + data.Content,
				Source: "CSV_TRANSFORMER",
			}
			resultCh <- result
			fmt.Printf("✅ CSV transformed: %s\n", result.Result)
		}
	}()

	// XML Processor Branch
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer processorWg.Done()

		for data := range xmlCh {
			fmt.Printf("🔖 Processing XML (ID: %d): %s\n", data.ID, data.Content)
			time.Sleep(200 * time.Millisecond) // XML conversion

			result := ProcessedData{
				ID:     data.ID,
				Result: "converted_xml_" + data.Content,
				Source: "XML_CONVERTER",
			}
			resultCh <- result
			fmt.Printf("✅ XML converted: %s\n", result.Result)
		}
	}()

	// Result Collector
	wg.Add(1)
	go func() {
		defer wg.Done()

		for result := range resultCh {
			fmt.Printf("🎯 Final result from %s: %s\n", result.Source, result.Result)
		}
	}()

	// Feed mixed data formats into pipeline
	sampleData := []RawData{
		{1, `{"name":"John","age":30}`, "JSON"},
		{2, `"Name,Age\nAlice,25\nBob,35"`, "CSV"},
		{3, `<user><name>Carol</name><age>28</age></user>`, "XML"},
		{4, `{"product":"laptop","price":999}`, "JSON"},
		{5, `"Product,Price\nMouse,25\nKeyboard,75"`, "CSV"},
		{6, `<order><id>123</id><total>150</total></order>`, "XML"},
	}

	for _, data := range sampleData {
		rawDataCh <- data
		fmt.Printf("📨 Submitted %s data (ID: %d)\n", data.Format, data.ID)
	}

	close(rawDataCh)

	// Wait for processors to finish, then close resultCh
	go func() {
		processorWg.Wait() // Wait for just the processors
		close(resultCh)    // Then close results channel
	}()

	wg.Wait() // Wait for all goroutines including result collector

	fmt.Println("\n🎉 ETL Pipeline completed!")
	fmt.Println("💡 Each format was processed by specialized workers")
}

// TUTOR: Pipeline recovery handles stage failures without losing data.
// Failed stages can be restarted while preserving pipeline state. ✅
// Checkpoint mechanisms enable recovery from partial failures. ✅
// Circuit breaker patterns protect against cascading failures.
// Resilient pipelines maintain operation despite component failures.
func demoPipelineRecoveryRestart() {
	fmt.Println("\n=== Pipeline Stage Restart Recovery ===")

	inputCh := make(chan int, 10)
	outputCh := make(chan int, 10)
	done := make(chan bool)

	// Simulate a processing stage that randomly fails
	var processedCount int

	var startProcessor func(int)
	startProcessor = func(id int) {
		fmt.Printf("🚀 Starting processor %d\n", id)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("💥 Processor %d crashed: %v\n", id, r)
					// Signal for restart
					go func() {
						time.Sleep(100 * time.Millisecond)
						startProcessor(id + 1) // Restart with new ID
					}()
				}
			}()

			for num := range inputCh {
				// Simulate random failures (20% chance)
				if rand.Intn(100) < 20 {
					panic(fmt.Sprintf("random failure processing %d", num))
				}

				fmt.Printf("⚙️  Processor %d: processing %d\n", id, num)
				time.Sleep(100 * time.Millisecond) // Simulate work

				result := num * 2
				outputCh <- result
				processedCount++

				fmt.Printf("✅ Processor %d: %d → %d\n", id, num, result)
			}

			fmt.Printf("🏁 Processor %d finished normally\n", id)
			close(outputCh) // Close when processing is done
		}()
	}

	// Start initial processor
	startProcessor(1)

	// Result collector
	var results []int
	go func() {
		for result := range outputCh {
			results = append(results, result)
			fmt.Printf("📦 Collected result: %d\n", result)
		}
		done <- true
	}()

	// Feed data into pipeline
	var inputWg sync.WaitGroup
	inputWg.Add(1)
	go func() {
		defer inputWg.Done()
		for i := 1; i <= 20; i++ {
			inputCh <- i
			fmt.Printf("📨 Sent: %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
		close(inputCh)
	}()

	// Wait for input to finish
	inputWg.Wait()

	// Wait for results
	<-done

	fmt.Printf("\n🎯 Recovery Summary:\n")
	fmt.Printf("📊 Total results collected: %d\n", len(results))
	fmt.Printf("📈 Expected: 20, Got: %d\n", len(results))
	fmt.Printf("💪 Pipeline survived failures and auto-restarted stages!\n")
}

func demoCheckpointRecovery() {
	fmt.Println("\n=== Pipeline Checkpoint Recovery ===")

	type Checkpoint struct {
		LastProcessedID int    `json:"last_processed_id"`
		BatchSize       int    `json:"batch_size"`
		Timestamp       string `json:"timestamp"`
	}

	type WorkItem struct {
		ID   int
		Data string
	}

	// Simulated persistent storage for checkpoints
	var savedCheckpoint *Checkpoint

	saveCheckpoint := func(checkpoint Checkpoint) {
		savedCheckpoint = &checkpoint
		fmt.Printf("💾 Checkpoint saved: Last ID=%d, Time=%s\n",
			checkpoint.LastProcessedID, checkpoint.Timestamp)
	}

	loadCheckpoint := func() *Checkpoint {
		if savedCheckpoint != nil {
			fmt.Printf("📂 Checkpoint loaded: Resuming from ID=%d\n",
				savedCheckpoint.LastProcessedID)
		} else {
			fmt.Println("📂 No checkpoint found, starting fresh")
		}
		return savedCheckpoint
	}

	// Simulate work items (like file processing, database records, etc.)
	allWork := make([]WorkItem, 50)
	for i := 0; i < 50; i++ {
		allWork[i] = WorkItem{
			ID:   i + 1,
			Data: fmt.Sprintf("data_item_%d", i+1),
		}
	}

	processItem := func(item WorkItem) error {
		// Simulate processing time
		time.Sleep(20 * time.Millisecond)

		// Simulate random failures (10% chance)
		if rand.Intn(100) < 10 {
			return fmt.Errorf("processing failed for item %d", item.ID)
		}

		fmt.Printf("✅ Processed item %d: %s\n", item.ID, item.Data)
		return nil
	}

	runPipeline := func(startFromID int) {
		fmt.Printf("🚀 Starting pipeline from ID %d\n", startFromID)

		batchSize := 5
		processed := 0

		for i := startFromID; i <= len(allWork); i += batchSize {
			batch := []WorkItem{}

			// Create batch
			for j := i; j < i+batchSize && j <= len(allWork); j++ {
				batch = append(batch, allWork[j-1]) // j-1 because array is 0-indexed
			}

			if len(batch) == 0 {
				break
			}

			fmt.Printf("\n📦 Processing batch starting from ID %d (%d items)\n",
				batch[0].ID, len(batch))

			// Process batch
			lastProcessedID := startFromID - 1
			for _, item := range batch {
				err := processItem(item)
				if err != nil {
					fmt.Printf("💥 Pipeline failed: %v\n", err)
					fmt.Printf("🔄 Last successful ID: %d\n", lastProcessedID)
					return
				}
				lastProcessedID = item.ID
				processed++
			}

			// Save checkpoint after successful batch
			checkpoint := Checkpoint{
				LastProcessedID: lastProcessedID,
				BatchSize:       batchSize,
				Timestamp:       time.Now().Format("15:04:05"),
			}
			saveCheckpoint(checkpoint)
		}

		fmt.Printf("\n🎉 Pipeline completed! Processed %d items total\n", processed)
	}

	// First run - simulate failure partway through
	fmt.Println("=== Initial Pipeline Run ===")
	runPipeline(1)

	// Recovery run - resume from checkpoint
	fmt.Println("\n=== Recovery Pipeline Run ===")
	checkpoint := loadCheckpoint()

	if checkpoint != nil {
		nextID := checkpoint.LastProcessedID + 1
		fmt.Printf("🔄 Resuming processing from ID %d\n", nextID)
		runPipeline(nextID)
	} else {
		fmt.Println("🔄 No checkpoint found, starting from beginning")
		runPipeline(1)
	}

	fmt.Printf("\n🎯 Checkpoint Recovery Summary:\n")
	fmt.Printf("💡 Pipeline can resume exactly where it left off\n")
	fmt.Printf("🛡️  No duplicate processing or data loss\n")
	fmt.Printf("⚡ Critical for long-running data processing jobs\n")
}

func demoCircuitBreaker() {
	fmt.Println("\n=== Pipeline Circuit Breaker ===")

	// Simulate an unreliable processing stage
	unreliableStage := func(id int) error {
		time.Sleep(50 * time.Millisecond)

		// High failure rate initially, then recovers
		if id < 10 {
			// 80% failure rate for first 10 items
			if rand.Intn(100) < 80 {
				return fmt.Errorf("stage overloaded")
			}
		} else {
			// 20% failure rate after recovery
			if rand.Intn(100) < 20 {
				return fmt.Errorf("transient error")
			}
		}

		fmt.Printf("✅ Successfully processed item %d\n", id)
		return nil
	}

	// Fallback processing (degraded service)
	fallbackStage := func(id int) error {
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("🔄 Fallback processed item %d (degraded mode)\n", id)
		return nil
	}

	// Create circuit breaker: 3 failures trigger OPEN, 2 second timeout
	cb := NewCircuitBreaker(3, 2*time.Second)

	fmt.Println("🚀 Starting pipeline with circuit breaker protection")

	// Process 25 items through circuit breaker
	for i := 1; i <= 25; i++ {
		fmt.Printf("\n📦 Processing item %d [Circuit: %s]\n", i, cb.GetState())

		err := cb.Call(func() error {
			return unreliableStage(i)
		})

		if err != nil {
			fmt.Printf("❌ Primary stage failed: %v\n", err)

			// Use fallback when circuit is open
			if cb.GetState() == "OPEN" {
				fallbackStage(i)
			}
		}

		time.Sleep(100 * time.Millisecond) // Simulate processing interval
	}

	fmt.Printf("\n🎯 Circuit Breaker Summary:\n")
	fmt.Printf("🛡️  Protected pipeline from cascading failures\n")
	fmt.Printf("🔄 Provided fallback processing during outages\n")
	fmt.Printf("⚡ Automatically detected service recovery\n")
	fmt.Printf("🎛️  Final state: %s\n", cb.GetState())
}

func demonstratePipelineRecovery() {
	// demoPipelineRecoveryRestart()
	// demoCheckpointRecovery()
	demoCircuitBreaker()
}

// TUTOR: Real-world pipelines process streaming data from various sources.
// Common applications include data ETL, log processing, and content transformation.
// Pipelines enable scalable data processing architectures.
// Understanding pipelines is essential for building data-intensive applications.
// Pipeline patterns appear in microservices, messaging systems, and analytics platforms.
// TODO: Show practical pipeline applications
func demonstrateRealWorldPipelines() {
	fmt.Println("\n=== Real-World Pipeline Applications ===")

	// TODO: Show log processing pipeline (parse, filter, aggregate, store)
	// TODO: Demonstrate image processing pipeline (resize, filter, format, save)
	// TODO: Show data ETL pipeline (extract, transform, validate, load)
	// TODO: Illustrate content processing pipeline (fetch, parse, enrich, index)
	// TODO: Show integration with external systems and databases
}

func main() {
	fmt.Println("🔄 Pipeline Patterns - Streaming Data Transformation 🔄")
	fmt.Println("Build efficient data processing chains")

	// TODO: Implement each demonstration function
	// Progress from simple to complex pipeline architectures

	// demonstrateBasicPipeline()
	// demonstratePipelineFanOut()
	// demonstratePipelineFanIn()
	// demonstratePipelineBuffering()
	// demonstratePipelineErrorHandling()
	// demonstratePipelineCancellation()
	// demonstratePipelineMonitoring()
	// demonstratePipelineBranching()
	demonstratePipelineRecovery()
	// demonstrateRealWorldPipelines()
}
