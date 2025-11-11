package main

import (
	"errors"
	"fmt"
	"math"
)

// ===== BASIC INTERFACE DEFINITION =====

// TODO: Define a Speaker interface with one method: Speak() string
// YOUR CODE HERE
type Speaker interface {
	Speak() string
}

// TODO: Define a Dog struct with Name field
// YOUR CODE HERE
type Dog struct {
	Name string
}

// TODO: Implement Speak() method for Dog - should return "[name] says Woof!"
// YOUR CODE HERE
func (d Dog) Speak() string {
	return fmt.Sprintf("%s says Woof!", d.Name)
}

// TODO: Define a Cat struct with Name field
// YOUR CODE HERE
type Cat struct {
	Name string
}

// TODO: Implement Speak() method for Cat - should return "[name] says Meow!"
// YOUR CODE HERE
func (c Cat) Speak() string {
	return fmt.Sprintf("%s says Meow!", c.Name)
}

// TODO: Define a Robot struct with Model field
// YOUR CODE HERE
type Robot struct {
	Model string
}

// TODO: Implement Speak() method for Robot - should return "Robot [model] says: BEEP BOOP!"
// YOUR CODE HERE
func (r Robot) Speak() string {
	return fmt.Sprintf("Robot %s says: Beep!", r.Model)
}

func demonstrateBasicInterfaces() {
	fmt.Println("=== Basic Interface Implementation ===")

	// TODO: Create a slice of Speaker containing Dog, Cat, and Robot instances
	// YOUR CODE HERE
	speakers := []Speaker{Dog{Name: "Rex"}, Cat{"Whiskers"}, Robot{Model: "R2D2"}}

	// TODO: Loop through the speakers and call Speak() on each
	// YOUR CODE HERE
	for _, speaker := range speakers {
		fmt.Println(speaker.Speak())
	}

	// TODO: Create a Speaker variable and assign different types to it
	// Demonstrate that the same variable can hold different implementing types
	// YOUR CODE HERE
	var speaker Speaker
	speaker = speakers[0] // Dog
	fmt.Println(speaker.Speak())

	speaker = speakers[1] // Cat
	fmt.Println(speaker.Speak())

	speaker = speakers[2] // Robot
	fmt.Println(speaker.Speak())
}

// ===== INTERFACE WITH MULTIPLE METHODS =====

// TODO: Define a Shape interface with two methods: Area() float64 and Perimeter() float64
// YOUR CODE HERE
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TODO: Define a Rectangle struct with Width and Height fields (both float64)
// YOUR CODE HERE
type Rectangle struct {
	Width  float64
	Height float64
}

// TODO: Implement Area() method for Rectangle - returns width * height
// YOUR CODE HERE
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// TODO: Implement Perimeter() method for Rectangle - returns 2 * (width + height)
// YOUR CODE HERE
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// TODO: Define a Circle struct with Radius field (float64)
// YOUR CODE HERE
type Circle struct {
	Radius float64
}

// TODO: Implement Area() method for Circle - returns π * radius²
// YOUR CODE HERE
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// TODO: Implement Perimeter() method for Circle - returns 2 * π * radius
// YOUR CODE HERE
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func demonstrateMultiMethodInterface() {
	fmt.Println("\n=== Multi-Method Interface ===")

	// TODO: Create a slice of Shape containing Rectangle and Circle instances
	// Example: Rectangle{Width: 10, Height: 5}, Circle{Radius: 3}
	// YOUR CODE HERE
	shapes := []Shape{Rectangle{Width: 10, Height: 5}, Circle{Radius: 3}}

	// TODO: Loop through shapes, calculate area and perimeter for each
	// Keep track of total area and print individual shape info
	// YOUR CODE HERE
	for _, shape := range shapes {
		fmt.Printf("Shape: %T\n", shape)
		fmt.Printf("Area: %f\n", shape.Area())
		fmt.Printf("Perimeter: %f\n", shape.Perimeter())
	}

	// TODO: Print the total area of all shapes
	// YOUR CODE HERE
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}

	fmt.Printf("Total area: %f\n", totalArea)
}

// ===== INTERFACE SATISFACTION =====

// TODO: Define a Walker interface with Walk() string method
// YOUR CODE HERE
type Walker interface {
	Walk() string
}

// TODO: Define a Runner interface with Run() string method
// YOUR CODE HERE
type Runner interface {
	Run() string
}

// TODO: Define an Athlete struct with Name field
// YOUR CODE HERE
type Athlete struct {
	Name string
}

