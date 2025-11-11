// Exercise 1: Variable Declarations
// Master Go's variable declaration syntax and scoping

package main

import (
	"fmt"
	"reflect"
)

// Package-level variables (can only use var, not :=)
var globalCounter int = 0
var appName string = "Go Learning App"

func main() {
	fmt.Println("=== Variable Declarations ===")

	// TODO: Complete these functions to demonstrate variable declarations
	basicDeclarations()
	shortDeclarations()
	multipleDeclarations()
	variableScope()
	zeroValues()
	typeInference()

	fmt.Println("\n📦 Variable declarations mastery complete!")
}

// TODO: Implement basic variable declarations using var keyword
func basicDeclarations() {
	fmt.Println("\n1. Basic var Declarations:")
	// Your code here:
	// - Declare variables using var keyword
	// - Use explicit types
	// - Initialize with values
	// - Show different syntax variations

	// Example:
	// var name string = "Alice"
	// var age int = 25
	// var height float64 = 5.6
	// var isStudent bool = true

	// Print the variables and their types
	// fmt.Printf("Name: %s (type: %T)\n", name, name)
	// fmt.Printf("Age: %d (type: %T)\n", age, age)
	// fmt.Printf("Height: %.1f (type: %T)\n", height, height)
	// fmt.Printf("Is student: %t (type: %T)\n", isStudent, isStudent)

	var name string = "Alice"
	var age int = 25
	var height float64 = 5.6
	var isStudent bool = true

	fmt.Println(name, age, height, isStudent)
}

// TODO: Implement short variable declarations using :=
func shortDeclarations() {
	fmt.Println("\n2. Short Variable Declarations (:=):")
	// Your code here:
	// - Use := for short declarations
	// - Let Go infer the types
	// - Compare with Python variable assignment

	// Example:
	// name := "Bob"        // string inferred
	// age := 30           // int inferred
	// salary := 50000.50  // float64 inferred
	// isManager := false  // bool inferred

	// Print variables with type information
	// fmt.Printf("Name: %s (type: %T)\n", name, name)
	// fmt.Printf("Age: %d (type: %T)\n", age, age)
	// fmt.Printf("Salary: %.2f (type: %T)\n", salary, salary)
	// fmt.Printf("Is manager: %t (type: %T)\n", isManager, isManager)

	// Python comparison:
	// In Python: name = "Bob"  # type determined at runtime
	// In Go: name := "Bob"     # type determined at compile time

	name := "Bob"
	age := 30
	salary := 50000.50
	isManager := false // that's pretty high salary for a non-manager.

	fmt.Println(name, age, salary, isManager)
	fmt.Printf("Name: %s (type: %T)\n", name, name)
	fmt.Printf("Salary: %f (type: %T)\n", salary, salary)
	fmt.Printf("Is Manager: %t (type: %T)\n", isManager, isManager)
	fmt.Printf("Age: %d (type: %T)\n", age, age)

}

// TODO: Implement multiple variable declarations
func multipleDeclarations() {
	fmt.Println("\n3. Multiple Variable Declarations:")
	// Your code here:
	// - Declare multiple variables at once
	// - Use different declaration styles
	// - Show grouped declarations

	var x, y, z int = 1, 2, 3
	var (
		firstName string  = "John"
		salary    float64 = 50000.50
		isManager bool    = false
	)

	fmt.Printf("x: %d (type: %T)\n", x, x)

	for _, v := range []any{x, y, z, firstName, salary, isManager} {
		fmt.Printf("%v (type: %T)\n", v, v)
	}

	var myAny any

	myAny = "hello"
	fmt.Printf("myAny: %v (type: %T)\n", myAny, myAny)

	myAny = 10
	fmt.Printf("myAny: %v (type: %T)\n", myAny, myAny)

	myAny = true
	fmt.Printf("myAny: %v (type: %T)\n", myAny, myAny)

	myAny = 10.5
	fmt.Printf("myAny: %v (type: %T)\n", myAny, myAny)

	myAny = []any{1, 2, 3}
	fmt.Printf("myAny: %v (type: %T)\n", myAny, myAny)

	myAny = map[string]any{"name": "John", "age": 30}
	fmt.Printf("myAny: %v (type: %T)\n", myAny, myAny)

	myAny = struct {
		name string
		age  int
	}{"John", 30}
	fmt.Printf("myAny: %v (type: %T)\n", myAny, myAny)

}

