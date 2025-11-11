// Week 8: Select Statements
// This file demonstrates select statements for channel multiplexing and coordination

package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// TUTOR: Basic select statement is Go's multiplexer for channels - it waits for multiple channel operations simultaneously.
// Select picks the first channel operation that becomes ready, enabling true non-blocking concurrent programming.
// Like Unix select() for file descriptors, but for Go channels. Think of it as "wait for any of these events".
// Use select for: handling multiple channels, implementing timeouts, non-blocking operations, event-driven programming.
// Key behavior: If multiple cases are ready simultaneously, Go chooses one randomly (fairness).
// TODO: Demonstrate basic select statement
func demonstrateBasicSelect() {
	fmt.Println("=== Basic Select Statement ===")

	// TODO: Create multiple channels
	// TODO: Use select to handle whichever channel is ready first
	// Show non-deterministic behavior when multiple channels are ready

	// TODO: Create two channels that send at different times
	// TODO: Use select to receive from whichever is ready
	// TODO: Show that select picks the first ready operation

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		defer close(ch1)
		for {
			random := rand.Intn(10)
			fmt.Println("ch1 will send in", random, "s")
			time.Sleep(time.Duration(random) * time.Second)
			ch1 <- "Hello from ch1 that waited " + strconv.Itoa(random) + "s"
		}
	}()

	go func() {
		defer close(ch2)
		for {
			random := rand.Intn(10)
			fmt.Println("ch2 will send in", random, "s")
			time.Sleep(time.Duration(random) * time.Second)
			ch2 <- "Hello from ch2 that waited " + strconv.Itoa(random) + "s"
		}
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}

	fmt.Println("Basic select demonstration completed!")
}

