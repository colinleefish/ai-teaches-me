package main

import (
	"fmt"
	"runtime"
	"time"
)

func politeTalker(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s: Message %d\n", name, i+1)
		runtime.Gosched() // "I'll let others talk now"
		time.Sleep(100 * time.Millisecond)
	}
}

func rudeTalker(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s: Message %d\n", name, i+1)
		// No Gosched() - keeps talking without giving others a turn
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== Polite Talkers (with Gosched) ===")
	go politeTalker("Alice")
	go politeTalker("Bob")
	time.Sleep(1 * time.Second)

	fmt.Println("\n=== Rude Talkers (without Gosched) ===")
	go rudeTalker("Charlie")
	go rudeTalker("Diana")
	time.Sleep(1 * time.Second)
}
