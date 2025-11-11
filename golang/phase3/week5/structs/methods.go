package main

import (
	"fmt"
	"math"
	"strings"
)

// ===== BASIC METHOD DEFINITION =====

type Rectangle struct {
	Width  float64
	Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with value receiver
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Method with pointer receiver (for modification)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Method with pointer receiver (returns self for chaining)
func (r *Rectangle) SetWidth(width float64) *Rectangle {
	r.Width = width
	return r
}

func (r *Rectangle) SetHeight(height float64) *Rectangle {
	r.Height = height
	return r
}

func demonstrateBasicMethods() {
	fmt.Println("=== Basic Methods ===")

	rect := Rectangle{Width: 10, Height: 5}

	// Calling methods with value receivers
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	// Calling method with pointer receiver
	rect.Scale(2)
	fmt.Printf("After scaling by 2: %+v\n", rect)
	fmt.Printf("New area: %.2f\n", rect.Area())

	// Method chaining
	rect2 := Rectangle{}
	rect2.SetWidth(8).SetHeight(6)
	fmt.Printf("Chained rectangle: %+v\n", rect2)
}

// ===== VALUE VS POINTER RECEIVERS =====

type Counter struct {
	value int
}

// Value receiver - operates on a copy
func (c Counter) ValueIncrement() {
	c.value++ // This modifies the copy, not the original
	fmt.Printf("Inside ValueIncrement: %d\n", c.value)
}

// Value receiver - reads value
func (c Counter) Value() int {
	return c.value
}

// Pointer receiver - modifies the original
func (c *Counter) PointerIncrement() {
	c.value++
	fmt.Printf("Inside PointerIncrement: %d\n", c.value)
}

// Pointer receiver - resets value
func (c *Counter) Reset() {
	c.value = 0
}

func demonstrateReceiverTypes() {
	fmt.Println("\n=== Value vs Pointer Receivers ===")

	counter := Counter{value: 10}
	fmt.Printf("Initial counter: %+v\n", counter)

	// Value receiver - doesn't modify original
	counter.ValueIncrement()
	fmt.Printf("After ValueIncrement: %+v\n", counter) // Still 10

	// Pointer receiver - modifies original
	counter.PointerIncrement()
	fmt.Printf("After PointerIncrement: %+v\n", counter) // Now 11

	// Go automatically handles address-of and dereference
	counterPtr := &counter
	counterPtr.PointerIncrement()             // Go automatically dereferences
	fmt.Printf("Via pointer: %+v\n", counter) // Now 12

	// Value receiver also works with pointers (Go auto-dereferences)
	fmt.Printf("Value via pointer: %d\n", counterPtr.Value())
}

// ===== METHODS ON DIFFERENT TYPES =====

// Methods on basic types (with type alias)
type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t)
}

func (t Temperature) Fahrenheit() float64 {
	return float64(t)*9/5 + 32
}

func (t Temperature) Kelvin() float64 {
	return float64(t) + 273.15
}

// String method - implements fmt.Stringer interface
func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°C", t.Celsius())
}

// Methods on slice types
type Numbers []int

func (n Numbers) Sum() int {
	total := 0
	for _, num := range n {
		total += num
	}
	return total
}

func (n Numbers) Average() float64 {
	if len(n) == 0 {
		return 0
	}
	return float64(n.Sum()) / float64(len(n))
}

func (n *Numbers) Add(num int) {
	*n = append(*n, num)
}

func (n Numbers) String() string {
	return fmt.Sprintf("Numbers%v", []int(n))
}

// Methods on map types
type Inventory map[string]int

func (inv Inventory) Total() int {
	total := 0
	for _, quantity := range inv {
		total += quantity
	}
	return total
}

func (inv Inventory) Add(item string, quantity int) {
	inv[item] += quantity
}

func (inv Inventory) Remove(item string, quantity int) bool {
	current, exists := inv[item]
	if !exists || current < quantity {
		return false
	}

	if current == quantity {
		delete(inv, item)
	} else {
		inv[item] = current - quantity
	}
	return true
}

func demonstrateMethodsOnTypes() {
	fmt.Println("\n=== Methods on Different Types ===")

	// Temperature methods
	temp := Temperature(25.0)
	fmt.Printf("Temperature: %s\n", temp) // Uses String() method
	fmt.Printf("Fahrenheit: %.1f°F\n", temp.Fahrenheit())
	fmt.Printf("Kelvin: %.1fK\n", temp.Kelvin())

	// Number slice methods
	numbers := Numbers{1, 2, 3, 4, 5}
	fmt.Printf("Numbers: %s\n", numbers)
	fmt.Printf("Sum: %d\n", numbers.Sum())
	fmt.Printf("Average: %.2f\n", numbers.Average())

	numbers.Add(6)
	fmt.Printf("After adding 6: %s\n", numbers)

	// Map methods
	inventory := Inventory{
		"apples":  10,
		"bananas": 5,
		"oranges": 8,
	}

	fmt.Printf("Inventory total: %d items\n", inventory.Total())

	inventory.Add("apples", 5)
	fmt.Printf("After adding apples: %v\n", inventory)

	success := inventory.Remove("bananas", 3)
	fmt.Printf("Removed bananas: %t, inventory: %v\n", success, inventory)
}

// ===== METHOD SETS AND INTERFACES =====

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

// Value receiver methods
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Pointer receiver method
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

// Interface that requires pointer receiver
type Scalable interface {
	Scale(factor float64)
}

