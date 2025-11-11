# How to Run the Variables & Types Exercises 🚀

## 📁 Available Files

- `variables.go` - Variable declarations and scoping
- `types.go` - Basic types and operations  
- `collections.go` - Arrays, slices, and maps
- `conversions.go` - Type conversions and casting

## 🎯 Running Individual Exercises

### Option 1: Run Each File Separately
```bash
# Run variables exercises
go run variables.go

# Run types exercises  
go run types.go

# Run collections exercises (in separate directory)
cd collections-practice
go run main.go

# Run conversions exercises (in separate directory)
cd conversions-practice
go run main.go
```

### Option 2: Run Specific Functions
Since all files are in the same package, you can run specific functions by modifying the main function in any file.

For example, to run collections exercises, add this to any main function:
```go
RunCollectionsExercises()
```

### Option 3: Use the Pre-created Directories
The directories are already set up for you:
```bash
# Collections exercises
cd collections-practice
go run main.go

# Conversions exercises  
cd conversions-practice
go run main.go
```

## 🧪 Testing Your Solutions

Each file contains TODO comments with hints. Complete the exercises by:

1. **Read the TODO comments** - They contain hints and instructions
2. **Implement the code** - Fill in the missing parts
3. **Run the file** - See your output
4. **Compare with solutions** - Check `../solutions/` folder

## 📚 Exercise Overview

### Collections (`collections.go`)
- Arrays vs Slices
- Map operations (Go's dictionaries)
- String and slice operations
- Advanced collections (2D slices, maps with slices)
- Python vs Go comparisons

### Conversions (`conversions.go`)
- Basic type conversions (int ↔ float ↔ string)
- String conversions with error handling
- Interface{} and type assertions
- Custom type conversions
- Advanced conversions with reflection

## 🎓 Learning Goals

By completing these exercises, you'll master:
- ✅ Go's collection types (arrays, slices, maps)
- ✅ Type conversion patterns and error handling
- ✅ String manipulation and byte operations
- ✅ Interface{} usage and type assertions
- ✅ Python to Go translation patterns

## 💡 Tips

- **Start with variables.go** - Build your foundation
- **Practice collections** - These are fundamental to Go
- **Master conversions** - Essential for real-world programming
- **Compare with Python** - Use your Python knowledge as reference

Happy coding! 🐹 