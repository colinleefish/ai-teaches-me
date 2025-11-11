// Week 8: Channel Basics
// This file demonstrates the fundamentals of channels in Go

package main

import (
	"fmt"
	"sync"
	"time"
)

// TUTOR: Basic channels are Go's foundational communication mechanism between goroutines.
// They implement CSP (Communicating Sequential Processes) - "Don't communicate by sharing memory; share memory by communicating".
// Unbuffered channels are synchronous - they block until both sender and receiver are ready, providing natural synchronization.
// Use channels for: coordinating goroutines, passing data safely between concurrent operations, signaling completion.
// Characteristics: Type-safe, first-class values, can be passed to functions, stored in data structures.
// TODO: Demonstrate basic channel creation and usage
func demonstrateBasicChannels() {
	fmt.Println("=== Basic Channel Operations ===")

	// TODO: Create an unbuffered channel
	// Use make(chan type) syntax
	ch := make(chan int)

	// TODO: Send and receive in separate goroutines
	// Show that unbuffered channels block until both sender and receiver are ready
	go func() {
		ch <- 1
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch)
	}()

	// TODO: Demonstrate the synchronization property
	// Show how channels coordinate goroutines
	time.Sleep(1 * time.Second)

	fmt.Println("Basic channel operations completed!")
}

// TUTOR: Channel blocking is the key to Go's synchronization model. Understanding blocking behavior prevents deadlocks.
// Unbuffered channels block on send until someone receives, and block on receive until someone sends.
// This creates natural "handshake" synchronization - operations complete together or not at all.
// Use blocking for: ensuring operations happen in order, creating barriers between goroutines, implementing semaphores.
// Critical: Always ensure there's a corresponding receiver for every sender (and vice versa) to avoid deadlocks.
// TODO: Demonstrate channel blocking behavior
func demonstrateChannelBlocking() {
	fmt.Println("\n=== Channel Blocking Behavior ===")

	// TODO: Show blocking send on unbuffered channel
	// Create a channel and try to send without a receiver
	// Use a goroutine to avoid deadlock and show the blocking
	ch := make(chan int)
	// go func() {
	// 	fmt.Println("Sending 6 in 1 second")
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println("Sending 6")
	// 	ch <- 6
	// 	fmt.Println("Sent 6")
	// }()

	// time.Sleep(3 * time.Second)

	// TODO: Show blocking receive on empty channel
	// Try to receive from channel with no sender
	// go func() {
	// 	fmt.Println("Receiving from channel")
	// 	fmt.Println(<-ch)
	// 	fmt.Println("Received from channel")
	// }()

	// time.Sleep(3 * time.Second)

	// TODO: Show synchronous nature of unbuffered channels
	// Demonstrate that send doesn't complete until receive happens

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Sending 6")
		ch <- 6
		fmt.Println("Sent 6")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Receiving from channel")
		fmt.Println(<-ch)
		fmt.Println("Received from channel")
	}()

	wg.Wait()

	fmt.Println("Channel blocking demonstration completed!")
}

// TUTOR: Channel directions enforce communication patterns at compile-time, preventing bugs and clarifying API contracts.
// Send-only (chan<-) channels can only send values, receive-only (<-chan) can only receive values.
// Go automatically converts bidirectional channels to directional ones when passed to functions with directional parameters.
// Use directional channels for: API design clarity, preventing misuse, documenting data flow intentions.
// Pattern: Create bidirectional channel, pass as directional to specialized functions (producers get send-only, consumers get receive-only).
// TODO: Demonstrate channel directions (send-only, receive-only)
func demonstrateChannelDirections() {
	fmt.Println("\n=== Channel Directions ===")

	// TODO: Create functions with directional channel parameters

	// Send-only channel parameter

	ch := make(chan string)
	sender := func(ch chan<- string, message string) {
		ch <- message
	}

	// Receive-only channel parameter
	// receiver := func(ch <-chan string) string {
	//     // TODO: Receive message from channel
	//     // Show that you can only receive, not send
	//     return ""
	// }
	receiver := func(ch <-chan string) string {
		return <-ch
	}

	// TODO: Create a bidirectional channel
	// Pass it to sender and receiver functions
	// Show how Go automatically converts to directional channels

	go sender(ch, "Hello")

	fmt.Println(receiver(ch))

	fmt.Println("Channel directions demonstration completed!")
}