// TUTOR: Select with default case transforms select from blocking to non-blocking operation.
// Default case executes immediately if no other cases are ready, preventing goroutine blocking.
// This enables polling patterns, non-blocking channel operations, and responsive event loops.
// Use default for: checking channel status without waiting, polling patterns, preventing deadlocks, responsive UIs.
// Critical difference: without default = blocking (waits), with default = non-blocking (never waits).
// TODO: Demonstrate select with default case
func demonstrateSelectDefault() {
	fmt.Println("\n=== Select with Default Case ===")

	// TODO: Show non-blocking channel operations using default
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		defer close(ch1)
		for {
			random := rand.Intn(10)
			fmt.Println("ch1 will send in", random, "s")
			time.Sleep(time.Duration(random) * time.Second)
			ch1 <- "Hello from ch1 that waited " + strconv.Itoa(random) + "s"
		}
	}()

	go func() {
		defer close(ch2)
		for {
			random := rand.Intn(10)
			fmt.Println("ch2 will send in", random, "s")
			time.Sleep(time.Duration(random) * time.Second)
			ch2 <- "Hello from ch2 that waited " + strconv.Itoa(random) + "s"
		}
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("No channel ready")
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// TODO: Try to send to full channel with default fallback
	ch := make(chan string, 1) // Capacity = 1
	ch <- "first"              // Channel now full

	// This is essentially what Go is doing:
	select {
	case ch <- "second": // if (!ch.full()) { insert it }
		fmt.Println("Inserted")
	default: // else { default behavior }
		fmt.Println("Channel full, can't insert")
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	// TODO: Try to receive from empty channel with default fallback

	ch = make(chan string, 1)
	select {
	case msg := <-ch:
		fmt.Println("Received", msg)
	default:
		fmt.Println("No channel ready, default case executed")
	}

	// TODO: Demonstrate polling pattern with default
	// Check channel status without blocking

	mailboxChannel := make(chan string, 1)

	go func() {
		sendingIn := rand.Intn(10)
		fmt.Println("Sending in", sendingIn, "s")
		time.Sleep(time.Duration(sendingIn) * time.Second)
		mailboxChannel <- "Hello from mailbox that waited " + strconv.Itoa(sendingIn) + "s"
	}()
polling:
	for {
		select {
		case msg := <-mailboxChannel:
			fmt.Println("Received", msg)
			break polling
		default:
			fmt.Println("No message received, default case executed")
			time.Sleep(500 * time.Millisecond)
		}
	}

	// TODO: Show difference between blocking and non-blocking select

	fmt.Println("Select default demonstration completed!")
}

// TUTOR: Timeout patterns with select prevent operations from hanging indefinitely - essential for robust systems.
// Combines operation channels with timer channels (time.After, context, ticker) in a race condition.
// The first to complete wins: either the operation succeeds or timeout triggers graceful failure.
// Use timeouts for: HTTP requests, database queries, user input, any operation that might hang.
// Patterns: simple timeout (time.After), context timeout (professional), periodic (ticker), deadline-based (business logic).
// TODO: Demonstrate timeout patterns with select
func demonstrateTimeoutPatterns() {
	fmt.Println("\n=== Timeout Patterns ===")

	// TODO: Pattern 1: Simple timeout with time.After
	simpleTimeout := func() {
		// TODO: Create operation that might take too long
		// TODO: Use select with time.After for timeout
		goingToTake := rand.Intn(10)
		fmt.Println("Going to take", goingToTake, "s")
		respChannel := timedChannel("Hello from respChannel that waited "+strconv.Itoa(goingToTake)+"s", time.Duration(goingToTake)*time.Second)

		select {
		case msg := <-respChannel:
			fmt.Println(msg)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout")
		}
	}

	// TODO: Pattern 2: Timeout with context
	contextTimeout := func() {
		// Create context that times out after 3 seconds
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel() // Important: prevent resource leak

		// Simulate slow operation
		goingToTake := rand.Intn(10)
		fmt.Println("Going to take", goingToTake, "s")
		resultCh := timedChannel("Operation completed", time.Duration(goingToTake)*time.Second)

		// Race: operation vs timeout
		select {
		case result := <-resultCh:
			fmt.Println("✅ Success:", result)

		case <-ctx.Done(): // Context's timeout channel
			fmt.Println("⏰ Context timeout:", ctx.Err())
			// ctx.Err() returns: "context deadline exceeded"
		}
	}

	// TODO: Pattern 3: Periodic timeout (ticker)
	periodicTimeout := func() {
		// TODO: Use time.Ticker for repeated timeouts
		// TODO: Show how to handle periodic events

		ticker := time.NewTicker(1 * time.Second) // this ticks every second
		timeout := time.After(5 * time.Second)
		defer ticker.Stop()

	periodicTimeout:
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker ticked")
			case <-timeout:
				fmt.Println("Timeout")
				break periodicTimeout
			}
		}

	}

	// TODO: Pattern 4: Deadline-based timeout
	deadlineTimeout := func() {
		// TODO: Calculate remaining time to deadline
		// TODO: Use dynamic timeout based on deadline

		now := time.Now()
		endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, now.Location())
		timeUntilEndOfDay := time.Until(endOfDay)

		fmt.Println("Time until end of day:", timeUntilEndOfDay)
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

	deadlineTimeout:
		for {
			select {
			case <-time.After(timeUntilEndOfDay):
				fmt.Println("Time until end of day reached")
				break deadlineTimeout
			case <-ticker.C:
				fmt.Println("Ticker ticked")
			}
		}
	}

	// TODO: Execute timeout patterns
	simpleTimeout()
	contextTimeout()
	periodicTimeout()
	deadlineTimeout()

	fmt.Println("Timeout patterns demonstration completed!")
}

