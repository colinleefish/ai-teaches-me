// Exercise 4: Package System Understanding
// Learn how Go organizes code into packages and modules

package main

import (
	"fmt"
	"math"
	"math/rand"
	"mypackage"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Go Package System ===")

	// TODO: Complete these functions to demonstrate Go package usage
	demonstrateStandardPackages()
	demonstrateImportStatements()
	demonstratePackageNaming()
	demonstrateModuleSystem()
	createSimplePackage()

	fmt.Println("\n📦 Package system mastery complete!")
}

// TODO: Implement demonstration of standard library packages
func demonstrateStandardPackages() {
	fmt.Println("\n1. Standard Library Packages:")
	// Your code here:
	// - Use fmt package for formatted output
	// - Use math package for mathematical operations
	// - Use strings package for string manipulation
	// - Use time package for time operations
	// - Use strconv package for string conversions

	// Example:
	// fmt.Println("Using fmt package for formatted output")
	// fmt.Printf("Square root of 16: %.2f\n", math.Sqrt(16))
	// fmt.Printf("Uppercase: %s\n", strings.ToUpper("hello"))
	// fmt.Printf("Current time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	// fmt.Printf("String to int: %s\n", strconv.Itoa(42))

	// Try these functions:
	// - math.Pow(2, 3)
	// - strings.Contains("hello", "ell")
	// - time.Since(time.Now())
	result, err := strconv.ParseInt("123", 10, 16)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("String to int: %d\n", result)
	}
	fmt.Printf("Current time: %s\n", time.Now().Format("2006-01-02Y15:04"))
	fmt.Printf("Square root of 16: %.2f\n", math.Sqrt(16))
	fmt.Printf("Pi: %.2f\n", math.Pi)

	for _, v := range []int{65, 76, 82} {
		fmt.Printf("%s\n", strconv.AppendInt(nil, int64(v), 10))
	}

	mypackage.PublicFunction()

	mypackage.privateFunction()

}

// TODO: Implement demonstration of import statements
func demonstrateImportStatements() {
	fmt.Println("\n2. Import Statement Variations:")
	// Your code here:
	// - Show different ways to import packages
	// - Demonstrate package aliases
	// - Explain dot imports (though not recommended)

	// Examples of import styles:
	// import "fmt"                    // Standard import
	// import f "fmt"                  // Aliased import
	// import . "fmt"                  // Dot import (not recommended)
	// import _ "some/package"          // Blank import (for side effects)

	// For this exercise, we'll use the already imported packages
	fmt.Println("Standard import: fmt.Println()")

	// Show how to use imported packages
	randomNum := rand.Intn(100)
	fmt.Printf("Random number (0-99): %d\n", randomNum)

	// String manipulation
	text := "Go Package System"
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Lowercase: %s\n", strings.ToLower(text))
	fmt.Printf("Word count: %d\n", len(strings.Fields(text)))
}

// TODO: Implement demonstration of package naming conventions
func demonstratePackageNaming() {
	fmt.Println("\n3. Package Naming Conventions:")
	// Your code here:
	// - Explain Go package naming rules
	// - Show examples of good package names
	// - Demonstrate how package names relate to imports

	// Package naming rules:
	// 1. Use lowercase letters only
	// 2. No underscores or mixed caps
	// 3. Short, concise names
	// 4. Avoid stuttering (e.g., don't name a package "util" if it's in "myapp/util")

	// Examples:
	fmt.Println("Good package names:")
	fmt.Println("- fmt (format)")
	fmt.Println("- http (HTTP protocol)")
	fmt.Println("- json (JSON encoding)")
	fmt.Println("- time (time operations)")
	fmt.Println("- strconv (string conversions)")

	fmt.Println("\nPackage import paths:")
	fmt.Println("- Standard library: \"fmt\", \"math\", \"strings\"")
	fmt.Println("- Subpackages: \"math/rand\", \"net/http\"")
	fmt.Println("- External: \"github.com/user/repo\"")
}

