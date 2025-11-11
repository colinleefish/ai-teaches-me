package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// Exercise 1: Basic Functions with Multiple Returns
func exerciseBasicFunctions() {
	fmt.Println("=== Exercise 1: Basic Functions ===")

	// TODO: Create a function add(a, b int) int that returns the sum
	// YOUR CODE HERE

	// TODO: Create a function divide(a, b float64) (float64, error)
	// Return an error if b is 0, otherwise return the division result
	// YOUR CODE HERE

	// Test the functions
	fmt.Printf("Add(5, 3) = %d\n", add(5, 3))

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Divide(10, 2) = %.2f\n", result)
	}

	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("Expected error: %v\n", err2)
	} else {
		fmt.Printf("Unexpected result: %.2f\n", result2)
	}
}

// TODO: Implement add function here
// YOUR CODE HERE
// func add(a, b int) int {
// 	return a + b
// }

// TODO: Implement divide function here
// YOUR CODE HERE
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Exercise 2: Named Returns
func exerciseNamedReturns() {
	fmt.Println("\n=== Exercise 2: Named Returns ===")

	// TODO: Create a function calculateStats(numbers []int) (min, max, avg int)
	// Use named returns to make the function self-documenting
	// YOUR CODE HERE

	numbers := []int{3, 7, 2, 9, 1, 5}
	min, max, avg := calculateStats(numbers)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Min: %d, Max: %d, Average: %d\n", min, max, avg)

	// TODO: Create a function validateUser(name, email string) (isValid bool, reason string)
	// Return true if both name and email are non-empty, false otherwise
	// YOUR CODE HERE

	// Test validation
	testCases := []struct{ name, email string }{
		{"Alice", "alice@example.com"},
		{"", "bob@example.com"},
		{"Charlie", ""},
		{"", ""},
	}

	for _, test := range testCases {
		valid, reason := validateUser(test.name, test.email)
		fmt.Printf("User(%s, %s) -> Valid: %t, Reason: %s\n",
			test.name, test.email, valid, reason)
	}
}

// TODO: Implement calculateStats function with named returns
// YOUR CODE HERE
func calculateStats(numbers []int) (min, max, avg int) {
	// YOUR CODE HERE
	min, max, sum := numbers[0], numbers[0], 0
	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		sum += n
	}
	return min, max, sum / len(numbers) // min, max, avg
}

// TODO: Implement validateUser function with named returns
// YOUR CODE HERE
func validateUser(name, email string) (isValid bool, reason string) {
	// YOUR CODE HERE
	switch {
	case name == "" && email == "":
		return false, "Name and email are empty"
	case name == "":
		return false, "Name is empty"
	case email == "":
		return false, "Email is empty"
	}
	return true, ""
}

