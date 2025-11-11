// Exercise 3: Go Commands Mastery
// Learn to use go run, go build, gofmt, and other essential Go commands

package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("=== Go Commands Demonstration ===")

	// TODO: Complete these functions to demonstrate Go command usage
	demonstrateGoRun()
	demonstrateGoBuild()
	demonstrateGoFormat()
	demonstrateGoModules()
	demonstratePerformance()

	fmt.Println("\n🛠️ Go commands mastery complete!")
}

// TODO: Implement demonstration of 'go run' command
func demonstrateGoRun() {
	fmt.Println("\n1. Go Run Command:")
	// Your code here:
	// - Explain what 'go run' does
	// - Show current execution method
	// - Demonstrate temporary compilation

	// Example:
	// fmt.Println("This program was executed using: go run commands.go")
	// fmt.Println("'go run' compiles and runs the program in one step")
	// fmt.Println("No permanent executable file is created")

	fmt.Printf("This program was executed using: %s\n", os.Args[0])
	fmt.Printf("'go run' compiles and runs the program in one step\n")
	fmt.Printf("No permanent executable file is created\n")
}

// TODO: Implement demonstration of 'go build' command
func demonstrateGoBuild() {
	fmt.Println("\n2. Go Build Command:")
	// Your code here:
	// - Explain what 'go build' does
	// - Show how to check if an executable exists
	// - Demonstrate the difference from 'go run'

	// Instructions for testing:
	// 1. Run: go build commands.go
	// 2. This creates an executable file
	// 3. Run: ./commands (on Unix/Mac) or commands.exe (on Windows)

	// Example:
	fmt.Println("To build this program: go build commands.go")
	fmt.Println("This creates an executable file you can run directly")
	fmt.Println("The executable is platform-specific")

	// Check if we're running from a built executable
	if len(os.Args) > 0 && os.Args[0] == "./commands" {
		fmt.Println("🎉 You're running from a built executable!")
	} else {
		fmt.Println("💡 Try: go build commands.go && ./commands")
	}
}

// TODO: Implement demonstration of 'gofmt' command
func demonstrateGoFormat() {
	fmt.Println("\n3. Go Format Command:")
	// Your code here:
	// - Explain what 'gofmt' does
	// - Show examples of formatting rules
	// - Demonstrate the importance of consistent formatting

	// Example:
	// fmt.Println("gofmt automatically formats Go code")
	// fmt.Println("Usage: gofmt -w filename.go")
	// fmt.Println("All Go code should be formatted with gofmt")

	// Show some intentionally poorly formatted code (in comments)
	// Then show how it should look after gofmt

	// Before gofmt:
	// var x    int=10
	// var y       string     ="hello"

	// After gofmt:
	// var x int = 10
	// var y string = "hello"

	// Create some variables to demonstrate proper formatting
	var number int = 42
	var message string = "Go formatting is important"
	fmt.Printf("Properly formatted: number=%d, message=%s\n", number, message)
}

// TODO: Implement demonstration of Go modules
func demonstrateGoModules() {
	fmt.Println("\n4. Go Modules:")
	// Your code here:
	// - Explain what Go modules are
	// - Show how to initialize a module
	// - Demonstrate go.mod file importance

	// Example:
	// fmt.Println("Go modules manage dependencies")
	// fmt.Println("Initialize with: go mod init module-name")
	// fmt.Println("Dependencies are tracked in go.mod file")

	// Check if we're in a Go module
	if _, err := os.Stat("go.mod"); err == nil {
		fmt.Println("✅ Running in a Go module")
	} else {
		fmt.Println("❌ Not in a Go module. Run: go mod init phase1-foundation")
	}
}

// TODO: Implement performance demonstration
func demonstratePerformance() {
	fmt.Println("\n5. Performance Comparison:")
	// Your code here:
	// - Compare Go's compilation speed to other languages
	// - Demonstrate Go's execution speed
	// - Show memory usage basics

	// Example:
	fmt.Println("Go compiles very fast compared to other compiled languages")
	fmt.Printf("Current Go version: %s\n", runtime.Version())
	fmt.Printf("Running on: %s/%s\n", runtime.GOOS, runtime.GOARCH)

	// Demonstrate a simple performance test
	start := time.Now()

	// Simple computational task
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}

	elapsed := time.Since(start)
	fmt.Printf("Calculated sum of 1M numbers in: %v\n", elapsed)
	fmt.Printf("Final sum: %d\n", sum)
}

// 🎯 LEARNING GOALS:
// 1. Master essential Go command-line tools
// 2. Understand compilation vs interpretation
// 3. Practice Go code formatting standards
// 4. Learn about Go modules and dependency management
// 5. Appreciate Go's performance characteristics

// 🐍 PYTHON COMPARISON:
// Python: python script.py (interpreted)
// Go: go run main.go (compiled then executed)
//
// Python: No built-in formatter (use black, yapf)
// Go: gofmt built-in and standardized
//
// Python: pip install package
// Go: go get package (with modules)
//
// Python: requirements.txt
// Go: go.mod

// 🚀 COMMANDS TO PRACTICE:
// 1. go run commands.go          - Run without building
// 2. go build commands.go        - Build executable
// 3. ./commands                  - Run the executable
// 4. gofmt -w commands.go        - Format the code
// 5. go mod init module-name     - Initialize module
// 6. go version                  - Check Go version
// 7. go env                      - Show Go environment

// 🔧 EXERCISE INSTRUCTIONS:
// 1. Complete all TODO functions
// 2. Test with 'go run commands.go'
// 3. Build with 'go build commands.go'
// 4. Run the executable './commands'
// 5. Format with 'gofmt -w commands.go'
// 6. Compare the experience with Python development
