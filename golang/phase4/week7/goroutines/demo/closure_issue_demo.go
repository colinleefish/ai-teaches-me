package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== PROBLEMATIC: Closure capturing loop variable ===")

	var wg sync.WaitGroup

	// This is the WRONG way - all goroutines will likely print the same number
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)           // Small delay to let loop finish
			fmt.Printf("Goroutine captured i = %d\n", i) // BUG: captures reference to i
		}()
	}
	wg.Wait()

	fmt.Println("\n=== CORRECT: Passing value as parameter ===")

	// This is the RIGHT way - each goroutine gets its own copy
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine received val = %d\n", val) // CORRECT: own copy of value
		}(i)
	}
	wg.Wait()

	fmt.Println("\n=== ALSO CORRECT: Creating local copy ===")

	// Another way to fix it - create a local copy
	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i // Create a new variable in this iteration's scope
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine captured local i = %d\n", i) // CORRECT: captures local copy
		}()
	}
	wg.Wait()
}
