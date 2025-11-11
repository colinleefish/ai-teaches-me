package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TUTOR: Directional channels restrict operations to either send-only or receive-only.
// This creates type safety and clear API contracts for concurrent functions.
// Send-only channels use 'chan<- type' syntax, receive-only use '<-chan type'.
// Directional channels prevent misuse and make intent explicit.
// Function signatures with directional channels document data flow direction.
// TODO: Demonstrate basic directional channel types
func demonstrateDirectionalTypes() {
	fmt.Println("=== Directional Channel Types ===")

	// TODO: Create a bidirectional channel
	// TODO: Convert to send-only and receive-only channels
	// TODO: Show that directional channels prevent wrong operations
	// TODO: Demonstrate type safety with directional channels

	twoWayChannel := make(chan int)

	var wg sync.WaitGroup

	go func() {
		twoWayChannel <- 42 // this gets blocked until the channel is read
	}()

	fmt.Println(<-twoWayChannel) // this gets blocked until the channel is written

	receiveAndPrint := func(ch <-chan int) {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}

	sendToChannel := func(ch chan<- int) {
		defer wg.Done()
		ticker := time.NewTicker(1 * time.Second)
		defer close(ch)
		defer ticker.Stop()
		for i := 0; i < 5; i++ {
			select {
			case <-ticker.C:
				ch <- rand.Intn(100)
			}
		}
	}

	wg.Add(2)

	go receiveAndPrint(twoWayChannel)

	go sendToChannel(twoWayChannel)

	wg.Wait()
}

// TUTOR: Send-only channels can only send values, not receive them.
// They're used in functions that produce data for other goroutines.
// Attempting to receive from a send-only channel is a compile error.
// Send-only channels make producer functions' intent clear and safe.
// Use send-only parameters to prevent functions from accidentally reading.
// TODO: Demonstrate send-only channel usage patterns
func demonstrateSendOnlyChannels() {
	fmt.Println("\n=== Send-Only Channels ===")

	// TODO: Create a function that takes a send-only channel parameter
	// TODO: Show how to send values through send-only channels
	// TODO: Demonstrate compile-time protection against receiving
	// TODO: Show typical producer function patterns
}

// TUTOR: Receive-only channels can only receive values, not send them.
// They're used in functions that consume data from other goroutines.
// Attempting to send to a receive-only channel is a compile error.
// Receive-only channels make consumer functions' intent clear and safe.
// Use receive-only parameters to prevent functions from accidentally writing.
// TODO: Demonstrate receive-only channel usage patterns
func demonstrateReceiveOnlyChannels() {
	fmt.Println("\n=== Receive-Only Channels ===")

	// TODO: Create a function that takes a receive-only channel parameter
	// TODO: Show how to receive values from receive-only channels
	// TODO: Demonstrate compile-time protection against sending
	// TODO: Show typical consumer function patterns
}

// TUTOR: Bidirectional channels can be implicitly converted to directional ones.
// You can pass a 'chan T' where 'chan<- T' or '<-chan T' is expected.
// This conversion is automatic and safe - no explicit casting needed.
// The reverse conversion (directional to bidirectional) is not allowed.
// This enables flexible function design with strict type safety.
// TODO: Demonstrate automatic channel type conversion
func demonstrateChannelConversion() {
	fmt.Println("\n=== Channel Type Conversion ===")

	// TODO: Create a bidirectional channel
	// TODO: Pass it to functions expecting directional channels
	// TODO: Show that conversion happens automatically
	// TODO: Demonstrate that reverse conversion is not allowed
}

// TUTOR: Directional channels enable clear API design for concurrent functions.
// Producers take send-only channels, consumers take receive-only channels.
// This documents data flow and prevents API misuse at compile time.
// Complex systems benefit from explicit data flow direction specification.
// Good channel APIs are self-documenting through their type signatures.
// TODO: Demonstrate API design with directional channels
func demonstrateAPIDesign() {
	fmt.Println("\n=== API Design with Directional Channels ===")

	// TODO: Create producer function with send-only parameter
	// TODO: Create consumer function with receive-only parameter
	// TODO: Create coordinator function that connects them
	// TODO: Show how type signatures document data flow
}

// TUTOR: Channel closing behavior differs between directional channels.
// Only send-only channels can be closed (close() requires send permission).
// Receive-only channels can detect closure but cannot initiate it.
// This enforces the principle that senders control channel lifecycle.
// Proper closing patterns prevent panics and resource leaks.
// TODO: Demonstrate closing behavior with directional channels
func demonstrateClosingBehavior() {
	fmt.Println("\n=== Closing Behavior ===")

	// TODO: Show that only send-only channels can be closed
	// TODO: Demonstrate closure detection on receive-only channels
	// TODO: Show proper sender-closes pattern
	// TODO: Illustrate why receivers cannot close channels
}

// TUTOR: Range loops work with receive-only channels for clean consumption.
// The range loop automatically detects channel closure and exits.
// This creates clean, readable consumer code without explicit closure checks.
// Range loops with directional channels express intent clearly.
// Combining range with receive-only channels is a common pattern.
// TODO: Demonstrate range loops with directional channels
func demonstrateRangeWithDirectional() {
	fmt.Println("\n=== Range Loops with Directional Channels ===")

	// TODO: Create a producer that sends values and closes channel
	// TODO: Create a consumer that uses range on receive-only channel
	// TODO: Show clean consumption without explicit closure checks
	// TODO: Demonstrate typical producer-consumer patterns
}

// TUTOR: Complex systems can chain directional channels for data pipelines.
// Each stage takes input and output channels with appropriate directions.
// This creates type-safe data processing pipelines with clear interfaces.
// Pipeline stages are composable and can be tested independently.
// Directional channels make pipeline architecture explicit and safe.
// TODO: Demonstrate pipeline design with directional channels
func demonstratePipelineDesign() {
	fmt.Println("\n=== Pipeline Design ===")

	// TODO: Create multiple pipeline stages with directional channels
	// TODO: Show how stages connect through type-safe interfaces
	// TODO: Demonstrate composable pipeline architecture
	// TODO: Show testing benefits of clear channel directions

	double := func(in <-chan int, out chan<- int) {
		defer close(out)
		for v := range in {
			out <- v * 2
		}
	}

	square := func(in <-chan int, out chan<- int) {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}
	step1 := make(chan int)
	step2 := make(chan int)
	step3 := make(chan int)

	go double(step1, step2)

	go square(step2, step3)

	go func() {
		for i := 0; i < 5; i++ {
			step1 <- i
		}
		close(step1)
	}()

	for v := range step3 {
		fmt.Println(v)
	}

	fmt.Println("Done")
}

func main() {
	fmt.Println("📡 Go Concurrency: Directional Channels")

	// Build understanding of channel type safety
	// demonstrateDirectionalTypes()
	// demonstrateSendOnlyChannels()
	// demonstrateReceiveOnlyChannels()
	// demonstrateChannelConversion()
	// demonstrateAPIDesign()
	// demonstrateClosingBehavior()
	// demonstrateRangeWithDirectional()
	demonstratePipelineDesign()

	fmt.Println("\n✅ Directional channel fundamentals complete!")
	fmt.Println("Next: Learn about buffered channels for flow control")
}
