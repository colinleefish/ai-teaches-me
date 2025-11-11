package main

import (
	"fmt"
	"time"
)

// ===== STRUCT DEFINITION =====

// Basic struct definition
type Person struct {
	Name string
	Age  int
}

// Struct with various field types
type User struct {
	ID       int
	Name     string
	Email    string
	Active   bool
	Created  time.Time
	Tags     []string
	Settings map[string]string
}

// Struct with unexported (private) fields
type BankAccount struct {
	Owner   string
	balance float64 // unexported (lowercase)
	pin     int     // unexported
}

// Empty struct (zero memory)
type Signal struct{}

func demonstrateStructDefinition() {
	fmt.Println("=== Struct Definition ===")

	// Zero value of struct (all fields get their zero values)
	var person Person
	fmt.Printf("Zero value person: %+v\n", person) // {Name: Age:0}

	var user User
	fmt.Printf("Zero value user: %+v\n", user)
	// {ID:0 Name: Email: Active:false Created:0001-01-01 00:00:00 +0000 UTC Tags:[] Settings:map[]}

	// Empty struct
	var signal Signal
	fmt.Printf("Empty struct: %+v\n", signal) // {}
}

// ===== STRUCT INSTANTIATION =====

func demonstrateInstantiation() {
	fmt.Println("\n=== Struct Instantiation ===")

	// Method 1: Struct literal with positional values
	person1 := Person{"Alice", 30}
	fmt.Printf("Positional: %+v\n", person1)

	// Method 2: Struct literal with named fields (recommended)
	person2 := Person{
		Name: "Bob",
		Age:  25,
	}
	fmt.Printf("Named fields: %+v\n", person2)

	// Method 3: Partial initialization (other fields get zero values)
	person3 := Person{Name: "Charlie"}
	fmt.Printf("Partial init: %+v\n", person3) // {Name:Charlie Age:0}

	// Method 4: Using new() - returns pointer to zero value
	person4 := new(Person)
	fmt.Printf("Using new(): %+v\n", person4) // &{Name: Age:0}

	// Method 5: Address of struct literal - returns pointer
	person5 := &Person{
		Name: "Diana",
		Age:  28,
	}
	fmt.Printf("Address of literal: %+v\n", person5) // &{Name:Diana Age:28}

	// Method 6: Declare then assign
	var person6 Person
	person6.Name = "Eve"
	person6.Age = 35
	fmt.Printf("Declare then assign: %+v\n", person6)
}

// ===== ACCESSING STRUCT FIELDS =====

func demonstrateFieldAccess() {
	fmt.Println("\n=== Field Access ===")

	user := User{
		ID:    1,
		Name:  "Alice",
		Email: "alice@example.com",
		Tags:  []string{"admin", "active"},
		Settings: map[string]string{
			"theme": "dark",
			"lang":  "en",
		},
	}

	// Reading fields
	fmt.Printf("User ID: %d\n", user.ID)
	fmt.Printf("User Name: %s\n", user.Name)
	fmt.Printf("User Tags: %v\n", user.Tags)

	// Modifying fields
	user.Active = true
	user.Created = time.Now()
	fmt.Printf("Updated user: %+v\n", user)

	// Working with slice fields
	user.Tags = append(user.Tags, "verified")
	fmt.Printf("Updated tags: %v\n", user.Tags)

	// Working with map fields
	user.Settings["notifications"] = "enabled"
	fmt.Printf("Updated settings: %v\n", user.Settings)
}

// ===== POINTER ACCESS =====

func demonstratePointerAccess() {
	fmt.Println("\n=== Pointer Access ===")

	// Direct struct
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Direct access: %s\n", person.Name)

	// Pointer to struct
	personPtr := &person

	// Both syntaxes work (Go auto-dereferences)
	fmt.Printf("Pointer access (auto): %s\n", personPtr.Name)
	fmt.Printf("Pointer access (explicit): %s\n", (*personPtr).Name)

	// Modifying through pointer
	personPtr.Age = 31
	fmt.Printf("Modified through pointer: %+v\n", person) // Original is modified

	// Creating pointer with new
	newPersonPtr := new(Person)
	newPersonPtr.Name = "Bob"
	newPersonPtr.Age = 25
	fmt.Printf("New person via pointer: %+v\n", newPersonPtr)
}

// ===== STRUCT COMPARISON =====

func demonstrateComparison() {
	fmt.Println("\n=== Struct Comparison ===")

	person1 := Person{Name: "Alice", Age: 30}
	person2 := Person{Name: "Alice", Age: 30}
	person3 := Person{Name: "Bob", Age: 25}

	// Structs are comparable if all fields are comparable
	fmt.Printf("person1 == person2: %t\n", person1 == person2) // true
	fmt.Printf("person1 == person3: %t\n", person1 == person3) // false

	// Structs with uncomparable fields (slices, maps) can't be compared
	user1 := User{Name: "Alice", Tags: []string{"admin"}}
	user2 := User{Name: "Alice", Tags: []string{"admin"}}

	// This would cause a compile error:
	// fmt.Printf("user1 == user2: %t\n", user1 == user2) // Error!

	// But you can compare individual comparable fields
	fmt.Printf("Same name: %t\n", user1.Name == user2.Name) // true
}

