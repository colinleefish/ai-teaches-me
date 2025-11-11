package main

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// ===== EXERCISE 1: BASIC INTERFACE IMPLEMENTATION =====

// TODO: Define a Vehicle interface with the following methods:
// - Start() string
// - Stop() string
// - GetSpeed() int
// YOUR CODE HERE
type Vehicle interface {
	Start() string
	Stop() string
	GetSpeed() int
}

// TODO: Implement the Vehicle interface for:
// 1. Car struct with fields: Brand, Model, Speed
// 2. Bicycle struct with fields: Type, Speed
// 3. Airplane struct with fields: Model, Speed, Altitude
// Each should have appropriate Start(), Stop(), and GetSpeed() implementations
// YOUR CODE HERE
type Car struct {
	Brand string
	Model string
	Speed int
}

type Bicycle struct {
	Type  string
	Speed int
}

type Airplane struct {
	Model    string
	Speed    int
	Altitude int
}

func (c Car) Start() string {
	return fmt.Sprintf("Car %s %s started", c.Brand, c.Model)
}

func (c Car) Stop() string {
	return fmt.Sprintf("Car %s %s stopped", c.Brand, c.Model)
}

func (c Car) GetSpeed() int {
	return c.Speed
}

func (b Bicycle) Start() string {
	return fmt.Sprintf("Bicycle %s started", b.Type)
}

func (b Bicycle) Stop() string {
	return fmt.Sprintf("Bicycle %s stopped", b.Type)
}

func (b Bicycle) GetSpeed() int {
	return b.Speed
}

func (a Airplane) Start() string {
	return fmt.Sprintf("Airplane %s started", a.Model)
}

func (a Airplane) Stop() string {
	return fmt.Sprintf("Airplane %s stopped", a.Model)
}

func (a Airplane) GetSpeed() int {
	return a.Speed
}

func exerciseBasicInterface() {
	fmt.Println("=== Exercise 1: Basic Interface Implementation ===")

	// TODO: Create instances of Car, Bicycle, and Airplane
	// YOUR CODE HERE
	car := Car{
		Brand: "Toyota",
		Model: "Corolla",
		Speed: 100,
	}
	bicycle := Bicycle{
		Type:  "Mountain",
		Speed: 20,
	}
	airplane := Airplane{
		Model:    "Boeing 747",
		Speed:    500,
		Altitude: 10000,
	}

	// TODO: Store them in a slice of Vehicle interfaces
	// YOUR CODE HERE
	vehicles := []Vehicle{car, bicycle, airplane}

	// TODO: Loop through and call Start(), Stop(), and GetSpeed() on each
	// YOUR CODE HERE
	for _, vehicle := range vehicles {
		fmt.Println(vehicle.Start())
		fmt.Println(vehicle.GetSpeed())
		fmt.Println(vehicle.Stop())
	}
}

// ===== EXERCISE 2: INTERFACE COMPOSITION =====

// TODO: Define these interfaces:
// - Flyer interface with Fly() string method
// - Swimmer interface with Swim() string method
// - AmphibiousVehicle interface that combines Flyer and Swimmer
// YOUR CODE HERE
type Flyer interface {
	Fly() string
}

type Swimmer interface {
	Swim() string
}

type AmphibiousVehicle interface {
	Flyer
	Swimmer
}

// TODO: Implement a SeaPlane struct that satisfies AmphibiousVehicle
// It should have Name, Speed, and Altitude fields
// YOUR CODE HERE

type SeaPlane struct {
	Name     string
	Speed    int
	Altitude int
}

func (s SeaPlane) Fly() string {
	return fmt.Sprintf("SeaPlane %s is flying at an altitude of %d km", s.Name, s.Altitude)
}

func (s SeaPlane) Swim() string {
	return fmt.Sprintf("SeaPlane %s is swimming at %d km/h", s.Name, s.Speed)
}