func demonstrateMethodSets() {
	fmt.Println("\n=== Method Sets and Interfaces ===")

	circle := Circle{Radius: 5}
	circlePtr := &circle

	// Value receivers can be called on both values and pointers
	fmt.Printf("Circle area (value): %.2f\n", circle.Area())
	fmt.Printf("Circle area (pointer): %.2f\n", circlePtr.Area())

	// Pointer receivers need addressable values
	circlePtr.Scale(2)
	fmt.Printf("After scaling: radius = %.2f\n", circle.Radius)

	// Interface satisfaction
	var shape Shape
	shape = circle // Works - value implements interface with value receiver methods
	fmt.Printf("Shape area: %.2f\n", shape.Area())

	shape = circlePtr // Also works - pointer can access value receiver methods
	fmt.Printf("Shape area via pointer: %.2f\n", shape.Area())

	// For interfaces requiring pointer receivers
	var scalable Scalable
	scalable = circlePtr // Works - pointer implements interface
	scalable.Scale(0.5)
	fmt.Printf("After scaling via interface: radius = %.2f\n", circle.Radius)

	// scalable = circle // This would NOT work - value doesn't implement Scalable
}

// ===== METHOD VALUES AND EXPRESSIONS =====

type Calculator struct {
	total float64
}

func (c *Calculator) Add(value float64) {
	c.total += value
}

func (c *Calculator) Multiply(value float64) {
	c.total *= value
}

func (c Calculator) Total() float64 {
	return c.total
}

func demonstrateMethodValues() {
	fmt.Println("\n=== Method Values and Expressions ===")

	calc := Calculator{total: 10}

	// Method value - binds the receiver
	addMethod := calc.Add
	addMethod(5) // Equivalent to calc.Add(5)
	fmt.Printf("After method value add: %.2f\n", calc.Total())

	// Method expression - receiver becomes first parameter
	addExpr := (*Calculator).Add
	addExpr(&calc, 3) // Equivalent to calc.Add(3)
	fmt.Printf("After method expression add: %.2f\n", calc.Total())

	// Useful for passing methods as functions
	operations := []func(*Calculator, float64){
		(*Calculator).Add,
		(*Calculator).Multiply,
	}

	calc.total = 5
	for i, op := range operations {
		op(&calc, 2)
		fmt.Printf("After operation %d: %.2f\n", i+1, calc.Total())
	}
}

// ===== METHODS WITH COMPLEX LOGIC =====

type StringProcessor struct {
	text string
}

func NewStringProcessor(text string) *StringProcessor {
	return &StringProcessor{text: text}
}

// Method returning modified copy (functional style)
func (sp StringProcessor) ToUpper() StringProcessor {
	return StringProcessor{text: strings.ToUpper(sp.text)}
}

func (sp StringProcessor) ToLower() StringProcessor {
	return StringProcessor{text: strings.ToLower(sp.text)}
}

func (sp StringProcessor) Reverse() StringProcessor {
	runes := []rune(sp.text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return StringProcessor{text: string(runes)}
}

// Method modifying in place
func (sp *StringProcessor) ReplaceInPlace(old, new string) {
	sp.text = strings.ReplaceAll(sp.text, old, new)
}

func (sp StringProcessor) String() string {
	return sp.text
}

// Fluent interface pattern
func (sp StringProcessor) Chain() *StringChain {
	return &StringChain{processor: sp}
}

type StringChain struct {
	processor StringProcessor
}

func (sc *StringChain) ToUpper() *StringChain {
	sc.processor = sc.processor.ToUpper()
	return sc
}

func (sc *StringChain) ToLower() *StringChain {
	sc.processor = sc.processor.ToLower()
	return sc
}

func (sc *StringChain) Reverse() *StringChain {
	sc.processor = sc.processor.Reverse()
	return sc
}

func (sc *StringChain) Result() string {
	return sc.processor.String()
}

func demonstrateComplexMethods() {
	fmt.Println("\n=== Complex Method Patterns ===")

	// Functional style (immutable)
	processor := StringProcessor{text: "Hello World"}
	upper := processor.ToUpper()
	reversed := upper.Reverse()

	fmt.Printf("Original: %s\n", processor)
	fmt.Printf("Upper: %s\n", upper)
	fmt.Printf("Reversed: %s\n", reversed)

	// Mutable style
	mutableProcessor := NewStringProcessor("Go Programming")
	mutableProcessor.ReplaceInPlace("Go", "Golang")
	fmt.Printf("After replacement: %s\n", mutableProcessor)

	// Fluent interface
	result := StringProcessor{text: "hello world"}.
		Chain().
		ToUpper().
		Reverse().
		Result()
	fmt.Printf("Fluent result: %s\n", result)
}

// ===== MAIN DEMO FUNCTION =====

func runMethodsDemo() {
	fmt.Println("⚙️ Go Methods and Receivers Tutorial")
	fmt.Println("====================================")

	demonstrateBasicMethods()
	demonstrateReceiverTypes()
	demonstrateMethodsOnTypes()
	demonstrateMethodSets()
	demonstrateMethodValues()
	demonstrateComplexMethods()

	fmt.Println("\n✅ Methods and receivers concepts covered!")
	fmt.Println("\n🎯 Key Points:")
	fmt.Println("- Methods are functions with receivers")
	fmt.Println("- Value receivers work on copies, pointer receivers on originals")
	fmt.Println("- Go automatically handles address-of and dereference")
	fmt.Println("- Method sets determine interface satisfaction")
	fmt.Println("- Methods can be defined on any named type")
	fmt.Println("- Use value receivers for small data and immutable operations")
	fmt.Println("- Use pointer receivers for modification or large structs")
	fmt.Println("- Method values and expressions enable functional patterns")
}