// ===== ANONYMOUS STRUCTS =====

func demonstrateAnonymousStructs() {
	fmt.Println("\n=== Anonymous Structs ===")

	// Anonymous struct - defined and used inline
	config := struct {
		Host string
		Port int
		SSL  bool
	}{
		Host: "localhost",
		Port: 8080,
		SSL:  false,
	}

	fmt.Printf("Config: %+v\n", config)

	// Anonymous struct in slice
	servers := []struct {
		Name string
		URL  string
	}{
		{Name: "prod", URL: "https://api.example.com"},
		{Name: "dev", URL: "http://dev.example.com"},
	}

	for _, server := range servers {
		fmt.Printf("Server %s: %s\n", server.Name, server.URL)
	}

	// Anonymous struct for temporary grouping
	response := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}{
		Success: true,
		Message: "Operation completed",
		Data:    map[string]int{"count": 42},
	}

	fmt.Printf("Response: %+v\n", response)
}

// ===== CONSTRUCTOR FUNCTIONS =====

// Constructor function with validation
func NewUser(name, email string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}

	return &User{
		Name:     name,
		Email:    email,
		Active:   true,
		Created:  time.Now(),
		Tags:     []string{},
		Settings: make(map[string]string),
	}, nil
}

// Constructor with default values
func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		Owner:   owner,
		balance: initialBalance,
		pin:     1234, // default PIN
	}
}

// Getter methods for unexported fields
func (ba *BankAccount) Balance() float64 {
	return ba.balance
}

func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.balance += amount
	}
}

func demonstrateConstructors() {
	fmt.Println("\n=== Constructor Functions ===")

	// Using constructor with validation
	user, err := NewUser("Alice", "alice@example.com")
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}
	fmt.Printf("Created user: %+v\n", user)

	// Invalid user
	_, err = NewUser("", "invalid@example.com")
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
	}

	// Bank account with encapsulation
	account := NewBankAccount("Bob", 1000.0)
	fmt.Printf("Initial balance: %.2f\n", account.Balance())

	account.Deposit(500.0)
	fmt.Printf("After deposit: %.2f\n", account.Balance())
}

// ===== STRUCT COPYING =====

func demonstrateCopying() {
	fmt.Println("\n=== Struct Copying ===")

	original := Person{Name: "Alice", Age: 30}

	// Value copy - creates independent copy
	copy1 := original
	copy1.Age = 31
	fmt.Printf("Original: %+v\n", original) // {Name:Alice Age:30}
	fmt.Printf("Copy: %+v\n", copy1)        // {Name:Alice Age:31}

	// Pointer copy - shares same data
	ptr1 := &original
	ptr2 := ptr1
	ptr2.Age = 32
	fmt.Printf("Original after pointer modification: %+v\n", original) // {Name:Alice Age:32}
	fmt.Printf("Both pointers point to same data: %t\n", ptr1 == ptr2)

	// Complex struct with reference types
	user1 := User{
		Name: "Bob",
		Tags: []string{"admin"},
		Settings: map[string]string{
			"theme": "dark",
		},
	}

	// Shallow copy - reference fields are shared!
	user2 := user1
	user2.Name = "Charlie"            // Independent
	user2.Tags[0] = "user"            // Shared!
	user2.Settings["theme"] = "light" // Shared!

	fmt.Printf("User1 after copy modification: %+v\n", user1)
	fmt.Printf("User2: %+v\n", user2)
	// Notice: Tags and Settings are modified in both!
}

// ===== MAIN DEMO FUNCTION =====

func runStructBasicsDemo() {
	fmt.Println("🏗️ Go Struct Basics Tutorial")
	fmt.Println("=============================")

	demonstrateStructDefinition()
	demonstrateInstantiation()
	demonstrateFieldAccess()
	demonstratePointerAccess()
	demonstrateComparison()
	demonstrateAnonymousStructs()
	demonstrateConstructors()
	demonstrateCopying()

	fmt.Println("\n✅ Struct basics concepts covered!")
	fmt.Println("\n🎯 Key Points:")
	fmt.Println("- Structs define custom types with named fields")
	fmt.Println("- Zero values are useful - design around them")
	fmt.Println("- Multiple instantiation patterns available")
	fmt.Println("- Go auto-dereferences pointers to structs")
	fmt.Println("- Use constructor functions for validation")
	fmt.Println("- Be careful with shallow copying of reference types")
	fmt.Println("- Anonymous structs are great for temporary data structures")
}
