# Go Reflection 🪞

## What is Reflection?

**Reflection** is Go's built-in system for **examining and manipulating values at runtime** when you don't know their types at compile time. Think of it as a "mirror" that lets your program look at itself.

## 🎪 The Problem Reflection Solves

### Without Reflection (Limited):
```go
// You must know the type at compile time
var x int = 42
var y string = "hello"

// How do you write a generic function that works with any type?
func printValue(???) {  // What type to use?
    fmt.Println(???)    // How to handle different types?
}
```

### With Reflection (Flexible):
```go
// Works with ANY type at runtime
func printValue(v interface{}) {
    val := reflect.ValueOf(v)
    fmt.Printf("Type: %v, Value: %v\n", val.Type(), val.Interface())
}

printValue(42)      // Type: int, Value: 42
printValue("hello") // Type: string, Value: hello
```

## 🔧 Core Reflection Types

### 1. **reflect.Type** (Type Information)
```go
var x int = 42
t := reflect.TypeOf(x)

fmt.Println(t.Name())       // "int"
fmt.Println(t.Kind())       // reflect.Int
fmt.Println(t.Size())       // 8 (on 64-bit systems)
```

### 2. **reflect.Value** (Value Manipulation)
```go
var x int = 42
v := reflect.ValueOf(x)

fmt.Println(v.Kind())       // reflect.Int
fmt.Println(v.Int())        // 42
fmt.Println(v.Interface())  // 42 (as interface{})
```

### 3. **reflect.Kind** (Fundamental Types)
```go
// All types have a fundamental "Kind"
reflect.Int, reflect.String, reflect.Slice, reflect.Map,
reflect.Struct, reflect.Ptr, reflect.Chan, reflect.Func, ...
```

## ⚡ Basic Usage Patterns

### Type Inspection
```go
func inspectType(x interface{}) {
    t := reflect.TypeOf(x)
    v := reflect.ValueOf(x)
    
    fmt.Printf("Type: %v\n", t)
    fmt.Printf("Kind: %v\n", t.Kind())
    fmt.Printf("Value: %v\n", v.Interface())
    
    // Check specific kinds
    switch v.Kind() {
    case reflect.Int:
        fmt.Printf("Integer: %d\n", v.Int())
    case reflect.String:
        fmt.Printf("String: %s\n", v.String())
    case reflect.Slice:
        fmt.Printf("Slice length: %d\n", v.Len())
    }
}
```

### Setting Values (Requires Pointers!)
```go
func modifyValue(x interface{}) {
    v := reflect.ValueOf(x)
    
    // Must pass pointer to modify original value
    if v.Kind() != reflect.Ptr {
        fmt.Println("Need pointer to modify value")
        return
    }
    
    elem := v.Elem()
    if !elem.CanSet() {
        return
    }
    
    // Set based on type
    switch elem.Kind() {
    case reflect.Int:
        elem.SetInt(100)
    case reflect.String:
        elem.SetString("modified")
    }
}

// Usage: Must pass pointer!
var x int = 42
modifyValue(&x)
fmt.Println(x)  // 100
```

## 🏗️ Working with Structs

### Struct Field Inspection
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    Email string `json:"email,omitempty"`
}

func inspectStruct(s interface{}) {
    t := reflect.TypeOf(s)
    v := reflect.ValueOf(s)
    
    // Handle pointer to struct
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
        v = v.Elem()
    }
    
    fmt.Printf("Struct: %s\n", t.Name())
    
    // Iterate through fields
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        
        fmt.Printf("Field: %s\n", field.Name)
        fmt.Printf("  Type: %s\n", field.Type)
        fmt.Printf("  Value: %v\n", value.Interface())
        fmt.Printf("  Tag: %s\n", field.Tag.Get("json"))
    }
}
```

### Dynamic Select (Our Fan-in Example)
```go
func dynamicSelect(channels ...interface{}) (int, interface{}, bool) {
    cases := make([]reflect.SelectCase, len(channels))
    
    for i, ch := range channels {
        cases[i] = reflect.SelectCase{
            Dir:  reflect.SelectRecv,  // We want to receive
            Chan: reflect.ValueOf(ch), // Convert to reflect.Value
        }
    }
    
    // Dynamic select at runtime!
    chosen, value, ok := reflect.Select(cases)
    return chosen, value.Interface(), ok
}

// This enables our fanInSelectScalable function!
```

### Generic Function Calls
```go
func callFunction(fn interface{}, args ...interface{}) []interface{} {
    fnVal := reflect.ValueOf(fn)
    if fnVal.Kind() != reflect.Func {
        panic("Not a function")
    }
    
    // Convert arguments to reflect.Value
    in := make([]reflect.Value, len(args))
    for i, arg := range args {
        in[i] = reflect.ValueOf(arg)
    }
    
    // Call function
    results := fnVal.Call(in)
    
    // Convert results back to interface{}
    out := make([]interface{}, len(results))
    for i, result := range results {
        out[i] = result.Interface()
    }
    
    return out
}

// Usage:
add := func(a, b int) int { return a + b }
results := callFunction(add, 5, 3)
fmt.Println(results[0]) // 8
```

## 🌐 Real-World Examples

### Simple JSON Marshal
```go
import (
    "fmt"
    "reflect"
    "strings"
)

