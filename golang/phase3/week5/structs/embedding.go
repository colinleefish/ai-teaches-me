package main

import (
	"fmt"
	"log"
	"time"
)

// ===== BASIC EMBEDDING =====

// Base type for embedding examples
type PersonEmbed struct {
	Name string
	Age  int
}

func (p PersonEmbed) Greet() string {
	return fmt.Sprintf("Hello, I'm %s and I'm %d years old", p.Name, p.Age)
}

func (p PersonEmbed) String() string {
	return fmt.Sprintf("PersonEmbed{Name: %s, Age: %d}", p.Name, p.Age)
}

// Embedded type
type Employee struct {
	PersonEmbed // Embedded field - promotes PersonEmbed's fields and methods
	ID          int
	Position    string
	Salary      float64
}

func (e Employee) Work() string {
	return fmt.Sprintf("%s is working as a %s", e.Name, e.Position)
}

// Method with same name - shadows the embedded method
func (e Employee) String() string {
	return fmt.Sprintf("Employee{ID: %d, Name: %s, Position: %s}", e.ID, e.Name, e.Position)
}

func demonstrateBasicEmbedding() {
	fmt.Println("=== Basic Embedding ===")

	// Creating embedded struct
	emp := Employee{
		PersonEmbed: PersonEmbed{Name: "Alice", Age: 30},
		ID:          123,
		Position:    "Software Engineer",
		Salary:      75000,
	}

	// Direct access to embedded fields
	fmt.Printf("Employee name: %s\n", emp.Name) // Promoted from PersonEmbed
	fmt.Printf("Employee age: %d\n", emp.Age)   // Promoted from PersonEmbed
	fmt.Printf("Employee ID: %d\n", emp.ID)

	// Access to embedded methods
	fmt.Printf("Greeting: %s\n", emp.Greet()) // PersonEmbed's method

	// Own methods
	fmt.Printf("Work: %s\n", emp.Work())

	// Method shadowing
	fmt.Printf("String (Employee): %s\n", emp.String())                // Employee's String method
	fmt.Printf("String (PersonEmbed): %s\n", emp.PersonEmbed.String()) // PersonEmbed's String method explicitly
}

// ===== MULTIPLE EMBEDDING =====

