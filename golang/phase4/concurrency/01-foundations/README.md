# Level 1: Foundations - The 4 Cornerstones

This directory contains the foundational building blocks of Go concurrency. Master these 4 cornerstones before moving to advanced patterns.

## 🧱 The 4 Cornerstones

| File            | Cornerstone    | Description                            |
| --------------- | -------------- | -------------------------------------- |
| `goroutines.go` | **Goroutines** | Lightweight concurrent execution units |
| `waitgroups.go` | **WaitGroups** | Coordination and synchronization       |
| `channels.go`   | **Channels**   | Safe communication between goroutines  |
| `select.go`     | **Select**     | Channel control flow and choice        |

## 🎯 How to Practice

Each file is a standalone learning module. Run them individually:

```bash
# Start with goroutines
go run goroutines.go

# Then coordination
go run waitgroups.go

# Then communication
go run channels.go

# Finally control flow
go run select.go
```

## 📚 Learning Order

Follow this exact sequence - each builds on the previous:

1. **goroutines.go** - Learn to create concurrent execution
2. **waitgroups.go** - Learn to coordinate multiple goroutines
3. **channels.go** - Learn to communicate safely between goroutines
4. **select.go** - Learn to control channel operations

## 🛠️ Implementation Approach

Each file contains:

- **TUTOR comments**: Explain concepts and when to use them
- **TODO comments**: Guide your implementation
- **Function signatures**: You complete the implementations

Start with the first function, implement it, test it, then move to the next.

## ✅ Success Criteria

You've mastered Level 1 when you can:

- Create goroutines confidently
- Coordinate them with WaitGroups
- Pass data safely through channels
- Use select for channel control flow

## 🔗 What's Next

After completing all 4 cornerstone files, advance to **Level 2: Core Concepts** to learn safety and advanced channel types!

---

**Note**: The linter may show "main redeclared" warnings because each file has its own main function. This is intentional - these files are meant to be run individually, not together as a package.
