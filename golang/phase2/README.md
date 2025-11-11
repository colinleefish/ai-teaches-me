# Phase 2: Control Flow, Functions & Error Handling 🚀

Welcome to Phase 2! Now that you've mastered Go's type system and collections, it's time to learn how to control program flow, write powerful functions, and handle errors like a pro.

## 📋 Phase 2 Overview

**Goal**: Master Go's control structures, function patterns, and error handling strategies.

### Week 3: Control Flow & Functions
- **Day 15-18**: Control Flow Structures (`week3/control-flow/`)
- **Day 19-21**: Functions & Methods (`week3/functions/`)

### Week 4: Error Handling & Advanced Patterns
- **Day 22-25**: Error Handling Mastery (`week4/error-handling/`)
- **Day 26-28**: Advanced Function Patterns (`week4/advanced-patterns/`)

## 🎯 Learning Objectives

By the end of Phase 2, you should be able to:
- [ ] Use all of Go's control flow structures effectively
- [ ] Write functions with multiple return values and proper error handling
- [ ] Create custom error types and handle complex error scenarios
- [ ] Understand the difference between `panic`/`recover` and normal error handling
- [ ] Use `defer` for resource cleanup and function timing
- [ ] Write recursive functions and use function types as first-class values
- [ ] Apply Go's error handling idioms in real-world scenarios

## 📁 Directory Structure

```
phase2/
├── week3/
│   ├── control-flow/        # if/else, for, switch, defer
│   └── functions/           # Function syntax, methods, closures
├── week4/
│   ├── error-handling/      # Error types, wrapping, patterns
│   └── advanced-patterns/   # Recursion, function types, advanced patterns
├── quizzes/                 # Knowledge checks for Phase 2
├── projects/                # Hands-on projects
└── solutions/               # Reference solutions
```

## 🧪 How to Use This Phase

### 1. Complete Week 3 First
Start with control flow structures, then move to functions. Each builds on the previous concepts.

### 2. Practice with Real Examples
Each section includes:
- Concept explanations with Python comparisons
- Hands-on exercises with `// YOUR CODE HERE` placeholders
- Real-world examples and patterns
- Common pitfalls and how to avoid them

### 3. Build Projects
Apply your knowledge with practical projects:
- **Simple Calculator** (control flow + functions)
- **File Processor** (error handling + functions)
- **Command-Line Tool** (all concepts combined)

### 4. Test Your Understanding
- Complete quizzes after each week
- Solve coding challenges
- Compare your solutions with provided references

## 🚀 Getting Started

1. **Prerequisites**: Complete Phase 1 (variables, types, collections)

2. **Start with Week 3**:
   ```bash
   cd phase2/week3/control-flow
   cat README.md  # Read instructions
   ```

3. **Work Through Systematically**:
   - Read explanations and examples
   - Fill in `// YOUR CODE HERE` sections
   - Run and test your code
   - Move to next section

## 📊 Progress Tracking

### Week 3 Progress
- [ ] Control Flow: if/else statements
- [ ] Control Flow: for loops (all variants)
- [ ] Control Flow: switch statements
- [ ] Control Flow: defer keyword
- [ ] Functions: Basic syntax and returns
- [ ] Functions: Methods and receivers
- [ ] Functions: Anonymous functions and closures
- [ ] **Week 3 Quiz**: Score ___/100

### Week 4 Progress
- [ ] Error Handling: Error interface
- [ ] Error Handling: Custom errors
- [ ] Error Handling: Error wrapping
- [ ] Error Handling: panic/recover
- [ ] Advanced: Recursive functions
- [ ] Advanced: Function types
- [ ] Advanced: Higher-order functions
- [ ] **Week 4 Quiz**: Score ___/100

### Phase 2 Completion
- [ ] **Control Flow Project**: Complete calculator
- [ ] **Functions Project**: File processor
- [ ] **Error Handling Project**: Robust CLI tool
- [ ] **Final Quiz**: Score ___/100
- [ ] **All Tests Pass**: ✅
- [ ] **Ready for Phase 3**: ✅

## 🎓 Key Concepts You'll Master

### Control Flow Mastery
- Go's unique `for` loop patterns (C-style, while, range, infinite)
- `switch` without fallthrough and type switches
- `defer` for cleanup and function instrumentation
- Proper use of `goto` (rarely needed but good to know)

### Function Excellence
- Multiple return values and the `(result, error)` pattern
- Named returns and when to use them
- Variadic functions for flexible APIs
- Methods vs functions and receiver types
- Closures and function scope

### Error Handling Expertise
- Go's explicit error handling philosophy
- Creating meaningful custom error types
- Error wrapping for context preservation
- Recovery strategies and when to use `panic`
- Validation patterns and early returns

## ⚡ Python vs Go Highlights

### Control Flow
```python
# Python
while condition:
    do_something()

for item in items:
    process(item)

# Go
for condition {
    doSomething()
}

for _, item := range items {
    process(item)
}
```

### Error Handling
```python
# Python - Exception based
try:
    result = risky_operation()
except Exception as e:
    logger.error(f"Failed: {e}")
    raise

# Go - Explicit error checking
result, err := riskyOperation()
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
```

### Functions
```python
# Python
def calculate(a, b):
    if b == 0:
        raise ValueError("division by zero")
    return a / b

# Go
func calculate(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

## 🆘 Getting Help

If you get stuck:
1. Review the Python vs Go comparisons
2. Check the `solutions/` folder for reference implementations
3. Read error messages carefully - Go's compiler is helpful
4. Practice with the interactive exercises

## 🎯 Ready to Begin?

Start your Phase 2 journey:
```bash
cd week3/control-flow
```

Let's master Go's program flow and function patterns! 🐹 