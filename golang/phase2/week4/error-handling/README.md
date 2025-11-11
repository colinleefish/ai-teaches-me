# Week 4: Error Handling Mastery 🛡️

Welcome to Go's error handling! Go takes a unique approach to error handling that's explicit, predictable, and forces you to think about what can go wrong.

## 📚 What You'll Learn

- **Error interface**: Understanding Go's fundamental error type
- **Custom errors**: Creating meaningful error types for your applications
- **Error wrapping**: Adding context while preserving original errors
- **Error strategies**: Patterns for handling errors in different scenarios
- **Panic/Recover**: Emergency exits and how to handle them
- **Validation patterns**: Early returns and input validation

## 🎯 Learning Objectives

After completing this section, you'll be able to:
- [ ] Create custom error types that implement the error interface
- [ ] Wrap errors to add context while preserving the original error
- [ ] Choose appropriate error handling strategies for different scenarios
- [ ] Use panic and recover appropriately (sparingly!)
- [ ] Implement robust validation patterns
- [ ] Handle errors gracefully in real-world applications
- [ ] Debug and trace errors through error wrapping chains

## 📁 Files in This Section

- `basic_errors.go` - Error interface and basic error handling
- `custom_errors.go` - Creating custom error types
- `error_wrapping.go` - Error wrapping and unwrapping (Go 1.13+)
- `error_strategies.go` - Different patterns for handling errors
- `panic_recover.go` - When and how to use panic/recover
- `validation.go` - Input validation and early returns
- `error_practice.go` - **YOUR PRACTICE FILE** - Fill in the blanks!

## ⚡ Key Differences from Python

### Error Handling Philosophy
```python
# Python - Exception-based (implicit)
def divide(a, b):
    return a / b  # Might raise ZeroDivisionError

try:
    result = divide(10, 0)
except ZeroDivisionError:
    print("Cannot divide by zero")

# Go - Explicit error returns
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Printf("Error: %v\n", err)
}
```

### Custom Errors
```python
# Python - Custom exception classes
class ValidationError(Exception):
    def __init__(self, field, message):
        self.field = field
        self.message = message
        super().__init__(f"{field}: {message}")

raise ValidationError("email", "invalid format")

# Go - Types that implement error interface
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

return ValidationError{Field: "email", Message: "invalid format"}
```

### Error Wrapping
```python
# Python - Exception chaining
try:
    process_data()
except ValueError as e:
    raise ProcessingError("Failed to process") from e

# Go - Error wrapping
if err := processData(); err != nil {
    return fmt.Errorf("failed to process: %w", err)
}
```

## 🚀 Getting Started

1. Read through the example files to understand the concepts
2. Open `error_practice.go`
3. Fill in the `// YOUR CODE HERE` sections
4. Run with: `go run error_practice.go`
5. Test your understanding with different error scenarios

## 💡 Pro Tips

1. **Always check errors**: `if err != nil` is the Go way
2. **Wrap errors for context**: Use `fmt.Errorf("context: %w", err)`
3. **Custom errors for your domain**: Create meaningful error types
4. **Early returns**: Validate input and return errors early
5. **Don't ignore errors**: Handle them explicitly or document why you're ignoring
6. **Use panic sparingly**: Only for truly exceptional circumstances

## 🧪 Exercises to Complete

Each exercise builds your error handling expertise:

1. **Basic Error Handling** - Working with the error interface
2. **Custom Error Types** - Creating domain-specific errors
3. **Error Wrapping** - Adding context to errors
4. **Error Strategies** - Different patterns for different scenarios
5. **Panic and Recover** - Emergency handling mechanisms
6. **Validation Patterns** - Input validation and early returns
7. **Real-World Error Handling** - Combining all patterns

## 🎯 Success Criteria

You'll know you've mastered this section when you can:
- Create meaningful custom error types for your applications
- Wrap errors appropriately to add context
- Choose the right error handling strategy for different scenarios
- Use panic and recover appropriately (rarely!)
- Implement robust validation patterns
- Debug errors through wrapping chains
- Write error-resilient code that fails gracefully

## 🔗 What's Next

After mastering error handling, you'll move on to advanced function patterns where you'll learn about recursion, higher-order functions, and functional programming concepts in Go.

## 📖 Error Handling Best Practices

### The Error Interface
```go
type error interface {
    Error() string
}
```

### Common Patterns
```go
// 1. Simple error creation
errors.New("something went wrong")

// 2. Formatted errors
fmt.Errorf("failed to process %s: %v", filename, originalErr)

// 3. Error wrapping (Go 1.13+)
fmt.Errorf("failed to process %s: %w", filename, originalErr)

// 4. Custom error types
type NetworkError struct {
    Operation string
    Addr      string
    Err       error
}

func (e NetworkError) Error() string {
    return fmt.Sprintf("%s %s: %v", e.Operation, e.Addr, e.Err)
}

// 5. Error checking and unwrapping
if errors.Is(err, ErrNotFound) {
    // Handle specific error
}

var netErr NetworkError
if errors.As(err, &netErr) {
    // Handle specific error type
}
```

Let's master Go's explicit error handling! 🐹 