// TODO: Implement variable scope demonstration
func variableScope() {
	fmt.Println("\n4. Variable Scope:")
	// Your code here:
	// - Show function-level scope
	// - Demonstrate block scope
	// - Access package-level variables
	// - Show variable shadowing

	// Function-level variable
	functionVar := "I'm function scoped"

	// Block scope example
	if true {
		blockVar := "I'm block scoped"
		// Variable shadowing
		functionVar := "I shadow the function variable"

		fmt.Printf("Inside block - blockVar: %s\n", blockVar)
		fmt.Printf("Inside block - functionVar: %s\n", functionVar)
	}

	// blockVar is not accessible here
	fmt.Printf("Outside block - functionVar: %s\n", functionVar)

	// Access package-level variables
	fmt.Printf("Global counter: %d\n", globalCounter)
	fmt.Printf("App name: %s\n", appName)

	// Modify global variable
	globalCounter++
	fmt.Printf("Incremented counter: %d\n", globalCounter)
}

// TODO: Implement zero values demonstration
func zeroValues() {
	fmt.Println("\n5. Zero Values:")
	// Your code here:
	// - Declare variables without initialization
	// - Show zero values for different types
	// - Compare with Python's None/default values

	// Example:
	// var num int
	// var text string
	// var flag bool
	// var decimal float64

	// fmt.Printf("Zero int: %d\n", num)
	// fmt.Printf("Zero string: '%s'\n", text)
	// fmt.Printf("Zero bool: %t\n", flag)
	// fmt.Printf("Zero float64: %f\n", decimal)

	// Python comparison:
	// In Python: variables must be assigned before use
	// In Go: variables have zero values by default

	var num int
	var text string
	var flag bool
	var decimal float64

	fmt.Printf("Zero int: %d\n", num)
	fmt.Printf("Zero string: %s\n", text)
	fmt.Printf("Zero bool: %t\n", flag)
	fmt.Printf("Zero float64: %f\n", decimal)
}

// TODO: Implement type inference demonstration
func typeInference() {
	fmt.Println("\n6. Type Inference:")
	// Your code here:
	// - Show how Go infers types
	// - Use reflect package to show actual types
	// - Compare different literal types

	// Example:
	// a := 42          // int
	// b := 42.0        // float64
	// c := "hello"     // string
	// d := true        // bool
	// e := 'A'         // rune (int32)

	// Use reflection to show types
	// fmt.Printf("a = %v, type: %s\n", a, reflect.TypeOf(a))
	// fmt.Printf("b = %v, type: %s\n", b, reflect.TypeOf(b))
	// fmt.Printf("c = %v, type: %s\n", c, reflect.TypeOf(c))
	// fmt.Printf("d = %v, type: %s\n", d, reflect.TypeOf(d))
	// fmt.Printf("e = %v, type: %s\n", e, reflect.TypeOf(e))

	a := 10
	// b := 10.5
	// c := "hello"
	// d := true
	// e := 'A'

	fmt.Printf("a =%v, type: %T\n, reflect.TypeOf(a) = %s\n", a, a, reflect.TypeOf(a))
}

// 🎯 LEARNING GOALS:
// 1. Master var vs := declaration syntax
// 2. Understand Go's static type system
// 3. Learn about variable scoping rules
// 4. Understand zero values concept
// 5. Practice type inference

// 🐍 PYTHON COMPARISON:
// Python: x = 10                    # Dynamic typing
// Go: var x int = 10 or x := 10    # Static typing
//
// Python: No variable declaration needed
// Go: Variables must be declared before use
//
// Python: Variables can change type
// Go: Variables have fixed types
//
// Python: No concept of zero values
// Go: All variables have zero values

// 🚀 TESTING YOUR WORK:
// 1. Run: go run variables.go
// 2. Observe the output for each section
// 3. Experiment with different variable declarations
// 4. Try causing compile errors to understand type safety

// 🔧 COMMON MISTAKES TO AVOID:
// 1. Using := at package level (only var works there)
// 2. Forgetting that := declares new variables
// 3. Variable shadowing in nested blocks
// 4. Assuming variables can change types (they can't)
// 5. Not understanding zero values vs nil

// 📝 EXERCISES TO TRY:
// 1. Create variables of different types
// 2. Try variable shadowing in nested blocks
// 3. Access global variables from functions
// 4. Use type inference with different literals
// 5. Compare Go's approach with Python's dynamic typing
