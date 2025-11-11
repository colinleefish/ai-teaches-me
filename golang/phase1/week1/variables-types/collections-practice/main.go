package main

import (
	"fmt"
	"sort"
	"strings"
)

// Exercise 1: Arrays vs Slices
func exerciseArraysVsSlices() {
	fmt.Println("=== Exercise 1: Arrays vs Slices ===")

	// TODO: Create a fixed-size array of 5 integers
	// Hint: Use [5]int syntax
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v, Type: %T, Length: %d\n", numbers, numbers, len(numbers))

	// TODO: Create a slice of integers (dynamic size)
	// Hint: Use []int syntax
	var slice []int = []int{10, 20, 30, 40, 50}
	fmt.Printf("Slice: %v, Type: %T, Length: %d, Capacity: %d\n", slice, slice, len(slice), cap(slice))

	// TODO: Create a slice using make() with length 3 and capacity 5
	// Hint: make([]int, 3, 5)
	dynamicSlice := make([]int, 3, 5)
	fmt.Printf("Dynamic Slice: %v, Length: %d, Capacity: %d\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	// TODO: Append elements to the slice
	// Hint: Use append() function
	dynamicSlice = append(dynamicSlice, 100, 200)
	fmt.Printf("After append: %v, Length: %d, Capacity: %d\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))
}

// Exercise 2: Slice Operations
func exerciseSliceOperations() {
	fmt.Println("\n=== Exercise 2: Slice Operations ===")

	// TODO: Create a slice with numbers 1-10
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original: %v\n", numbers)

	// TODO: Get a slice from index 2 to 5 (exclusive)
	// Hint: numbers[2:5]
	subSlice := numbers[2:5]
	fmt.Printf("Sub-slice [2:5]: %v\n", subSlice)

	// TODO: Get first 3 elements
	// Hint: numbers[:3]
	firstThree := numbers[:3]
	fmt.Printf("First 3: %v\n", firstThree)

	// TODO: Get last 3 elements
	// Hint: numbers[7:] or numbers[len(numbers)-3:]
	lastThree := numbers[len(numbers)-3:]
	fmt.Printf("Last 3: %v\n", lastThree)

	// TODO: Copy a slice
	// Hint: Use copy() function
	copied := make([]int, len(numbers))
	copy(copied, numbers)
	fmt.Printf("Copied: %v\n", copied)

	// TODO: Sort the slice in descending order
	// Hint: Use sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Printf("Sorted descending: %v\n", numbers)
}

// Exercise 3: Maps (Go's Dictionaries)
func exerciseMaps() {
	fmt.Println("\n=== Exercise 3: Maps ===")

	// TODO: Create a map of string to int (like Python dict)
	// Hint: map[string]int{}
	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}
	fmt.Printf("Scores: %v\n", scores)

	// TODO: Add a new key-value pair
	scores["David"] = 88
	fmt.Printf("After adding David: %v\n", scores)

	// TODO: Update an existing value
	scores["Alice"] = 98
	fmt.Printf("After updating Alice: %v\n", scores)

	// TODO: Check if a key exists and get its value
	// Hint: Use the comma ok idiom: value, exists := map[key]
	if score, exists := scores["Eve"]; exists {
		fmt.Printf("Eve's score: %d\n", score)
	} else {
		fmt.Println("Eve not found in scores")
	}

	// TODO: Delete a key
	delete(scores, "Bob")
	fmt.Printf("After deleting Bob: %v\n", scores)

	// TODO: Create a map using make()
	// Hint: make(map[string]string)
	contacts := make(map[string]string)
	contacts["Alice"] = "alice@email.com"
	contacts["Bob"] = "bob@email.com"
	fmt.Printf("Contacts: %v\n", contacts)
}

// Exercise 4: String and Slice Operations
func exerciseStringSliceOperations() {
	fmt.Println("\n=== Exercise 4: String and Slice Operations ===")

	// TODO: Create a string
	text := "Hello, Go World!"
	fmt.Printf("Original text: %s\n", text)

	// TODO: Convert string to slice of bytes
	bytes := []byte(text)
	fmt.Printf("As bytes: %v\n", bytes)

	// TODO: Convert string to slice of runes (Unicode characters)
	runes := []rune(text)
	fmt.Printf("As runes: %v\n", runes)

	// TODO: Split string into slice
	// Hint: Use strings.Split()
	words := strings.Split(text, " ")
	fmt.Printf("Words: %v\n", words)

	// TODO: Join slice back into string
	// Hint: Use strings.Join()
	joined := strings.Join(words, "-")
	fmt.Printf("Joined with '-': %s\n", joined)

	// TODO: Create a slice of strings
	names := []string{"Alice", "Bob", "Charlie", "David"}
	fmt.Printf("Names: %v\n", names)

	// TODO: Filter names that start with 'A'
	var filteredNames []string
	for _, name := range names {
		if strings.HasPrefix(name, "A") {
			filteredNames = append(filteredNames, name)
		}
	}
	fmt.Printf("Names starting with 'A': %v\n", filteredNames)
}

// Exercise 5: Advanced Collections
func exerciseAdvancedCollections() {
	fmt.Println("\n=== Exercise 5: Advanced Collections ===")

	// TODO: Create a 2D slice (slice of slices)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Matrix: %v\n", matrix)

	// TODO: Access element at row 1, column 2
	element := matrix[1][2]
	fmt.Printf("Element at [1][2]: %d\n", element)

	// TODO: Create a map with slice values
	studentScores := map[string][]int{
		"Alice":   {95, 87, 92},
		"Bob":     {78, 85, 90},
		"Charlie": {88, 92, 85},
	}
	fmt.Printf("Student scores: %v\n", studentScores)

	// TODO: Calculate average score for each student
	for student, scores := range studentScores {
		sum := 0
		for _, score := range scores {
			sum += score
		}
		average := float64(sum) / float64(len(scores))
		fmt.Printf("%s's average: %.2f\n", student, average)
	}

	// TODO: Create a map with struct-like data
	type Person struct {
		Name string
		Age  int
		City string
	}

	people := map[int]Person{
		1: {Name: "Alice", Age: 25, City: "New York"},
		2: {Name: "Bob", Age: 30, City: "San Francisco"},
		3: {Name: "Charlie", Age: 28, City: "Boston"},
	}

	fmt.Printf("People: %v\n", people)

	// TODO: Find person by ID
	if person, exists := people[2]; exists {
		fmt.Printf("Person with ID 2: %s, %d years old, from %s\n",
			person.Name, person.Age, person.City)
	}
}

// Exercise 6: Collection Utilities
func exerciseCollectionUtilities() {
	fmt.Println("\n=== Exercise 6: Collection Utilities ===")

	// TODO: Create a slice and demonstrate various operations
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Printf("Original: %v\n", numbers)

	// TODO: Find the length and capacity
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// TODO: Check if slice is nil
	var nilSlice []int
	fmt.Printf("Nil slice is nil: %v\n", nilSlice == nil)

	// TODO: Create a slice with specific capacity
	// Hint: make([]int, 0, 10)
	preallocated := make([]int, 0, 10)
	fmt.Printf("Preallocated: %v, Length: %d, Capacity: %d\n",
		preallocated, len(preallocated), cap(preallocated))

	// TODO: Demonstrate slice growth
	for i := 0; i < 15; i++ {
		preallocated = append(preallocated, i)
		if i%5 == 0 {
			fmt.Printf("After %d appends: Length=%d, Capacity=%d\n",
				i+1, len(preallocated), cap(preallocated))
		}
	}

	// TODO: Create a map and demonstrate operations
	config := map[string]interface{}{
		"port":    8080,
		"host":    "localhost",
		"debug":   true,
		"timeout": 30.5,
	}

	fmt.Printf("Config: %v\n", config)

	// TODO: Iterate over map keys and values
	fmt.Println("Config entries:")
	for key, value := range config {
		fmt.Printf("  %s: %v (%T)\n", key, value, value)
	}

	// TODO: Get all keys from map
	var keys []string
	for key := range config {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Printf("Sorted keys: %v\n", keys)
}

// Exercise 7: Python vs Go Collections Comparison
func exercisePythonVsGoComparison() {
	fmt.Println("\n=== Exercise 7: Python vs Go Collections Comparison ===")

	fmt.Println("Python vs Go Collections:")
	fmt.Println("Python: my_list = [1, 2, 3]")
	fmt.Println("Go:     mySlice := []int{1, 2, 3}")

	fmt.Println("\nPython: my_dict = {'key': 'value'}")
	fmt.Println("Go:     myMap := map[string]string{'key': 'value'}")

	fmt.Println("\nPython: len(my_list)")
	fmt.Println("Go:     len(mySlice)")

	fmt.Println("\nPython: my_list.append(4)")
	fmt.Println("Go:     mySlice = append(mySlice, 4)")

	fmt.Println("\nPython: my_dict['new_key'] = 'new_value'")
	fmt.Println("Go:     myMap['new_key'] = 'new_value'")

	fmt.Println("\nPython: if 'key' in my_dict:")
	fmt.Println("Go:     if value, exists := myMap['key']; exists {")

	// TODO: Demonstrate the differences with actual code
	pythonStyleList := []int{1, 2, 3}
	pythonStyleDict := map[string]int{"a": 1, "b": 2}

	fmt.Printf("\nPython-style list: %v\n", pythonStyleList)
	fmt.Printf("Python-style dict: %v\n", pythonStyleDict)

	// Go-style operations
	pythonStyleList = append(pythonStyleList, 4) // Like list.append()
	if value, exists := pythonStyleDict["a"]; exists {
		fmt.Printf("Key 'a' exists with value: %d\n", value)
	}
}

// RunCollectionsExercises runs all collection exercises
func RunCollectionsExercises() {
	fmt.Println("🎯 Go Collections Practice")
	fmt.Println("========================\n")

	exerciseArraysVsSlices()
	exerciseSliceOperations()
	exerciseMaps()
	exerciseStringSliceOperations()
	exerciseAdvancedCollections()
	exerciseCollectionUtilities()
	exercisePythonVsGoComparison()

	fmt.Println("\n✅ Collections exercises completed!")
	fmt.Println("\n💡 Key Takeaways:")
	fmt.Println("- Arrays have fixed size, slices are dynamic")
	fmt.Println("- Maps are Go's equivalent of Python dictionaries")
	fmt.Println("- Use make() to preallocate slices and maps")
	fmt.Println("- Slices can grow automatically with append()")
	fmt.Println("- Maps use comma-ok idiom for safe access")
}

func main() {
	RunCollectionsExercises()
}