func exerciseInterfaceComposition() {
	fmt.Println("\n=== Exercise 2: Interface Composition ===")

	// TODO: Create a SeaPlane instance
	// YOUR CODE HERE
	seaPlane := SeaPlane{
		Name:     "SeaPlane 1",
		Speed:    100,
		Altitude: 1000,
	}
	// TODO: Use it as different interface types (Flyer, Swimmer, AmphibiousVehicle)
	// YOUR CODE HERE

	var flyer Flyer = seaPlane
	var swimmer Swimmer = seaPlane
	var amphibious AmphibiousVehicle = seaPlane

	fmt.Println(flyer.Fly())
	fmt.Println(swimmer.Swim())
	fmt.Println(amphibious.Fly())
	fmt.Println(amphibious.Swim())

	// TODO: Create functions that accept each interface type and test them
	// YOUR CODE HERE

	fmt.Println("Printing SeaPlane as different interface types:")
	printFly := func(f Flyer) {
		fmt.Println(f.Fly())
	}
	printSwim := func(s Swimmer) {
		fmt.Println(s.Swim())
	}
	printAmphibious := func(a AmphibiousVehicle) {
		fmt.Printf("AmphibiousVehicle: %v\n", a.Fly())
		fmt.Printf("AmphibiousVehicle: %v\n", a.Swim())
	}
	printFly(seaPlane)
	printSwim(seaPlane)
	printAmphibious(seaPlane)
}

// ===== EXERCISE 3: IMPLEMENTING STANDARD INTERFACES =====

// TODO: Create a Product struct with:
// - ID (int)
// - Name (string)
// - Price (float64)
// - Category (string)
// YOUR CODE HERE
type Product struct {
	ID       int
	Name     string
	Price    float64
	Category string
}

// TODO: Implement fmt.Stringer for Product
// Format: "Product: [Name] ($[Price]) - Category: [Category]"
// YOUR CODE HERE
func (p Product) String() string {
	return fmt.Sprintf("Product: %s ($%.2f) - Category: %s", p.Name, p.Price, p.Category)
}

// TODO: Create a ProductList type ([]Product) and implement sort.Interface
// Sort by Price in ascending order
// YOUR CODE HERE
type ProductList []Product

func (p ProductList) Len() int {
	return len(p)
}

func (p ProductList) Less(i, j int) bool {
	return p[i].Price < p[j].Price
}

func (p ProductList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func exerciseStandardInterfaces() {
	fmt.Println("\n=== Exercise 3: Standard Interfaces ===")

	// TODO: Create a slice of products with sample data
	// YOUR CODE HERE
	products := ProductList{
		{ID: 1, Name: "Laptop", Price: 1000.0, Category: "Electronics"},
		{ID: 2, Name: "Smartphone", Price: 800.0, Category: "Electronics"},
		{ID: 3, Name: "Shirt", Price: 20.0, Category: "Clothing"},
		{ID: 4, Name: "Jeans", Price: 50.0, Category: "Clothing"},
		{ID: 5, Name: "Coffee Maker", Price: 150.0, Category: "Kitchen"},
		{ID: 6, Name: "Toaster", Price: 80.0, Category: "Kitchen"},
	}

	// TODO: Print products before sorting (using Stringer)
	// YOUR CODE HERE
	fmt.Println("Products before sorting:")
	for _, product := range products {
		fmt.Println(product)
	}

	// TODO: Sort products using sort.Sort()
	// YOUR CODE HERE
	sort.Sort(products)

	// TODO: Print products after sorting
	// YOUR CODE HERE
	fmt.Println("Products after sorting:")
	for _, product := range products {
		fmt.Println(product)
	}
}

// ===== EXERCISE 4: CUSTOM READER/WRITER =====

// TODO: Implement a ROT13Reader that wraps an io.Reader
// It should apply ROT13 encoding to the text as it's read
// ROT13: A→N, B→O, C→P, ..., N→A, O→B, P→C, etc.
// YOUR CODE HERE
type ROT13Reader struct {
	Reader io.Reader
}

func (r ROT13Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	for i := range p {
		if p[i] >= 'A' && p[i] <= 'Z' {
			p[i] = 'A' + (p[i]-'A'+13)%26
		} else if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = 'a' + (p[i]-'a'+13)%26
		}
	}
	return n, err
}

