# Phase 5: Standard Library & Packages 📚

Welcome to Go's rich ecosystem! This phase explores Go's comprehensive standard library and teaches you to build complete applications using battle-tested packages. You'll learn the essential tools every Go developer needs.

## 📚 What You'll Learn

### Week 9: Essential Packages
- **Core packages**: `fmt`, `strings`, `strconv`, `time`, `math`
- **JSON handling**: Encoding, decoding, custom marshaling
- **HTTP client**: Making requests, handling responses
- **Context package**: Cancellation, timeouts, values
- **Logging**: Structured logging and best practices

### Week 10: File I/O & System Programming
- **File operations**: Reading, writing, streaming
- **Directory management**: Walking, watching, permissions
- **Environment variables**: Configuration management
- **Command-line tools**: Flag parsing, argument handling
- **System interfaces**: OS interaction, signal handling

## 🎯 Learning Objectives

After completing this phase, you'll be able to:
- [ ] Use Go's standard library packages effectively
- [ ] Handle JSON data and HTTP communications
- [ ] Manage files, directories, and system resources
- [ ] Build command-line applications with proper configuration
- [ ] Apply context for cancellation and timeouts
- [ ] Implement proper logging and error handling
- [ ] Write portable cross-platform Go programs

## 📁 Phase Structure

```
phase5/
├── week9/          # Essential Packages
│   └── packages/
│       ├── README.md
│       ├── fmt_strings.go          # Text formatting and manipulation
│       ├── time_math.go            # Time handling and mathematics
│       ├── json_handling.go        # JSON encoding/decoding
│       ├── http_client.go          # HTTP client operations
│       ├── context_usage.go        # Context for cancellation
│       └── packages_practice.go    # Practice exercises
│
└── week10/         # File I/O & System Programming
    └── fileio/
        ├── README.md
        ├── file_operations.go      # Reading and writing files
        ├── directory_management.go # Directory operations
        ├── environment_config.go   # Environment variables
        ├── cli_tools.go           # Command-line applications
        ├── system_interfaces.go   # OS and system interaction
        └── fileio_practice.go     # Practice exercises
```

## ⚡ Key Differences from Python

### JSON Handling
```python
# Python - Dynamic JSON
import json

data = {"name": "John", "age": 30}
json_str = json.dumps(data)
parsed = json.loads(json_str)
print(parsed["name"])  # Direct access
```

```go
// Go - Structured JSON with types
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

data := Person{Name: "John", Age: 30}
jsonBytes, _ := json.Marshal(data)
var parsed Person
json.Unmarshal(jsonBytes, &parsed)
fmt.Println(parsed.Name)  // Type-safe access
```

### HTTP Client
```python
# Python - requests library
import requests

response = requests.get("https://api.github.com/users/octocat")
data = response.json()
print(data["name"])
```

```go
// Go - Built-in HTTP client
import (
    "encoding/json"
    "net/http"
)

resp, err := http.Get("https://api.github.com/users/octocat")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var data map[string]interface{}
json.NewDecoder(resp.Body).Decode(&data)
fmt.Println(data["name"])
```

### File Operations
```python
# Python - Simple file operations
with open("file.txt", "r") as f:
    content = f.read()

with open("output.txt", "w") as f:
    f.write("Hello, World!")
```

```go
// Go - Explicit error handling
import (
    "io"
    "os"
)

// Reading
content, err := os.ReadFile("file.txt")
if err != nil {
    log.Fatal(err)
}

// Writing
err = os.WriteFile("output.txt", []byte("Hello, World!"), 0644)
if err != nil {
    log.Fatal(err)
}
```

## 🚀 Getting Started

1. **Week 9**: Start with `week9/packages/README.md`
2. **Explore packages**: Learn essential standard library packages
3. **Build HTTP clients**: Practice with real APIs
4. **Master JSON**: Handle structured data effectively
5. **Week 10**: Move to file I/O and system programming
6. **Build CLI tools**: Create command-line applications

## 💡 Go Standard Library Philosophy