type Address struct {
	Street  string
	City    string
	Country string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

type Contact struct {
	Email string
	Phone string
}

func (c Contact) ContactInfo() string {
	return fmt.Sprintf("Email: %s, Phone: %s", c.Email, c.Phone)
}

// Multiple embeddings
type Customer struct {
	PersonEmbed // Embedded
	Address     // Embedded
	Contact     // Embedded
	ID          int
	Active      bool
}

func (c Customer) Profile() string {
	return fmt.Sprintf("Customer %s (%d): %s | %s",
		c.Name, c.ID, c.FullAddress(), c.ContactInfo())
}

func demonstrateMultipleEmbedding() {
	fmt.Println("\n=== Multiple Embedding ===")

	customer := Customer{
		PersonEmbed: PersonEmbed{Name: "Bob", Age: 35},
		Address:     Address{Street: "123 Main St", City: "New York", Country: "USA"},
		Contact:     Contact{Email: "bob@example.com", Phone: "+1-555-0123"},
		ID:          456,
		Active:      true,
	}

	// Access fields from all embedded types
	fmt.Printf("Name: %s\n", customer.Name)      // From Person
	fmt.Printf("City: %s\n", customer.City)      // From Address
	fmt.Printf("Email: %s\n", customer.Email)    // From Contact
	fmt.Printf("Customer ID: %d\n", customer.ID) // Own field

	// Access methods from all embedded types
	fmt.Printf("Greeting: %s\n", customer.Greet())      // From Person
	fmt.Printf("Address: %s\n", customer.FullAddress()) // From Address
	fmt.Printf("Contact: %s\n", customer.ContactInfo()) // From Contact
	fmt.Printf("Profile: %s\n", customer.Profile())     // Own method
}

// ===== EMBEDDING WITH CONFLICTS =====

type Writer struct {
	Name string
}

func (w Writer) Write() string {
	return fmt.Sprintf("%s is writing", w.Name)
}

type Reader struct {
	Name string
}

func (r Reader) Read() string {
	return fmt.Sprintf("%s is reading", r.Name)
}

// This would cause a conflict!
// type Conflicted struct {
//     Writer
//     Reader  // Both have Name field - ambiguous!
// }

// Solution: Use named embedding
type Editor struct {
	Writer Writer
	Reader Reader
	ID     int
}

func (e Editor) Edit() string {
	return fmt.Sprintf("Editor %d: %s and %s", e.ID, e.Writer.Write(), e.Reader.Read())
}

// Alternative: Embed one, include other as named field
type Journalist struct {
	Writer               // Embedded - promotes fields/methods
	ReadingSkills Reader // Named field - no promotion
	ID            int
}

func demonstrateEmbeddingConflicts() {
	fmt.Println("\n=== Embedding Conflicts ===")

	// Named embedding to avoid conflicts
	editor := Editor{
		Writer: Writer{Name: "Alice"},
		Reader: Reader{Name: "Alice"},
		ID:     1,
	}

	fmt.Printf("Writer name: %s\n", editor.Writer.Name)
	fmt.Printf("Reader name: %s\n", editor.Reader.Name)
	fmt.Printf("Edit: %s\n", editor.Edit())

	// Mixed approach
	journalist := Journalist{
		Writer:        Writer{Name: "Bob"},
		ReadingSkills: Reader{Name: "Bob"},
		ID:            2,
	}

	fmt.Printf("Journalist name: %s\n", journalist.Name)      // Promoted from Writer
	fmt.Printf("Write: %s\n", journalist.Write())             // Promoted from Writer
	fmt.Printf("Read: %s\n", journalist.ReadingSkills.Read()) // Explicit access
}

// ===== INTERFACE SATISFACTION THROUGH EMBEDDING =====

type Printer interface {
	Print() string
}

type Scanner interface {
	Scan() string
}

type Device interface {
	Printer
	Scanner
	Status() string
}

// Basic implementations
type BasicPrinter struct {
	Model string
}

func (bp BasicPrinter) Print() string {
	return fmt.Sprintf("%s is printing", bp.Model)
}

type BasicScanner struct {
	Model string
}

func (bs BasicScanner) Scan() string {
	return fmt.Sprintf("%s is scanning", bs.Model)
}

// Combined device using embedding
type MultiFunction struct {
	BasicPrinter // Embeds Printer interface implementation
	BasicScanner // Embeds Scanner interface implementation
	Name         string
}

func (mf MultiFunction) Status() string {
	return fmt.Sprintf("MultiFunction %s is ready", mf.Name)
}

// The MultiFunction automatically satisfies Device interface!

func demonstrateInterfaceSatisfaction() {
	fmt.Println("\n=== Interface Satisfaction Through Embedding ===")

	mf := MultiFunction{
		BasicPrinter: BasicPrinter{Model: "HP LaserJet"},
		BasicScanner: BasicScanner{Model: "Epson Scanner"},
		Name:         "Office Pro 3000",
	}

	// Can be used as Device interface
	var device Device = mf

	fmt.Printf("Print: %s\n", device.Print())   // From embedded BasicPrinter
	fmt.Printf("Scan: %s\n", device.Scan())     // From embedded BasicScanner
	fmt.Printf("Status: %s\n", device.Status()) // Own method

	// Individual interface satisfaction
	var printer Printer = mf
	var scanner Scanner = mf

	fmt.Printf("As Printer: %s\n", printer.Print())
	fmt.Printf("As Scanner: %s\n", scanner.Scan())
}

// ===== COMPOSITION PATTERNS =====

// Strategy pattern using embedding
type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Printf("[CONSOLE] %s\n", message)
}

type FileLogger struct {
	Filename string
}

func (fl FileLogger) Log(message string) {
	fmt.Printf("[FILE:%s] %s\n", fl.Filename, message)
}

// Service with embedded logger strategy
type UserService struct {
	Logger Logger // Strategy composition
	users  map[int]string
}

func NewUserService(logger Logger) *UserService {
	return &UserService{
		Logger: logger,
		users:  make(map[int]string),
	}
}

func (us *UserService) CreateUser(id int, name string) {
	us.users[id] = name
	us.Logger.Log(fmt.Sprintf("Created user: %d - %s", id, name))
}

func (us *UserService) GetUser(id int) (string, bool) {
	name, exists := us.users[id]
	if exists {
		us.Logger.Log(fmt.Sprintf("Retrieved user: %d - %s", id, name))
	} else {
		us.Logger.Log(fmt.Sprintf("User not found: %d", id))
	}
	return name, exists
}

