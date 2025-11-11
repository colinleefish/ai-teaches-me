# Week 1 - Basics: Go Setup & First Programs 🚀

Welcome to your first week of Go! This section covers the fundamentals of Go development environment and your first programs.

## 📋 Learning Objectives

By the end of this section, you will:
- [ ] Have Go properly installed and configured
- [ ] Understand Go's project structure and modules
- [ ] Write, compile, and run Go programs
- [ ] Use essential Go commands (`go run`, `go build`, `gofmt`)
- [ ] Understand the difference between compiled and interpreted languages

## 🎯 Exercises Overview

### Exercise 1: Installation & Setup
**File**: `setup.go`
- Verify Go installation
- Create your first Go module
- Understand workspace structure

### Exercise 2: Hello World Variations
**File**: `hello.go`
- Basic Hello World
- Command-line arguments
- Formatted output

### Exercise 3: Go Commands Mastery
**File**: `commands.go`
- Using `go run` vs `go build`
- Code formatting with `gofmt`
- Understanding compilation

### Exercise 4: Package System
**File**: `packages.go`
- Creating and using packages
- Import statements
- Module system basics

## 🔧 Prerequisites

### Install Go
```bash
# macOS with Homebrew
brew install go

# Or download from https://golang.org/dl/
```

### Verify Installation
```bash
go version
# Should output: go version go1.21.x darwin/amd64 (or similar)
```

### Set up your workspace
```bash
cd phase1
go mod init phase1-foundation
```

## 💻 Exercise Instructions

### Exercise 1: Setup Verification
1. Open `setup.go`
2. Complete the functions to demonstrate Go installation
3. Run with: `go run setup.go`
4. Expected output: Your Go version and environment info

### Exercise 2: Hello World Evolution
1. Open `hello.go`
2. Complete the progression from simple to advanced hello world
3. Test with: `go run hello.go`
4. Try with arguments: `go run hello.go Alice Bob`

### Exercise 3: Go Commands Practice
1. Open `commands.go`
2. Follow the instructions to practice different Go commands
3. Build the program: `go build commands.go`
4. Run the executable: `./commands`

### Exercise 4: Package Understanding
1. Open `packages.go`
2. Complete the package import and usage examples
3. Run and observe the output

## 🧪 Testing Your Work

Run the test suite:
```bash
go test
```

All tests should pass before moving to the next section.

## 📊 Self-Assessment

Rate your understanding (1-5 scale):
- [ ] Go installation and setup: ___/5
- [ ] Basic Go program structure: ___/5
- [ ] Go commands (run, build, fmt): ___/5
- [ ] Package system basics: ___/5

**Target**: All items should be 4/5 or higher before proceeding.

## 🐍 Python vs Go Quick Reference

| Concept | Python | Go |
|---------|--------|-----|
| **Run Program** | `python script.py` | `go run main.go` |
| **Import** | `import math` | `import "math"` |
| **Package** | `__init__.py` | `package main` |
| **Entry Point** | `if __name__ == "__main__"` | `func main()` |
| **Formatting** | `black script.py` | `gofmt -w main.go` |

## 🚀 Ready to Start?

1. Complete each exercise file in order
2. Run tests to verify your solutions
3. Check your self-assessment scores
4. Move to `../variables-types/` when ready

Let's get coding! 🐹 