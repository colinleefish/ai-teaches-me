# Week 9: Essential Packages 📦

Welcome to Go's standard library ecosystem! This week covers the most essential packages that every Go developer uses daily. You'll learn to work with text, time, JSON, HTTP, and context - the building blocks of modern Go applications.

## 🎯 Learning Objectives

By the end of this week, you'll understand:
- Core text manipulation with `fmt` and `strings`
- Time handling and mathematical operations
- JSON encoding/decoding for data exchange
- HTTP client operations for API integration
- Context usage for cancellation and timeouts
- Best practices for package usage and error handling

## 📚 Topics Covered

### 1. Text & Formatting (`fmt_strings.go`)
- String formatting with `fmt` package
- String manipulation with `strings` package
- String conversion with `strconv` package
- Unicode and regular expressions
- Performance considerations

### 2. Time & Math (`time_math.go`)
- Time parsing, formatting, and arithmetic
- Timezone handling and duration calculations
- Mathematical functions and random numbers
- Sorting and searching utilities
- Performance-critical calculations

### 3. JSON Handling (`json_handling.go`)
- Marshaling and unmarshaling JSON
- Custom JSON tags and field control
- Streaming JSON for large datasets
- Error handling and validation
- Working with dynamic JSON

### 4. HTTP Client (`http_client.go`)
- Making HTTP requests (GET, POST, PUT, DELETE)
- Request/response handling and headers
- Authentication and cookies
- Timeouts and retries
- Connection pooling and performance

### 5. Context Usage (`context_usage.go`)
- Creating and using contexts
- Cancellation and timeout patterns
- Context values and best practices
- Integration with HTTP and other packages
- Avoiding context anti-patterns

### 6. Practice Exercises (`packages_practice.go`)
- Real-world integration scenarios
- Building CLI tools with packages
- API clients and data processing
- Error handling patterns

## 📦 Essential Package Overview

| Package | Purpose | Key Functions |
|---------|---------|---------------|
| **fmt** | Formatted I/O | `Printf`, `Sprintf`, `Scanf` |
| **strings** | String operations | `Split`, `Join`, `Replace`, `Contains` |
| **strconv** | String conversions | `Atoi`, `Itoa`, `ParseFloat`, `FormatInt` |
| **time** | Time operations | `Now`, `Parse`, `Format`, `Add` |
| **math** | Mathematics | `Abs`, `Max`, `Min`, `Sqrt`, `Round` |
| **encoding/json** | JSON handling | `Marshal`, `Unmarshal`, `NewDecoder` |
| **net/http** | HTTP operations | `Get`, `Post`, `NewRequest`, `Client` |
| **context** | Cancellation | `Background`, `WithTimeout`, `WithCancel` |

## 🚀 Quick Start Examples

### String Operations
```go
package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    // String formatting
    name := "Go"
    version := 1.21
    fmt.Printf("Language: %s, Version: %.2f\n", name, version)
    
    // String manipulation
    text := "Hello, World!"
    fmt.Println(strings.ToUpper(text))
    fmt.Println(strings.Split(text, ", "))
    
    // String conversion
    numStr := "42"
    num, _ := strconv.Atoi(numStr)
    fmt.Printf("Number: %d\n", num)
}
```

### JSON Handling
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Marshal to JSON
person := Person{Name: "Alice", Age: 30}
jsonData, _ := json.Marshal(person)
fmt.Println(string(jsonData))

// Unmarshal from JSON
var decoded Person
json.Unmarshal(jsonData, &decoded)
fmt.Printf("%+v\n", decoded)
```

### HTTP Client
```go
// Simple GET request
resp, err := http.Get("https://api.github.com/users/octocat")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, _ := io.ReadAll(resp.Body)
fmt.Println(string(body))
```

## 🛠️ How to Practice

1. **Start with basics**: Read `fmt_strings.go` for text operations
2. **Learn time/math**: Study `time_math.go` for calculations
3. **Master JSON**: Practice with `json_handling.go`
4. **Build HTTP clients**: Work through `http_client.go`
5. **Use context**: Understand `context_usage.go`
6. **Apply knowledge**: Complete `packages_practice.go`

## 🧪 Testing Your Code

```bash
# Run individual files
go run fmt_strings.go
go run json_handling.go

# Run with race detector
go run -race http_client.go

# Run tests
go test -v ./...

# Benchmark performance
go test -bench=. -benchmem
```

## ⚠️ Common Package Pitfalls

### 1. String Building Performance
```go
// ❌ SLOW - String concatenation in loop
var result string
for i := 0; i < 1000; i++ {
    result += "hello"  // Creates new string each time
}

// ✅ FAST - Use strings.Builder
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("hello")
}
result := builder.String()
```

### 2. JSON Field Visibility
```go
// ❌ WRONG - Private fields not marshaled
type User struct {
    name string  // Won't appear in JSON
    age  int     // Won't appear in JSON
}

// ✅ CORRECT - Public fields with tags
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

### 3. HTTP Resource Leaks
```go
// ❌ WRONG - Body not closed
resp, _ := http.Get(url)
data, _ := io.ReadAll(resp.Body)  // Leak!

// ✅ CORRECT - Always close body
resp, err := http.Get(url)
if err != nil {
    return err
}
defer resp.Body.Close()
data, _ := io.ReadAll(resp.Body)
```

### 4. Context Misuse
```go
// ❌ WRONG - Empty context
ctx := context.TODO()
result := longOperation(ctx)  // No cancellation

// ✅ CORRECT - Timeout context
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
result := longOperation(ctx)
```

## 🎯 Key Concepts to Master

### Time Operations
```go
// Parsing and formatting
layout := "2006-01-02 15:04:05"
t, _ := time.Parse(layout, "2023-12-25 10:30:00")
fmt.Println(t.Format(time.RFC3339))

// Time arithmetic
now := time.Now()
tomorrow := now.Add(24 * time.Hour)
duration := tomorrow.Sub(now)
```

### JSON with Custom Types
```go
type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
    return json.Marshal(time.Time(t).Unix())
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
    var unix int64
    if err := json.Unmarshal(data, &unix); err != nil {
        return err
    }
    *t = Timestamp(time.Unix(unix, 0))
    return nil
}
```

### HTTP with Context
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
client := &http.Client{}
resp, err := client.Do(req)
```

## 📈 Performance Tips

1. **String Operations**: Use `strings.Builder` for concatenation
2. **JSON**: Use streaming for large datasets
3. **HTTP**: Reuse clients and enable connection pooling
4. **Time**: Cache formatted strings for repeated use
5. **Context**: Don't create contexts in hot paths

## 🔍 Debugging Tools

```go
// JSON debugging
func prettyPrint(v interface{}) {
    b, _ := json.MarshalIndent(v, "", "  ")
    fmt.Println(string(b))
}

// HTTP debugging
func debugRequest(req *http.Request) {
    dump, _ := httputil.DumpRequest(req, true)
    fmt.Println(string(dump))
}

// Time debugging
func logDuration(name string, start time.Time) {
    fmt.Printf("%s took %v\n", name, time.Since(start))
}
```

## 🔗 What's Next

After mastering essential packages, you'll learn **File I/O & System Programming** in Week 10 - working with files, directories, environment variables, and building command-line tools!

## 📊 Package Usage Statistics

Based on Go community surveys:
1. **fmt** - 99% of Go projects
2. **strings** - 95% of Go projects  
3. **net/http** - 85% of Go projects
4. **encoding/json** - 80% of Go projects
5. **time** - 75% of Go projects
6. **context** - 70% of Go projects

Ready to master Go's essential packages! 📦⚡🐹