// Exercise 3: Variadic Functions
func exerciseVariadicFunctions() {
	fmt.Println("\n=== Exercise 3: Variadic Functions ===")

	// TODO: Create a function sum(numbers ...int) int
	// Sum all the provided numbers
	// YOUR CODE HERE

	// TODO: Create a function concatenate(separator string, words ...string) string
	// Join all words with the separator
	// YOUR CODE HERE

	// Test variadic functions
	fmt.Printf("Sum() = %d\n", sum())
	fmt.Printf("Sum(1) = %d\n", sum(1))
	fmt.Printf("Sum(1, 2, 3) = %d\n", sum(1, 2, 3))
	fmt.Printf("Sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))

	fmt.Printf("Concatenate('-') = '%s'\n", concatenate("-"))
	fmt.Printf("Concatenate('-', 'hello') = '%s'\n", concatenate("-", "hello"))
	fmt.Printf("Concatenate('-', 'hello', 'world') = '%s'\n", concatenate("-", "hello", "world"))

	// TODO: Create a function findMax(first int, rest ...int) int
	// Find the maximum among all numbers (at least one required)
	// YOUR CODE HERE

	fmt.Printf("FindMax(5) = %d\n", findMax(5))
	fmt.Printf("FindMax(5, 2, 8, 1, 9) = %d\n", findMax(5, 2, 8, 1, 9))
}

// TODO: Implement sum variadic function
// YOUR CODE HERE
func sum(numbers ...int) int {
	// YOUR CODE HERE
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// TODO: Implement concatenate variadic function
// YOUR CODE HERE
func concatenate(separator string, words ...string) string {
	// YOUR CODE HERE
	concat := ""
	for i, w := range words {
		if i > 0 {
			concat += separator
		}
		concat += w
	}
	return concat
}

// TODO: Implement findMax variadic function
// YOUR CODE HERE
func findMax(first int, rest ...int) int {
	// YOUR CODE HERE
	max := first
	for _, n := range rest {
		if n > max {
			max = n
		}
	}
	return max
}

// Exercise 4: Anonymous Functions and Closures
func exerciseAnonymousFunctions() {
	fmt.Println("\n=== Exercise 4: Anonymous Functions and Closures ===")

	// TODO: Create an anonymous function that squares a number
	// YOUR CODE HERE
	square := func(x int) int {
		// YOUR CODE HERE
		return x * x
	}

	fmt.Printf("Square of 5: %d\n", square(5))

	// TODO: Create a closure that captures a counter variable
	// Return a function that increments and returns the counter
	// YOUR CODE HERE
	counter := func() func() int {
		// YOUR CODE HERE
		count := 0
		return func() int {
			// YOUR CODE HERE
			count++
			return count
		}
	}()

	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())
	fmt.Printf("Counter: %d\n", counter())

	// TODO: Create a function that returns a multiplier closure
	createMultiplier := func(factor int) func(int) int {
		// YOUR CODE HERE
		return func(x int) int {
			return x * factor
		}
	}

	double := createMultiplier(2)
	triple := createMultiplier(3)

	fmt.Printf("Double 5: %d\n", double(5))
	fmt.Printf("Triple 5: %d\n", triple(5))

	// TODO: Use an anonymous function to filter even numbers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evens []int
	filter := func(ns []int) []int {
		even := []int{}
		for _, n := range ns {
			if n%2 == 0 {
				even = append(even, n)
			}
		}
		return even
	}
	evens = filter(numbers)

	// Create and immediately call an anonymous function
	// YOUR CODE HERE

	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Evens: %v\n", evens)
}

// Exercise 5: Methods on Types
func exerciseMethods() {
	fmt.Println("\n=== Exercise 5: Methods on Types ===")

	// TODO: Define a Circle struct with Radius field
	// YOUR CODE HERE

	// TODO: Create a method Area() float64 for Circle (value receiver)
	// YOUR CODE HERE

	// TODO: Create a method Scale(factor float64) for Circle (pointer receiver)
	// This should modify the circle's radius
	// YOUR CODE HERE

	// Test Circle methods
	circle := Circle{Radius: 5.0}
	fmt.Printf("Circle radius: %.2f\n", circle.Radius)
	fmt.Printf("Circle area: %.2f\n", circle.Area())

	circle.Scale(2.0)
	fmt.Printf("After scaling by 2:\n")
	fmt.Printf("Circle radius: %.2f\n", circle.Radius)
	fmt.Printf("Circle area: %.2f\n", circle.Area())

	// TODO: Define a Person struct with Name and Age fields
	// YOUR CODE HERE

	// TODO: Create a method String() string for Person (value receiver)
	// This makes Person implement fmt.Stringer interface
	// YOUR CODE HERE

	// TODO: Create a method Birthday() for Person (pointer receiver)
	// This should increment the person's age
	// YOUR CODE HERE

	person := Person{Name: "Alice", Age: 25}
	fmt.Printf("Person: %s\n", person) // Uses String() method
	person.Birthday()
	fmt.Printf("After birthday: %s\n", person)
}

// TODO: Define Circle struct
// YOUR CODE HERE
type Circle struct {
	Radius float64
}

// TODO: Implement Circle methods
// YOUR CODE HERE
func (c Circle) Area() float64 {
	// YOUR CODE HERE - use math.Pi * r * r
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Scale(factor float64) {
	// YOUR CODE HERE
	c.Radius *= factor
}

// TODO: Define Person struct
// YOUR CODE HERE
type Person struct {
	Name string
	Age  int
}

// TODO: Implement Person methods
// YOUR CODE HERE
func (p Person) String() string {
	// YOUR CODE HERE
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

func (p *Person) Birthday() {
	// YOUR CODE HERE
	p.Age++
}

// Exercise 6: Function Types and Higher-Order Functions
func exerciseFunctionTypes() {
	fmt.Println("\n=== Exercise 6: Function Types ===")

	// TODO: Define a function type MathOperation
	// Should take two ints and return an int
	// YOUR CODE HERE

	// TODO: Create functions that match the MathOperation type
	// YOUR CODE HERE

	// Test function variables
	var operation MathOperation

	operation = add
	fmt.Printf("5 + 3 = %d\n", operation(5, 3))

	operation = multiply
	fmt.Printf("5 * 3 = %d\n", operation(5, 3))

	// TODO: Create a function calculate(a, b int, op MathOperation) int
	// This demonstrates functions as parameters
	// YOUR CODE HERE

	fmt.Printf("Calculate(10, 4, add) = %d\n", calculate(10, 4, add))
	fmt.Printf("Calculate(10, 4, multiply) = %d\n", calculate(10, 4, multiply))

	// TODO: Create a slice of operations and apply them
	operations := []MathOperation{add, multiply}
	values := [][]int{{2, 3}, {4, 5}}

	for i, op := range operations {
		a, b := values[i][0], values[i][1]
		result := op(a, b)
		fmt.Printf("Operation %d: %d, %d -> %d\n", i+1, a, b, result)
	}

	// TODO: Create a function applyToAll(numbers []int, fn func(int) int) []int
	// Apply the function to all numbers in the slice
	// YOUR CODE HERE

	numbers := []int{1, 2, 3, 4, 5}
	squared := applyToAll(numbers, func(x int) int { return x * x })
	doubled := applyToAll(numbers, func(x int) int { return x * 2 })

	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Squared: %v\n", squared)
	fmt.Printf("Doubled: %v\n", doubled)
}

// TODO: Define MathOperation function type
// YOUR CODE HERE
type MathOperation func(int, int) int

// TODO: Implement functions that match MathOperation
// YOUR CODE HERE
func add(a, b int) int {
	// YOUR CODE HERE
	return a + b
}

func multiply(a, b int) int {
	// YOUR CODE HERE
	return a * b
}

// TODO: Implement calculate function
// YOUR CODE HERE
func calculate(a, b int, op MathOperation) int {
	// YOUR CODE HERE
	return op(a, b)
}

// TODO: Implement applyToAll function
// YOUR CODE HERE
func applyToAll(numbers []int, fn func(int) int) []int {
	// YOUR CODE HERE
	results := []int{}
	for _, n := range numbers {
		results = append(results, fn(n))
	}
	return results
}

// Exercise 7: Real-World Function Patterns
func exerciseRealWorld() {
	fmt.Println("\n=== Exercise 7: Real-World Patterns ===")

	// TODO: Create a function processData(data []string, validator
	// func(string) bool, transformer func(string) string) []string
	// Filter data using validator, then transform valid items
	// YOUR CODE HERE

	data := []string{"hello", "", "world", "go", "", "programming"}

	processData := func(data []string, validator func(string) bool, transformer func(string) string) []string {
		filtered := []string{}

		for _, item := range data {
			if validator(item) {
				filtered = append(filtered, transformer(item))
			}
		}

		return filtered
	}

	// Validator: non-empty strings
	isNonEmpty := func(s string) bool {
		return s != ""
	}

	// Transformer: uppercase
	toUpper := func(s string) string {
		return strings.ToUpper(s)
	}

	processed := processData(data, isNonEmpty, toUpper)
	fmt.Printf("Original: %v\n", data)
	fmt.Printf("Processed: %v\n", processed)

	// TODO: Create a function retry(operation func() error, maxAttempts int) error
	// Retry an operation up to maxAttempts times
	// YOUR CODE HERE
	retry := func(operation func() error, maxAttempts int) error {
		for i := 0; i < maxAttempts; i++ {
			if err := operation(); err == nil {
				return nil
			}
		}
		return errors.New("operation failed after max attempts")
	}

	// Simulate a flaky operation
	attemptCount := 0
	flakyOperation := func() error {
		attemptCount++
		if attemptCount < 3 {
			return errors.New("temporary failure")
		}
		return nil
	}

	err := retry(flakyOperation, 5)
	if err != nil {
		fmt.Printf("Operation failed after retries: %v\n", err)
	} else {
		fmt.Printf("Operation succeeded after %d attempts\n", attemptCount)
	}

	// TODO: Create a function memoize(fn func(int) int) func(int) int
	// Return a memoized version of the function (caches results)
	// YOUR CODE HERE
	type IntIntFunc func(int) int

	memoize := func(fn IntIntFunc) IntIntFunc {
		cache := make(map[int]int)

		return func(n int) int {
			if val, ok := cache[n]; ok {
				return val
			}
			result := fn(n)
			cache[n] = result
			return result
		}
	}

	// Expensive fibonacci function
	var fibonacci func(int) int
	fibonacci = func(n int) int {
		fmt.Printf("Computing fibonacci(%d)\n", n)
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	memoizedFib := memoize(fibonacci)
	fmt.Println("First call to memoized fibonacci(5):")
	result1 := memoizedFib(5)
	fmt.Printf("Result: %d\n", result1)

	fmt.Println("Second call to memoized fibonacci(5):")
	result2 := memoizedFib(5)
	fmt.Printf("Result: %d\n", result2)
}

// Main function to run all exercises
func main() {
	fmt.Println("🚀 Go Functions Practice")
	fmt.Println("========================")

	exerciseBasicFunctions()
	exerciseNamedReturns()
	exerciseVariadicFunctions()
	exerciseAnonymousFunctions()
	exerciseMethods()
	exerciseFunctionTypes()
	exerciseRealWorld()

	fmt.Println("\n✅ Function exercises completed!")
	fmt.Println("\n💡 Key Takeaways:")
	fmt.Println("- Functions can return multiple values")
	fmt.Println("- Named returns make functions self-documenting")
	fmt.Println("- Variadic functions provide flexible APIs")
	fmt.Println("- Closures capture variables from outer scope")
	fmt.Println("- Methods attach behavior to types")
	fmt.Println("- Function types enable higher-order programming")
	fmt.Println("- The (result, error) pattern is idiomatic Go")
}
