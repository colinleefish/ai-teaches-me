# Week 1 - Variables & Types: Go's Type System 🔢

Welcome to Go's type system! This section covers variables, basic types, and how they differ from Python's dynamic typing.

## 📋 Learning Objectives

By the end of this section, you will:
- [ ] Master Go's variable declaration syntax
- [ ] Understand Go's static type system
- [ ] Work with basic Go types (int, float, string, bool)
- [ ] Use arrays, slices, and maps effectively
- [ ] Understand zero values and type conversions
- [ ] Compare Go's typing to Python's dynamic typing

## 🎯 Exercises Overview

### Exercise 1: Variable Declarations
**File**: `variables.go`
- `var` keyword declarations
- Short variable declarations with `:=`
- Multiple variable declarations
- Variable scope

### Exercise 2: Basic Types
**File**: `types.go`
- Numeric types (int, float64, etc.)
- String operations
- Boolean logic
- Zero values

### Exercise 3: Collections
**File**: `collections.go`
- Arrays vs Slices
- Maps (Go's dictionaries)
- String and slice operations

### Exercise 4: Type Conversions
**File**: `conversions.go`
- Type casting
- String conversions
- Interface{} usage

## 💡 Key Concepts

### Static vs Dynamic Typing
- **Python**: `x = 5` (type inferred at runtime)
- **Go**: `var x int = 5` (type checked at compile time)

### Zero Values
Go initializes variables to their zero values:
- `int`: 0
- `float64`: 0.0
- `string`: ""
- `bool`: false
- `slice`: nil
- `map`: nil

### Declaration Patterns
```go
// Long form
var name string = "Alice"
var age int = 30

// Short form (inside functions only)
name := "Alice"
age := 30

// Multiple declarations
var (
    name string = "Alice"
    age  int    = 30
)
```

## 🐍 Python vs Go Reference

| Concept | Python | Go |
|---------|--------|-----|
| **Variables** | `x = 10` | `var x int = 10` or `x := 10` |
| **Lists** | `[1, 2, 3]` | `[]int{1, 2, 3}` (slice) |
| **Dictionaries** | `{"key": "value"}` | `map[string]string{"key": "value"}` |
| **Type Check** | `type(x)` | `reflect.TypeOf(x)` |
| **String Format** | `f"Hello {name}"` | `fmt.Sprintf("Hello %s", name)` |

## 💻 Exercise Instructions

### Exercise 1: Variable Declarations
1. Open `variables.go`
2. Complete variable declaration examples
3. Test different scopes and declarations
4. Run: `go run variables.go`

### Exercise 2: Basic Types
1. Open `types.go`
2. Implement type demonstrations
3. Test zero values and operations
4. Run: `go run types.go`

### Exercise 3: Collections
1. Open `collections.go`
2. Practice with arrays, slices, and maps
3. Compare with Python equivalents
4. Run: `go run collections.go`

### Exercise 4: Type Conversions
1. Open `conversions.go`
2. Implement type casting examples
3. Handle conversion errors
4. Run: `go run conversions.go`

## 🧪 Testing Your Work

Run the test suite:
```bash
go test
```

Check specific exercises:
```bash
go test -v -run TestVariables
go test -v -run TestTypes
go test -v -run TestCollections
go test -v -run TestConversions
```

## 📊 Self-Assessment

Rate your understanding (1-5 scale):
- [ ] Variable declarations (var vs :=): ___/5
- [ ] Basic types and zero values: ___/5
- [ ] Arrays, slices, and maps: ___/5
- [ ] Type conversions: ___/5
- [ ] Static vs dynamic typing: ___/5

**Target**: All items should be 4/5 or higher before proceeding.

## 🔧 Common Pitfalls

### 1. Short Declaration Scope
```go
// ❌ Wrong: := only works inside functions
var x := 10  // Error!

// ✅ Correct:
var x int = 10
// or inside a function:
func main() {
    x := 10  // OK
}
```

### 2. Slice vs Array
```go
// Array (fixed size)
var arr [3]int = [3]int{1, 2, 3}

// Slice (dynamic)
var slice []int = []int{1, 2, 3}
```

### 3. Map Initialization
```go
// ❌ Wrong: using uninitialized map
var m map[string]int
m["key"] = 1  // Runtime panic!

// ✅ Correct:
var m map[string]int = make(map[string]int)
m["key"] = 1  // OK
```

## 🚀 Ready to Start?

1. Complete each exercise file in order
2. Run tests to verify your solutions
3. Compare with Python equivalents
4. Move to `../../week2/control-flow/` when ready

Let's master Go's type system! 🎯 