// TUTOR: Channel multiplexing enables complex data flow patterns by routing messages between multiple channels.
// Fan-in: merge multiple input channels into one output. Fan-out: distribute from one input to multiple outputs.
// Multiplexing: route messages to different channels based on content or conditions.
// Use multiplexing for: load balancing, message routing, pipeline stages, distributed system coordination.
// These patterns are building blocks for scalable concurrent architectures and microservice communication.
// TODO: Demonstrate channel multiplexing
func demonstrateChannelMultiplexing() {
	fmt.Println("\n=== Channel Multiplexing ===")

	// TODO: Fan-in pattern: Merge multiple input channels
	fanInChannelOne := make(chan string)
	fanInChannelTwo := make(chan string)

	fanIn := func(input1, input2 <-chan string) <-chan string {
		combinedChannel := make(chan string)

		// Start a goroutine to continuously merge channels
		go func() {
			defer close(combinedChannel)

			for {
				select {
				case msg, ok := <-input1:
					if !ok {
						input1 = nil // Disable this channel
					} else {
						combinedChannel <- "FROM CH1: " + msg
					}
				case msg, ok := <-input2:
					if !ok {
						input2 = nil // Disable this channel
					} else {
						combinedChannel <- "FROM CH2: " + msg
					}
				}

				// Exit when both channels are closed
				if input1 == nil && input2 == nil {
					break
				}
			}
		}()

		return combinedChannel
	}

	// Producer goroutine - sends to both channels randomly
	go func() {
		defer close(fanInChannelOne)
		defer close(fanInChannelTwo)

		for counter := 0; counter < 8; counter++ {
			// Random delay to simulate real work
			time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)

			// Send to random channel
			picker := rand.Intn(2)
			if picker == 0 {
				fanInChannelOne <- fmt.Sprintf("Message %d", counter)
			} else {
				fanInChannelTwo <- fmt.Sprintf("Message %d", counter)
			}
		}

		fmt.Println("🎭 Producer finished sending messages")
	}()

	fmt.Println("🔀 Starting Fan-In demonstration...")
	fanInChannel := fanIn(fanInChannelOne, fanInChannelTwo)

	fmt.Println("📡 Listening for merged messages...")
	for msg := range fanInChannel {
		fmt.Printf("✅ Merged: %s\n", msg)
	}

	fmt.Println("🎉 Fan-In completed - all messages merged!")

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// TODO: Fan-out pattern: Distribute to multiple output channels
	fanOut := func(input <-chan string, output1, output2 chan<- string) {
		defer close(output1) // ✅ Producer closes outputs
		defer close(output2)

		for msg := range input { // ✅ Clean termination when input closes
			// Fan-out: send to whichever output is available first
			select {
			case output1 <- "TO-OUT1: " + msg:
				fmt.Printf("📤 Distributed to output1: %s\n", msg)
			case output2 <- "TO-OUT2: " + msg:
				fmt.Printf("📤 Distributed to output2: %s\n", msg)
			default:
				// Both outputs full - handle backpressure
				fmt.Printf("⚠️  Both outputs busy, dropping: %s\n", msg)
			}
		}
		fmt.Println("🎭 Fan-out completed")
	}

	output1 := make(chan string)
	output2 := make(chan string)

	wg := sync.WaitGroup{}

	wg.Add(3)

	// Consumer 1 - processes messages from output1
	go func() {
		defer wg.Done()
		for msg := range output1 { // ✅ Range exits when fanOut closes channel
			fmt.Printf("🔵 Consumer1 processing: %s\n", msg)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("🔵 Consumer1 finished: %s\n", msg)
		}
		fmt.Println("🔵 Consumer1 shut down")
	}()

	// Consumer 2 - processes messages from output2
	go func() {
		defer wg.Done()
		for msg := range output2 { // ✅ Range exits when fanOut closes channel
			fmt.Printf("🔴 Consumer2 processing: %s\n", msg)
			time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
			fmt.Printf("🔴 Consumer2 finished: %s\n", msg)
		}
		fmt.Println("🔴 Consumer2 shut down")
	}()

	inputChannel := make(chan string)

	// Start the fan-out processor
	go fanOut(inputChannel, output1, output2)

	// Producer - generates input for fan-out
	fmt.Println("🔀 Starting Fan-Out demonstration...")
	go func() {
		defer wg.Done()
		defer close(inputChannel)

		for i := 0; i < 6; i++ {
			msg := fmt.Sprintf("FanOut-Message-%d", i)
			fmt.Printf("📥 Producing: %s\n", msg)
			inputChannel <- msg
			time.Sleep(800 * time.Millisecond)
		}
		fmt.Println("🎭 Fan-out producer finished")
	}()

	// Wait for everything to complete (consumers will shut down when channels close)
	wg.Wait()
	fmt.Println("🎉 Fan-Out demonstration completed!")

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// TODO: Multiplexer pattern: Route based on message content
	multiplexer := func(input <-chan Message, outputs map[string]chan Message) {
		// TODO: Route messages to appropriate output channel
		// TODO: Use select for non-blocking sends

		defer func() {
			for name, ch := range outputs {
				close(ch)
				fmt.Println("Closed", name)
			}
		}()

		for msg := range input {
			select {
			case outputs[msg.Type] <- msg:
				fmt.Println("Sent to", msg.Type)
			default:
				fmt.Println("No output channel ready")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}

	// TODO: Test multiplexing patterns
	// TODO: Create test data and show routing behavior
	messages := []Message{
		{Type: "sea", Content: "Sea mail", From: "China"},
		{Type: "air", Content: "Air mail", From: "USA"},
		{Type: "land", Content: "Land mail", From: "Canada"},
		{Type: "sea", Content: "Sea mail", From: "China"},
		{Type: "air", Content: "Air mail", From: "USA"},
		{Type: "land", Content: "Land mail", From: "Canada"},
		{Type: "sea", Content: "Sea mail", From: "China"},
	}

	outputs := map[string]chan Message{
		"sea":  make(chan Message),
		"air":  make(chan Message),
		"land": make(chan Message),
	}

	msgChannel := make(chan Message)

	wg.Add(5)

	go func() {
		defer wg.Done()
		multiplexer(msgChannel, outputs)
	}()

	go func() {
		defer wg.Done()
		for msg := range outputs["sea"] {
			fmt.Println("Received", msg)
		}
	}()

	go func() {
		defer wg.Done()
		for msg := range outputs["air"] {
			fmt.Println("Received", msg)
		}
	}()

	go func() {
		defer wg.Done()
		for msg := range outputs["land"] {
			fmt.Println("Received", msg)
		}
	}()

	go func() {
		defer wg.Done()
		defer close(msgChannel)
		for _, msg := range messages {
			msgChannel <- msg
		}
	}()

	wg.Wait()

	fmt.Println("Channel multiplexing demonstration completed!")
}

// TUTOR: Priority channels ensure important messages are processed before less critical ones.
// Select normally chooses randomly when multiple cases are ready, but we can implement priority using nested selects.
// High-priority channels are checked first, low-priority only when high-priority is empty.
// Use priority for: urgent vs normal messages, admin vs user operations, critical vs batch processing.
// Essential for systems where some operations are more time-sensitive than others.
// TODO: Demonstrate priority channels
func demonstratePriorityChannels() {
	fmt.Println("\n=== Priority Channels ===")

	// TODO: Show how select doesn't guarantee priority
	// Multiple cases ready simultaneously are chosen randomly

	// TODO: Implement priority using nested selects
	// prioritySelect := func(highPriority, lowPriority <-chan string) {
	// TODO: Always check high priority first
	// TODO: Only check low priority if high priority not ready
	// }

	// TODO: Implement priority queue using multiple channels
	type PriorityQueue struct {
		// TODO: Define channels for different priority levels
	}

	// TODO: Test priority handling
	// Send to both high and low priority channels simultaneously
	// Show that high priority is always processed first

	fmt.Println("Priority channels demonstration completed!")
}

// TUTOR: Coordination patterns use select to implement complex synchronization between multiple goroutines.
// Barrier sync: wait for all goroutines to reach a point.
// Leader election: choose one goroutine to lead.
// Work stealing: idle workers take work from busy ones.
// Circuit breaker: fail fast when system is unhealthy.
// Use coordination for: distributed algorithms, fault tolerance, load balancing, consensus protocols.
// These patterns are fundamental to building reliable distributed systems and parallel algorithms.
// TODO: Demonstrate coordination patterns
func demonstrateCoordinationPatterns() {
	fmt.Println("\n=== Coordination Patterns ===")

	// TODO: Pattern 1: Barrier synchronization
	barrierSync := func() {
		// TODO: Wait for multiple goroutines to reach barrier
		// TODO: Use select to coordinate release
	}

	// TODO: Pattern 2: Leader election
	leaderElection := func() {
		// TODO: Multiple goroutines compete to become leader
		// TODO: Use select for non-blocking leadership attempts
	}

	// TODO: Pattern 3: Work stealing
	workStealing := func() {
		// TODO: Workers try to steal work from each other
		// TODO: Use select to attempt stealing without blocking
	}

	// TODO: Pattern 4: Circuit breaker
	circuitBreaker := func() {
		// TODO: Use select to implement circuit breaker logic
		// TODO: Handle success, failure, and timeout cases
	}

	// TODO: Execute coordination patterns
	barrierSync()
	leaderElection()
	workStealing()
	circuitBreaker()

	fmt.Println("Coordination patterns demonstration completed!")
}

// TUTOR: Dynamic channel selection allows runtime control of which channels participate in select operations.
// By setting channels to nil, you can enable/disable select cases based on runtime conditions.
// This creates powerful state machines and conditional communication patterns without complex if-else logic.
// Use dynamic selection for: feature toggles, protocol state machines, adaptive systems, configuration changes.
// Pattern: set channel to nil to "mask" it from select, assign real channel to "unmask" it.
// TODO: Demonstrate dynamic channel selection
func demonstrateDynamicSelection() {
	fmt.Println("\n=== Dynamic Channel Selection ===")

	// TODO: Show how to enable/disable channels using nil
	dynamicChannels := func() {
		// TODO: Create multiple channels
		// TODO: Set channels to nil to disable them in select
		// TODO: Show how select ignores nil channels
	}

	// TODO: Implement conditional channel operations
	conditionalOps := func() {
		// TODO: Enable/disable channels based on conditions
		// TODO: Show dynamic behavior based on state
	}

	// TODO: Round-robin channel selection
	// roundRobin := func(channels []chan string) {
	// TODO: Select from channels in round-robin fashion
	// TODO: Use index tracking and nil channels
	// }

	// TODO: Test dynamic selection
	dynamicChannels()
	conditionalOps()

	fmt.Println("Dynamic selection demonstration completed!")
}

// TUTOR: Select statement gotchas are common misunderstandings that can cause bugs in concurrent programs.
// Random selection, nil channel behavior, default case implications, and closed channel handling.
// Understanding these behaviors is crucial for writing correct concurrent code and debugging concurrency issues.
// Use this knowledge to: avoid race conditions, design robust communication patterns, debug select-related bugs.
// These aren't bugs - they're features! But they can surprise developers coming from other languages.
// TODO: Demonstrate select statement gotchas
func demonstrateSelectGotchas() {
	fmt.Println("\n=== Select Statement Gotchas ===")

	// TODO: Gotcha 1: Random selection when multiple cases ready
	randomSelection := func() {
		// TODO: Show that select is non-deterministic
		// TODO: Multiple ready channels chosen randomly
	}

	// TODO: Gotcha 2: Empty select blocks forever
	// emptySelect := func() {
	// TODO: Show what happens with select{}
	// Use timeout to avoid hanging
	// }

	// TODO: Gotcha 3: Nil channel cases are ignored
	nilChannelCases := func() {
		// TODO: Show that nil channels don't participate in select
		// TODO: Demonstrate this as feature, not bug
	}

	// TODO: Gotcha 4: Default case prevents blocking
	defaultPreventsBlocking := func() {
		// TODO: Show that default always makes select non-blocking
		// TODO: Sometimes this isn't what you want
	}

	// TODO: Gotcha 5: Channel closing behavior in select
	closingBehavior := func() {
		// TODO: Show how closed channels behave in select
		// TODO: Closed channels are always ready to receive
	}

	// TODO: Demonstrate each gotcha safely
	randomSelection()
	// emptySelect() // Skip to avoid hanging
	nilChannelCases()
	defaultPreventsBlocking()
	closingBehavior()

	fmt.Println("Select gotchas demonstration completed!")
}

// TUTOR: Performance considerations help choose the right concurrency patterns for performance-critical code.
// Select has overhead: case evaluation, random selection, channel state checking - but scales well.
// Buffer sizes, number of cases, polling frequency all affect performance characteristics.
// Use performance knowledge for: optimizing hot paths, choosing buffer sizes, balancing responsiveness vs throughput.
// Measure first, optimize second - premature optimization is often unnecessary for select operations.
// TODO: Demonstrate performance considerations
func demonstratePerformanceConsiderations() {
	fmt.Println("\n=== Performance Considerations ===")

	// TODO: Measure select performance vs if-else chains
	// selectVsIfElse := func() {
	// TODO: Compare performance of select vs nested if-else
	// }

	// TODO: Impact of number of cases in select
	// selectCaseCount := func() {
	// TODO: Measure performance with different number of cases
	// TODO: Show that select scales well
	// }

	// TODO: Default case performance impact
	// defaultCaseImpact := func() {
	// 	// TODO: Compare select with and without default case
	// 	// TODO: Show polling vs blocking performance
	// }

	// TODO: Channel buffer size impact on select
	// bufferSizeImpact := func() {
	// TODO: Show how buffer size affects select performance
	// }

	// TODO: Run performance tests
	// selectVsIfElse()
	// selectCaseCount()
	// defaultCaseImpact()
	// bufferSizeImpact()

	fmt.Println("Performance testing completed!")
}

// TUTOR: Real-world select patterns show how select enables production-grade concurrent systems.
// HTTP servers with timeouts, distributed system coordination, message queue consumers, background job processors.
// These patterns combine multiple select concepts to solve complex real-world concurrency challenges.
// Use these patterns for: building web services, microservice communication, job processing systems, monitoring tools.
// These are proven patterns used in production systems - learn them to build robust concurrent applications.
// TODO: Demonstrate real-world select patterns
func demonstrateRealWorldPatterns() {
	fmt.Println("\n=== Real-World Select Patterns ===")

	// TODO: HTTP server with timeout
	// httpServerPattern := func() {
	// TODO: Handle HTTP requests with timeout using select
	// TODO: Show graceful degradation
	// }

	// TODO: Distributed system coordination
	// distributedPattern := func() {
	// TODO: Coordinate multiple nodes using select
	// TODO: Handle node failures and timeouts
	// }

	// TODO: Message queue consumer
	// messageQueuePattern := func() {
	// TODO: Consume from multiple queues with priority
	// TODO: Handle queue failures gracefully
	// }

	// TODO: Background job processor
	// jobProcessorPattern := func() {
	// TODO: Process jobs with cancellation support
	// TODO: Handle shutdown signals
	// }

	// TODO: Execute real-world patterns
	// httpServerPattern()
	// distributedPattern()
	// messageQueuePattern()
	// jobProcessorPattern()

	fmt.Println("Real-world patterns demonstration completed!")
}

// Message type for examples
type Message struct {
	Type    string
	Content string
	From    string
}

// Helper function to create timed channel
func timedChannel(msg string, delay time.Duration) <-chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(delay)
		ch <- msg
	}()
	return ch
}

