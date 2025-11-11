// Exercise 2: Hello World Evolution
// Progress from simple to advanced Hello World programs

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("=== Hello World Evolution ===")

	// TODO: Complete each function to demonstrate different Hello World variations
	basicHello()
	formattedHello()
	greetWithArgs()
	interactiveGreeting()
	multilingualGreeting()

	fmt.Println("\n🎉 Hello World evolution complete!")
}

// TODO: Implement the most basic Hello World
func basicHello() {
	fmt.Println("\n1. Basic Hello World:")
	// Your code here:
	// - Print "Hello, World!" to the console
	// - Use fmt.Println()

	// Expected output: Hello, World!
	fmt.Println("Hello, World!")
}

// TODO: Implement formatted Hello World with variables
func formattedHello() {
	fmt.Println("\n2. Formatted Hello World:")
	// Your code here:
	// - Create variables for name and language
	// - Use fmt.Printf() with formatting verbs
	// - Try different formatting options

	// Example:
	// name := "Go"
	// language := "programming language"
	// fmt.Printf("Hello from %s, the %s!\n", name, language)

	// Expected output: Hello from Go, the programming language!

	name := "Go"
	language := "programming language"
	fmt.Printf("Hello from %s, the %s!\n", name, language)
}

// TODO: Implement greeting with command-line arguments
func greetWithArgs() {
	fmt.Println("\n3. Greeting with Command-line Arguments:")
	// Your code here:
	// - Check if command-line arguments are provided
	// - Greet each person passed as an argument
	// - Handle the case when no arguments are provided

	// Hint: Use os.Args to access command-line arguments
	// Note: os.Args[0] is the program name itself

	// Example usage: go run hello.go Alice Bob Charlie
	// Expected output:
	// Hello, Alice!
	// Hello, Bob!
	// Hello, Charlie!

	if len(os.Args) < 2 {
		fmt.Println("Hello, World!")
		return
	}

	for _, name := range os.Args[1:] {
		fmt.Printf("Hello, %s!\n", name)
	}

}

// TODO: Implement interactive greeting
func interactiveGreeting() {
	fmt.Println("\n4. Interactive Greeting:")
	// Your code here:
	// - Get current time and show appropriate greeting
	// - Use time.Now() to get current time
	// - Show different greetings based on time of day

	// Example logic:
	// hour := time.Now().Hour()
	// if hour < 12 {
	//     fmt.Println("Good morning!")
	// } else if hour < 17 {
	//     fmt.Println("Good afternoon!")
	// } else {
	//     fmt.Println("Good evening!")
	// }

	hour := time.Now().Hour()

	if hour < 12 {
		fmt.Println("Good morning!")
	} else if hour < 17 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}
}

// TODO: Implement multilingual greeting
func multilingualGreeting() {
	fmt.Println("\n5. Multilingual Greeting:")
	// Your code here:
	// - Create a map of languages to greetings
	// - Display greetings in different languages
	// - Use a for loop to iterate through the map

	// Example:
	// greetings := map[string]string{
	//     "English": "Hello",
	//     "Spanish": "Hola",
	//     "French":  "Bonjour",
	//     "German":  "Hallo",
	//     "Japanese": "こんにちは",
	// }

	// Expected output:
	// English: Hello
	// Spanish: Hola
	// French: Bonjour
	// German: Hallo
	// Japanese: こんにちは

	greetings := map[string]string{
		"English":  "Hello",
		"Spanish":  "Hola",
		"French":   "Bonjour",
		"German":   "Hallo",
		"Japanese": "こんにちは",
	}

	for language, greeting := range greetings {
		fmt.Printf("%s: %s\n", language, greeting)
	}
}

// 🎯 LEARNING GOALS:
// 1. Understand basic Go program structure
// 2. Practice fmt package for output formatting
// 3. Learn to handle command-line arguments
// 4. Work with time package
// 5. Use maps and loops for data iteration

// 🐍 PYTHON COMPARISON:
// Python: print("Hello, World!")
// Go: fmt.Println("Hello, World!")
//
// Python: import sys; sys.argv
// Go: import "os"; os.Args
//
// Python: f"Hello, {name}!"
// Go: fmt.Sprintf("Hello, %s!", name)
//
// Python: from datetime import datetime; datetime.now()
// Go: import "time"; time.Now()

// 🚀 TESTING YOUR WORK:
// 1. Run: go run hello.go
// 2. Run: go run hello.go Alice Bob
// 3. Observe different outputs for different times of day
// 4. Verify multilingual greetings display correctly

// 🔧 COMMON MISTAKES TO AVOID:
// - Don't forget to import necessary packages
// - Remember that os.Args[0] is the program name
// - Use fmt.Printf() for formatted output, fmt.Println() for simple output
// - Maps in Go are unordered, so output order may vary