func marshal(v interface{}) string {
    val := reflect.ValueOf(v)
    typ := reflect.TypeOf(v)
    
    switch val.Kind() {
    case reflect.String:
        return fmt.Sprintf(`"%s"`, val.String())
    case reflect.Int:
        return fmt.Sprintf("%d", val.Int())
    case reflect.Struct:
        var fields []string
        for i := 0; i < val.NumField(); i++ {
            field := typ.Field(i)
            if !field.IsExported() {
                continue  // Skip private fields
            }
            value := val.Field(i)
            key := field.Name
            val := marshal(value.Interface())
            fields = append(fields, fmt.Sprintf(`"%s":%s`, key, val))
        }
        return fmt.Sprintf("{%s}", strings.Join(fields, ","))
    default:
        return `null`
    }
}
```

### Generic Deep Copy
```go
func deepCopy(src interface{}) interface{} {
    srcVal := reflect.ValueOf(src)
    
    switch srcVal.Kind() {
    case reflect.Struct:
        newStruct := reflect.New(srcVal.Type()).Elem()
        for i := 0; i < srcVal.NumField(); i++ {
            field := srcVal.Field(i)
            if field.CanInterface() {
                copied := deepCopy(field.Interface())
                newStruct.Field(i).Set(reflect.ValueOf(copied))
            }
        }
        return newStruct.Interface()
        
    case reflect.Slice:
        newSlice := reflect.MakeSlice(srcVal.Type(), srcVal.Len(), srcVal.Cap())
        for i := 0; i < srcVal.Len(); i++ {
            copied := deepCopy(srcVal.Index(i).Interface())
            newSlice.Index(i).Set(reflect.ValueOf(copied))
        }
        return newSlice.Interface()
        
    default:
        return src  // Basic types
    }
}
```

## ⚠️ Reflection Best Practices

### ✅ DO:
```go
// Always check Kind before type-specific operations
if v.Kind() == reflect.Int {
    val := v.Int()  // Safe - we checked first
}

// Check if value can be set before setting
if v.CanSet() {
    v.SetString("new value")
}

// Check for nil pointers
if v.Kind() == reflect.Ptr && v.IsNil() {
    return // Handle nil case
}

// Use type assertions when type is known
if s, ok := x.(string); ok {
    return s  // Faster than reflection
}
```

### ❌ DON'T:
```go
// Don't use reflection for simple type checks
if reflect.TypeOf(x).Kind() == reflect.String {  // ❌ Slow
    // Use type assertion instead
}
if _, ok := x.(string); ok {  // ✅ Fast
    // Better approach
}

// Don't ignore panic risks
v.Int()  // ❌ Panics if not an int

// Don't use reflection in hot paths
for i := 0; i < 1000000; i++ {
    reflect.ValueOf(data).String()  // ❌ Very slow
}
```

## 🚀 Performance Considerations

| Operation | Reflection | Direct Code | Performance Ratio |
|-----------|------------|-------------|------------------|
| **Type Check** | `reflect.TypeOf()` | Type assertion | ~10x slower |
| **Field Access** | `v.FieldByName()` | Direct access | ~20x slower |
| **Function Call** | `v.Call()` | Direct call | ~50x slower |

## 🎯 When to Use Reflection

### ✅ GOOD Use Cases:
- **Serialization/Deserialization** (JSON, XML)
- **ORM frameworks** (database mapping)
- **Generic utilities** (deep copy, validation)
- **Testing frameworks** (mock generation)
- **Dynamic select** (like our fan-in example!)

### ❌ AVOID Reflection When:
- **Performance is critical** (hot paths)
- **Type is known at compile time** (use interfaces)
- **Simple operations** (use type assertions)
- **Code clarity matters** (reflection is complex)

## 💡 Reflection vs Alternatives

### Use Interfaces Instead
```go
// Instead of reflection:
func process(data interface{}) {
    v := reflect.ValueOf(data)
    if v.Kind() == reflect.String {
        processString(v.String())
    }
}

// Use interfaces:
type Processor interface {
    Process()
}

func process(p Processor) {
    p.Process()  // Cleaner and faster
}
```

### Use Type Assertions
```go
// Instead of reflection:
if reflect.TypeOf(x).Kind() == reflect.String {
    // Complex reflection code
}

// Use type assertion:
if s, ok := x.(string); ok {
    // Use s - much faster!
}
```

## 🛡️ Safety Guidelines

```go
func safeReflection(x interface{}) {
    v := reflect.ValueOf(x)
    
    // Always check Kind first
    switch v.Kind() {
    case reflect.Int:
        val := v.Int()  // Safe
    case reflect.Ptr:
        if !v.IsNil() {  // Check nil
            elem := v.Elem()
        }
    case reflect.Slice:
        if v.Len() > 0 {  // Check bounds
            first := v.Index(0)
        }
    }
}
```

## 🔗 Key Takeaways

1. **Reflection is powerful but slow** - use sparingly
2. **Always check `Kind()`** before operations to avoid panics
3. **Use pointers for modifying values** (`&variable`)
4. **Prefer interfaces and type assertions** when possible
5. **Perfect for frameworks** - avoid in business logic
6. **Test thoroughly** - reflection bypasses compile-time safety

## 💭 Rob Pike's Quote

> *"Clear is better than clever. Reflection is never clear."*

**Reflection is Go's escape hatch for ultimate flexibility - but use it wisely!** 🪞⚡

Perfect for our `fanInSelectScalable` example, but always consider simpler alternatives first!