// Helper function to simulate work with random duration
func simulateWork(id string, minMs, maxMs int) <-chan string {
	ch := make(chan string, 1)
	go func() {
		duration := time.Duration(rand.Intn(maxMs-minMs)+minMs) * time.Millisecond
		time.Sleep(duration)
		ch <- fmt.Sprintf("Work %s completed in %v", id, duration)
	}()
	return ch
}

// Helper function to safely demonstrate with timeout
func safelyDemonstrate(name string, fn func(), timeout time.Duration) {
	fmt.Printf("--- %s ---\n", name)
	done := make(chan bool)

	go func() {
		fn()
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Completed successfully")
	case <-time.After(timeout):
		fmt.Println("Timed out")
	}
}

func main() {
	fmt.Println("🎛️  Welcome to Select Statements! 🎛️")
	fmt.Println("This file teaches you channel multiplexing and coordination")

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// TODO: Implement each demonstration function
	// Start with basic select and progress to advanced patterns

	// demonstrateBasicSelect()
	// demonstrateSelectDefault()
	// demonstrateTimeoutPatterns()
	demonstrateChannelMultiplexing()
	// demonstratePriorityChannels()
	// demonstrateCoordinationPatterns()
	// demonstrateDynamicSelection()
	// demonstrateSelectGotchas()
	// demonstratePerformanceConsiderations()
	// demonstrateRealWorldPatterns()

	fmt.Println("\n🎉 Congratulations! You've mastered select statements!")
	fmt.Println("Next: Learn advanced patterns in channel_patterns.go")
}

