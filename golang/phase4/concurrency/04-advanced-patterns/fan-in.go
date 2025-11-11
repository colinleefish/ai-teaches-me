// Fan-In Pattern: Merging multiple input streams into single output
// Key concepts: Stream merging, fairness, select multiplexing

package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

// Basic fan-in: merge multiple channels into one
func fanIn(inputs ...<-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup

	// Start goroutine for each input channel
	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for val := range ch {
				output <- val
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// Fan-in with select (non-blocking, fair multiplexing)
func fanInSelect(input1, input2 <-chan string) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)
		for {
			select {
			case val, ok := <-input1:
				if !ok {
					input1 = nil
				} else {
					output <- val
				}
			case val, ok := <-input2:
				if !ok {
					input2 = nil
				} else {
					output <- val
				}
			}

			// Both channels closed
			if input1 == nil && input2 == nil {
				return
			}
		}
	}()

	return output
}

// Fan-in with scalable select using reflection
func fanInSelectScalable(inputs ...<-chan string) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)

		// Build reflect.SelectCase slice from input channels
		cases := make([]reflect.SelectCase, len(inputs))
		for i, ch := range inputs {
			cases[i] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			}
		}

		activeChannels := len(inputs)

		for activeChannels > 0 {
			// Use reflect.Select for dynamic select
			chosen, value, ok := reflect.Select(cases)

			if !ok {
				// Channel closed, disable this case
				cases[chosen].Chan = reflect.ValueOf(nil)
				activeChannels--
			} else {
				// Send received value to output
				output <- value.String()
			}
		}
	}()

	return output
}

// Producer that generates messages
func producer(name string, delay time.Duration) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)
		for i := 0; i < 5; i++ {
			time.Sleep(delay)
			output <- fmt.Sprintf("%s: message %d", name, i)
		}
	}()

	return output
}

func main() {
	fmt.Println("=== Fan-In Pattern ===")

	// Create multiple producers
	alice := producer("Alice", 100*time.Millisecond)
	bob := producer("Bob", 150*time.Millisecond)
	charlie := producer("Charlie", 200*time.Millisecond)

	// Example 1: Basic fan-in
	fmt.Println("\n1. Basic Fan-In:")
	merged := fanIn(alice, bob, charlie)

	// Read merged stream
	for i := 0; i < 15; i++ {
		fmt.Println(<-merged)
	}

	// Example 2: Select-based fan-in (more control)
	fmt.Println("\n2. Select-based Fan-In:")
	david := producer("David", 80*time.Millisecond)
	eve := producer("Eve", 120*time.Millisecond)

	selectMerged := fanInSelect(david, eve)

	for msg := range selectMerged {
		fmt.Println(msg)
	}

	// Example 3: Scalable select-based fan-in
	fmt.Println("\n3. Scalable Select-based Fan-In:")
	frank := producer("Frank", 90*time.Millisecond)
	grace := producer("Grace", 130*time.Millisecond)
	henry := producer("Henry", 110*time.Millisecond)
	iris := producer("Iris", 170*time.Millisecond)

	scalableMerged := fanInSelectScalable(frank, grace, henry, iris)

	for msg := range scalableMerged {
		fmt.Println(msg)
	}
}
