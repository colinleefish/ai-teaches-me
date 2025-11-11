// Exercise 2: Basic Types
// Explore Go's fundamental types and their operations

package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== Basic Types ===")

	// TODO: Complete these functions to demonstrate Go types
	numericTypes()
	stringOperations()
	booleanLogic()
	typeConversions()
	constantsVsVariables()
	demonstrateTypeBehavior()
	fmt.Println("\n🔢 Basic types mastery complete!")
}

// TODO: Implement numeric types demonstration
func numericTypes() {
	fmt.Println("\n1. Numeric Types:")
	// Your code here:
	// - Show different integer types (int, int8, int16, int32, int64)
	// - Show floating-point types (float32, float64)
	// - Demonstrate type sizes and ranges
	// - Show basic arithmetic operations

	// Example:
	// var smallInt int8 = 127
	// var bigInt int64 = 9223372036854775807
	// var pi float64 = 3.14159
	// var smallFloat float32 = 3.14

	// fmt.Printf("int8: %d (size: %d bytes)\n", smallInt, 1)
	// fmt.Printf("int64: %d (size: %d bytes)\n", bigInt, 8)
	// fmt.Printf("float64: %.5f (size: %d bytes)\n", pi, 8)
	// fmt.Printf("float32: %.2f (size: %d bytes)\n", smallFloat, 4)

	// Arithmetic operations
	// a, b := 10, 3
	// fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	// fmt.Printf("Division: %d / %d = %d\n", a, b, a/b)
	// fmt.Printf("Modulo: %d %% %d = %d\n", a, b, a%b)
	// fmt.Printf("Float division: %.2f / %.2f = %.2f\n", float64(a), float64(b), float64(a)/float64(b))

	var smallInt int8 = 127
	fmt.Printf("smallInt = %v, type: %s\n", smallInt, reflect.TypeOf(smallInt))

	var bigInt int64 = 1234567890123456789
	fmt.Printf("bigInt = %v, type: %s\n", bigInt, reflect.TypeOf(bigInt))

	var biggerInt int64 = int64(smallInt) + bigInt
	fmt.Printf("biggerInt = %v, type: %s\n", biggerInt, reflect.TypeOf(biggerInt))

}

// TODO: Implement string operations
func stringOperations() {
	fmt.Println("\n2. String Operations:")
	// Your code here:
	// - Create and manipulate strings
	// - Show string concatenation
	// - Use strings package functions
	// - Demonstrate string formatting
	// - Show the difference between strings and runes

	// Example:
	// firstName := "John"
	// lastName := "Doe"
	// fullName := firstName + " " + lastName

	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName

	fmt.Printf("Full name: %v\n, type: %s, reflected value: %T\n", fullName, reflect.TypeOf(fullName), reflect.ValueOf(fullName))
	fmt.Printf("Length: %d\n", len(fullName))
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(fullName))
	fmt.Printf("Uppercase v2: %v\n", reflect.ValueOf(strings.ToUpper(fullName)))
	fmt.Printf("Contains 'John': %t\n", strings.Contains(fullName, "John"))

	// String formatting
	age := 30
	fmt.Printf("Formatted: %s is %d years old\n", fullName, age)

	// Runes (Unicode code points)
	// text := "Hello 世界"
	// fmt.Printf("String: %s\n", text)
	// fmt.Printf("Length in bytes: %d\n", len(text))
	// fmt.Printf("Length in runes: %d\n", len([]rune(text)))

	text := "Hello 塞卡旖旎"
	fmt.Printf("String: %s\n", text)
	fmt.Printf("Length in bytes: %d\n", len(text))
	fmt.Printf("Length in runes: %d\n", len([]rune(text)))
}

// TODO: Implement boolean logic
func booleanLogic() {
	fmt.Println("\n3. Boolean Logic:")
	// Your code here:
	// - Show boolean values and operations
	// - Demonstrate logical operators (&&, ||, !)
	// - Show comparison operators
	// - Use booleans in conditional logic

	// Example:
	// isAdult := true
	// hasLicense := false
	// age := 25

	// fmt.Printf("Is adult: %t\n", isAdult)
	// fmt.Printf("Has license: %t\n", hasLicense)
	// fmt.Printf("Can drive: %t\n", isAdult && hasLicense)
	// fmt.Printf("Needs training: %t\n", isAdult && !hasLicense)

	// Comparisons
	// fmt.Printf("Age >= 18: %t\n", age >= 18)
	// fmt.Printf("Age == 25: %t\n", age == 25)
	// fmt.Printf("Age != 30: %t\n", age != 30)

	isAdult := true
	hasLicense := false
	age := 25

	fmt.Printf("Is adult: %t\n", isAdult)
	fmt.Printf("Has license: %v\n", hasLicense)
	fmt.Printf("Can drive: %t\n", isAdult && hasLicense)

	fmt.Printf("Age >= 18: %t\n", age >= 18)
	fmt.Printf("Age == 25: %t\n", age == 25)

}