func (r ROT13Reader) String() string {
	return fmt.Sprintf("ROT13Reader: %v", r.Reader)
}

// TODO: Implement a CountingWriter that wraps an io.Writer
// It should count the number of bytes written and provide a Count() method
// YOUR CODE HERE
type CountingWriter struct {
	Writer io.Writer
	count  int
}

func (c *CountingWriter) Write(p []byte) (n int, err error) {
	n, err = c.Writer.Write(p)
	c.count += n
	return n, err
}

func (c *CountingWriter) Count() int {
	return c.count
}

func NewCountingWriter(w io.Writer) (*CountingWriter, error) {
	return &CountingWriter{Writer: w}, nil
}

func exerciseCustomReaderWriter() {
	fmt.Println("\n=== Exercise 4: Custom Reader/Writer ===")

	// TODO: Test ROT13Reader with a string input
	// Example: "Hello World" should become "Uryyb Jbeyq"
	// YOUR CODE HERE
	rot13Reader := ROT13Reader{
		Reader: strings.NewReader("Hello World"),
	}
	rot13Reader.Read([]byte("Hello World"))
	fmt.Println(rot13Reader.String())

	// TODO: Test CountingWriter by writing some data and checking count
	// YOUR CODE HERE

	buf := bytes.Buffer{}
	cw, err := NewCountingWriter(&buf)
	if err != nil {
		fmt.Println("Error creating CountingWriter:", err)
		return
	}
	cw.Write([]byte("Hello World"))
	fmt.Println("Bytes written:", cw.Count())

	// TODO: Chain them together: write ROT13 encoded data and count bytes
	// YOUR CODE HERE
	newRot13Reader := ROT13Reader{
		Reader: strings.NewReader("Hello World"),
	}
	newBuf := bytes.Buffer{}
	newWriter := &CountingWriter{
		Writer: &newBuf,
	}
	io.Copy(newWriter, newRot13Reader)
	fmt.Println("Bytes written:", newWriter.Count())
	fmt.Println("Data written:", newBuf.String())
}

// ===== EXERCISE 5: ERROR INTERFACE =====

// TODO: Create custom error types:
// 1. ValidationError with Field and Reason
// 2. NetworkError with URL, StatusCode, and Message
// 3. TimeoutError with Operation and Duration
// All should implement the error interface
// YOUR CODE HERE

// TODO: Create a function processRequest(url string, timeout time.Duration) error
// It should return different error types based on conditions:
// - ValidationError if url is empty
// - TimeoutError if timeout < 1 second
// - NetworkError if url doesn't start with "http"
// - nil if everything is ok
// YOUR CODE HERE

func exerciseErrorInterface() {
	fmt.Println("\n=== Exercise 5: Error Interface ===")

	// TODO: Test processRequest with different inputs to trigger each error type
	// YOUR CODE HERE

	// TODO: Use type assertions or type switches to handle each error type differently
	// YOUR CODE HERE
}

// ===== EXERCISE 6: POLYMORPHISM PATTERN =====

// TODO: Design a payment system with:
// - PaymentProcessor interface with Process(amount float64) error method
// - CreditCardProcessor struct with CardNumber field
// - PayPalProcessor struct with Email field
// - BankTransferProcessor struct with AccountNumber field
// YOUR CODE HERE

// TODO: Create a PaymentService struct that uses a PaymentProcessor
// It should have a SetProcessor method to change processors dynamically
// YOUR CODE HERE

func exercisePolymorphism() {
	fmt.Println("\n=== Exercise 6: Polymorphism Pattern ===")

	// TODO: Create different payment processors
	// YOUR CODE HERE

	// TODO: Create a payment service and test with different processors
	// YOUR CODE HERE

	// TODO: Show that the same PaymentService can use different processors
	// YOUR CODE HERE
}

// ===== EXERCISE 7: INTERFACE DESIGN PATTERNS =====

// TODO: Implement the Observer pattern:
// - Observer interface with Update(message string) method
// - Subject interface with Attach/Detach/Notify methods
// - NewsPublisher struct implementing Subject
// - Subscriber struct implementing Observer with Name field
// YOUR CODE HERE

