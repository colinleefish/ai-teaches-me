# Week 10: File I/O & System Programming 📁

Welcome to system-level programming in Go! This week covers file operations, directory management, environment configuration, and building command-line tools. You'll learn to interact with the operating system effectively.

## 🎯 Learning Objectives

By the end of this week, you'll understand:
- File reading, writing, and streaming operations
- Directory traversal and management
- Environment variable handling and configuration
- Command-line argument parsing and flag handling
- System interfaces and OS interaction
- Cross-platform file path handling

## 📚 Topics Covered

### 1. File Operations (`file_operations.go`)
- Reading and writing files
- File streaming and buffered I/O
- File permissions and metadata
- Temporary files and directories
- File locking and atomic operations

### 2. Directory Management (`directory_management.go`)
- Directory creation and traversal
- File system walking and searching
- Path manipulation and validation
- File watching and monitoring
- Archive and compression operations

### 3. Environment Configuration (`environment_config.go`)
- Environment variable handling
- Configuration file management
- Secrets and security considerations
- Configuration validation and defaults
- Hot-reloading configuration

### 4. CLI Tools (`cli_tools.go`)
- Command-line argument parsing
- Flag package and custom flags
- Subcommands and nested commands
- Interactive CLI interfaces
- CLI testing and validation

### 5. System Interfaces (`system_interfaces.go`)
- Process management and signals
- System information and resources
- Network interfaces and system calls
- Cross-platform compatibility
- Error handling and logging

### 6. Practice Exercises (`fileio_practice.go`)
- Building complete CLI applications
- File processing pipelines
- System monitoring tools
- Configuration management systems

## 🛠️ Essential Packages

| Package | Purpose | Key Functions |
|---------|---------|---------------|
| **os** | Operating system interface | `Open`, `Create`, `Remove`, `Getenv` |
| **io** | I/O primitives | `Reader`, `Writer`, `Copy`, `ReadAll` |
| **bufio** | Buffered I/O | `Scanner`, `Reader`, `Writer` |
| **path/filepath** | File path manipulation | `Join`, `Clean`, `Walk`, `Match` |
| **flag** | Command-line flag parsing | `String`, `Int`, `Bool`, `Parse` |
| **os/exec** | External command execution | `Command`, `Run`, `Output` |

## 🚀 Quick Start Examples

### File Operations
```go
// Reading a file
content, err := os.ReadFile("config.txt")
if err != nil {
    log.Fatal(err)
}

// Writing a file
err = os.WriteFile("output.txt", []byte("Hello, World!"), 0644)
if err != nil {
    log.Fatal(err)
}

// Streaming large files
file, err := os.Open("large_file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

### Command-Line Flags
```go
import "flag"

var (
    name    = flag.String("name", "World", "Name to greet")
    count   = flag.Int("count", 1, "Number of greetings")
    verbose = flag.Bool("verbose", false, "Enable verbose output")
)

func main() {
    flag.Parse()
    
    for i := 0; i < *count; i++ {
        fmt.Printf("Hello, %s!\n", *name)
    }
}
```

### Directory Walking
```go
err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    
    if info.IsDir() {
        fmt.Printf("Directory: %s\n", path)
    } else {
        fmt.Printf("File: %s (%d bytes)\n", path, info.Size())
    }
    
    return nil
})
```

## 🧪 How to Practice

1. **Start with files**: Read `file_operations.go` for basic I/O
2. **Explore directories**: Study `directory_management.go` 
3. **Handle configuration**: Work through `environment_config.go`
4. **Build CLI tools**: Practice with `cli_tools.go`
5. **System interaction**: Understand `system_interfaces.go`
6. **Apply knowledge**: Complete `fileio_practice.go`

## ⚠️ Common Pitfalls

### 1. Resource Leaks
```go
// ❌ WRONG - File not closed
file, err := os.Open("file.txt")
data, err := io.ReadAll(file)  // Leak!

// ✅ CORRECT - Always close files
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close()
data, err := io.ReadAll(file)
```

### 2. Path Handling
```go
// ❌ WRONG - Hardcoded path separators
path := "data/files/config.txt"  // Won't work on Windows

// ✅ CORRECT - Use filepath.Join
path := filepath.Join("data", "files", "config.txt")
```

### 3. Permission Issues
```go
// ❌ WRONG - Wrong permissions
os.WriteFile("secret.txt", data, 0777)  // Too permissive

// ✅ CORRECT - Appropriate permissions
os.WriteFile("secret.txt", data, 0600)  // Owner read/write only
```

## 📁 File Permissions Reference

| Permission | Octal | Description |
|------------|-------|-------------|
| `---` | 0 | No permissions |
| `r--` | 4 | Read only |
| `rw-` | 6 | Read and write |
| `rwx` | 7 | Read, write, and execute |

Common combinations:
- `0644` - Owner: rw-, Group: r--, Others: r--
- `0755` - Owner: rwx, Group: r-x, Others: r-x  
- `0600` - Owner: rw-, Group: ---, Others: ---

## 🎯 Key Concepts to Master

### Error Handling Patterns
```go
// Check if file exists
if _, err := os.Stat("file.txt"); os.IsNotExist(err) {
    // File doesn't exist
}

// Handle different error types
file, err := os.Open("file.txt")
if err != nil {
    if os.IsPermission(err) {
        // Permission denied
    } else if os.IsNotExist(err) {
        // File not found
    } else {
        // Other error
    }
    return err
}
```

### Atomic File Operations
```go
// Write atomically by writing to temp file first
tmpFile, err := os.CreateTemp("", "config-*.tmp")
if err != nil {
    return err
}
defer os.Remove(tmpFile.Name())

_, err = tmpFile.Write(data)
if err != nil {
    return err
}

err = tmpFile.Close()
if err != nil {
    return err
}

// Atomic rename
return os.Rename(tmpFile.Name(), "config.txt")
```

## 📊 Performance Considerations

| Operation | Best For | Performance |
|-----------|----------|-------------|
| `os.ReadFile` | Small files (<1MB) | Fast, simple |
| `bufio.Scanner` | Line-by-line reading | Memory efficient |
| `io.Copy` | Large file copying | Streaming, low memory |
| `filepath.Walk` | Directory traversal | Recursive, can be slow |
| `os.ReadDir` | Directory listing | Fast, non-recursive |

## 🔗 What's Next

After mastering file I/O and system programming, you'll advance to **Phase 6: Web Development** where you'll build HTTP servers and database-driven applications!

## 🏗️ Real-World Applications

By the end of this week, you'll be able to build:
- **Log processors**: Parse and analyze log files
- **File converters**: Transform data between formats
- **System monitors**: Track file changes and system resources
- **CLI utilities**: Command-line tools with flags and configuration
- **Backup tools**: File synchronization and archiving
- **Configuration managers**: Environment-aware configuration systems

Ready to master Go's system programming capabilities! 📁⚡🐹
