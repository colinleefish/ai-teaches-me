# Week 6: Interfaces & Polymorphism 🔌

Welcome to Go's powerful interface system! Interfaces in Go are implemented implicitly, enabling flexible and composable designs through polymorphism.

## 📚 What You'll Learn

- **Interface Basics**: Definition, implicit implementation, and satisfaction
- **Polymorphism**: Dynamic behavior through interface types
- **Standard Interfaces**: `io.Reader`, `io.Writer`, `fmt.Stringer`, `error`
- **Type Assertions**: Runtime type checking and conversion
- **Interface Composition**: Building complex contracts
- **Empty Interface**: Working with `interface{}` and `any`

## 🎯 Learning Objectives

After completing this week, you'll be able to:
- [ ] Define and implement interfaces implicitly
- [ ] Use polymorphism for flexible code design
- [ ] Work with standard library interfaces effectively
- [ ] Perform type assertions and type switches safely
- [ ] Compose interfaces for complex behaviors
- [ ] Apply interface design principles
- [ ] Debug interface satisfaction issues

## 📁 Files in This Section

- `interface_basics.go` - Interface definition, implementation, and basics
- `polymorphism.go` - Polymorphic behavior and design patterns
- `common_interfaces.go` - Standard library interfaces and usage
- `interface_practice.go` - **YOUR PRACTICE FILE** - Hands-on exercises

## ⚡ Key Differences from Python

### Interface Implementation
```python
# Python - Explicit inheritance required
from abc import ABC, abstractmethod

class Animal(ABC):
    @abstractmethod
    def speak(self):
        pass

class Dog(Animal):  # Must explicitly inherit
    def speak(self):
        return "Woof!"

# Go - Implicit implementation (duck typing)
type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {  // Automatically satisfies Animal
    return "Woof!"
}
```

### Polymorphism
```python
# Python - Class-based polymorphism
def make_sound(animal: Animal):
    return animal.speak()

dog = Dog()
sound = make_sound(dog)  # Works because Dog inherits from Animal

# Go - Interface-based polymorphism
func makeSound(animal Animal) string {
    return animal.Speak()
}

dog := Dog{Name: "Buddy"}
sound := makeSound(dog)  // Works because Dog implements Animal
```

### Multiple Interface Implementation
```python
# Python - Multiple inheritance
class Swimmer(ABC):
    @abstractmethod
    def swim(self):
        pass

class Duck(Animal, Swimmer):  # Multiple inheritance
    def speak(self):
        return "Quack!"
    
    def swim(self):
        return "Swimming..."

# Go - Automatic satisfaction of multiple interfaces
type Swimmer interface {
    Swim() string
}

type Duck struct {
    Name string
}

func (d Duck) Speak() string { return "Quack!" }
func (d Duck) Swim() string { return "Swimming..." }

// Duck automatically satisfies both Animal and Swimmer
```

## 🚀 Getting Started

1. **Read**: Start with `interface_basics.go` to understand interface fundamentals
2. **Study**: Work through `polymorphism.go` for design patterns
3. **Explore**: Check out `common_interfaces.go` for standard library usage
4. **Practice**: Complete exercises in `interface_practice.go`

## 💡 Interface Design Principles

1. **Keep interfaces small**: "The bigger the interface, the weaker the abstraction"
2. **Define interfaces where they're used**: Not where they're implemented
3. **Accept interfaces, return structs**: Enable flexibility for callers
4. **Use composition**: Combine small interfaces for complex behaviors
5. **Favor many small interfaces**: Over few large ones
6. **Name interfaces by capability**: Use `-er` suffix (Reader, Writer, Closer)

## 🧪 Core Concepts to Master

### 1. Implicit Interface Satisfaction
```go
type Writer interface {
    Write([]byte) (int, error)
}

type FileWriter struct {
    filename string
}

// This method automatically makes FileWriter implement Writer
func (fw FileWriter) Write(data []byte) (int, error) {
    fmt.Printf("Writing to %s: %s\n", fw.filename, string(data))
    return len(data), nil
}

// No explicit "implements" keyword needed!
var w Writer = FileWriter{filename: "output.txt"}
```

### 2. Interface Composition
```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader  // Embedded interface
    Writer  // Embedded interface
}

// Any type implementing both Read and Write automatically satisfies ReadWriter
```

### 3. Empty Interface and Type Assertions
```go
func process(data interface{}) {
    // Type assertion
    if str, ok := data.(string); ok {
        fmt.Printf("Got string: %s\n", str)
    } else if num, ok := data.(int); ok {
        fmt.Printf("Got number: %d\n", num)
    }
}

// Type switch
func processSwitch(data interface{}) {
    switch v := data.(type) {
    case string:
        fmt.Printf("String: %s\n", v)
    case int:
        fmt.Printf("Integer: %d\n", v)
    case []byte:
        fmt.Printf("Bytes: %s\n", string(v))
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

### 4. Standard Library Integration
```go
import (
    "fmt"
    "io"
    "strings"
)

// Your type implementing standard interfaces
type UpperCaseReader struct {
    reader io.Reader
}

func (ucr UpperCaseReader) Read(p []byte) (n int, err error) {
    n, err = ucr.reader.Read(p)
    for i := 0; i < n; i++ {
        if p[i] >= 'a' && p[i] <= 'z' {
            p[i] = p[i] - 'a' + 'A'
        }
    }
    return
}

// Usage with any io.Reader
func demo() {
    original := strings.NewReader("hello world")
    upper := UpperCaseReader{reader: original}
    
    data, _ := io.ReadAll(upper)
    fmt.Println(string(data)) // "HELLO WORLD"
}
```

## 🎯 Success Criteria

You'll know you've mastered interfaces when you can:
- Design clean, small interfaces for your domain
- Implement polymorphic behavior naturally
- Work seamlessly with standard library interfaces
- Use type assertions and switches effectively
- Compose interfaces for complex behaviors
- Apply the "accept interfaces, return structs" principle
- Debug interface satisfaction issues confidently

## 🔗 What's Next

After mastering interfaces and polymorphism, you'll advance to **Phase 4: Concurrency** where you'll learn Go's most distinctive feature - goroutines and channels for concurrent programming!

Let's build flexible, polymorphic Go applications! 🐹 