# Goroutine Loop Variable Capture Problem

## Problem Description

**Status: 🔍 Under Investigation**

Classic Go concurrency bug where goroutines in loops capture loop variables by reference, not by value.

## The Issue

```go
// ❌ WRONG - All goroutines see the same variable
for i := 0; i < 3; i++ {
    go func() {
        fmt.Printf("Wrong: %d\n", i) // Race condition!
    }()
}
```

**Expected:** Prints `0, 1, 2` in some order  
**Actual:** Usually prints `3, 3, 3` or random values due to race condition

## Why It Happens

1. **Closure captures variable reference**, not value
2. **Goroutines are scheduled**, not immediate 
3. **Loop continues** while goroutines are still pending
4. **All goroutines share the same `i` variable**

## Race Condition Complexity

Adding timing (like `time.Sleep`) inside goroutines creates even more complex race conditions:

```go
for i := 0; i < 3; i++ {
    go func() {
        time.Sleep(1 * time.Second)  // Changes everything!
        fmt.Printf("Wrong: %d\n", i)
    }()
}
```

**Observed output:** `2, 0, 1` (random values as goroutines wake up during loop execution)

## Solutions

### 1. Pass as Parameter
```go
for i := 0; i < 3; i++ {
    go func(val int) {
        fmt.Printf("Correct: %d\n", val)
    }(i)
}
```

### 2. Capture in Local Variable
```go
for i := 0; i < 3; i++ {
    i := i // Create new variable in loop scope
    go func() {
        fmt.Printf("Correct: %d\n", i)
    }()
}
```

## Debugging Notes

- Use VS Code debugger with breakpoints to see the race condition in real-time
- Set breakpoints on: loop variable, goroutine creation, and goroutine execution
- Observe timing differences between goroutine scheduling and loop execution

## Key Insights

- **Goroutines + Closures = Dangerous** when not handled carefully
- **Timing matters** - small changes completely alter race condition behavior
- **Always pass data explicitly** to goroutines rather than relying on closures
- **This bug appears in real code** - URL processing, file handling, etc.

## Location

Code example: `phase4/concurrency/01-foundations/goroutines.go:169-180`