// TUTOR: Channel closing signals "no more data will be sent" and is crucial for graceful shutdown patterns.
// Only senders should close channels (receivers detect closure). Closing enables range loops to exit naturally.
// A closed channel: allows receives (returns zero value), panics on send, panics on re-close.
// Use closing for: signaling completion to range loops, broadcasting shutdown to multiple goroutines, resource cleanup.
// Pattern: defer close(ch) in producer functions ensures channel always closes, even on panic.
// TODO: Demonstrate channel closing
func demonstrateChannelClosing() {
	fmt.Println("\n=== Channel Closing ===")

	// TODO: Show how to close a channel
	// Demonstrate that closing is done by sender, not receiver
	wg := sync.WaitGroup{}

	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// ch <- 1
		close(ch)
	}()

	// TODO: Show receiving from closed channel
	// Demonstrate that receives from closed channel return zero value
	// Show the two-value receive to check if channel is closed

	time.Sleep(1 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()

	// TODO: Show what happens when sending to closed channel
	// This should panic - demonstrate with recover

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		ch <- 1
	}()

	// TODO: Show multiple closes cause panic
	// Use recover to catch the panic

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		close(ch)
		close(ch)
	}()

	wg.Wait()

	fmt.Println("Channel closing demonstration completed!")
}

// TUTOR: Range over channels provides elegant producer-consumer patterns with automatic termination.
// 'for value := range ch' automatically exits when channel is closed - no manual termination logic needed.
// Range blocks waiting for values and only exits when sender calls close(ch).
// Use range for: processing all values from a channel, clean producer-consumer patterns, avoiding manual loop management.
// Critical: Producer MUST close the channel or range will block forever. Use defer close(ch) for safety.
// Demonstrate range over channels with proper synchronization
func demonstrateChannelRange() {
	fmt.Println("\n=== Range Over Channels ===")

	var wg sync.WaitGroup

	// Producer: sends values and closes channel
	producer := func(ch chan<- int) {
		defer wg.Done() // Signal producer is done
		defer close(ch)
		fmt.Println("Producer: Starting to send values...")

		for i := 1; i <= 5; i++ {
			fmt.Printf("Producer: Sending %d\n", i)
			ch <- i
			// time.Sleep(100 * time.Millisecond) // Simulate work
		}

		fmt.Println("Producer: Function ending (defer will close channel)")
	}

	// Consumer: uses range to receive ALL values until channel closes
	consumer := func(ch <-chan int) {
		defer wg.Done() // Signal consumer is done
		fmt.Println("Consumer: Starting to receive values...")

		// Range automatically exits when channel is closed
		for value := range ch {
			fmt.Printf("Consumer: Received %d\n", value)
			// time.Sleep(50 * time.Millisecond) // Simulate processing
		}

		fmt.Println("Consumer: Channel closed, exiting range loop")
	}

	// Set up synchronization
	ch := make(chan int)
	wg.Add(2) // We have 2 goroutines to wait for

	// Launch producer and consumer
	go producer(ch)
	go consumer(ch)

	// Wait for both to complete
	wg.Wait()
	fmt.Println("Range demonstration completed - both goroutines finished!")
}

