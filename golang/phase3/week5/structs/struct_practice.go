package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ===== EXERCISE 1: BASIC STRUCT DEFINITION =====

// TODO: Define a Book struct with the following fields:
// - Title (string)
// - Author (string)
// - Pages (int)
// - PublishedYear (int)
// - InStock (bool)
// YOUR CODE HERE

type Book struct {
	Title         string
	Author        string
	Pages         int
	PublishedYear int
	InStock       bool
}

func exerciseBasicStruct() {
	fmt.Println("=== Exercise 1: Basic Struct Definition ===")

	// TODO: Create a Book instance with sample data
	// YOUR CODE HERE
	sampleBook := Book{
		Title:         "The Three Body Problem",
		Author:        "Liu Cixin",
		Pages:         300,
		PublishedYear: 2008,
		InStock:       true,
	}

	// TODO: Print the book information
	// YOUR CODE HERE
	fmt.Printf("Book: %+v\n", sampleBook)

	// TODO: Modify some fields and print again
	// YOUR CODE HERE
	sampleBook.InStock = false
	fmt.Printf("Book: %+v\n", sampleBook)
}

// ===== EXERCISE 2: METHODS AND RECEIVERS =====

// TODO: Add methods to your Book struct:
// 1. String() string - returns formatted book information
// 2. IsClassic() bool - returns true if published before 1980
// 3. UpdateStock(inStock bool) - updates the InStock field (use pointer receiver)
// 4. PageCategory() string - returns "Short" (<200), "Medium" (200-500), "Long" (>500)
// YOUR CODE HERE

func (b Book) String() string {
	return fmt.Sprintf(
		"Title: %s, Author: %s, Pages: %d, Published Year: %d, In Stock: %t",
		b.Title, b.Author, b.Pages, b.PublishedYear, b.InStock,
	)
}

func (b Book) IsClassic() bool {
	return b.PublishedYear < 1980
}

func (b *Book) UpdateStock(inStock bool) *Book {
	b.InStock = inStock
	return b
}

func (b *Book) PageCategory() string {
	switch {
	case b.Pages < 200:
		return "Short"
	case b.Pages < 500:
		return "Medium"
	case b.Pages > 500:
		return "Long"
	default:
		return "Unknown"
	}
}

func exerciseMethodsAndReceivers() {
	fmt.Println("\n=== Exercise 2: Methods and Receivers ===")

	// TODO: Create a book and test all your methods
	// YOUR CODE HERE

	sampleBook := Book{
		Title:         "The Three Body Problem",
		Author:        "Liu Cixin",
		Pages:         300,
		PublishedYear: 2008,
		InStock:       true,
	}

	fmt.Printf("Book: %+v\n", sampleBook)

	// TODO: Test the UpdateStock method (should modify the original)
	// YOUR CODE HERE
	sampleBook.UpdateStock(false)
	fmt.Printf("Book: %+v\n", sampleBook)
}

// ===== EXERCISE 3: STRUCT EMBEDDING =====

// TODO: Define a MediaItem struct with common fields:
// - Title (string)
// - Creator (string)
// - Year (int)
// YOUR CODE HERE

type MediaItem struct {
	Title   string
	Creator string
	Year    int
}

// TODO: Define a Movie struct that embeds MediaItem and adds:
// - Duration (int) // in minutes
// - Genre (string)
// YOUR CODE HERE

type Movie struct {
	MediaItem
	Duration int
	Genre    string
}

// TODO: Define a Song struct that embeds MediaItem and adds:
// - Duration (int) // in seconds
// - Album (string)
// YOUR CODE HERE

type Song struct {
	MediaItem
	Duration int
	Album    string
}

// TODO: Add methods:
// - (MediaItem) Info() string - returns basic info
// - (Movie) Play() string - returns "Playing movie: [title]"
// - (Song) Play() string - returns "Playing song: [title]"
// YOUR CODE HERE

func (m MediaItem) Info() string {
	return fmt.Sprintf("Title: %s, Creator: %s, Year: %d", m.Title, m.Creator, m.Year)
}