// TODO: Implement type conversions
func typeConversions() {
	fmt.Println("\n4. Type Conversions:")
	// Your code here:
	// - Show explicit type conversions
	// - Convert between numeric types
	// - Convert strings to numbers and vice versa
	// - Handle conversion errors

	// Example:
	// var i int = 42
	// var f float64 = float64(i)
	// var u uint = uint(f)

	// fmt.Printf("int to float64: %d -> %.2f\n", i, f)
	// fmt.Printf("float64 to uint: %.2f -> %d\n", f, u)

	// String conversions
	// str := "123"
	// num, err := strconv.Atoi(str)
	// if err != nil {
	//     fmt.Printf("Error converting '%s' to int: %v\n", str, err)
	// } else {
	//     fmt.Printf("String to int: '%s' -> %d\n", str, num)
	// }

	// fmt.Printf("Int to string: %d -> '%s'\n", num, strconv.Itoa(num))

	var i int = 42
	var iToF = float64(i)
	fmt.Printf("i = %v, type: %T\n", iToF, iToF)

	var f float64 = 3.14
	fmt.Printf("f = %v, type: %T\n", f, f)
	var fToU = uint(f)
	fmt.Printf("fToU = %v, type: %T\n", fToU, fToU)

	str := "123"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Printf("Error converting str: %s to int\n", str)
	} else {
		fmt.Printf("String to int: %s -> %d\n, type of num: %T\n, type of str: %T\n", str, num, num, str)
	}

	fmt.Printf("Int to string: %d -> %s\n", num, strconv.Itoa(num))
	fmt.Printf("type of num: %T, type of strconv.Itoa(num): %T\n", num, strconv.Itoa(num))

}

// TODO: Implement constants vs variables
func constantsVsVariables() {
	fmt.Println("\n5. Constants vs Variables:")
	// Your code here:
	// - Define constants with const keyword
	// - Show typed and untyped constants
	// - Demonstrate iota for enumeration
	// - Compare with variables

	// Example:
	const pi = "3.14159"
	const greeting string = "Hello"

	fmt.Printf("Pi (constant): %.5f\n", pi)
	fmt.Printf("Greeting (constant): %s\n", greeting)

	num, err := strconv.ParseFloat(pi, 64)
	if err != nil {
		fmt.Printf("Error converting pi: %s to float64\n", pi)
	} else {
		fmt.Printf("Pi (constant)======: %.5f\n", num)
	}

	// Iota example
	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
	)

	fmt.Printf("Monday: %d, Tuesday: %d, Wednesday: %d\n", Monday, Tuesday, Wednesday)

	// Variable vs constant
	var variable int = 10
	variable = 20 // OK
	fmt.Printf("Variable changed: %d\n", variable)

}

// Helper function to demonstrate type behavior
func demonstrateTypeBehavior() {
	fmt.Println("\n6. Type Behavior:")

	// Zero values
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Printf("Zero int: %d\n", zeroInt)
	fmt.Printf("Zero float: %.2f\n", zeroFloat)
	fmt.Printf("Zero string: '%s'\n", zeroString)
	fmt.Printf("Zero bool: %t\n", zeroBool)

	// Type limits
	fmt.Printf("Max int8: %d\n", math.MaxInt8)
	fmt.Printf("Min int8: %d\n", math.MinInt8)
	fmt.Printf("Max float64: %g\n", math.MaxFloat64)
}

// 🎯 LEARNING GOALS:
// 1. Understand Go's type system
// 2. Master numeric operations and conversions
// 3. Work with strings and Unicode
// 4. Use boolean logic effectively
// 5. Distinguish between constants and variables

// 🐍 PYTHON COMPARISON:
// Python: 5 / 2 = 2.5 (float division)
// Go: 5 / 2 = 2 (integer division), 5.0 / 2.0 = 2.5
//
// Python: str(42) converts to string
// Go: strconv.Itoa(42) converts to string
//
// Python: int("123") converts to int
// Go: strconv.Atoi("123") converts to int with error handling
//
// Python: No explicit type declarations
// Go: Must declare types or let compiler infer

// 🚀 TESTING YOUR WORK:
// 1. Run: go run types.go
// 2. Experiment with different type operations
// 3. Try invalid conversions to see error handling
// 4. Compare behavior with Python

// 🔧 COMMON MISTAKES:
// 1. Forgetting explicit type conversions
// 2. Integer division vs float division
// 3. String vs []byte vs []rune confusion
// 4. Ignoring error returns from strconv functions
// 5. Assuming automatic type promotion