// TUTOR: State machines with nil channels demonstrate how to control program flow dynamically.
// By setting channels to nil, we can enable/disable different behaviors in select statements.
// This creates clean, readable state machines without complex if-else logic.
// Use this pattern for: game states, protocol handlers, workflow management, UI state management.
// Pattern: Create channels for each state, set to nil to disable, create/assign to enable.
func demonstrateChannelStateMachine() {
	fmt.Println("\n=== Channel State Machine ===")
	fmt.Println("Simulating a simple download manager with pause/resume/cancel")

	// State channels - nil means state is disabled
	var startCh chan bool  // Start download
	var pauseCh chan bool  // Pause download
	var resumeCh chan bool // Resume download
	var cancelCh chan bool // Cancel download

	// Status
	downloading := false
	paused := false
	cancelled := false

	// Helper to print current state
	printState := func() {
		if cancelled {
			fmt.Println("📛 Status: CANCELLED")
		} else if downloading {
			fmt.Println("⬇️  Status: DOWNLOADING")
		} else if paused {
			fmt.Println("⏸️  Status: PAUSED")
		} else {
			fmt.Println("⏹️  Status: IDLE")
		}
	}

	// State machine function
	updateState := func(action string) {
		// Reset all channels to nil (disable all actions)
		startCh = nil
		pauseCh = nil
		resumeCh = nil
		cancelCh = nil

		switch {
		case cancelled:
			// Terminal state - no actions possible
			fmt.Println("🚫 Download cancelled - no actions available")

		case downloading:
			// Can pause or cancel while downloading
			pauseCh = make(chan bool, 1)
			cancelCh = make(chan bool, 1)
			fmt.Println("⚡ Available actions: PAUSE, CANCEL")

		case paused:
			// Can resume or cancel while paused
			resumeCh = make(chan bool, 1)
			cancelCh = make(chan bool, 1)
			fmt.Println("⚡ Available actions: RESUME, CANCEL")

		default: // idle
			// Can start download
			startCh = make(chan bool, 1)
			fmt.Println("⚡ Available actions: START")
		}
	}

	// Simulate user actions
	actions := []struct {
		name    string
		trigger func()
	}{
		{"START", func() {
			if startCh != nil {
				startCh <- true
			}
		}},
		{"PAUSE", func() {
			if pauseCh != nil {
				pauseCh <- true
			}
		}},
		{"RESUME", func() {
			if resumeCh != nil {
				resumeCh <- true
			}
		}},
		{"CANCEL", func() {
			if cancelCh != nil {
				cancelCh <- true
			}
		}},
	}

	// Initial state
	printState()
	updateState("init")

	// Simulate a sequence of user actions
	actionSequence := []int{0, 1, 2, 1, 3} // START, PAUSE, RESUME, PAUSE, CANCEL

	for i, actionIdx := range actionSequence {
		fmt.Printf("\n--- Step %d: User clicks %s ---\n", i+1, actions[actionIdx].name)

		// Trigger the action
		actions[actionIdx].trigger()

		// Process state changes with select
		select {
		case <-startCh:
			fmt.Println("🎬 Starting download...")
			downloading = true
			paused = false

		case <-pauseCh:
			fmt.Println("⏸️  Pausing download...")
			downloading = false
			paused = true

		case <-resumeCh:
			fmt.Println("▶️  Resuming download...")
			downloading = true
			paused = false

		case <-cancelCh:
			fmt.Println("❌ Cancelling download...")
			downloading = false
			paused = false
			cancelled = true

		default:
			fmt.Println("❗ Action not available in current state")
		}

		printState()
		updateState("update")
	}

	fmt.Println("\n💡 Notice how nil channels elegantly disable unavailable actions!")
	fmt.Println("Channel state machine demonstration completed!")
}

