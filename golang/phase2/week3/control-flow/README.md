# Week 3: Control Flow Structures 🔄

Welcome to Go's control flow! Unlike Python, Go has a simpler but more explicit approach to controlling program execution. Let's master all the ways to control your program's flow.

## 📚 What You'll Learn

- **if/else statements**: Go's condition checking (no parentheses needed!)
- **for loops**: The only loop in Go, but it's incredibly flexible
- **switch statements**: Clean multi-way branching without fallthrough
- **defer statements**: Cleanup and resource management
- **goto/labels**: When you absolutely need them (rarely!)

## 🎯 Learning Objectives

After completing this section, you'll be able to:
- [x] Write conditional logic with `if/else` statements
- [x] Use all four patterns of Go's `for` loop
- [x] Create clean multi-way branches with `switch`
- [x] Properly manage resources with `defer`
- [x] Understand when and how to use `goto` (though you probably won't need it)

## 📁 Files in This Section

- `if_else.go` - Conditional statements and scope
- `loops.go` - All patterns of Go's flexible for loop
- `switch.go` - Switch statements and type switches
- `defer.go` - Resource cleanup and defer patterns
- `goto.go` - Labels and goto (for completeness)
- `control_flow_practice.go` - **YOUR PRACTICE FILE** - Fill in the blanks!

## ⚡ Key Differences from Python

### Conditions (No Parentheses!)
```python
# Python
if (x > 0 and y < 10):
    print("valid")

# Go
if x > 0 && y < 10 {
    fmt.Println("valid")
}
```

### Loops (Only `for`, But Flexible)
```python
# Python has while, for
while condition:
    do_something()

for i in range(10):
    print(i)

for item in items:
    process(item)

# Go only has for, but it does everything
for condition {
    doSomething()
}

for i := 0; i < 10; i++ {
    fmt.Println(i)
}

for _, item := range items {
    process(item)
}
```

### Switch (No Fallthrough by Default)
```python
# Python doesn't have switch - uses if/elif
if value == "a":
    handle_a()
elif value == "b" or value == "c":
    handle_bc()
else:
    handle_default()

# Go
switch value {
case "a":
    handleA()
case "b", "c":
    handleBC()
default:
    handleDefault()
}
```

## 🚀 Getting Started

1. Read through the example files to understand the concepts
2. Open `control_flow_practice.go`
3. Fill in the `// YOUR CODE HERE` sections
4. Run with: `go run control_flow_practice.go`
5. Test your understanding with different inputs

## 💡 Pro Tips

1. **No parentheses needed** around conditions in `if` statements
2. **Opening brace must be on same line** as `if`, `for`, `switch`
3. **`defer` executes in LIFO order** (Last In, First Out)
4. **Use `:=` for new variables** in if conditions: `if x := getValue(); x > 0 {}`
5. **Range loops** give you index and value: `for i, v := range slice {}`

## 🧪 Exercises to Complete

Each exercise builds your understanding step by step:

1. **Basic Conditions** - if/else with various operators
2. **Loop Patterns** - All four types of for loops
3. **Switch Logic** - Regular and type switches
4. **Resource Management** - Using defer properly
5. **Real-World Scenarios** - Combining control flow patterns

## 🎯 Success Criteria

You'll know you've mastered this section when you can:
- Write clean conditional logic without parentheses
- Choose the right loop pattern for any situation
- Use switch statements for clean multi-way branching
- Properly clean up resources with defer
- Combine control flow structures effectively

Let's dive in and master Go's control flow! 🐹

---

## ✅ Progress Report - COMPLETED!

**Date Completed:** December 2024  
**Status:** All exercises completed successfully! 🎉

### 📋 What Was Accomplished

#### ✅ Exercise 1: Basic If/Else Statements
- ✅ Age validation (adult/minor check)
- ✅ Grade assignment using if/else if chains
- ✅ If with initialization pattern: `if x := getValue(); x > 50 {}`
- ✅ Implemented `getValue()` function with random number generation

#### ✅ Exercise 2: For Loop Patterns  
- ✅ C-style for loop (traditional 3-part loop)
- ✅ While-style loop using bit shifting for powers of 2
- ✅ Infinite loop with break statement
- ✅ Range loop over slice (colors)
- ✅ Range loop over map (student grades)
- ✅ Range loop with continue (filtering even numbers)

#### ✅ Exercise 3: Switch Statements
- ✅ Basic switch for weekday/weekend classification
- ✅ Switch with multiple values for season determination
- ✅ Switch without expression (replacing if/else chains)
- ✅ Type switch with `interface{}` handling different types

#### ✅ Exercise 4: Defer Statements
- ✅ Basic defer showing LIFO execution order
- ✅ Defer with variables (understanding value capture)
- ✅ Defer for resource cleanup (file operations)

#### ✅ Exercise 5: Real-World Scenarios
- ✅ Number guessing game with conditional logic
- ✅ Grade calculator with average computation
- ✅ FizzBuzz implementation with switch statements
- ✅ Email validation using `strings.Contains()`

#### ✅ Exercise 6: Combined Control Flow
- ✅ Menu system with user input and switch
- ✅ Data processing pipeline (filter evens, square, sum)
- ✅ Goto retry mechanism (understanding labels)

### 🎯 Key Concepts Mastered

- **Bit Shifting**: Learned `1 << j` for calculating 2^j efficiently
- **Type Assertions**: Understood `value.(type)` for type switching
- **Pointers Basics**: Used `&variable` for input functions like `fmt.Scanln()`
- **Defer Patterns**: Resource cleanup and LIFO execution order
- **Switch Advantages**: No break needed, automatic fallthrough prevention
- **Go vs Python**: Explicit error handling, no parentheses in conditions

### 💡 Key Insights Gained

1. **Go's `for` loop is incredibly versatile** - replaces while, do-while, and foreach
2. **Switch statements are safer than C/Java** - no accidental fallthrough
3. **Defer is perfect for cleanup** - ensures resources are released
4. **Bit operations are first-class** - `1 << j` is idiomatic for powers of 2
5. **Type safety is explicit** - type assertions require explicit checking
6. **Error handling is visible** - prefer `result, ok` patterns over risky single-value forms

### 🔍 Areas for Future Reference

- **Pointers**: Will be covered in detail in upcoming chapters
- **Error Handling**: More sophisticated patterns coming up
- **Concurrency**: Control flow with goroutines and channels
- **Performance**: When to choose different control flow patterns

**Next Steps:** Ready to move on to Functions and Error Handling! 🚀 