func exerciseObserverPattern() {
	fmt.Println("\n=== Exercise 7: Observer Pattern ===")

	// TODO: Create a news publisher
	// YOUR CODE HERE

	// TODO: Create several subscribers and attach them
	// YOUR CODE HERE

	// TODO: Publish some news and show all subscribers receive it
	// YOUR CODE HERE

	// TODO: Detach one subscriber and publish again
	// YOUR CODE HERE
}

// ===== EXERCISE 8: ADVANCED INTERFACE USAGE =====

// TODO: Create a generic data transformer system:
// - Transformer interface with Transform(data interface{}) (interface{}, error)
// - StringUpperTransformer that converts strings to uppercase
// - NumberDoubler that doubles numeric values (int, float64)
// - SliceReverser that reverses any slice
// Use type assertions and type switches
// YOUR CODE HERE

// TODO: Create a Pipeline struct that chains multiple transformers
// It should have Add(transformer Transformer) and Process(data interface{}) methods
// YOUR CODE HERE

func exerciseAdvancedInterfaces() {
	fmt.Println("\n=== Exercise 8: Advanced Interface Usage ===")

	// TODO: Create individual transformers
	// YOUR CODE HERE

	// TODO: Create a pipeline and add transformers
	// YOUR CODE HERE

	// TODO: Test the pipeline with different data types
	// YOUR CODE HERE
}

// ===== EXERCISE 9: INTERFACE BEST PRACTICES =====

// TODO: Design a logging system following interface best practices:
// - Small, focused interfaces (Logger, Formatter, Writer)
// - Interface composition for complex behaviors
// - Accept interfaces, return structs principle
// YOUR CODE HERE

func exerciseBestPractices() {
	fmt.Println("\n=== Exercise 9: Interface Best Practices ===")

	// TODO: Demonstrate small, focused interfaces
	// YOUR CODE HERE

	// TODO: Show interface composition
	// YOUR CODE HERE

	// TODO: Apply "accept interfaces, return structs" principle
	// YOUR CODE HERE
}

// ===== EXERCISE 10: REAL-WORLD SCENARIO =====

// TODO: Build a simple HTTP-like system:
// - Handler interface with Handle(request Request) Response
// - Request struct with Method, URL, Body
// - Response struct with StatusCode, Body
// - Different handlers: StaticFileHandler, APIHandler, NotFoundHandler
// - Router that selects appropriate handler based on URL pattern
// YOUR CODE HERE

func exerciseRealWorldScenario() {
	fmt.Println("\n=== Exercise 10: Real-World Scenario ===")

	// TODO: Create different handlers
	// YOUR CODE HERE

	// TODO: Create a router and register handlers
	// YOUR CODE HERE

	// TODO: Simulate handling different requests
	// YOUR CODE HERE
}

// ===== HELPER FUNCTIONS =====

// TODO: Implement helper functions as needed for your exercises
// YOUR CODE HERE

// ===== MAIN FUNCTION =====

func main() {
	fmt.Println("🔌 Week 6: Interfaces & Polymorphism Practice")
	fmt.Println("==============================================")

	// exerciseBasicInterface()
	// exerciseInterfaceComposition()
	// exerciseStandardInterfaces()
	exerciseCustomReaderWriter()
	// exerciseErrorInterface()
	// exercisePolymorphism()
	// exerciseObserverPattern()
	// exerciseAdvancedInterfaces()
	// exerciseBestPractices()
	// exerciseRealWorldScenario()

	fmt.Println("\n✅ Week 6 exercises completed!")
	fmt.Println("\n💡 Key Learnings:")
	fmt.Println("- Interfaces enable polymorphism and flexible design")
	fmt.Println("- Implicit implementation makes interfaces natural in Go")
	fmt.Println("- Standard library interfaces integrate seamlessly")
	fmt.Println("- Small, focused interfaces are easier to implement")
	fmt.Println("- Interface composition builds complex behaviors")
	fmt.Println("- Type assertions and switches handle runtime types")
	fmt.Println("- Design patterns become natural with interfaces")
	fmt.Println("- 'Accept interfaces, return structs' enables flexibility")
}