// TUTOR: Nil channels have special blocking behavior useful for disabling communication paths dynamically.
// Send/receive on nil channels blocks forever - they never succeed. This is by design, not a bug.
// Use nil channels for: disabling select cases dynamically, implementing timeouts, creating optional communication paths.
// Pattern: Set channel to nil to "turn off" that case in a select statement - very powerful for state machines.
// Common in multiplexers where you want to temporarily disable certain inputs/outputs.
// TODO: Demonstrate nil channels
func demonstrateNilChannels() {
	fmt.Println("\n=== Nil Channels ===")

	// TODO: Show nil channel behavior
	// var nilCh chan int
	var nilCh chan int

	// TODO: Show that sending to nil channel blocks forever
	// Use select with default to avoid blocking

	select {
	case nilCh <- 1:
		fmt.Println("Sent to nil channel")
	default:
		fmt.Println("Default case executed")
	}

	// TODO: Show that receiving from nil channel blocks forever
	// Use select with default to avoid blocking
	select {
	case value := <-nilCh:
		fmt.Printf("Received from nil channel: %d\n", value)
	default:
		fmt.Println("Receive from nil channel blocked (default case executed)")
	}

	// TODO: Show practical use of nil channels
	// Use nil channels to disable cases in select statements
	fmt.Println("\n--- Practical Use: Disabling Select Cases ---")

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ch1 <- "from ch1"
	ch2 <- "from ch2"

	// Demonstrate disabling ch2 by setting it to nil
	fmt.Println("Before disabling ch2:")
	select {
	case msg := <-ch1:
		fmt.Printf("Received: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No channels ready")
	}

	// FIX: Drain ch1 to make sure it's empty
	select {
	case msg := <-ch1:
		fmt.Printf("Drained remaining from ch1: %s\n", msg)
	default:
		// ch1 was already empty
	}

	// Disable ch2 by setting to nil
	ch2 = nil
	ch1 <- "from ch1 again"

	fmt.Println("After disabling ch2 (set to nil):")
	select {
	case msg := <-ch1:
		fmt.Printf("Received: %s (ch2 was skipped because it's nil)\n", msg)
	case msg := <-ch2: // This case will never execute because ch2 is nil
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No channels ready")
	}

	fmt.Println("Nil channels demonstration completed!")
}

// TUTOR: Channels are first-class values in Go - they can be stored, passed, returned, and manipulated like any other value.
// This enables powerful patterns: channel factories, channel registries, dynamic channel creation/routing.
// Channels can be stored in slices, maps, structs - enabling complex communication topologies.
// Use channels as values for: building message routers, creating channel pools, implementing pub-sub systems.
// Pattern: Functions can return channels, allowing callers to receive results asynchronously.
// TODO: Demonstrate channel as first-class values
func demonstrateChannelsAsValues() {
	fmt.Println("\n=== Channels as First-Class Values ===")

	wg := sync.WaitGroup{}

	// TODO: Show channels can be stored in variables

	// TODO: Show channels can be passed as function parameters
	// pubFuncThatAcceptsChannel := func(ch chan<- int) {
	// 	defer wg.Done()
	// 	defer close(ch)
	// 	ch <- 1
	// }

	subFuncThatAcceptsChannel := func(ch <-chan int) int {
		defer wg.Done()
		return <-ch
	}

	// wg.Add(2)
	// chPassedAsValue := make(chan int)
	// go pubFuncThatAcceptsChannel(chPassedAsValue)
	// time.Sleep(1 * time.Second)
	// fmt.Println(subFuncThatAcceptsChannel(chPassedAsValue))
	// wg.Wait()

	// TODO: Show channels can be returned from functions
	wg.Add(2)
	chGenerator := func() chan int {
		ch := make(chan int)
		go func() {
			defer wg.Done()
			ch <- 15
			close(ch)
		}()
		return ch
	}

	chPassedAsValue := chGenerator()
	fmt.Println(subFuncThatAcceptsChannel(chPassedAsValue))

	// TODO: Show channels can be stored in slices/maps

	// TODO: Create a channel factory function
	createIntChannel := func(bufferSize int) chan int {
		// TODO: Return a new channel with given buffer size
		return make(chan int, bufferSize)
	}

	ch := createIntChannel(10)
	ch <- 7
	fmt.Println(<-ch)

	// TODO: Create a channel registry
	channelRegistry := make(map[string]chan int)
	channelRegistry["ch10"] = createIntChannel(10)
	channelRegistry["ch20"] = createIntChannel(20)

	// TODO: Store and retrieve channels from registry
	// Show how channels can be managed dynamically
	channelRegistry["ch10"] <- 10
	channelRegistry["ch20"] <- 20
	channelRegistry["ch10"] <- 30
	channelRegistry["ch20"] <- 40
	fmt.Println(<-channelRegistry["ch10"])
	fmt.Println(<-channelRegistry["ch20"])
	fmt.Println(<-channelRegistry["ch10"])
	fmt.Println(<-channelRegistry["ch20"])

	fmt.Println("Channels as values demonstration completed!")
}

