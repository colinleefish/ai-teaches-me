// Week 1 Solutions - Reference Implementations
// Use these solutions to check your work or get unstuck

package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("🔍 Week 1 Reference Solutions")
	fmt.Println("Choose a solution to view:")
	fmt.Println("1. setup.go solutions")
	fmt.Println("2. hello.go solutions")
	fmt.Println("3. variables.go solutions")
	fmt.Println("4. types.go solutions")
	
	// For demonstration, we'll show all solutions
	setupSolutions()
	helloSolutions()
	variablesSolutions()
	typesSolutions()
}

// ===== SETUP.GO SOLUTIONS =====
func setupSolutions() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📝 SETUP.GO SOLUTIONS")
	fmt.Println(strings.Repeat("=", 50))
	
	checkGoVersionSolution()
	checkGoEnvironmentSolution()
	demonstrateBasicSyntaxSolution()
	exploreWorkspaceSolution()
}

func checkGoVersionSolution() {
	fmt.Println("\n1. Go Version Information:")
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
}

func checkGoEnvironmentSolution() {
	fmt.Println("\n2. Go Environment:")
	fmt.Printf("GOROOT: %s\n", os.Getenv("GOROOT"))
	fmt.Printf("GOPATH: %s\n", os.Getenv("GOPATH"))
	if dir, err := os.Getwd(); err == nil {
		fmt.Printf("Current directory: %s\n", dir)
	}
}

func demonstrateBasicSyntaxSolution() {
	fmt.Println("\n3. Basic Go Syntax Demo:")
	var name string = "Go"
	age := 14  // Go was first released in 2009
	const creator = "Google"
	fmt.Printf("Language: %s, Age: %d, Creator: %s\n", name, age, creator)
}

func exploreWorkspaceSolution() {
	fmt.Println("\n4. Workspace Structure:")
	if entries, err := os.ReadDir("."); err == nil {
		fmt.Printf("Directory contents:\n")
		for _, entry := range entries {
			fmt.Printf("  %s\n", entry.Name())
		}
	}
	
	if _, err := os.Stat("go.mod"); err == nil {
		fmt.Println("✅ go.mod file exists")
	} else {
		fmt.Println("❌ go.mod file not found")
	}
}

// ===== HELLO.GO SOLUTIONS =====
func helloSolutions() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📝 HELLO.GO SOLUTIONS")
	fmt.Println(strings.Repeat("=", 50))
	
	basicHelloSolution()
	formattedHelloSolution()
	greetWithArgsSolution()
	interactiveGreetingSolution()
	multilingualGreetingSolution()
}

func basicHelloSolution() {
	fmt.Println("\n1. Basic Hello World:")
	fmt.Println("Hello, World!")
}

func formattedHelloSolution() {
	fmt.Println("\n2. Formatted Hello World:")
	name := "Go"
	language := "programming language"
	fmt.Printf("Hello from %s, the %s!\n", name, language)
}

func greetWithArgsSolution() {
	fmt.Println("\n3. Greeting with Command-line Arguments:")
	args := os.Args[1:] // Skip program name
	if len(args) == 0 {
		fmt.Println("Hello, anonymous user!")
	} else {
		for _, name := range args {
			fmt.Printf("Hello, %s!\n", name)
		}
	}
}

func interactiveGreetingSolution() {
	fmt.Println("\n4. Interactive Greeting:")
	hour := time.Now().Hour()
	if hour < 12 {
		fmt.Println("Good morning!")
	} else if hour < 17 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}
}

func multilingualGreetingSolution() {
	fmt.Println("\n5. Multilingual Greeting:")
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

// ===== VARIABLES.GO SOLUTIONS =====
func variablesSolutions() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📝 VARIABLES.GO SOLUTIONS")
	fmt.Println(strings.Repeat("=", 50))
	
	basicDeclarationsSolution()
	shortDeclarationsSolution()
	multipleDeclarationsSolution()
	zeroValuesSolution()
	typeInferenceSolution()
}

func basicDeclarationsSolution() {
	fmt.Println("\n1. Basic var Declarations:")
	var name string = "Alice"
	var age int = 25
	var height float64 = 5.6
	var isStudent bool = true
	
	fmt.Printf("Name: %s (type: %T)\n", name, name)
	fmt.Printf("Age: %d (type: %T)\n", age, age)
	fmt.Printf("Height: %.1f (type: %T)\n", height, height)
	fmt.Printf("Is student: %t (type: %T)\n", isStudent, isStudent)
}

func shortDeclarationsSolution() {
	fmt.Println("\n2. Short Variable Declarations (:=):")
	name := "Bob"        // string inferred
	age := 30            // int inferred
	salary := 50000.50   // float64 inferred
	isManager := false   // bool inferred
	
	fmt.Printf("Name: %s (type: %T)\n", name, name)
	fmt.Printf("Age: %d (type: %T)\n", age, age)
	fmt.Printf("Salary: %.2f (type: %T)\n", salary, salary)
	fmt.Printf("Is manager: %t (type: %T)\n", isManager, isManager)
}

func multipleDeclarationsSolution() {
	fmt.Println("\n3. Multiple Variable Declarations:")
	
	// Style 1: Multiple var declarations
	var x, y, z int = 1, 2, 3
	
	// Style 2: Grouped var declarations
	var (
		firstName string = "John"
		lastName  string = "Doe"
		age       int    = 35
	)
	
	// Style 3: Multiple short declarations
	a, b, c := 10, 20, 30
	
	fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z)
	fmt.Printf("Full name: %s %s, age: %d\n", firstName, lastName, age)
	fmt.Printf("a: %d, b: %d, c: %d\n", a, b, c)
}

