// Exercise 1: Installation & Setup Verification
// This file helps you verify that Go is properly installed and configured.

package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("=== Go Installation Verification ===")

	// Complete these functions to demonstrate Go installation
	checkGoVersion()
	checkGoEnvironment()
	demonstrateBasicSyntax()
	exploreWorkspace()

	fmt.Println("\n✅ Setup verification complete!")
}

// TODO: Implement this function to display Go version information
// Hint: Use runtime.Version() and runtime.GOOS, runtime.GOARCH
func checkGoVersion() {
	fmt.Println("\n1. Go Version Information:")
	// Your code here:
	// - Print Go version using runtime.Version()
	// - Print operating system using runtime.GOOS
	// - Print architecture using runtime.GOARCH

	// Example output:
	// Go version: go1.21.0
	// OS: darwin
	// Architecture: amd64

	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
}

// TODO: Implement this function to show Go environment variables
// Hint: Use os.Getenv() to get environment variables
func checkGoEnvironment() {
	fmt.Println("\n2. Go Environment:")
	// Your code here:
	// - Print GOROOT using os.Getenv("GOROOT")
	// - Print GOPATH using os.Getenv("GOPATH")
	// - Print current working directory using os.Getwd()

	// Example output:
	// GOROOT: /usr/local/go
	// GOPATH: /Users/username/go
	// Current directory: /path/to/your/project

	fmt.Printf("GOROOT: %s\n", os.Getenv("GOROOT"))
	fmt.Printf("GOPATH: %s\n", os.Getenv("GOPATH"))
	dir, _ := os.Getwd()
	fmt.Printf("Current directory: %s\n", dir)
}

// TODO: Implement this function to demonstrate basic Go syntax
func demonstrateBasicSyntax() {
	fmt.Println("\n3. Basic Go Syntax Demo:")
	// Your code here:
	// - Declare a variable using var keyword
	// - Declare a variable using short declaration :=
	// - Create a constant
	// - Print all three values

	// Example:
	// var name string = "Go"
	// age := 14  // Go was first released in 2009, so about 14 years old
	// const creator = "Google"
	// fmt.Printf("Language: %s, Age: %d, Creator: %s\n", name, age, creator)

	var name string = "Go"
	age := 14
	const creator = "Google"
	fmt.Printf("Language: %s, Age: %d, Creator: %s\n", name, age, creator)
}

// TODO: Implement this function to explore the workspace structure
func exploreWorkspace() {
	fmt.Println("\n4. Workspace Structure:")
	// Your code here:
	// - List current directory contents
	// - Check if go.mod exists
	// - Print working directory

	// Hint: Use os.ReadDir(".") to read current directory
	// Use os.Stat("go.mod") to check if go.mod exists

	dirContents, _ := os.ReadDir(".")
	fmt.Println("Directory contents:")
	for _, file := range dirContents {
		fmt.Printf("  %s\n", file.Name())
	}

	if _, err := os.Stat("go.mod"); err == nil {
		fmt.Println("go.mod exists")
	} else {
		fmt.Println("go.mod does not exist")
	}
	currentDir, _ := os.Getwd()
	fmt.Println("Current directory: ", currentDir)
}

// 🎯 LEARNING GOALS:
// 1. Verify Go is properly installed on your system
// 2. Understand Go's runtime environment
// 3. Practice basic Go syntax and conventions
// 4. Explore Go module system

// 🐍 PYTHON COMPARISON:
// Python: import sys; print(sys.version)
// Go: import "runtime"; runtime.Version()
//
// Python: os.getcwd()
// Go: os.Getwd()
//
// Python: os.environ.get("PATH")
// Go: os.Getenv("PATH")

// 🚀 WHEN YOU'RE DONE:
// 1. Run: go run setup.go
// 2. Verify all information is displayed correctly
// 3. Compare with Python's sys module functionality
// 4. Move to hello.go for the next exercise
