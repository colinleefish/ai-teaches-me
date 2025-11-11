package main

import (
	"fmt"
	"sync"
	"time"
)

// TUTOR: Channel closing signals "no more values will be sent" to receivers.
// close() function closes a channel, making it ready-only for senders.
// Closing is a one-way operation - once closed, channels cannot be reopened.
// Only senders should close channels, never receivers.
// Understanding closure semantics is essential for proper resource management.
// TODO: Demonstrate basic channel closing behavior
func demonstrateBasicClosing() {
	fmt.Println("=== Basic Channel Closing ===")

	// TODO: Create channel, send values, then close it
	// TODO: Show that closed channels can still be read from
	// TODO: Demonstrate zero value return after channel is drained
	// TODO: Show panic behavior when sending to closed channel

	ch := make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("channel closed")
		defer close(ch)
		ch <- 1
		ch <- 2
	}()

	wg.Wait()

	for v := range ch {
		fmt.Println(v)
	}

	v, ok := <-ch
	fmt.Println("value from channel", v, "ok", ok)

	defer func() {

		if recover() != nil {
			fmt.Println("recovered from panic")
			return
		}
		fmt.Println("done")
	}()

	ch <- 3

}

// TUTOR: The comma ok idiom detects whether a channel is closed.
// value, ok := <-ch returns the value and a boolean status.
// ok is true for normal receives, false for closed+empty channels.
// This idiom enables proper handling of channel lifecycle.
// Always check ok when channel closure timing is uncertain.
// TODO: Demonstrate comma ok idiom for closure detection
func demonstrateCommaOkIdiom() {
	fmt.Println("\n=== Comma Ok Idiom ===")
	fmt.Println(" implemented in demonstrateBasicClosing")
	// TODO: Create channel and send some values
	// TODO: Use comma ok idiom to receive with status
	// TODO: Show different ok values before and after closing
	// TODO: Demonstrate proper closure detection in loops
}

// TUTOR: Range loops automatically handle channel closure detection.
// for value := range ch continues until channel is closed and drained.
// Range loops exit cleanly when channels close, no manual checking needed.
// This creates elegant consumer code that handles closure automatically.
// Range is the preferred way to consume from channels until closure.
// TODO: Demonstrate range loops with channel closure
func demonstrateRangeWithClosure() {
	fmt.Println("\n=== Range Loops and Closure ===")
	fmt.Println(" implemented in demonstrateBasicClosing")
	// TODO: Create producer that sends values then closes channel
	// TODO: Use range loop to consume all values until closure
	// TODO: Show automatic loop termination on channel close
	// TODO: Compare range loop with manual comma ok checking
}

// TUTOR: Only senders should close channels, establishing clear ownership.
// Receivers closing channels can cause "send on closed channel" panics.
// This principle prevents race conditions and undefined behavior.
// Sender-closes pattern makes resource lifecycle predictable.
// When multiple senders exist, use coordination to designate closer.
// TODO: Demonstrate sender-closes principle
func demonstrateSenderCloses() {
	fmt.Println("\n=== Sender-Closes Principle ===")
	fmt.Println(" implemented in demonstrateBasicClosing")
	// TODO: Show proper pattern where sender closes channel
	// TODO: Demonstrate what happens when receiver tries to close
	// TODO: Show coordination patterns for multiple senders
	// TODO: Illustrate clear ownership and responsibility
}

// TUTOR: Closing nil channels causes panic, so check before closing.
// Multiple closes on same channel also cause panic.
// Use sync.Once or flags to ensure channels are closed exactly once.
// Defensive programming prevents closure-related panics in complex systems.
// Always validate channel state before attempting to close.
// TODO: Demonstrate safe closing patterns and common pitfalls
func demonstrateSafeClosing() {
	fmt.Println("\n=== Safe Closing Patterns ===")

	// TODO: Show panic scenarios with nil and already-closed channels
	// TODO: Demonstrate sync.Once pattern for safe closing
	// TODO: Show flag-based approaches to prevent double closing
	// TODO: Illustrate defensive programming practices

}

// TUTOR: defer statements ensure channels are closed even on early returns.
// defer close(ch) at function start guarantees cleanup.
// This pattern prevents resource leaks in error scenarios.
// Defer works well with the sender-closes principle.
// Proper cleanup patterns improve program reliability.
// TODO: Demonstrate defer patterns for channel closing
func demonstrateDeferClosing() {
	fmt.Println("\n=== Defer Closing Patterns ===")

	// TODO: Show defer close() at function start
	// TODO: Demonstrate closure on early returns and errors
	// TODO: Show how defer ensures cleanup in all code paths
	// TODO: Illustrate best practices for resource management
}

