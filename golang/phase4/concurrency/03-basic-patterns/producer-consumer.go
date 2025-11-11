package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// 🎯 PRODUCER-CONSUMER PATTERN
// One of the most fundamental concurrency patterns in Go
// Producers generate data, consumers process it, channels coordinate the flow
// This pattern decouples data generation from data processing
// Perfect for scenarios where production and consumption rates differ

// TUTOR: The producer-consumer pattern separates data creation from data processing.
// Producers focus solely on generating work items or data.
// Consumers focus solely on processing work items or data.
// Channels act as the buffer and coordination mechanism between them.
// This separation enables independent scaling of producers and consumers.
// TODO: Demonstrate basic single producer, single consumer pattern
func demonstrateBasicProducerConsumer() {
	fmt.Println("=== Basic Producer-Consumer ===")

	// TODO: Create a channel for work items
	// TODO: Start a producer goroutine that generates work
	// TODO: Start a consumer goroutine that processes work
	// TODO: Show how channel acts as buffer between them
	// TODO: Demonstrate proper shutdown coordination

	ch := make(chan int, 30)

	// producer
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- rand.Intn(100)
		}
	}()

	// consumer, sum all the results and return the sum.
	wg.Add(1)
	go func() {
		sum := 0
		defer wg.Done()
		for i := range ch {
			sum += i
		}
		fmt.Println("sum:", sum)
	}()

	wg.Wait()
}

