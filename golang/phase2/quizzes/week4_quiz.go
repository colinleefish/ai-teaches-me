package main

import (
	"errors"
	"fmt"
)

// Week 4 Quiz: Error Handling & Advanced Function Patterns
// Complete this quiz to test your understanding of Phase 2 concepts

func main() {
	fmt.Println("📝 Week 4 Quiz: Error Handling & Advanced Function Patterns")
	fmt.Println("==========================================================")

	runQuiz()
}

func runQuiz() {
	fmt.Println("\n=== Part 1: Error Handling ===")

	// Question 1: Custom Error Types
	fmt.Println("\n1. Which interface must a type implement to be used as an error?")
	fmt.Println("   a) fmt.Stringer")
	fmt.Println("   b) error")
	fmt.Println("   c) fmt.Error")
	fmt.Println("   d) runtime.Error")

	// TODO: Implement your answer checking logic
	// Correct answer: b) error

	// Question 2: Error Wrapping
	fmt.Println("\n2. Which verb should you use with fmt.Errorf to wrap an error (Go 1.13+)?")
	fmt.Println("   a) %v")
	fmt.Println("   b) %s")
	fmt.Println("   c) %w")
	fmt.Println("   d) %e")

	// Correct answer: c) %w

	// Question 3: Error Checking
	fmt.Println("\n3. What's the idiomatic way to check for errors in Go?")
	fmt.Println("   a) try/catch blocks")
	fmt.Println("   b) if err != nil")
	fmt.Println("   c) switch err.(type)")
	fmt.Println("   d) errors.Check(err)")

	// Correct answer: b) if err != nil

	// Question 4: Panic vs Error
	fmt.Println("\n4. When should you use panic instead of returning an error?")
	fmt.Println("   a) For any error condition")
	fmt.Println("   b) For user input validation")
	fmt.Println("   c) For programming errors that shouldn't happen")
	fmt.Println("   d) For network timeouts")

	// Correct answer: c) For programming errors that shouldn't happen

	// Question 5: Error Unwrapping
	fmt.Println("\n5. Which function checks if an error is or wraps a specific error?")
	fmt.Println("   a) errors.Is()")
	fmt.Println("   b) errors.As()")
	fmt.Println("   c) errors.Unwrap()")
	fmt.Println("   d) errors.Check()")

	// Correct answer: a) errors.Is()

	fmt.Println("\n=== Part 2: Advanced Function Patterns ===")

	// Question 6: Recursion
	fmt.Println("\n6. What's essential for every recursive function?")
	fmt.Println("   a) A loop")
	fmt.Println("   b) A base case")
	fmt.Println("   c) Error handling")
	fmt.Println("   d) Memoization")

	// Correct answer: b) A base case

	// Question 7: Higher-Order Functions
	fmt.Println("\n7. What is a higher-order function?")
	fmt.Println("   a) A function with multiple return values")
	fmt.Println("   b) A function that accepts or returns other functions")
	fmt.Println("   c) A function with named returns")
	fmt.Println("   d) A function that handles errors")

	// Correct answer: b) A function that accepts or returns other functions

	// Question 8: Memoization
	fmt.Println("\n8. What is the purpose of memoization?")
	fmt.Println("   a) To handle errors")
	fmt.Println("   b) To cache function results for performance")
	fmt.Println("   c) To compose functions")
	fmt.Println("   d) To log function calls")

	// Correct answer: b) To cache function results for performance

	// Question 9: Function Composition
	fmt.Println("\n9. In function composition f(g(x)), which function is applied first?")
	fmt.Println("   a) f")
	fmt.Println("   b) g")
	fmt.Println("   c) Both simultaneously")
	fmt.Println("   d) It depends on the implementation")

	// Correct answer: b) g

	// Question 10: Closures
	fmt.Println("\n10. What can closures capture from their surrounding scope?")
	fmt.Println("    a) Only parameters")
	fmt.Println("    b) Only return values")
	fmt.Println("    c) Variables from the enclosing function")
	fmt.Println("    d) Global variables only")

	// Correct answer: c) Variables from the enclosing function

	fmt.Println("\n=== Practical Questions ===")

	// Practical Question 1
	fmt.Println("\n11. Complete this custom error type:")
	fmt.Println("```go")
	fmt.Println("type ValidationError struct {")
	fmt.Println("    Field string")
	fmt.Println("    Message string")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("func (e ValidationError) _____ _____ {")
	fmt.Println("    return fmt.Sprintf(\"%s: %s\", e.Field, e.Message)")
	fmt.Println("}")
	fmt.Println("```")
	fmt.Println("Fill in the blanks: _____ _____")

	// Answer: Error() string

	// Practical Question 2
	fmt.Println("\n12. Complete this memoization function:")
	fmt.Println("```go")
	fmt.Println("func memoize(fn func(int) int) func(int) int {")
	fmt.Println("    cache := make(map[int]int)")
	fmt.Println("    return func(n int) int {")
	fmt.Println("        if result, exists := cache[n]; _____ {")
	fmt.Println("            return result")
	fmt.Println("        }")
	fmt.Println("        result := fn(n)")
	fmt.Println("        cache[n] = result")
	fmt.Println("        return result")
	fmt.Println("    }")
	fmt.Println("}")
	fmt.Println("```")
	fmt.Println("Fill in the blank: _____")

	// Answer: exists

	fmt.Println("\n=== Code Analysis ===")

	// Code Analysis Question
	fmt.Println("\n13. What will this code print?")
	fmt.Println("```go")
	fmt.Println("func createCounter() func() int {")
	fmt.Println("    count := 0")
	fmt.Println("    return func() int {")
	fmt.Println("        count++")
	fmt.Println("        return count")
	fmt.Println("    }")
	fmt.Println("}")
	fmt.Println("")
	fmt.Println("counter := createCounter()")
	fmt.Println("fmt.Println(counter())")
	fmt.Println("fmt.Println(counter())")
	fmt.Println("fmt.Println(counter())")
	fmt.Println("```")
	fmt.Println("Output: ___, ___, ___")

	// Answer: 1, 2, 3

	// Error Handling Analysis
	fmt.Println("\n14. Is this error handling correct?")
	fmt.Println("```go")
	fmt.Println("func processFile(filename string) error {")
	fmt.Println("    data, err := readFile(filename)")
	fmt.Println("    if err != nil {")
	fmt.Println("        return fmt.Errorf(\"failed to read %s: %w\", filename, err)")
	fmt.Println("    }")
	fmt.Println("    return processData(data)")
	fmt.Println("}")
	fmt.Println("```")
	fmt.Println("a) Yes, it properly wraps the error with context")
	fmt.Println("b) No, it should use %v instead of %w")
	fmt.Println("c) No, it should panic instead")
	fmt.Println("d) No, it should ignore the error")

	// Answer: a) Yes, it properly wraps the error with context

	fmt.Println("\n15. Which recursive function implementation is better?")
	fmt.Println("\nOption A:")
	fmt.Println("```go")
	fmt.Println("func factorial(n int) int {")
	fmt.Println("    if n <= 1 {")
	fmt.Println("        return 1")
	fmt.Println("    }")
	fmt.Println("    return n * factorial(n-1)")
	fmt.Println("}")
	fmt.Println("```")
	fmt.Println("\nOption B:")
	fmt.Println("```go")
	fmt.Println("func factorial(n int) (int, error) {")
	fmt.Println("    if n < 0 {")
	fmt.Println("        return 0, errors.New(\"factorial undefined for negative numbers\")")
	fmt.Println("    }")
	fmt.Println("    if n <= 1 {")
	fmt.Println("        return 1, nil")
	fmt.Println("    }")
	fmt.Println("    result, err := factorial(n-1)")
	fmt.Println("    if err != nil {")
	fmt.Println("        return 0, err")
	fmt.Println("    }")
	fmt.Println("    return n * result, nil")
	fmt.Println("}")
	fmt.Println("```")
	fmt.Println("Answer: a) Option A  b) Option B")

	// Answer: b) Option B (handles edge cases and errors properly)

	fmt.Println("\n=== Quiz Complete ===")
	fmt.Println("Review your answers and check the solutions below!")

	printAnswers()
}