/*
🔍 Key Concepts to Remember:

1. **Select Syntax**: select { case <-ch1: ... case ch2 <- val: ... default: ... }
2. **Non-blocking**: Default case makes select non-blocking
3. **Random Selection**: When multiple cases ready, Go chooses randomly
4. **Nil Channels**: Nil channels in select are ignored
5. **Closed Channels**: Always ready to receive (returns zero value)
6. **Empty Select**: select{} blocks forever
7. **Timeout Pattern**: Use time.After() for timeouts

🎛️ Select Patterns:
```go
// Basic timeout
select {
case result := <-resultCh:
    return result
case <-time.After(5*time.Second):
    return errors.New("timeout")
}

// Non-blocking send
select {
case ch <- value:
    // Sent successfully
default:
    // Channel full, handle accordingly
}

// Fan-in
select {
case msg := <-ch1:
    output <- msg
case msg := <-ch2:
    output <- msg
}

// Dynamic channel disabling
if condition {
    ch = nil // Disable channel in select
}
select {
case <-ch: // Ignored if ch is nil
    // Handle message
default:
    // Handle other cases
}
```

🎯 When to Use Select:
- **Timeouts**: Operations that might hang
- **Multiplexing**: Handle multiple channels
- **Non-blocking**: Avoid goroutine blocking
- **Coordination**: Complex synchronization patterns
- **Priority**: Prefer certain channels over others

🚨 Select Gotchas:
- Random selection when multiple cases ready
- Default case always makes select non-blocking
- Nil channels are completely ignored
- Closed channels always ready to receive
- Empty select{} blocks forever

🎯 Next Steps:
- Master advanced channel patterns
- Build complex concurrent systems
- Apply patterns to real-world problems
- Complete practice exercises
*/