// TODO: Implement Walk() method for Athlete - returns "[name] is walking"
// YOUR CODE HERE
func (a Athlete) Walk() string {
	return fmt.Sprintf("%s is walking", a.Name)
}

// TODO: Implement Run() method for Athlete - returns "[name] is running"
// YOUR CODE HERE
func (a Athlete) Run() string {
	return fmt.Sprintf("%s is running", a.Name)
}

func demonstrateInterfaceSatisfaction() {
	fmt.Println("\n=== Interface Satisfaction ===")

	// TODO: Create an Athlete instance
	// YOUR CODE HERE
	alanWalker := Athlete{Name: "Alan Walker"}
	// TODO: Assign the athlete to Walker and Runner interface variables
	// Demonstrate that one type can satisfy multiple interfaces
	// YOUR CODE HERE
	var walker Walker = alanWalker
	var runner Runner = alanWalker

	// TODO: Call the interface methods
	// YOUR CODE HERE
	fmt.Println(walker.Walk())
	fmt.Println(runner.Run())

	// TODO: Pass the athlete to functions that accept interfaces
	// YOUR CODE HERE
	makeWalk(alanWalker)
	makeRun(alanWalker)
}

// TODO: Implement makeWalk function that accepts Walker interface
// Should print "Making someone walk: [result of Walk()]"
// YOUR CODE HERE
func makeWalk(walker Walker) {
	fmt.Println("Making someone walk:", walker.Walk())
}

// TODO: Implement makeRun function that accepts Runner interface
// Should print "Making someone run: [result of Run()]"
// YOUR CODE HERE
func makeRun(runner Runner) {
	fmt.Println("Making someone run:", runner.Run())
}

// ===== EMPTY INTERFACE =====

func demonstrateEmptyInterface() {
	fmt.Println("\n=== Empty Interface (interface{}) ===")

	// TODO: Declare an interface{} variable
	// YOUR CODE HERE
	var void interface{}

	// TODO: Assign different types to it (int, string, slice, struct)
	// Print the value and type using %v and %T
	// YOUR CODE HERE
	void = 10
	fmt.Printf("Value: %v, Type: %T\n", void, void)
	void = "Hello"
	fmt.Printf("Value: %v, Type: %T\n", void, void)
	void = []int{1, 2, 3}
	fmt.Printf("Value: %v, Type: %T\n", void, void)

	// TODO: Create a slice of interface{} containing mixed types
	// (int, string, bool, float64, Dog)
	// YOUR CODE HERE
	voidCollection := []interface{}{10, "Hello", []int{1, 2, 3}, Dog{Name: "Rex"}}

	// TODO: Loop through the mixed slice and print each item's value and type
	// YOUR CODE HERE
	for _, item := range voidCollection {
		fmt.Printf("Value: %v, Type: %T\n", item, item)
	}
}

// ===== TYPE ASSERTIONS =====

func demonstrateTypeAssertions() {
	fmt.Println("\n=== Type Assertions ===")

	// TODO: Create an interface{} variable with a string value
	// YOUR CODE HERE
	var void interface{} = "虚空"

	// TODO: Perform a safe type assertion using the "comma ok" idiom
	// Check if the data is a string and print accordingly
	// YOUR CODE HERE
	if str, ok := void.(string); ok {
		fmt.Println("Got string:", str)
	}

	// TODO: Perform an unsafe type assertion (without checking)
	// Note: This can panic if the type is wrong!
	// YOUR CODE HERE
	// num := void.(int)
	// fmt.Println("Got int:", num)

	// TODO: Test type assertions with different values
	// Call testTypeAssertion with int, string, float64, and Dog
	// YOUR CODE HERE
	testTypeAssertion(10)
	testTypeAssertion("Hello")
	testTypeAssertion(3.14)
	testTypeAssertion(Dog{Name: "Rex"})
}

// TODO: Implement testTypeAssertion function that takes interface{}
// Use multiple if statements with type assertions to handle:
// - string: print "Got string: [value]"
// - int: print "Got int: [value]"
// - float64: print "Got float64: [value]"
// - default: print "Got unknown type: [type] with value [value]"
// YOUR CODE HERE

func testTypeAssertion(some interface{}) {
	if str, ok := some.(string); ok {
		fmt.Println("Got string:", str)
	} else if num, ok := some.(int); ok {
		fmt.Println("Got int:", num)
	} else if float, ok := some.(float64); ok {
		fmt.Println("Got float64:", float)
	} else {
		fmt.Printf("Got unknown type: %T with value %v\n", some, some)
	}
}