// TUTOR: Channel closure propagates through pipeline stages naturally.
// When input channel closes, processing stage can close output channel.
// This creates clean shutdown cascades through complex systems.
// Proper closure handling enables graceful system termination.
// Pipeline stages should handle and propagate closure signals.
// TODO: Demonstrate closure propagation in pipelines
func demonstrateClosurePropagation() {
	fmt.Println("\n=== Closure Propagation ===")

	// TODO: Create multi-stage pipeline with proper closure handling
	// TODO: Show how closure propagates through pipeline stages
	// TODO: Demonstrate graceful shutdown of complex systems
	// TODO: Show proper resource cleanup in pipeline teardown

	wg := sync.WaitGroup{}
	wg.Add(3)

	input := make(chan int, 10)
	intermediate := make(chan int, 10)
	output := make(chan int, 10)

	double := func(in <-chan int, out chan<- int) {
		defer wg.Done()
		defer close(out)
		for v := range in {
			out <- v * 2
		}
	}

	addFive := func(in <-chan int, out chan<- int) {
		defer wg.Done()
		defer close(out)
		for v := range in {
			out <- v + 5
		}
	}

	go addFive(input, intermediate)
	go double(intermediate, output)

	go func() {
		defer wg.Done()
		defer close(input)
		for i := 0; i < 10; i++ {
			input <- i
		}
	}()

	for v := range output {
		fmt.Println(v)
	}

	wg.Wait()
}

// TUTOR: Testing channel closure requires careful timing and coordination.
// Test cases should verify proper closure behavior and timing.
// Use timeouts to prevent tests from hanging on unclosed channels.
// Test both normal and error scenarios for closure patterns.
// Good tests validate channel lifecycle management thoroughly.
// TODO: Demonstrate testing patterns for channel closure
func demonstrateClosureTesting() {
	fmt.Println("\n=== Testing Channel Closure ===")

	// 1. TEST IF CHANNEL IS CLOSED
	fmt.Println("🔍 Testing channel closure detection:")
	ch := make(chan int)

	go func() {
		ch <- 42
		close(ch)
	}()

	// Receive and check if channel is closed
	value, ok := <-ch
	if ok {
		fmt.Printf("✅ Received value: %d (channel open)\n", value)
	}

	value, ok = <-ch // Try to receive again
	if !ok {
		fmt.Println("✅ Channel is closed (ok=false)")
	}

	// 2. TEST WITH TIMEOUT (prevent hanging tests)
	fmt.Println("\n⏰ Testing with timeout:")
	slowChannel := make(chan int)

	go func() {
		time.Sleep(50 * time.Millisecond) // Simulate work
		slowChannel <- 99
		close(slowChannel)
	}()

	select {
	case val := <-slowChannel:
		fmt.Printf("✅ Received: %d within timeout\n", val)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("❌ Test timed out!")
	}

	// Check if closed after timeout
	select {
	case val, ok := <-slowChannel:
		if !ok {
			fmt.Println("✅ Channel properly closed")
		} else {
			fmt.Printf("⚠️ Unexpected value: %d\n", val)
		}
	case <-time.After(10 * time.Millisecond):
		fmt.Println("❌ Channel not closed - potential leak!")
	}

	// 3. TEST RANGE BEHAVIOR WITH CLOSURE
	fmt.Println("\n🔄 Testing range with closure:")
	rangeChannel := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			rangeChannel <- i
		}
		close(rangeChannel)
		fmt.Println("✅ Producer closed channel")
	}()

	received := []int{}
	for val := range rangeChannel {
		received = append(received, val)
	}

	if len(received) == 3 {
		fmt.Printf("✅ Range loop ended correctly, received: %v\n", received)
	} else {
		fmt.Printf("❌ Expected 3 values, got %d: %v\n", len(received), received)
	}

	// 4. TEST DOUBLE CLOSE (PANIC SCENARIO)
	fmt.Println("\n🚨 Testing double close protection:")
	testChannel := make(chan int)

	// First close
	close(testChannel)
	fmt.Println("✅ First close successful")

	// Try to detect if already closed before second close
	select {
	case _, ok := <-testChannel:
		if !ok {
			fmt.Println("✅ Channel already closed - avoiding double close")
		}
	default:
		fmt.Println("⚠️ Channel state unclear")
	}

	// 5. TEST PIPELINE CLOSURE PROPAGATION
	fmt.Println("\n🔗 Testing pipeline closure propagation:")
	input := make(chan int)
	output := make(chan int)

	// Simple processor
	go func() {
		defer close(output)
		for val := range input {
			output <- val * 2
		}
		fmt.Println("✅ Processor closed output after input closed")
	}()

	// Send data and close input
	go func() {
		input <- 5
		input <- 10
		close(input)
		fmt.Println("✅ Producer closed input")
	}()

	// Collect results with timeout
	results := []int{}
	timeout := time.After(100 * time.Millisecond)

	for {
		select {
		case val, ok := <-output:
			if !ok {
				fmt.Printf("✅ Pipeline completed, results: %v\n", results)
				return
			}
			results = append(results, val)
		case <-timeout:
			fmt.Println("❌ Pipeline test timed out!")
			return
		}
	}
}

func main() {
	fmt.Println("🔒 Go Concurrency: Channel Closing")

	// Build understanding of resource management
	// demonstrateBasicClosing()
	// demonstrateCommaOkIdiom()
	// demonstrateRangeWithClosure()
	// demonstrateSenderCloses()
	// demonstrateSafeClosing()
	// demonstrateDeferClosing()
	demonstrateClosurePropagation()
	// demonstrateClosureTesting()

	fmt.Println("\n✅ Channel closing fundamentals complete!")
	fmt.Println("Next: Learn about error handling in concurrent systems")
}
