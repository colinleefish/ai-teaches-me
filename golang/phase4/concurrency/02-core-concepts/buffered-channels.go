package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// TUTOR: Buffered channels can hold multiple values before blocking senders.
// Buffer size is specified at creation: make(chan type, capacity).
// Senders block only when buffer is full, receivers when buffer is empty.
// Buffered channels provide asynchronous communication with flow control.
// Buffer size affects program behavior and performance characteristics.
// TODO: Demonstrate basic buffered channel behavior
func demonstrateBasicBuffering() {
	fmt.Println("=== Basic Buffered Channel Behavior ===")

	// TODO: Create channels with different buffer sizes
	// TODO: Show how buffer size affects blocking behavior
	// TODO: Demonstrate sending multiple values before receiving
	// TODO: Compare buffered vs unbuffered channel behavior

	// buffered example
	bufferedChannel := make(chan int, 2)

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(bufferedChannel)
		defer fmt.Println("bufferedChannel closed")
		bufferedChannel <- 1
		bufferedChannel <- 2
	}()

	go func() {
		defer wg.Done()
		for v := range bufferedChannel {
			fmt.Println(v)
		}
	}()

	wg.Wait()

	fmt.Println("Buffered channel: channel is closed before values are read")

	// unbuffered example
	fmt.Println("\n=== Unbuffered Channel Behavior ===")

	unbufferedChannel := make(chan int)

	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(unbufferedChannel)
		defer fmt.Println("unbufferedChannel closed")
		unbufferedChannel <- 1
		unbufferedChannel <- 2
	}()

	go func() {
		defer wg.Done()
		for v := range unbufferedChannel {
			fmt.Println(v)
		}
	}()

	wg.Wait()

	fmt.Println("Unbuffered channel: channel is closed after values are read")
}

// TUTOR: Buffer capacity determines how many values can be queued.
// cap() function returns the buffer capacity of a channel.
// len() function returns current number of values in buffer.
// Understanding capacity vs length is crucial for flow control.
// Buffer metrics help debug and optimize concurrent programs.
// TODO: Demonstrate buffer capacity and length inspection
func demonstrateBufferMetrics() {
	fmt.Println("\n=== Buffer Capacity and Length ===")

	// TODO: Create buffered channels with different capacities
	// TODO: Use cap() and len() to inspect buffer state
	// TODO: Show how metrics change as values are sent/received
	// TODO: Demonstrate buffer filling and draining patterns

	bufferedChannel := make(chan int, rand.Intn(30)+1)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < cap(bufferedChannel); i++ {
			bufferedChannel <- rand.Intn(100)
		}
		close(bufferedChannel)
	}()

	go func() {
		defer wg.Done()
		count := 0
		for v := range bufferedChannel {
			fmt.Printf("[%d] received %d, buffer length: %d, cap: %d\n", count, v, len(bufferedChannel), cap(bufferedChannel))
			count++
		}
	}()

	wg.Wait()
}

// TUTOR: Buffered channels enable producer-consumer decoupling.
// Producers can run ahead of consumers when buffer has space.
// Consumers can process bursts without blocking producers immediately.
// Buffer size affects system throughput and latency characteristics.
// Proper buffer sizing improves overall system performance.
// TODO: Demonstrate producer-consumer decoupling with buffers
func demonstrateProducerConsumerDecoupling() {
	fmt.Println("\n=== Producer-Consumer Decoupling ===")

	// TODO: Create producer that generates data in bursts
	// TODO: Create consumer that processes data at different rate
	// TODO: Show how buffer accommodates rate differences
	// TODO: Demonstrate system behavior with different buffer sizes

	bufferedChannel := make(chan int, 10000)

	wg := sync.WaitGroup{}

	// this prints the buffer length every 500ms
	wg.Add(1)
	go func() {
		defer wg.Done()
		initial := true
		ticker := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				if initial || len(bufferedChannel) > 0 {
					fmt.Printf("buffer length: %d\n", len(bufferedChannel))
				} else {
					ticker.Stop()
					return
				}
				initial = false
			}
		}
	}()

	// producer
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(100 * time.Millisecond)
		countdown := time.NewTimer(5 * time.Second)
		defer countdown.Stop()
		for {
			select {
			case <-ticker.C:
				bufferedChannel <- rand.Intn(100)
			case <-countdown.C:
				ticker.Stop()
				close(bufferedChannel)
				return
			}
		}
	}()

	wg.Add(1)
	// consumer
	go func() {
		defer wg.Done()
		for v := range bufferedChannel {
			time.Sleep(200 * time.Millisecond)
			_ = v
		}
	}()

	wg.Wait()
}