**"Batteries Included" with Go Characteristics:**
- **Comprehensive**: Rich standard library for common tasks
- **Consistent**: Uniform APIs across packages
- **Efficient**: High-performance implementations
- **Composable**: Packages work well together
- **Stable**: Strong backward compatibility guarantees

**Package Design Principles:**
- Small, focused interfaces
- Explicit error handling
- Zero-value usefulness
- Composability over complexity

## 🎯 Essential Packages Overview

### Core Text & Data
- **fmt**: Formatted I/O operations
- **strings**: String manipulation utilities
- **strconv**: String conversions
- **unicode**: Unicode support
- **regexp**: Regular expressions

### Time & Math
- **time**: Time operations and formatting
- **math**: Mathematical functions
- **math/rand**: Random number generation
- **sort**: Sorting utilities

### Network & Web
- **net/http**: HTTP client and server
- **net/url**: URL parsing
- **crypto/tls**: TLS encryption
- **encoding/json**: JSON handling

### System & I/O
- **os**: Operating system interface
- **io**: I/O primitives
- **bufio**: Buffered I/O
- **path/filepath**: File path manipulation
- **flag**: Command-line flag parsing

### Concurrency & Context
- **context**: Cancellation and timeouts
- **sync**: Synchronization primitives
- **runtime**: Go runtime interface

## 🔧 Common Patterns

### Error Handling Pattern
```go
result, err := someOperation()
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
// Use result
```

### Context Pattern
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

result, err := operationWithContext(ctx)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        // Handle timeout
    }
    return err
}
```

### Resource Management Pattern
```go
file, err := os.Open("filename.txt")
if err != nil {
    return err
}
defer file.Close()  // Always clean up

// Use file
```

## 📊 Performance Characteristics

| Package | Use Case | Performance | Memory |
|---------|----------|-------------|---------|
| **fmt** | Formatting | Moderate | Higher (reflection) |
| **strconv** | Conversions | Fast | Low |
| **encoding/json** | JSON | Fast | Moderate |
| **net/http** | HTTP | Fast | Efficient pooling |
| **bufio** | Buffered I/O | Very Fast | Configurable |
| **time** | Time ops | Fast | Low |

## ⚠️ Common Pitfalls

### JSON Marshaling
```go
// ❌ WRONG - Unexported fields ignored
type Person struct {
    name string  // Won't be marshaled
    age  int     // Won't be marshaled
}

// ✅ CORRECT - Exported fields with tags
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

### HTTP Client Resources
```go
// ❌ WRONG - Resource leak
resp, err := http.Get(url)
data, _ := io.ReadAll(resp.Body)  // Body not closed

// ✅ CORRECT - Proper cleanup
resp, err := http.Get(url)
if err != nil {
    return err
}
defer resp.Body.Close()
data, err := io.ReadAll(resp.Body)
```

### Time Zone Handling
```go
// ❌ WRONG - Local time assumptions
now := time.Now()
fmt.Println(now)  // Local timezone

// ✅ CORRECT - Explicit timezone
now := time.Now().UTC()
fmt.Println(now.Format(time.RFC3339))
```

## 🎯 Success Criteria

You'll know you've mastered this phase when you can:
- Navigate the standard library documentation effectively
- Choose appropriate packages for common tasks
- Handle JSON and HTTP operations confidently
- Build command-line tools with proper configuration
- Manage files and directories safely
- Use context for cancellation and timeouts
- Apply Go idioms consistently across packages

## 🔗 What's Next

After mastering the standard library, you'll advance to **Phase 6: Web Development** where you'll build complete web applications and APIs using Go's powerful HTTP capabilities!

## 📈 Real-World Applications

By the end of this phase, you'll be able to build:
- **CLI tools**: Command-line utilities with flags and configuration
- **HTTP clients**: API integration and data fetching
- **File processors**: Log analyzers, data converters
- **System monitors**: Resource monitoring and alerting
- **Configuration managers**: Environment-based configuration
- **Data pipelines**: ETL and batch processing tools

Ready to explore Go's treasure trove of packages! 📚🔍🐹