// TUTOR: Multiple producers can feed into the same channel simultaneously.
// Each producer operates independently, generating work at its own pace.
// The channel naturally serializes all incoming work items.
// Consumers don't need to know how many producers exist.
// This pattern scales production capacity easily.
// TODO: Demonstrate multiple producers feeding one consumer
func demonstrateMultipleProducers() {
	fmt.Println("\n=== Multiple Producers, Single Consumer ===")

	// TODO: Create shared work channel
	// TODO: Launch multiple producer goroutines
	// TODO: Each producer generates different types of work
	// TODO: Single consumer processes all work types
	// TODO: Show how work items are naturally interleaved
	// TODO: Demonstrate proper producer coordination and shutdown

	ch := make(chan int, 30)
	consumerWg := sync.WaitGroup{}
	producerWg := sync.WaitGroup{}

	consumerWg.Add(1)

	go func() {
		defer consumerWg.Done()
		sum := 0
		for i := range ch {
			sum += i
		}
		fmt.Println("sum:", sum)
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		producerWg.Add(1)
		go func() {
			defer producerWg.Done()
			for i := 0; i < 10; i++ {
				ch <- rand.Intn(100)
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}()
	}

	go func() {
		producerWg.Wait()
		close(ch)
	}()

	consumerWg.Wait()

}

// TUTOR: Multiple consumers can process from the same channel simultaneously.
// Work items are automatically distributed among available consumers.
// Each consumer operates independently, processing at its own pace.
// This pattern provides natural load balancing across consumers.
// Adding more consumers increases processing capacity.
// TODO: Demonstrate single producer feeding multiple consumers
func demonstrateMultipleConsumers() {
	fmt.Println("\n=== Single Producer, Multiple Consumers ===")

	// TODO: Create work channel for job distribution
	// TODO: Start single producer generating consistent work
	// TODO: Launch multiple consumer goroutines
	// TODO: Show how work is automatically distributed
	// TODO: Demonstrate consumer load balancing
	// TODO: Show proper consumer termination when work ends

	start := time.Now()
	numConsumers := runtime.NumCPU() // you can change this to tweak the time taken to consume the data

	ch := make(chan int, 30)
	consumerWg := sync.WaitGroup{}
	producerWg := sync.WaitGroup{}

	consumer := func(id int) {
		defer consumerWg.Done()
		rate := 100
		for i := range ch {
			time.Sleep(time.Duration(rate) * time.Millisecond)
			fmt.Println("consumer", id, "received:", i)
		}
	}

	for i := 0; i < numConsumers; i++ {
		consumerWg.Add(1)
		go consumer(i)
	}

	producerWg.Add(1)
	go func() {
		defer producerWg.Done()
		for i := 0; i < 100; i++ {
			ch <- rand.Intn(100)
		}
		close(ch)
	}()

	producerWg.Wait()
	consumerWg.Wait()
	fmt.Println("time taken:", time.Since(start))
}

// TUTOR: Fan-out/fan-in combines multiple producers with multiple consumers.
// This pattern provides maximum concurrency and throughput.
// Producers and consumers scale independently based on needs.
// Natural load balancing occurs across all participants.
// Channel buffers can smooth out rate differences.
// TODO: Demonstrate multiple producers with multiple consumers
func demonstrateFanOutFanIn() {
	fmt.Println("\n=== Fan-Out/Fan-In: Multiple Producers & Consumers ===")
	fmt.Println("we touch this topic later")
	// TODO: Create channels for work distribution
	// TODO: Launch multiple producers with different work types
	// TODO: Launch multiple consumers with different processing speeds
	// TODO: Show natural load balancing across consumers
	// TODO: Demonstrate rate smoothing with buffered channels
	// TODO: Show proper coordination of complex shutdown
}

// TUTOR: Buffered channels act as queues between producers and consumers.
// Buffer size affects system behavior and performance characteristics.
// Large buffers smooth rate differences but increase memory usage.
// Small buffers provide backpressure but may block producers.
// Buffer sizing is a key design decision in producer-consumer systems.
// TODO: Demonstrate buffer effects on producer-consumer dynamics
func demonstrateBufferedProducerConsumer() {
	fmt.Println("\n=== Buffered Channels in Producer-Consumer ===")

	// TODO: Compare unbuffered vs buffered channel behavior
	// TODO: Show how buffer size affects producer blocking
	// TODO: Demonstrate backpressure with small buffers
	// TODO: Show memory usage implications of large buffers
	// TODO: Illustrate optimal buffer sizing strategies

	// TODO: Use select with buffered channels for non-blocking sends
	// TODO: Show how to test buffer availability before operations
	// TODO: Demonstrate graceful degradation when buffers are full
	// TODO: Create responsive systems that don't block on channel ops

	bufferedChannel := make(chan int, 10)

	var counter int32

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		// we send 1000 values to the buffered channel before stopping
		for atomic.LoadInt32(&counter) < 70 {
			select {
			case bufferedChannel <- rand.Intn(100):
				atomic.AddInt32(&counter, 1)
			default:
				fmt.Println("bufferedChannel is full, waiting for 1 second")
				time.Sleep(1 * time.Second)
			}
		}
		close(bufferedChannel)
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case v, ok := <-bufferedChannel:
				if !ok {
					return
				}
				fmt.Println("received: ", v, "counter: ", atomic.LoadInt32(&counter))
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	wg.Wait()

	fmt.Println("All done")

}

// TUTOR: Graceful shutdown in producer-consumer requires coordination.
// Producers must signal when no more work will be generated.
// Consumers must process remaining work before terminating.
// Channel closing is the primary shutdown signaling mechanism.
// WaitGroups ensure all goroutines complete before program exits.
// TODO: Demonstrate proper shutdown coordination patterns
func demonstrateGracefulShutdown() {
	fmt.Println("\n=== Graceful Shutdown Patterns ===")
	fmt.Println("this is implemented earlier")
	// TODO: Show producer signaling completion by closing channels
	// TODO: Demonstrate consumer handling of closed channels
	// TODO: Show WaitGroup coordination for clean shutdown
	// TODO: Illustrate timeout-based shutdown for unresponsive components
	// TODO: Show cleanup of resources during shutdown

}

// TUTOR: Producer-consumer with work priority requires multiple channels.
// High-priority work gets processed before normal work.
// Multiple channels with select statements enable priority handling.
// This pattern ensures critical work isn't delayed by bulk processing.
// Priority queues can be implemented using multiple channels.
// TODO: Demonstrate priority-based work processing
func demonstratePriorityWork() {
	fmt.Println("\n=== Priority Work Processing ===")
	fmt.Println("we touch this topic later")
	// TODO: Create separate channels for different priority levels
	// TODO: Show producers categorizing work by priority
	// TODO: Demonstrate consumers processing high-priority work first
	// TODO: Show select statement priority handling
	// TODO: Illustrate backpressure management across priority levels

}

// TUTOR: Error handling in producer-consumer requires error channels.
// Producers can encounter errors during work generation.
// Consumers can fail while processing work items.
// Error channels separate error handling from normal work flow.
// Proper error handling prevents silent failures in concurrent systems.
// TODO: Demonstrate error handling patterns in producer-consumer
func demonstrateErrorHandling() {
	fmt.Println("\n=== Error Handling in Producer-Consumer ===")

	// TODO: Create separate channels for work and errors
	// TODO: Show producers handling generation errors
	// TODO: Demonstrate consumers handling processing errors
	// TODO: Show error aggregation and reporting
	// TODO: Illustrate recovery and retry mechanisms
}

// TUTOR: Monitoring producer-consumer systems requires tracking key metrics.
// Monitor work generation rates, processing rates, and queue depths.
// Track error rates and processing latencies for health assessment.
// Use goroutine counts to monitor resource usage.
// Metrics help identify bottlenecks and scaling needs.
// TODO: Demonstrate monitoring and metrics collection
func demonstrateProducerConsumerMonitoring() {
	fmt.Println("\n=== Producer-Consumer Monitoring ===")

	// TODO: Track work generation and consumption rates
	// TODO: Monitor channel buffer usage and depth
	// TODO: Collect processing latency metrics
	// TODO: Show error rate tracking and alerting
	// TODO: Demonstrate bottleneck identification techniques

	// my take:
	// 1. track gen and cons rate by channel size
	// 2. use runtime.NumGoroutine() to track the number of goroutines
	// 3. use dedicate errCh to collect emitted errors.
}

func main() {
	fmt.Println("🏭 Producer-Consumer Patterns - The Foundation of Concurrent Data Flow 🏭")
	fmt.Println("Learn to build scalable data processing pipelines")

	// TODO: Implement each demonstration function
	// Start with basic patterns and build understanding

	// demonstrateBasicProducerConsumer()
	// demonstrateMultipleProducers()
	// demonstrateMultipleConsumers()
	// demonstrateFanOutFanIn()
	// demonstrateBufferedProducerConsumer()
	// demonstrateGracefulShutdown()
	// demonstratePriorityWork()
	// demonstrateErrorHandling()
	// demonstrateProducerConsumerMonitoring()
}