// TUTOR: Channel communication patterns are reusable solutions for common concurrency problems.
// Request-Response: Send request on one channel, receive response on another - enables async RPC.
// Worker Notification: Use channels to signal when work is complete - better than polling.
// Broadcast via Closing: Closing a channel signals all receivers simultaneously - elegant shutdown pattern.
// Use these patterns for: building distributed systems, coordinating workflows, implementing event systems.
// These patterns compose well - combine them to build complex concurrent architectures.
// TODO: Demonstrate channel communication patterns
func demonstrateCommunicationPatterns() {
	fmt.Println("\n=== Communication Patterns ===")

	// TODO: Pattern 1: Request-Response
	requestResponse := func() {
		// TODO: Create request and response channels
		// TODO: Send request and wait for response
		// Show bidirectional communication
	}

	// TODO: Pattern 2: Worker notification
	workerNotification := func() {
		// TODO: Use channel to signal worker completion
		// Show how channels can coordinate work
	}

	// TODO: Pattern 3: Broadcast using channel closing
	broadcast := func() {
		// TODO: Use channel closing to signal multiple goroutines
		// Show how closing can be used for broadcast
	}

	// TODO: Execute each pattern
	requestResponse()
	workerNotification()
	broadcast()

	fmt.Println("Communication patterns demonstration completed!")
}

// TUTOR: Channels vs traditional synchronization (mutexes, locks) trade-offs and when to use each.
// Channels excel at: data passing, coordination, producer-consumer patterns, composable operations.
// Mutexes excel at: protecting shared state, low-overhead operations, traditional critical sections.
// Channels are higher-level, more composable but with slight overhead. Mutexes are lower-level, faster but less composable.
// Use channels for: communication between goroutines, data pipelines, event-driven architecture.
// Use mutexes for: protecting shared data structures, performance-critical sections, simple state protection.
// TODO: Demonstrate channel vs other synchronization
func demonstrateChannelVsOther() {
	fmt.Println("\n=== Channels vs Other Synchronization ===")

	// TODO: Compare channel-based solution vs mutex-based solution
	// for a simple counter problem

	// TODO: Channel-based counter
	// channelCounter := func(n int) int {
	//     // TODO: Use channel to coordinate counter increments
	//     // Show how channels can replace mutexes
	//     return 0
	// }

	// TODO: Show timing comparison
	// Measure performance of channel vs mutex approach

	fmt.Println("Channel comparison demonstration completed!")
}

// TUTOR: Common channel mistakes can cause deadlocks, panics, or goroutine leaks - learn to avoid them.
// Mistake 1: Forgetting to close channels breaks range loops (they wait forever).
// Mistake 2: Sending on closed channels causes panics - check channel state or use select with default.
// Mistake 3: Closing channels multiple times causes panics - use sync.Once or careful state management.
// Mistake 4: Wrong channel directions cause compile errors - understand the flow of your data.
// Prevention: Use defer close(), test thoroughly, understand blocking behavior, follow established patterns.
// TODO: Demonstrate common channel mistakes
func demonstrateCommonMistakes() {
	fmt.Println("\n=== Common Channel Mistakes ===")

	// TODO: Mistake 1: Forgetting to close channel in range
	// mistake1 := func() {
	//     // TODO: Show what happens when you forget to close
	//     // Use timeout to avoid infinite blocking
	// }

	// TODO: Mistake 2: Sending on closed channel
	// mistake2 := func() {
	//     // TODO: Show panic when sending to closed channel
	//     // Use recover to handle the panic
	// }

	// TODO: Mistake 3: Closing channel multiple times
	// mistake3 := func() {
	//     // TODO: Show panic when closing already closed channel
	//     // Use recover to handle the panic
	// }

	// TODO: Mistake 4: Wrong direction expectations
	// mistake4 := func() {
	//     // TODO: Show compilation errors with wrong channel directions
	//     // Comment out the incorrect code with explanations
	// }

	// TODO: Execute each mistake demonstration safely
	fmt.Println("Demonstrating mistakes (safely)...")

	fmt.Println("Common mistakes demonstration completed!")
}