func (m Movie) play() string {
	return fmt.Sprintf("Playing movie: %s", m.Title)
}

func (s Song) play() string {
	return fmt.Sprintf("Playing song: %s", s.Title)
}

func exerciseStructEmbedding() {
	fmt.Println("\n=== Exercise 3: Struct Embedding ===")

	// TODO: Create instances of Movie and Song
	// YOUR CODE HERE

	ironMan := Movie{
		MediaItem: MediaItem{Title: "Iron Man", Creator: "Marvel", Year: 2008},
		Duration:  126,
		Genre:     "Action",
	}

	youRaiseMeUp := Song{
		MediaItem: MediaItem{Title: "You Raise Me Up", Creator: "Westlife", Year: 2004},
		Duration:  240,
		Album:     "Unbreakable",
	}

	// TODO: Demonstrate accessing embedded fields directly
	// YOUR CODE HERE

	fmt.Printf("Movie: %s\n", ironMan.Title)
	fmt.Printf("Song: %s\n", youRaiseMeUp.Title)

	// TODO: Demonstrate calling embedded methods
	// YOUR CODE HERE
	fmt.Printf("Movie: %s\n", ironMan.Info())
	fmt.Printf("Song: %s\n", youRaiseMeUp.Info())

	// TODO: Demonstrate method shadowing if you implemented Play() differently
	// YOUR CODE HERE
	fmt.Printf("Movie: %s\n", ironMan.play())
	fmt.Printf("Song: %s\n", youRaiseMeUp.play())
}

// ===== EXERCISE 4: STRUCT TAGS AND JSON =====