// TUTOR: Buffer size significantly impacts program performance and behavior.
// Small buffers provide tight coupling and immediate backpressure.
// Large buffers enable higher throughput but increase memory usage.
// Infinite buffers (very large) can lead to memory exhaustion.
// Choosing optimal buffer size requires understanding your workload.
// TODO: Demonstrate buffer size impact on performance
func demonstrateBufferSizeImpact() {
	fmt.Println("\n=== Buffer Size Impact ===")

	// TODO: Test same workload with different buffer sizes
	// TODO: Measure timing and memory usage differences
	// TODO: Show optimal buffer size for specific scenarios
	// TODO: Demonstrate trade-offs between memory and performance

	smallBuffer := make(chan int, 3)
	largeBuffer := make(chan int, 1000)

	pub := func(buffer chan int, wg *sync.WaitGroup, done chan time.Time) {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			buffer <- rand.Intn(100)
			time.Sleep(500 * time.Microsecond) // Producer rate: 2 items/ms
		}
		done <- time.Now() // Report when producer finishes
		close(buffer)
	}

	sub := func(buffer chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		count := 0
		for v := range buffer {
			time.Sleep(1 * time.Millisecond) // Consumer rate: 1 item/ms
			_ = v
			count++
		}
		fmt.Println("Sub received: ", count)
	}

	// Test with small buffer
	producerDone := make(chan time.Time, 1)
	start := time.Now()
	fmt.Println("Pub/Sub with small buffer: ", start.Format(time.RFC3339))
	wgSmall := sync.WaitGroup{}
	wgSmall.Add(2)
	go pub(smallBuffer, &wgSmall, producerDone)
	go sub(smallBuffer, &wgSmall)

	producerFinish := <-producerDone
	wgSmall.Wait()
	totalTime := time.Since(start)

	fmt.Printf("Small buffer - Producer finished in: %v, Total time: %v\n", producerFinish.Sub(start), totalTime)

	// Test with large buffer
	producerDone = make(chan time.Time, 1)
	start = time.Now()
	fmt.Println("Pub/Sub with large buffer: ", start.Format(time.RFC3339))
	wgLarge := sync.WaitGroup{}
	wgLarge.Add(2)
	go pub(largeBuffer, &wgLarge, producerDone)
	go sub(largeBuffer, &wgLarge)

	producerFinish = <-producerDone
	wgLarge.Wait()
	totalTime = time.Since(start)

	fmt.Printf("Large buffer - Producer finished in: %v, Total time: %v\n", producerFinish.Sub(start), totalTime)
}