// TODO: Implement demonstration of module system
func demonstrateModuleSystem() {
	fmt.Println("\n4. Go Module System:")
	// Your code here:
	// - Explain what Go modules are
	// - Show how modules organize packages
	// - Demonstrate go.mod file structure

	// Example:
	// fmt.Println("Go modules manage dependencies and versioning")
	// fmt.Println("A module is a collection of packages")
	// fmt.Println("go.mod file defines the module")

	// Show module information
	fmt.Println("Module concepts:")
	fmt.Println("- Module: Collection of related packages")
	fmt.Println("- go.mod: Module definition file")
	fmt.Println("- go.sum: Dependency checksums")
	fmt.Println("- Semantic versioning: v1.2.3")

	// Example go.mod structure:
	fmt.Println("\nExample go.mod:")
	fmt.Println("module github.com/user/myproject")
	fmt.Println("")
	fmt.Println("go 1.21")
	fmt.Println("")
	fmt.Println("require (")
	fmt.Println("    github.com/gorilla/mux v1.8.0")
	fmt.Println(")")
}

// TODO: Implement a simple package demonstration
func createSimplePackage() {
	fmt.Println("\n5. Creating Your Own Package:")
	// Your code here:
	// - Explain how to create a new package
	// - Show package declaration
	// - Demonstrate exported vs unexported names

	// Package creation steps:
	// 1. Create a new directory
	// 2. Add package declaration at top of .go files
	// 3. Implement functions and types
	// 4. Export names by capitalizing first letter

	fmt.Println("Package creation steps:")
	fmt.Println("1. Create directory: mkdir mypackage")
	fmt.Println("2. Add package declaration: package mypackage")
	fmt.Println("3. Implement functions:")
	fmt.Println("   - Exported: func PublicFunction() {}")
	fmt.Println("   - Unexported: func privateFunction() {}")
	fmt.Println("4. Import in main: import \"./mypackage\"")

	// Demonstrate exported vs unexported names
	fmt.Println("\nExported vs Unexported:")
	fmt.Println("- Exported: starts with capital letter (PublicFunc)")
	fmt.Println("- Unexported: starts with lowercase letter (privateFunc)")
	fmt.Println("- Only exported names can be used by other packages")

	// Example of using a "package" (we'll simulate this)
	demonstratePackageUsage()
}

// Helper function to simulate package usage
func demonstratePackageUsage() {
	// This simulates what you might do in a separate package
	fmt.Println("\nSimulated package usage:")

	// "Exported" functions (these would be in another package)
	result := calculateArea(5.0)
	fmt.Printf("Area of circle with radius 5: %.2f\n", result)

	message := formatMessage("Hello", "World")
	fmt.Printf("Formatted message: %s\n", message)
}

// These functions simulate what you might put in a separate package
func calculateArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func formatMessage(greeting, name string) string {
	return fmt.Sprintf("%s, %s!", greeting, name)
}

// 🎯 LEARNING GOALS:
// 1. Understand Go's package system
// 2. Learn to use standard library packages
// 3. Master import statements and conventions
// 4. Understand Go modules
// 5. Learn to create your own packages

// 🐍 PYTHON COMPARISON:
// Python: import math; math.sqrt(16)
// Go: import "math"; math.Sqrt(16)
//
// Python: from math import sqrt; sqrt(16)
// Go: Not directly supported (use aliases instead)
//
// Python: import math as m; m.sqrt(16)
// Go: import m "math"; m.Sqrt(16)
//
// Python: __init__.py defines packages
// Go: package declaration in each file
//
// Python: requirements.txt, setup.py
// Go: go.mod, go.sum

// 🚀 NEXT STEPS:
// 1. Complete all TODO functions
// 2. Experiment with different standard library packages
// 3. Try creating a simple package in a subdirectory
// 4. Practice importing and using your own packages
// 5. Read about Go's module system in detail

// 🔧 USEFUL STANDARD LIBRARY PACKAGES:
// - fmt: Formatted I/O
// - strings: String manipulation
// - strconv: String conversions
// - math: Mathematical functions
// - time: Time and date operations
// - os: Operating system interface
// - io: I/O primitives
// - net/http: HTTP client and server
// - encoding/json: JSON encoding/decoding
// - regexp: Regular expressions