// Decorator pattern using embedding
type TimestampLogger struct {
	Logger Logger // Embedded logger to decorate
}

func (tl TimestampLogger) Log(message string) {
	timestamped := fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), message)
	tl.Logger.Log(timestamped)
}

func demonstrateCompositionPatterns() {
	fmt.Println("\n=== Composition Patterns ===")

	// Strategy pattern
	consoleService := NewUserService(ConsoleLogger{})
	consoleService.CreateUser(1, "Alice")
	consoleService.GetUser(1)

	fmt.Println()

	fileService := NewUserService(FileLogger{Filename: "users.log"})
	fileService.CreateUser(2, "Bob")
	fileService.GetUser(2)

	fmt.Println()

	// Decorator pattern - adding timestamp
	timestampService := NewUserService(TimestampLogger{
		Logger: ConsoleLogger{},
	})
	timestampService.CreateUser(3, "Charlie")

	// Multiple decorators
	timestampFileService := NewUserService(TimestampLogger{
		Logger: FileLogger{Filename: "timestamped.log"},
	})
	timestampFileService.CreateUser(4, "Diana")
}

// ===== EMBEDDED INTERFACES =====

type DatabaseConnection interface {
	Connect() error
	Close() error
	Query(sql string) ([]string, error)
}

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

// Service with embedded interfaces
type DataService struct {
	DatabaseConnection // Embedded interface
	Cache              // Embedded interface
	Name               string
}

func (ds DataService) FetchUserData(userID string) ([]string, error) {
	// Try cache first
	if cached, found := ds.Get("user:" + userID); found {
		log.Printf("Cache hit for user %s", userID)
		return cached.([]string), nil
	}

	// Fetch from database
	data, err := ds.Query("SELECT * FROM users WHERE id = " + userID)
	if err != nil {
		return nil, err
	}

	// Cache the result
	ds.Set("user:"+userID, data)
	return data, nil
}

// Mock implementations for demo
type MockDB struct{}

func (m MockDB) Connect() error {
	fmt.Println("Connected to database")
	return nil
}

func (m MockDB) Close() error {
	fmt.Println("Closed database connection")
	return nil
}

func (m MockDB) Query(sql string) ([]string, error) {
	fmt.Printf("Executing query: %s\n", sql)
	return []string{"user_data_1", "user_data_2"}, nil
}

type MockCache struct {
	data map[string]interface{}
}

func NewMockCache() *MockCache {
	return &MockCache{data: make(map[string]interface{})}
}

func (m *MockCache) Get(key string) (interface{}, bool) {
	value, exists := m.data[key]
	return value, exists
}

func (m *MockCache) Set(key string, value interface{}) {
	m.data[key] = value
	fmt.Printf("Cached data for key: %s\n", key)
}

func demonstrateEmbeddedInterfaces() {
	fmt.Println("\n=== Embedded Interfaces ===")

	dataService := DataService{
		DatabaseConnection: MockDB{},
		Cache:              NewMockCache(),
		Name:               "UserDataService",
	}

	// Connect to database
	dataService.Connect()

	// Fetch user data (will hit database and cache)
	data1, err := dataService.FetchUserData("123")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("First fetch: %v\n", data1)
	}

	// Fetch same user data (will hit cache)
	data2, err := dataService.FetchUserData("123")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Second fetch: %v\n", data2)
	}

	// Clean up
	dataService.Close()
}

// ===== MAIN DEMO FUNCTION =====

func runEmbeddingDemo() {
	fmt.Println("🔗 Go Struct Embedding and Composition Tutorial")
	fmt.Println("===============================================")

	demonstrateBasicEmbedding()
	demonstrateMultipleEmbedding()
	demonstrateEmbeddingConflicts()
	demonstrateInterfaceSatisfaction()
	demonstrateCompositionPatterns()
	demonstrateEmbeddedInterfaces()

	fmt.Println("\n✅ Embedding and composition concepts covered!")
	fmt.Println("\n🎯 Key Points:")
	fmt.Println("- Embedding promotes fields and methods from embedded types")
	fmt.Println("- Multiple embedding is possible but watch for conflicts")
	fmt.Println("- Method shadowing allows overriding embedded behavior")
	fmt.Println("- Embedding enables interface satisfaction through composition")
	fmt.Println("- Use composition patterns like Strategy and Decorator")
	fmt.Println("- Embedded interfaces provide flexible dependency injection")
	fmt.Println("- Prefer composition over inheritance for flexible design")
}
