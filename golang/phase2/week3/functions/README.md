# Week 3: Functions & Methods 📦

Welcome to Go's function system! Go's functions are powerful and explicit, with features like multiple return values that make error handling elegant and clear.

## 📚 What You'll Learn

- **Function syntax**: Go's clean function definition style
- **Multiple returns**: Return multiple values including errors
- **Named returns**: Self-documenting return values
- **Variadic functions**: Functions that accept variable arguments
- **Anonymous functions**: Functions without names and closures
- **Methods**: Functions attached to types
- **Function types**: Functions as first-class values

## 🎯 Learning Objectives

After completing this section, you'll be able to:
- [ ] Write functions with multiple return values
- [ ] Use named returns effectively
- [ ] Create variadic functions for flexible APIs
- [ ] Work with anonymous functions and closures
- [ ] Understand the difference between methods and functions
- [ ] Use function types as variables and parameters
- [ ] Apply Go's function patterns in real-world scenarios

## 📁 Files in This Section

- `basic_functions.go` - Function syntax and multiple returns
- `named_returns.go` - Named returns and naked returns
- `variadic.go` - Functions with variable arguments
- `anonymous.go` - Anonymous functions and closures
- `methods.go` - Methods on types and receivers
- `function_types.go` - Functions as first-class values
- `functions_practice.go` - **YOUR PRACTICE FILE** - Fill in the blanks!

## ⚡ Key Differences from Python

### Function Definitions
```python
# Python
def calculate_area(width, height):
    return width * height

# Go
func calculateArea(width, height float64) float64 {
    return width * height
}
```

### Multiple Return Values
```python
# Python - using tuples
def divide(a, b):
    if b == 0:
        return None, "division by zero"
    return a / b, None

result, error = divide(10, 2)

# Go - built-in multiple returns
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 2)
```

### Error Handling Pattern
```python
# Python - exceptions
def risky_operation():
    if something_wrong:
        raise ValueError("something went wrong")
    return "success"

try:
    result = risky_operation()
except ValueError as e:
    print(f"Error: {e}")

# Go - explicit error returns
func riskyOperation() (string, error) {
    if somethingWrong {
        return "", errors.New("something went wrong")
    }
    return "success", nil
}

result, err := riskyOperation()
if err != nil {
    fmt.Printf("Error: %v\n", err)
}
```

### Methods vs Functions
```python
# Python - methods on classes
class Rectangle:
    def __init__(self, width, height):
        self.width = width
        self.height = height
    
    def area(self):
        return self.width * self.height

rect = Rectangle(5, 3)
print(rect.area())

# Go - methods on types
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

rect := Rectangle{Width: 5, Height: 3}
fmt.Println(rect.Area())
```

## 🚀 Getting Started

1. Read through the example files to understand the concepts
2. Open `functions_practice.go`
3. Fill in the `// YOUR CODE HERE` sections
4. Run with: `go run functions_practice.go`
5. Test your understanding with different inputs

## 💡 Pro Tips

1. **Use multiple returns** for error handling: `(result, error)`
2. **Named returns** make functions self-documenting
3. **Pointer receivers** modify the original, value receivers don't
4. **Variadic functions** use `...` syntax: `func sum(nums ...int)`
5. **Anonymous functions** can capture variables from outer scope (closures)
6. **Function types** enable higher-order programming patterns

## 🧪 Exercises to Complete

Each exercise builds your function mastery:

1. **Basic Functions** - Syntax and multiple returns
2. **Named Returns** - Self-documenting functions
3. **Variadic Functions** - Flexible parameter lists
4. **Anonymous Functions** - Closures and inline functions
5. **Methods** - Functions attached to types
6. **Function Types** - Functions as values
7. **Real-World Applications** - Combining all patterns

## 🎯 Success Criteria

You'll know you've mastered this section when you can:
- Write functions that return multiple values effectively
- Choose between named and unnamed returns appropriately
- Create variadic functions for flexible APIs
- Use anonymous functions and understand closures
- Distinguish between methods and functions
- Use function types for callback patterns
- Apply Go's (result, error) pattern consistently

## 🔗 What's Next

After mastering functions, you'll move on to error handling where you'll learn to create custom error types and implement robust error handling strategies.

Let's master Go's function system! 🐹 