// ===== TYPE SWITCHES =====

func demonstrateTypeSwitches() {
	fmt.Println("\n=== Type Switches ===")

	// TODO: Create a slice of interface{} with mixed types:
	// int, string, float64, bool, Dog, []int, nil
	// YOUR CODE HERE

	stuffBox := []interface{}{10, "Hello", 3.14, true, Dog{Name: "Rex"}, []int{1, 2, 3}, nil}

	// TODO: Loop through the values and call describeType on each
	// YOUR CODE HERE
	for _, item := range stuffBox {
		describeType(item)
	}
}

// TODO: Implement describeType function using a type switch
// Handle these cases:
// - nil: print "nil value"
// - int: print "integer: [value]"
// - string: print "string: [value] (length: [length])"
// - float64: print "float: [value]"
// - bool: print "boolean: [value]"
// - Dog: print "dog: [name]"
// - []int: print "int slice: [value] (length: [length])"
// - default: print "unknown type: [type] with value [value]"
// YOUR CODE HERE
func describeType(some interface{}) {

	// valid way
	switch v := some.(type) {
	case nil:
		fmt.Println("nil value")
	case int:
		fmt.Printf("integer: %d\n", v)
	case string:
		fmt.Printf("string: %s (length: %d)\n", v, len(v))
	case float64:
		fmt.Printf("float: %f\n", v)
	case bool:
		fmt.Printf("boolean: %t\n", v)
	case Dog:
		fmt.Printf("dog: %s\n", v.Name)
	case []int:
		fmt.Printf("int slice: %v (length: %d)\n", v, len(v))
	default:
		fmt.Printf("Got unknown type: %T with value %v\n", v, v)
	}
}

// ===== INTERFACE COMPOSITION =====

// TODO: Define a Reader interface with Read() string method
// YOUR CODE HERE
type Reader interface {
	Read() string
}

// TODO: Define a Writer interface with Write(data string) error method
// YOUR CODE HERE
type Writer interface {
	Write(data string) error
}

// TODO: Define a Closer interface with Close() error method
// YOUR CODE HERE
type Closer interface {
	Close() error
}

// TODO: Define a ReadWriter interface that embeds Reader and Writer
// YOUR CODE HERE
type ReadWriter interface {
	Reader
	Writer
}

// TODO: Define a ReadWriteCloser interface that embeds Reader, Writer, and Closer
// YOUR CODE HERE
type ReadWriteCloser interface {
	ReadWriter
	Closer
}

// TODO: Define a File struct with fields: name (string), content (string), position (int), closed (bool)
// YOUR CODE HERE
type File struct {
	Name     string
	Content  string
	Position int
	Closed   bool
}

// TODO: Implement Read() method for File (use pointer receiver)
// Return empty string if closed, otherwise return content from position to end
// Set position to end of content after reading
// YOUR CODE HERE
func (f *File) Read() string {
	if f.Closed {
		return ""
	}

	content := f.Content[f.Position:]
	f.Position = len(f.Content)
	return content
}

// TODO: Implement Write(data string) error method for File (use pointer receiver)
// Return error if file is closed, otherwise append data to content
// YOUR CODE HERE
func (f *File) Write(data string) error {
	if f.Closed {
		return errors.New("file is closed")
	}

	f.Content += data
	return nil
}

// TODO: Implement Close() error method for File (use pointer receiver)
// Return error if already closed, otherwise set closed to true and print message
// YOUR CODE HERE
func (f *File) Close() error {
	if f.Closed {
		return errors.New("file is already closed")
	}

	f.Closed = true
	return nil
}

func demonstrateInterfaceComposition() {
	fmt.Println("\n=== Interface Composition ===")

	// TODO: Create a File instance with initial content
	// YOUR CODE HERE
	file := File{
		Name:     "test.txt",
		Content:  "Hello, world!",
		Position: 0,
		Closed:   false,
	}

	// TODO: Assign the file to individual interface variables (Reader, Writer, Closer)
	// YOUR CODE HERE
	var reader Reader = &file
	var writer Writer = &file
	// var closer Closer = &file

	// TODO: Assign the file to composed interface variables (ReadWriter, ReadWriteCloser)
	// YOUR CODE HERE
	var readWriter ReadWriter = &file
	var readWriteCloser ReadWriteCloser = &file

	// TODO: Use the individual interfaces to read and write
	// YOUR CODE HERE
	reader.Read()
	writer.Write("Hello, world!")

	// TODO: Use the composed interfaces by calling helper functions
	// YOUR CODE HERE
	processReadWriter(readWriter)
	processReadWriteCloser(readWriteCloser)

	// TODO: Close the file
	// YOUR CODE HERE
	readWriteCloser.Close()
}