func printAnswers() {
	fmt.Println("\n🔍 Answer Key:")
	fmt.Println("==============")
	fmt.Println("1. b) error")
	fmt.Println("2. c) %w")
	fmt.Println("3. b) if err != nil")
	fmt.Println("4. c) For programming errors that shouldn't happen")
	fmt.Println("5. a) errors.Is()")
	fmt.Println("6. b) A base case")
	fmt.Println("7. b) A function that accepts or returns other functions")
	fmt.Println("8. b) To cache function results for performance")
	fmt.Println("9. b) g")
	fmt.Println("10. c) Variables from the enclosing function")
	fmt.Println("11. Error() string")
	fmt.Println("12. exists")
	fmt.Println("13. 1, 2, 3")
	fmt.Println("14. a) Yes, it properly wraps the error with context")
	fmt.Println("15. b) Option B")

	fmt.Println("\n📊 Scoring:")
	fmt.Println("13-15: Excellent! You've mastered error handling and advanced patterns")
	fmt.Println("10-12: Good! Review the areas you missed")
	fmt.Println("7-9:   Fair - Practice more with the exercises")
	fmt.Println("0-6:   Review the material and try the exercises again")

	fmt.Println("\n🎯 Next Steps:")
	fmt.Println("- Complete all Phase 2 exercises")
	fmt.Println("- Build a project combining error handling and advanced patterns")
	fmt.Println("- Prepare for Phase 3: Concurrency and Goroutines!")
}

// Example implementations for reference

// Example 1: Custom Error Type
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Example 2: Error Wrapping
func processFile(filename string) error {
	// Simulate file reading
	if filename == "" {
		return ValidationError{Field: "filename", Message: "cannot be empty"}
	}

	// Simulate an error from a lower-level function
	err := errors.New("file not found")
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", filename, err)
	}

	return nil
}

// Example 3: Memoized Function
func createMemoizedFibonacci() func(int) int {
	cache := make(map[int]int)

	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}

		if result, exists := cache[n]; exists {
			return result
		}

		result := fib(n-1) + fib(n-2)
		cache[n] = result
		return result
	}

	return fib
}

// Example 4: Higher-Order Function
func createFilter(predicate func(int) bool) func([]int) []int {
	return func(slice []int) []int {
		var result []int
		for _, item := range slice {
			if predicate(item) {
				result = append(result, item)
			}
		}
		return result
	}
}

// Example 5: Function Composition
type Transform func(int) int

func (t Transform) Then(next Transform) Transform {
	return func(x int) int {
		return next(t(x))
	}
}

func addOne(x int) int { return x + 1 }
func double(x int) int { return x * 2 }
func square(x int) int { return x * x }

// Usage: Transform(addOne).Then(double).Then(square)(5) = ((5+1)*2)^2 = 144
