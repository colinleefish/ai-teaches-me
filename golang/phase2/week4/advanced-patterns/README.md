# Week 4: Advanced Function Patterns 🚀

Welcome to advanced Go function patterns! Now that you've mastered error handling, let's explore powerful function techniques including recursion, higher-order functions, and functional programming concepts in Go.

## 📚 What You'll Learn

- **Recursive functions**: Elegant solutions to divide-and-conquer problems
- **Higher-order functions**: Functions that accept or return other functions
- **Functional patterns**: Map, filter, reduce operations in Go
- **Function composition**: Building complex operations from simple functions
- **Memoization**: Caching function results for performance
- **Decorator patterns**: Wrapping functions to add behavior
- **Pipeline patterns**: Chaining operations for data processing

## 🎯 Learning Objectives

After completing this section, you'll be able to:
- [ ] Write recursive functions with proper base cases and error handling
- [ ] Create higher-order functions that accept function parameters
- [ ] Implement functional programming patterns like map, filter, reduce
- [ ] Compose functions to build complex operations from simple ones
- [ ] Use memoization to optimize expensive computations
- [ ] Apply decorator patterns to add cross-cutting concerns
- [ ] Build data processing pipelines with function chaining

## 📁 Files in This Section

- `recursion.go` - Recursive function patterns and examples
- `higher_order.go` - Functions as first-class values
- `functional_patterns.go` - Map, filter, reduce in Go
- `composition.go` - Function composition techniques
- `memoization.go` - Caching function results
- `decorators.go` - Function wrapping and decoration
- `pipelines.go` - Data processing pipelines
- `advanced_practice.go` - **YOUR PRACTICE FILE** - Fill in the blanks!

## ⚡ Key Differences from Python

### Recursion with Error Handling
```python
# Python - Simple recursion
def factorial(n):
    if n <= 1:
        return 1
    return n * factorial(n - 1)

# Go - Recursion with error handling
func factorial(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("factorial undefined for negative numbers")
    }
    if n <= 1 {
        return 1, nil
    }
    result, err := factorial(n - 1)
    if err != nil {
        return 0, err
    }
    return n * result, nil
}
```

### Higher-Order Functions
```python
# Python - Built-in higher-order functions
numbers = [1, 2, 3, 4, 5]
squared = list(map(lambda x: x**2, numbers))
evens = list(filter(lambda x: x % 2 == 0, numbers))

# Go - Custom higher-order functions
func mapInt(slice []int, fn func(int) int) []int {
    result := make([]int, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

squared := mapInt(numbers, func(x int) int { return x * x })
```

### Function Composition
```python
# Python - Function composition
def add_one(x):
    return x + 1

def double(x):
    return x * 2

def compose(f, g):
    return lambda x: f(g(x))

add_one_then_double = compose(double, add_one)

# Go - Function composition
func addOne(x int) int {
    return x + 1
}

func double(x int) int {
    return x * 2
}

func compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}

addOneThenDouble := compose(double, addOne)
```

## 🚀 Getting Started

1. Read through the example files to understand the concepts
2. Open `advanced_practice.go`
3. Fill in the `// YOUR CODE HERE` sections
4. Run with: `go run advanced_practice.go`
5. Test your understanding with different scenarios

## 💡 Pro Tips

1. **Recursive base cases**: Always define clear termination conditions
2. **Tail recursion**: Go doesn't optimize tail recursion, be mindful of stack depth
3. **Function types**: Define custom function types for clarity
4. **Closure capture**: Be careful what variables closures capture
5. **Performance**: Memoization can dramatically speed up recursive functions
6. **Error handling**: Don't forget error handling in recursive calls

## 🧪 Exercises to Complete

Each exercise builds your advanced function skills:

1. **Recursive Algorithms** - Classic recursive problems with Go patterns
2. **Higher-Order Functions** - Functions that manipulate other functions
3. **Functional Programming** - Map, filter, reduce implementations
4. **Function Composition** - Building complex operations
5. **Memoization Patterns** - Caching for performance optimization
6. **Decorator Patterns** - Cross-cutting concerns and function wrapping
7. **Data Pipelines** - Chaining operations for data processing
8. **Real-World Applications** - Combining all patterns

## 🎯 Success Criteria

You'll know you've mastered this section when you can:
- Write recursive functions with proper error handling and base cases
- Create and use higher-order functions effectively
- Implement functional programming patterns in Go
- Compose simple functions into complex operations
- Apply memoization to optimize recursive algorithms
- Use decorator patterns to add behavior to functions
- Build efficient data processing pipelines
- Choose the right pattern for different problem types

## 🔗 What's Next

After mastering these advanced patterns, you'll be ready for Phase 3 where you'll learn about concurrency, goroutines, and channels - Go's most distinctive features!

## 📖 Advanced Pattern Examples

### Recursive Tree Traversal
```go
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

func (t *TreeNode) InorderTraversal(visit func(int)) {
    if t == nil {
        return
    }
    t.Left.InorderTraversal(visit)
    visit(t.Value)
    t.Right.InorderTraversal(visit)
}
```

### Memoized Fibonacci
```go
func memoizedFibonacci() func(int) int {
    cache := make(map[int]int)
    
    var fib func(int) int
    fib = func(n int) int {
        if n <= 1 {
            return n
        }
        if result, exists := cache[n]; exists {
            return result
        }
        result := fib(n-1) + fib(n-2)
        cache[n] = result
        return result
    }
    
    return fib
}
```

### Function Pipeline
```go
type Pipeline func([]int) []int

func (p Pipeline) Then(next Pipeline) Pipeline {
    return func(input []int) []int {
        return next(p(input))
    }
}

// Usage
pipeline := Pipeline(filterEvens).
    Then(mapDouble).
    Then(filterGreaterThan(10))

result := pipeline([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
```

Let's master Go's advanced function patterns! 🐹 