// TODO: Implement processReadWriter function that takes ReadWriter interface
// Should write some data and then read and print the content
// YOUR CODE HERE
func processReadWriter(rw ReadWriter) {
	rw.Write("Hello, world!")
	fmt.Println(rw.Read())
}

// TODO: Implement processReadWriteCloser function that takes ReadWriteCloser interface
// Should write some data and then read and print the content
// Don't close the file in this function
// YOUR CODE HERE
func processReadWriteCloser(rwc ReadWriteCloser) {
	rwc.Write("Hello, world!")
	fmt.Println(rwc.Read())
}

// ===== INTERFACE VALUES =====

// TODO: Define a Printer interface with Print() string method
// YOUR CODE HERE
type Printer interface {
	Print() string
}

// TODO: Define a DocumentBasic struct with Title and Content fields (both string)
// YOUR CODE HERE
type DocumentBasic struct {
	Title   string
	Content string
}

// TODO: Implement Print() method for DocumentBasic
// Should return "Document: [title]\nContent: [content]"
// YOUR CODE HERE

func (d DocumentBasic) Print() string {
	return fmt.Sprintf("Document: [%s]\nContent: [%s]", d.Title, d.Content)
}

func demonstrateInterfaceValues() {
	fmt.Println("\n=== Interface Values ===")

	// TODO: Declare a Printer interface variable (zero value)
	// YOUR CODE HERE
	var printer Printer

	// TODO: Check and print if the interface is nil
	// YOUR CODE HERE
	if printer == nil {
		fmt.Println("Printer is nil")
	}

	// TODO: Assign a DocumentBasic value to the interface
	// YOUR CODE HERE
	printer = DocumentBasic{Title: "Hello", Content: "World"}

	// TODO: Check if it's nil now and print the type and value
	// YOUR CODE HERE
	if printer == nil {
		fmt.Println("Printer is nil")
	} else {
		fmt.Printf("Printer value: %v, type: %T\n", printer, printer)
	}

	// TODO: Call the Print method on the interface
	// YOUR CODE HERE
	fmt.Println(printer.Print())

	// TODO: Set the interface back to nil and verify
	// YOUR CODE HERE
	printer = nil
	if printer == nil {
		fmt.Println("Printer is nil")
	} else {
		fmt.Printf("Printer value: %v, type: %T\n", printer, printer)
	}

	// Note: Calling methods on nil interface would panic!
	// printer.Print()
}

// ===== INTERFACE BEST PRACTICES =====

func demonstrateBestPractices() {
	fmt.Println("\n=== Interface Best Practices ===")

	// TODO: Print the key interface design principles:
	// - Keep interfaces small and focused
	// - Define interfaces where they're used, not implemented
	// - Use composition to build complex behaviors
	// - Accept interfaces, return structs
	// - Favor many small interfaces over few large ones
	// - Name interfaces by capability (use -er suffix)
	fmt.Println("Keep interfaces small and focused")
	fmt.Println("Define interfaces where they're used, not implemented")
	fmt.Println("Use composition to build complex behaviors")
	fmt.Println("Accept interfaces, return structs")
	fmt.Println("Favor many small interfaces over few large ones")
	fmt.Println("Name interfaces by capability (use -er suffix)")
}

// ===== MAIN DEMO FUNCTION =====

func main() {
	fmt.Println("🔌 Go Interface Basics Practice")
	fmt.Println("===============================")

	demonstrateBasicInterfaces()
	demonstrateMultiMethodInterface()
	demonstrateInterfaceSatisfaction()
	demonstrateEmptyInterface()
	demonstrateTypeAssertions()
	demonstrateTypeSwitches()
	demonstrateInterfaceComposition()
	demonstrateInterfaceValues()
	demonstrateBestPractices()

	fmt.Println("\n✅ Interface basics practice completed!")
	fmt.Println("\n🎯 Key Learning Goals:")
	fmt.Println("- Understand implicit interface satisfaction")
	fmt.Println("- Work with single and multi-method interfaces")
	fmt.Println("- Use empty interface{} for generic programming")
	fmt.Println("- Master type assertions and type switches")
	fmt.Println("- Apply interface composition patterns")
	fmt.Println("- Understand interface values and nil handling")
}