func zeroValuesSolution() {
	fmt.Println("\n5. Zero Values:")
	var num int
	var text string
	var flag bool
	var decimal float64
	
	fmt.Printf("Zero int: %d\n", num)
	fmt.Printf("Zero string: '%s'\n", text)
	fmt.Printf("Zero bool: %t\n", flag)
	fmt.Printf("Zero float64: %f\n", decimal)
}

func typeInferenceSolution() {
	fmt.Println("\n6. Type Inference:")
	a := 42          // int
	b := 42.0        // float64
	c := "hello"     // string
	d := true        // bool
	e := 'A'         // rune (int32)
	
	fmt.Printf("a = %v, type: %s\n", a, reflect.TypeOf(a))
	fmt.Printf("b = %v, type: %s\n", b, reflect.TypeOf(b))
	fmt.Printf("c = %v, type: %s\n", c, reflect.TypeOf(c))
	fmt.Printf("d = %v, type: %s\n", d, reflect.TypeOf(d))
	fmt.Printf("e = %v, type: %s\n", e, reflect.TypeOf(e))
}

// ===== TYPES.GO SOLUTIONS =====
func typesSolutions() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📝 TYPES.GO SOLUTIONS")
	fmt.Println(strings.Repeat("=", 50))
	
	numericTypesSolution()
	stringOperationsSolution()
	booleanLogicSolution()
	typeConversionsSolution()
	constantsVsVariablesSolution()
}

func numericTypesSolution() {
	fmt.Println("\n1. Numeric Types:")
	var smallInt int8 = 127
	var bigInt int64 = 9223372036854775807
	var pi float64 = 3.14159
	var smallFloat float32 = 3.14
	
	fmt.Printf("int8: %d (size: %d bytes)\n", smallInt, 1)
	fmt.Printf("int64: %d (size: %d bytes)\n", bigInt, 8)
	fmt.Printf("float64: %.5f (size: %d bytes)\n", pi, 8)
	fmt.Printf("float32: %.2f (size: %d bytes)\n", smallFloat, 4)
	
	// Arithmetic operations
	a, b := 10, 3
	fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Division: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("Modulo: %d %% %d = %d\n", a, b, a%b)
	fmt.Printf("Float division: %.2f / %.2f = %.2f\n", float64(a), float64(b), float64(a)/float64(b))
}

func stringOperationsSolution() {
	fmt.Println("\n2. String Operations:")
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	
	fmt.Printf("Full name: %s\n", fullName)
	fmt.Printf("Length: %d characters\n", len(fullName))
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(fullName))
	fmt.Printf("Contains 'John': %t\n", strings.Contains(fullName, "John"))
	
	// String formatting
	age := 30
	fmt.Printf("Formatted: %s is %d years old\n", fullName, age)
	
	// Runes (Unicode code points)
	text := "Hello 世界"
	fmt.Printf("String: %s\n", text)
	fmt.Printf("Length in bytes: %d\n", len(text))
	fmt.Printf("Length in runes: %d\n", len([]rune(text)))
}

func booleanLogicSolution() {
	fmt.Println("\n3. Boolean Logic:")
	isAdult := true
	hasLicense := false
	age := 25
	
	fmt.Printf("Is adult: %t\n", isAdult)
	fmt.Printf("Has license: %t\n", hasLicense)
	fmt.Printf("Can drive: %t\n", isAdult && hasLicense)
	fmt.Printf("Needs training: %t\n", isAdult && !hasLicense)
	
	// Comparisons
	fmt.Printf("Age >= 18: %t\n", age >= 18)
	fmt.Printf("Age == 25: %t\n", age == 25)
	fmt.Printf("Age != 30: %t\n", age != 30)
}

func typeConversionsSolution() {
	fmt.Println("\n4. Type Conversions:")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	
	fmt.Printf("int to float64: %d -> %.2f\n", i, f)
	fmt.Printf("float64 to uint: %.2f -> %d\n", f, u)
	
	// String conversions
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error converting '%s' to int: %v\n", str, err)
	} else {
		fmt.Printf("String to int: '%s' -> %d\n", str, num)
	}
	
	fmt.Printf("Int to string: %d -> '%s'\n", num, strconv.Itoa(num))
}

func constantsVsVariablesSolution() {
	fmt.Println("\n5. Constants vs Variables:")
	const pi = 3.14159
	const greeting string = "Hello"
	
	fmt.Printf("Pi (constant): %.5f\n", pi)
	fmt.Printf("Greeting (constant): %s\n", greeting)
	
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
	variable = 20  // OK
	fmt.Printf("Variable changed: %d\n", variable)
	
	// pi = 3.14  // This would cause a compile error
}

// 📚 SOLUTION NOTES:
// - These are reference implementations
// - Your solutions may vary but should achieve the same results
// - Focus on understanding the concepts, not memorizing exact syntax
// - Practice writing your own versions before looking at solutions

// 🎯 LEARNING TIPS:
// - Compare your solutions with these references
// - Identify any gaps in your understanding
// - Practice the concepts you find challenging
// - Don't just copy - understand why each solution works

// 🚀 NEXT STEPS:
// - Complete all exercises on your own first
// - Use solutions to check your work
// - Identify areas for improvement
// - Move to Week 2 when confident 