// TUTOR: Channel performance characteristics help you choose the right tool for performance-critical code.
// Channels have overhead: goroutine scheduling, channel operations, memory allocation.
// Performance factors: buffered vs unbuffered, channel size, number of goroutines, contention levels.
// Generally: channels are fast enough for most use cases, but not for microsecond-critical operations.
// Optimization techniques: buffer sizing, reducing channel operations, batching, choosing right patterns.
// Measure first: profile your code to identify actual bottlenecks before optimizing.
// TODO: Demonstrate channel performance characteristics
func demonstrateChannelPerformance() {
	fmt.Println("\n=== Channel Performance ===")

	// TODO: Measure channel operation performance
	// Compare send/receive speeds with different patterns

	// TODO: Test 1: Simple send/receive
	simpleTest := func() {
		// TODO: Measure time for many send/receive operations
	}

	// TODO: Test 2: Goroutine creation overhead vs channel reuse
	goroutineTest := func() {
		// TODO: Compare creating new goroutines vs reusing with channels
	}

	// TODO: Test 3: Channel vs direct function call
	callTest := func() {
		// TODO: Compare channel communication vs direct function calls
	}

	// TODO: Run performance tests
	simpleTest()
	goroutineTest()
	callTest()

	fmt.Println("Channel performance demonstration completed!")
}

// Helper function to simulate work
func doWork(id int, duration time.Duration) {
	fmt.Printf("Worker %d: Starting work for %v\n", id, duration)
	time.Sleep(duration)
	fmt.Printf("Worker %d: Work completed\n", id)
}

// Helper function to safely demonstrate panics
func safelyDemonstratePanic(description string, fn func()) {
	fmt.Printf("--- %s ---\n", description)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Caught panic: %v\n", r)
		}
	}()
	fn()
}

func main() {
	fmt.Println("📡 Welcome to Channel Basics! 📡")
	fmt.Println("This file teaches you the fundamentals of Go channels")

	// TODO: Implement each demonstration function
	// Start with basic operations and progress to advanced concepts

	// demonstrateBasicChannels()
	// demonstrateChannelBlocking()
	// demonstrateChannelDirections()
	// demonstrateChannelClosing()
	// demonstrateChannelRange()
	// demonstrateChannelStateMachine() // 🎯 NEW: State machine with nil channels
	// demonstrateNilChannels()
	demonstrateChannelsAsValues()
	// demonstrateCommunicationPatterns()
	// demonstrateChannelVsOther()
	// demonstrateCommonMistakes()
	// demonstrateChannelPerformance()

	fmt.Println("\n🎉 Congratulations! You've learned channel basics!")
	fmt.Println("Next: Learn about buffered channels in buffered_channels.go")
}

/*
🔍 Key Concepts to Remember:

1. **Channel Creation**: make(chan Type) for unbuffered, make(chan Type, size) for buffered
2. **Send/Receive**: ch <- value to send, value := <-ch to receive
3. **Blocking**: Unbuffered channels block until both sender and receiver are ready
4. **Directions**: chan<- for send-only, <-chan for receive-only
5. **Closing**: close(ch) to close, value, ok := <-ch to check if closed
6. **Range**: for value := range ch automatically handles channel closing
7. **Nil Channels**: Sending/receiving on nil channels blocks forever

📋 Channel Patterns:
```go
// Basic send/receive
ch := make(chan int)
go func() { ch <- 42 }()
value := <-ch

// Producer-consumer with range
go func() {
    defer close(ch)
    for i := 0; i < 10; i++ {
        ch <- i
    }
}()
for value := range ch {
    process(value)
}

// Check if closed
value, ok := <-ch
if !ok {
    // Channel is closed
}
```

🚨 Common Mistakes:
- Deadlock on unbuffered channels (no receiver/sender)
- Forgetting to close channels when using range
- Sending on closed channels (causes panic)
- Closing channels multiple times (causes panic)
- Wrong channel direction expectations

🎯 Next Steps:
- Learn buffered channels for performance
- Master select statements for multiplexing
- Understand advanced channel patterns
- Practice with real-world scenarios
*/