// TUTOR: Non-blocking operations become possible with select and buffered channels.
// Buffered channels with space don't block on send operations.
// Empty buffered channels still block on receive operations.
// Select with default can test buffer availability non-destructively.
// Non-blocking patterns enable responsive and flexible programs.
// TODO: Demonstrate non-blocking operations with buffered channels
func demonstrateNonBlockingOperations() {
	fmt.Println("\n=== Non-Blocking Operations ===")

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

// TUTOR: Buffered channels can create backpressure for flow control.
// When buffer fills, senders block, creating natural backpressure.
// This prevents fast producers from overwhelming slow consumers.
// Backpressure propagates through pipeline stages automatically.
// Understanding backpressure helps design stable concurrent systems.
// TODO: Demonstrate backpressure mechanisms with buffered channels
func demonstrateBackpressure() {
	fmt.Println("\n=== Backpressure and Flow Control ===")

	// Create pipeline: Producer -> Stage1 -> Stage2 -> Consumer
	ch1 := make(chan int, 2)    // Producer to Stage1
	ch2 := make(chan int, 2)    // Stage1 to Stage2
	ch3 := make(chan string, 2) // Stage2 to Consumer

	wg := sync.WaitGroup{}
	wg.Add(4)

	// Producer: Fast (every 50ms)
	go func() {
		defer wg.Done()
		defer close(ch1)
		for i := 1; i <= 8; i++ {
			fmt.Printf("Producer: sending %d\n", i)
			ch1 <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Stage1: Fast processing (double the number)
	go func() {
		defer wg.Done()
		defer close(ch2)
		for v := range ch1 {
			result := v * 2
			fmt.Printf("Stage1: %d -> %d\n", v, result)
			ch2 <- result
			time.Sleep(20 * time.Millisecond) // Fast
		}
	}()

	// Stage2: SLOW processing (convert to string) - BOTTLENECK!
	go func() {
		defer wg.Done()
		defer close(ch3)
		for v := range ch2 {
			result := fmt.Sprintf("item_%d", v)
			fmt.Printf("Stage2: %d -> %s (slow processing...)\n", v, result)
			time.Sleep(200 * time.Millisecond) // VERY SLOW - creates backpressure!
			ch3 <- result
		}
	}()

	// Consumer: Fast
	go func() {
		defer wg.Done()
		for v := range ch3 {
			fmt.Printf("Consumer: received %s\n", v)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("Notice how Stage2's slowness backs up the entire pipeline!")
}

// TUTOR: Buffered channels can be used for semaphore patterns.
// Buffer capacity limits concurrent operations naturally.
// Sending to full buffer blocks, creating resource limiting.
// This pattern controls concurrent resource access elegantly.
// Semaphore pattern prevents resource exhaustion in bounded systems.
// TODO: Demonstrate semaphore pattern with buffered channels

func demonstrateSemaphorePattern() {
	fmt.Println("\n=== Semaphore Pattern ===")

	// TODO: Create buffered channel to limit concurrent operations
	// TODO: Show how buffer capacity enforces resource limits
	// TODO: Demonstrate controlled parallel processing
	// TODO: Show graceful handling of resource contention

	numWorkers := 3
	semaphore := make(chan struct{}, numWorkers)

	wg := sync.WaitGroup{}

	countTo3 := func(id int) {
		counter := 0
		interval := rand.Intn(1000)
		for counter < 10 {
			fmt.Printf("[%d] counter: %d\n", id, counter)
			counter++
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore <- struct{}{}
			countTo3(i)
			<-semaphore
		}(i)
	}

	wg.Wait()
}

// TUTOR: Buffered channels enable work queuing and batch processing.
// Multiple producers can queue work items for batch processing.
// Consumers can process items individually or in batches.
// Buffer size affects queuing capacity and system responsiveness.
// Work queues decouple work generation from work processing.
// TODO: Demonstrate work queue patterns with buffered channels
func demonstrateWorkQueue() {
	fmt.Println("\n=== Work Queue Patterns ===")

	// TODO: Create work queue with multiple producers
	// TODO: Show batch processing by consumers
	// TODO: Demonstrate queue depth management
	// TODO: Show graceful queue overflow handling

	workQueue := make(chan int, 20) // Queue buffer

	wg := sync.WaitGroup{}

	// Producer: Generates work continuously for 3 seconds
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(workQueue) // Signal no more work

		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		timeout := time.After(3 * time.Second)

		workID := 1
		for {
			select {
			case <-ticker.C:
				fmt.Printf("Producer: queuing work %d\n", workID)
				workQueue <- workID
				workID++
			case <-timeout:
				fmt.Println("Producer: stopping")
				return
			}
		}
	}()

	// 3 Workers competing for work from shared queue
	for workerID := 1; workerID <= runtime.NumCPU(); workerID++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for work := range workQueue {
				fmt.Printf("Worker%d: processing work %d\n", id, work)
				time.Sleep(200 * time.Millisecond) // Simulate work
				fmt.Printf("Worker%d: completed work %d\n", id, work)
			}
		}(workerID)
	}

	wg.Wait()
	fmt.Println("All work completed!")
}

func main() {
	fmt.Println("🪣 Go Concurrency: Buffered Channels")

	// Build understanding of flow control
	// demonstrateBasicBuffering()
	// demonstrateBufferMetrics()
	// demonstrateProducerConsumerDecoupling()
	// demonstrateBufferSizeImpact()
	// demonstrateNonBlockingOperations()
	// demonstrateBackpressure()
	// demonstrateSemaphorePattern()
	demonstrateWorkQueue()

	fmt.Println("\n✅ Buffered channel fundamentals complete!")
	fmt.Println("Next: Learn about proper channel closing patterns")
}