// TODO: Define a Product struct with JSON tags:
// - ID (int) -> "id"
// - Name (string) -> "product_name"
// - Price (float64) -> "price" with string conversion
// - Category (string) -> "category"
// - InStock (bool) -> "in_stock"
// - Description (string) -> "description" with omitempty
// - InternalNotes (string) -> excluded from JSON (use "-")
// YOUR CODE HERE

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"product_name"`
	Price         string `json:"price"`
	Category      string `json:"category"`
	InStock       bool   `json:"in_stock"`
	Description   string `json:"description,omitempty"`
	InternalNotes string `json:"-"`
}

func exerciseStructTags() {
	fmt.Println("\n=== Exercise 4: Struct Tags and JSON ===")

	// TODO: Create a Product instance with sample data
	// YOUR CODE HERE
	shampoo := Product{
		ID:            1,
		Name:          "Shampoo",
		Price:         "10.99",
		Category:      "Personal Care",
		InStock:       true,
		Description:   "Shampoo for all hair types",
		InternalNotes: "This is a test note",
	}

	// TODO: Marshal to JSON and print
	// YOUR CODE HERE
	jsonData, err := json.Marshal(shampoo)
	if err != nil {
		fmt.Printf("Error marshalling to JSON: %v\n", err)
	}
	fmt.Printf("JSON: %+s\n", string(jsonData))

	// TODO: Unmarshal from this JSON string:

	// YOUR CODE HERE

	jsonData2 := []byte(`{"id":2,"product_name":"Laptop","price":"999.99","category":"Electronics","in_stock":true}`)
	var product2 Product
	err = json.Unmarshal(jsonData2, &product2)
	if err != nil {
		fmt.Printf("Error unmarshalling from JSON: %v\n", err)
	}

	// TODO: Print the unmarshaled product
	// YOUR CODE HERE

	fmt.Printf("Unmarshaled product: %+v\n", product2)
}

// ===== EXERCISE 5: COMPOSITION PATTERN =====

// TODO: Define interfaces:
// - Writer interface with Write(content string) error method
// - Reader interface with Read() (string, error) method
// YOUR CODE HERE

type Writer interface {
	Write(content string) error
}

type Reader interface {
	Read() (string, error)
}

// TODO: Define a FileManager struct that embeds both Writer and Reader interfaces
// and adds a Filename field
// YOUR CODE HERE

type FileManager struct {
	Writer
	Reader
	Filename string
}

// TODO: Implement concrete types:
// - ConsoleWriter that prints to console
// - FileWriter that simulates writing to a file
// - MemoryReader that returns predefined content
// YOUR CODE HERE

type ConsoleWriter struct {
	Content string
}

func (c ConsoleWriter) Write(content string) error {
	c.Content = content

	fmt.Printf("Writing to console: %s\n", content)
	return nil
}

type FileWriter struct {
	Filename string
}

func (f FileWriter) Write(content string) error {
	fmt.Printf("Writing to file: %s\n", content)
	return nil
}

type MemoryReader struct {
	Content string
}

func (m MemoryReader) Read() (string, error) {
	return m.Content, nil
}

func exerciseComposition() {
	fmt.Println("\n=== Exercise 5: Composition Pattern ===")

	// TODO: Create a FileManager with ConsoleWriter and MemoryReader
	// YOUR CODE HERE
	fm := FileManager{
		Writer:   &ConsoleWriter{Content: "Hello, World!"},
		Reader:   &MemoryReader{Content: "This is a test content"},
		Filename: "test.txt",
	}

	// TODO: Test reading and writing
	// YOUR CODE HERE
	fm.Write("This is a test content")
	content, err := fm.Read()
	if err != nil {
		fmt.Printf("Error reading from file: %v\n", err)
	}
	fmt.Printf("Read content: %s\n", content)

	// TODO: Create another FileManager with FileWriter
	// YOUR CODE HERE
	fm2 := FileManager{
		Writer:   &FileWriter{Filename: "test2.txt"},
		Reader:   &MemoryReader{Content: "This is a test content 2"},
		Filename: "test2.txt",
	}

	fm2.Write("This is a test content 2")
	content2, err := fm2.Read()
	if err != nil {
		fmt.Printf("Error reading from file: %v\n", err)
	}
	fmt.Printf("Read content: %s\n", content2)
}

// ===== EXERCISE 6: ADVANCED STRUCT PATTERNS =====

// TODO: Define a Customer struct with:
// - ID (int)
// - Name (string)
// - Email (string)
// - CreatedAt (time.Time)
// - Address (embedded struct with Street, City, ZipCode)
// - Orders (slice of Order structs)
// YOUR CODE HERE

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Order struct {
	ID        int
	Total     float64
	Items     []string
	OrderDate time.Time
}

type Customer struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
	Address   Address
	Orders    []Order
}

// TODO: Define an Order struct with:
// - ID (int)
// - Total (float64)
// - Items (slice of strings)
// - OrderDate (time.Time)
// YOUR CODE HERE

// TODO: Add methods:
// - (Customer) AddOrder(order Order) - adds order to customer's orders
// - (Customer) TotalSpent() float64 - calculates total from all orders
// - (Customer) RecentOrders(days int) []Order - returns orders from last N days
// YOUR CODE HERE

func (c *Customer) AddOrder(order Order) {
	c.Orders = append(c.Orders, order)
}

func (c *Customer) TotalSpent() float64 {
	total := 0.0
	for _, order := range c.Orders {
		total += order.Total
	}
	return total
}

func (c *Customer) RecentOrders(days int) []Order {
	recentOrders := []Order{}
	for _, order := range c.Orders {
		if order.OrderDate.After(time.Now().AddDate(0, 0, -days)) {
			recentOrders = append(recentOrders, order)
		}
	}
	return recentOrders
}

func exerciseAdvancedPatterns() {
	fmt.Println("\n=== Exercise 6: Advanced Struct Patterns ===")

	// TODO: Create a customer with embedded address
	// YOUR CODE HERE
	customer := Customer{
		ID:        1,
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now(),
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			ZipCode: "12345",
		},
		Orders: []Order{},
	}

	// TODO: Create several orders and add them to the customer
	// YOUR CODE HERE
	order1 := Order{
		ID:        1,
		Total:     100.0,
		Items:     []string{"Item 1", "Item 2"},
		OrderDate: time.Now(),
	}

	order2 := Order{
		ID:        2,
		Total:     200.0,
		Items:     []string{"Item 3", "Item 4"},
		OrderDate: time.Now().AddDate(0, 0, -15),
	}

	customer.AddOrder(order1)
	customer.AddOrder(order2)

	// TODO: Calculate total spent
	// YOUR CODE HERE

	fmt.Printf("Total spent: %f\n", customer.TotalSpent())

	// TODO: Get recent orders (use time.Now().AddDate(0, 0, -7) for 7 days ago)
	// YOUR CODE HERE
	recentOrders := customer.RecentOrders(7)
	fmt.Printf("Recent orders: %+v\n", recentOrders)
}

// ===== EXERCISE 7: CONSTRUCTOR FUNCTIONS =====

// TODO: Create constructor functions:
// - NewBook(title, author string, pages, year int) (*Book, error)
//   - Validate that title and author are not empty
//   - Validate that pages > 0 and year > 1400
// - NewCustomer(name, email string) (*Customer, error)
//   - Validate email contains "@"
//   - Initialize empty orders slice and set CreatedAt to now
// YOUR CODE HERE

func NewBook(title, author string, pages, year int) (*Book, error) {
	if title == "" || author == "" {
		return nil, fmt.Errorf("title and author are required")
	}
	if pages <= 0 || year < 1400 {
		return nil, fmt.Errorf("invalid pages or year")
	}

	return &Book{
		Title:         title,
		Author:        author,
		Pages:         pages,
		PublishedYear: year,
		InStock:       true,
	}, nil
}

func NewCustomer(name, email string) (*Customer, error) {
	if !strings.Contains(email, "@") {
		return nil, fmt.Errorf("invalid email")
	}

	return &Customer{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		Orders:    []Order{},
	}, nil
}

func exerciseConstructors() {
	fmt.Println("\n=== Exercise 7: Constructor Functions ===")

	// TODO: Test valid book creation
	// YOUR CODE HERE
	book, err := NewBook("The Great Gatsby", "F. Scott Fitzgerald", 180, 1925)
	if err != nil {
		fmt.Printf("Error creating book: %v\n", err)
	}
	fmt.Printf("Book: %+v\n", book)

	// TODO: Test invalid book creation (should return error)
	// YOUR CODE HERE
	book2, err := NewBook("", "F. Scott Fitzgerald", 180, 1925)
	if err != nil {
		fmt.Printf("Error creating book: %v\n", err)
	}
	fmt.Printf("Book: %+v\n", book2)

	// TODO: Test valid customer creation
	// YOUR CODE HERE
	customer, err := NewCustomer("John Doe", "john.doe@example.com")
	if err != nil {
		fmt.Printf("Error creating customer: %v\n", err)
	}
	fmt.Printf("Customer: %+v\n", customer)

	// TODO: Test invalid customer creation (should return error)
	// YOUR CODE HERE
	customer2, err := NewCustomer("John Doe", "john.doe.com")
	if err != nil {
		fmt.Printf("Error creating customer: %v\n", err)
	}
	fmt.Printf("Customer: %+v\n", customer2)
}

// ===== MAIN FUNCTION =====

func main() {
	fmt.Println("🏗️ Week 5: Structs & Methods Practice")
	fmt.Println("=====================================")

	exerciseBasicStruct()
	exerciseMethodsAndReceivers()
	exerciseStructEmbedding()
	exerciseStructTags()
	exerciseComposition()
	exerciseAdvancedPatterns()
	exerciseConstructors()

	fmt.Println("\n✅ Week 5 exercises completed!")
	fmt.Println("\n💡 Key Learnings:")
	fmt.Println("- Structs define custom data types")
	fmt.Println("- Methods attach behavior to types")
	fmt.Println("- Pointer receivers modify originals, value receivers work on copies")
	fmt.Println("- Embedding enables composition over inheritance")
	fmt.Println("- Struct tags provide metadata for serialization")
	fmt.Println("- Constructor functions enable validation and initialization")
	fmt.Println("- Composition patterns create flexible, reusable designs")
}
