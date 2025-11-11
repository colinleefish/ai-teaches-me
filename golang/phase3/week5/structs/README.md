# Week 5: Structs & Methods 🏗️

Welcome to Go's approach to object-oriented programming! Structs are Go's way of creating custom data types, and methods attach behavior to those types.

## 📚 What You'll Learn

- **Struct Basics**: Definition, instantiation, and zero values
- **Methods**: Attaching functions to types  
- **Receivers**: Value vs pointer receivers and when to use each
- **Embedding**: Composition patterns for code reuse
- **Struct Tags**: Metadata for JSON, validation, and serialization

## 🎯 Learning Objectives

After completing this week, you'll be able to:
- [ ] Define and instantiate structs with various patterns
- [ ] Create methods with appropriate receiver types
- [ ] Understand when to use value vs pointer receivers
- [ ] Use struct embedding for composition
- [ ] Apply struct tags for serialization and validation
- [ ] Design clean, maintainable struct hierarchies

## 📁 Files in This Section

- `struct_basics.go` - Struct definition, instantiation, and basic usage
- `methods.go` - Methods, receivers, and method sets
- `embedding.go` - Struct embedding and composition patterns
- `struct_tags.go` - Tags for JSON, validation, and metadata
- `struct_practice.go` - **YOUR PRACTICE FILE** - Hands-on exercises

## ⚡ Key Differences from Python

### Struct Definition vs Classes
```python
# Python - Classes
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
        self.email = ""  # Optional field

    def greet(self):
        return f"Hello, I'm {self.name}"

# Go - Structs
type Person struct {
    Name  string
    Age   int
    Email string  // Zero value is ""
}

func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.Name)
}
```

### Instantiation Patterns
```python
# Python
person1 = Person("Alice", 30)
person2 = Person(name="Bob", age=25)

# Go - Multiple ways
person1 := Person{"Alice", 30, ""}           // Positional
person2 := Person{Name: "Bob", Age: 25}      // Named fields
person3 := Person{Name: "Charlie"}           // Partial (Age=0, Email="")
person4 := new(Person)                       // Pointer to zero value
person5 := &Person{Name: "Diana", Age: 28}   // Pointer with values
```

### Methods vs Functions
```python
# Python - Methods are bound to classes
class Calculator:
    def add(self, a, b):
        return a + b

calc = Calculator()
result = calc.add(5, 3)

# Go - Methods are functions with receivers
type Calculator struct{}

func (c Calculator) Add(a, b int) int {
    return a + b
}

calc := Calculator{}
result := calc.Add(5, 3)
```

## 🚀 Getting Started

1. **Read**: Start with `struct_basics.go` to understand struct fundamentals
2. **Study**: Work through `methods.go` to learn about receivers
3. **Explore**: Check out `embedding.go` for composition patterns  
4. **Practice**: Use struct tags in `struct_tags.go`
5. **Apply**: Complete exercises in `struct_practice.go`

## 💡 Pro Tips

1. **Zero values**: Structs have useful zero values - design around them
2. **Pointer receivers**: Use for modification or large structs
3. **Value receivers**: Use for small structs and immutable operations
4. **Embedding**: Prefer composition over complex inheritance hierarchies
5. **Exported fields**: Capitalize field names for public access
6. **Constructor functions**: Use `NewTypeName()` functions for validation

## 🧪 Core Concepts to Master

### 1. Struct Definition and Zero Values
```go
type User struct {
    ID       int       // Zero value: 0
    Name     string    // Zero value: ""
    Email    string    // Zero value: ""
    Active   bool      // Zero value: false
    Created  time.Time // Zero value: time.Time{}
}

// Zero value struct is immediately usable
var user User
fmt.Printf("%+v\n", user) // {ID:0 Name: Email: Active:false Created:0001-01-01...}
```

### 2. Value vs Pointer Receivers
```go
type Counter struct {
    value int
}

// Value receiver - receives a copy
func (c Counter) Value() int {
    return c.value
}

// Pointer receiver - receives the original
func (c *Counter) Increment() {
    c.value++
}

// Usage
counter := Counter{}
counter.Increment()  // Modifies original
fmt.Println(counter.Value())  // 1
```

### 3. Embedding for Composition
```go
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string {
    return fmt.Sprintf("Hi, I'm %s", p.Name)
}

type Employee struct {
    Person    // Embedded - promotes Person's fields and methods
    JobTitle  string
    Salary    int
}

// Usage
emp := Employee{
    Person:   Person{Name: "Alice", Age: 30},
    JobTitle: "Developer",
    Salary:   75000,
}

fmt.Println(emp.Name)     // Direct access to embedded field
fmt.Println(emp.Greet())  // Access to embedded method
```

### 4. Struct Tags for Metadata
```go
type User struct {
    ID    int    `json:"id" db:"user_id"`
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}

// JSON serialization uses the tags
data, _ := json.Marshal(User{ID: 1, Name: "Alice", Email: "alice@example.com"})
// Output: {"id":1,"name":"Alice","email":"alice@example.com"}
```

## 🎯 Success Criteria

You'll know you've mastered structs and methods when you can:
- Design struct hierarchies using embedding naturally
- Choose appropriate receiver types (value vs pointer)
- Use struct tags effectively for serialization
- Create clean, composable data models
- Apply zero values in your struct design
- Write idiomatic constructor functions

## 🔗 What's Next

After mastering structs and methods, you'll move on to **Week 6: Interfaces & Polymorphism** where you'll learn Go's powerful interface system and how to achieve polymorphism through implicit interface satisfaction.

Let's build some solid data structures! 🐹 