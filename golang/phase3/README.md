# Phase 3: Structs & Interfaces 🏗️

Welcome to Go's type system! This phase covers Go's approach to object-oriented programming through structs and interfaces. You'll learn composition over inheritance and Go's powerful interface system.

## 📚 What You'll Learn

### Week 5: Structs & Methods
- **Struct definition**: Creating custom data types
- **Methods**: Attaching behavior to types
- **Receivers**: Value vs pointer receivers
- **Embedding**: Composition patterns
- **Struct tags**: Metadata for serialization

### Week 6: Interfaces & Polymorphism  
- **Interface basics**: Implicit implementation
- **Common interfaces**: `io.Reader`, `io.Writer`, `fmt.Stringer`
- **Polymorphism**: Dynamic behavior through interfaces
- **Interface composition**: Building complex contracts
- **Type switches**: Runtime type checking

## 🎯 Learning Objectives

After completing this phase, you'll be able to:
- [ ] Design and implement struct-based data models
- [ ] Create methods with appropriate receiver types
- [ ] Use struct embedding for composition
- [ ] Implement and use interfaces effectively
- [ ] Apply polymorphism patterns in Go
- [ ] Work with common standard library interfaces
- [ ] Design clean, composable APIs

## 📁 Phase Structure

```
phase3/
├── week5/          # Structs & Methods
│   └── structs/
│       ├── README.md
│       ├── struct_basics.go       # Struct definition and instantiation
│       ├── methods.go             # Methods and receivers
│       ├── embedding.go           # Struct embedding patterns
│       ├── struct_tags.go         # Tags for JSON, validation, etc.
│       └── struct_practice.go     # Practice exercises
│
└── week6/          # Interfaces & Polymorphism
    └── interfaces/
        ├── README.md
        ├── interface_basics.go    # Interface definition and implementation
        ├── polymorphism.go        # Dynamic behavior patterns
        ├── common_interfaces.go   # Standard library interfaces
        └── interface_practice.go  # Practice exercises
```

## ⚡ Key Differences from Python

### Object-Oriented Programming
```python
# Python - Classes with inheritance
class Animal:
    def __init__(self, name):
        self.name = name
    
    def speak(self):
        pass

class Dog(Animal):
    def speak(self):
        return f"{self.name} says Woof!"

# Go - Structs with composition
type Animal struct {
    Name string
}

type Dog struct {
    Animal  // Embedding (composition)
}

func (d Dog) Speak() string {
    return fmt.Sprintf("%s says Woof!", d.Name)
}
```

### Interfaces
```python
# Python - Explicit inheritance
from abc import ABC, abstractmethod

class Speaker(ABC):
    @abstractmethod
    def speak(self):
        pass

class Dog(Speaker):  # Must inherit
    def speak(self):
        return "Woof!"

# Go - Implicit implementation
type Speaker interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {  // Automatically implements Speaker
    return "Woof!"
}
```

## 🚀 Getting Started

1. **Week 5**: Start with `week5/structs/README.md`
2. **Study examples**: Read through tutorial files
3. **Practice**: Complete exercises in practice files
4. **Week 6**: Move to interfaces and polymorphism
5. **Build projects**: Apply concepts in real applications

## 💡 Go Philosophy

**Composition over Inheritance:**
- Go favors embedding structs over class hierarchies
- Interfaces are satisfied implicitly (duck typing)
- Small, focused interfaces are preferred
- "Accept interfaces, return structs"

**Interface Design Principles:**
- Keep interfaces small and focused
- Define interfaces where they're used, not where they're implemented
- Use composition to build complex behaviors
- Prefer many small interfaces over few large ones

## 🎯 Success Criteria

You'll know you've mastered this phase when you can:
- Design struct hierarchies using embedding
- Choose between value and pointer receivers appropriately
- Create and implement interfaces naturally
- Use type assertions and type switches effectively
- Apply polymorphism patterns in Go code
- Work confidently with standard library interfaces

## 🔗 What's Next

After mastering structs and interfaces, you'll advance to **Phase 4: Concurrency** where you'll learn Go's most powerful feature - goroutines and channels!

Let's build some robust, composable Go applications! 🐹 