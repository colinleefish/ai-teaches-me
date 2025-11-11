package main

import (
	"fmt"
)

// Exercise 1: Arrays vs Slices
// func exerciseArraysVsSlices() {
// 	fmt.Println("=== Exercise 1: Arrays vs Slices ===")

// 	// TODO: Create a fixed-size array of 5 integers
// 	// Hint: Use [5]int syntax
// 	// YOUR CODE HERE
// 	numbers := [5]int{0, 1, 2, 3, 4}
// 	fmt.Printf("Array: %v, Type: %T, Length: %d\n", numbers, numbers, len(numbers))

// 	// TODO: Create a slice of integers (dynamic size)
// 	// Hint: Use []int syntax
// 	// YOUR CODE HERE
// 	slice := []int{1, 2, 3, 4, 5}
// 	fmt.Printf("Slice: %v, Type: %T, Length: %d, Capacity: %d\n", slice, slice, len(slice), cap(slice))

// 	// TODO: Create a slice using make() with length 3 and capacity 5
// 	// Hint: make([]int, 3, 5)
// 	// YOUR CODE HERE
// 	sliceMadeByMake := make([]int, 3, 5)
// 	fmt.Printf("Dynamic Slice: %v, Length: %d, Capacity: %d\n", sliceMadeByMake, len(sliceMadeByMake), cap(sliceMadeByMake))

// 	// TODO: Append elements to the slice
// 	// Hint: Use append() function
// 	// YOUR CODE HERE
// 	dynamicSlice := make([]int, 0, 5)
// 	dynamicSlice = append(dynamicSlice, 1)
// 	dynamicSlice = append(dynamicSlice, 2)
// 	dynamicSlice = append(dynamicSlice, 3)
// 	dynamicSlice = append(dynamicSlice, 4)
// 	dynamicSlice = append(dynamicSlice, 5)
// 	fmt.Printf("After append: %v, Length: %d, Capacity: %d\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))
// }

// Exercise 2: Slice Operations
// func exerciseSliceOperations() {
// 	fmt.Println("\n=== Exercise 2: Slice Operations ===")

// 	// TODO: Create a slice with numbers 1-10
// 	// YOUR CODE HERE
// 	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	fmt.Printf("Original: %v\n", numbers)

// 	// TODO: Get a slice from index 2 to 5 (exclusive)
// 	// Hint: numbers[2:5]
// 	// YOUR CODE HERE
// 	subSlice := numbers[2:5]
// 	fmt.Printf("Sub-slice [2:5]: %v\n", subSlice)

// 	// TODO: Get first 3 elements
// 	// Hint: numbers[:3]
// 	// YOUR CODE HERE
// 	firstThree := numbers[:3]
// 	fmt.Printf("First 3: %v\n", firstThree)

// 	// TODO: Get last 3 elements
// 	// Hint: numbers[7:] or numbers[len(numbers)-3:]
// 	// YOUR CODE HERE
// 	lastThree := numbers[len(numbers)-3:]
// 	fmt.Printf("Last 3: %v\n", lastThree)

// 	// TODO: Copy a slice
// 	// Hint: Use copy() function
// 	// YOUR CODE HERE
// 	fmt.Println("numbers before copy", numbers)

// 	copied := make([]int, 3)
// 	copy(copied, numbers[2:5])
// 	fmt.Printf("Copied: %v\n", copied)

// 	// TODO: Sort the slice in descending order
// 	// Hint: Use sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
// 	// YOUR CODE HERE

// 	// intSlice := sort.IntSlice(numbers)
// 	// reverse := sort.Reverse(intSlice)
// 	// fmt.Println("reverse before sort: ", reverse)
// 	// sort.Sort(reverse)

// 	// fmt.Println("intSlice:", intSlice, " after sort:", reverse)
// 	// fmt.Printf("Sorted descending: %v\n", numbers)

// 	// create a slice of random integers
// 	r := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

// 	randomSlice := make([]int, 10)
// 	for i := range randomSlice {
// 		randomSlice[i] = r.Intn(100)
// 	}
// 	fmt.Println("randomSlice:", randomSlice)

// 	// now I want to sort the slice in ascending order, don't tell me the answer.

// 	ascendingSlice := make([]int, len(randomSlice))
// 	copy(ascendingSlice, randomSlice)

// 	slices.Sort(ascendingSlice)

// 	fmt.Println("ascendingSlice:", ascendingSlice)

// 	descendingSlice := make([]int, len(randomSlice))
// 	copy(descendingSlice, randomSlice)

// 	slices.SortFunc(descendingSlice, func(a, b int) int {
// 		return b - a
// 	})

// 	fmt.Println("descendingSlice:", descendingSlice)
// }

// // Exercise 3: Maps (Go's Dictionaries)
// func exerciseMaps() {
// 	fmt.Println("\n=== Exercise 3: Maps ===")

// 	// TODO: Create a map of string to int (like Python dict)
// 	// Hint: map[string]int{}
// 	// YOUR CODE HERE

// 	scores := make(map[string]int)
// 	scores["Alice"] = 100
// 	scores["Bob"] = 90
// 	scores["Charlie"] = 80
// 	fmt.Printf("Scores: %v\n", scores)

// 	// TODO: Add a new key-value pair
// 	// YOUR CODE HERE
// 	scores["David"] = 70
// 	fmt.Printf("After adding David: %v\n", scores)

// 	// TODO: Update an existing value
// 	// YOUR CODE HERE
// 	scores["Alice"] = 101
// 	fmt.Printf("After updating Alice: %v\n", scores)

// 	// TODO: Check if a key exists and get its value
// 	// Hint: Use the comma ok idiom: value, exists := map[key]
// 	// YOUR CODE HERE
// 	if score, exists := scores["Eve"]; exists {
// 		fmt.Printf("Eve's score: %d\n", score)
// 	} else {
// 		fmt.Println("Eve not found in scores")
// 	}

// 	// TODO: Delete a key
// 	// YOUR CODE HERE
// 	delete(scores, "Bob")
// 	fmt.Printf("After deleting Bob: %v\n", scores)

// 	// TODO: Create a map using make()
// 	// Hint: make(map[string]string)
// 	// YOUR CODE HERE
// 	contacts := map[string]string{}
// 	contacts["Alice"] = "alice@email.com"
// 	contacts["Bob"] = "bob@email.com"
// 	fmt.Printf("Contacts: %v\n", contacts)

// 	fmt.Println("--------------------------------")

// 	// TODO: sort a map by key
// 	ownerDogs := map[string]string{}

// 	ownerDogs["Alice"] = "Zed"
// 	ownerDogs["Colin"] = "Xanda"
// 	ownerDogs["Bob"] = "Yella"

// 	ownerNamesSorted := make([]string, 0, len(ownerDogs))

// 	for owner := range ownerDogs {
// 		ownerNamesSorted = append(ownerNamesSorted, owner)
// 	}

// 	slices.Sort(ownerNamesSorted)

// 	for _, owner := range ownerNamesSorted {
// 		fmt.Printf("%s: %s\n", owner, ownerDogs[owner])
// 	}

// 	ownerNamesSortedReverse := make([]string, 0, len(ownerDogs))

// 	for _, owner := range ownerNamesSorted {
// 		ownerNamesSortedReverse = append(ownerNamesSortedReverse, owner)
// 	}

// 	slices.SortFunc(ownerNamesSortedReverse, func(a, b string) int {
// 		return strings.Compare(b, a)
// 	})

// 	fmt.Println("--------------------------------")
// 	fmt.Println("reversed name sort: ")

// 	for _, owner := range ownerNamesSortedReverse {
// 		fmt.Printf("%s: %s\n", owner, ownerDogs[owner])
// 	}
// }

// // Exercise 4: String and Slice Operations
// func exerciseStringSliceOperations() {
// 	fmt.Println("\n=== Exercise 4: String and Slice Operations ===")

// 	// TODO: Create a string
// 	// YOUR CODE HERE
// 	helloWorld := "Hello, 世界!"
// 	fmt.Printf("Original text: %s\n", helloWorld)

// 	// TODO: Convert string to slice of bytes
// 	// YOUR CODE HERE
// 	bytes := []byte(helloWorld)
// 	fmt.Printf("As bytes: %v\n", bytes)
// 	ints := []int{}
// 	for _, b := range bytes {
// 		ints = append(ints, int(b))
// 	}
// 	fmt.Printf("As ints: %v\n", ints)

// 	// TODO: Convert string to slice of runes (Unicode characters)
// 	// YOUR CODE HERE
// 	runes := []rune(helloWorld)
// 	fmt.Printf("As runes: %v\n", runes)

// 	// TODO: Split string into slice
// 	// Hint: Use strings.Split()
// 	// YOUR CODE HERE
// 	newWords := "Alice Bob Charlie David Eve Frank Grace Hank Ivy Jack"
// 	words := strings.Split(newWords, " ")
// 	fmt.Printf("Words: %v\n", words)

// 	// TODO: Join slice back into string
// 	// Hint: Use strings.Join()
// 	// YOUR CODE HERE
// 	joined := strings.Join(words, "-")
// 	fmt.Printf("Joined with '-': %s\n", joined)

// 	// TODO: Create a slice of strings
// 	// YOUR CODE HERE
// 	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack"}
// 	fmt.Printf("Names: %v\n", names)

// 	// TODO: Filter names that start with 'A'
// 	// YOUR CODE HERE
// 	filteredNames := []string{}
// 	for _, name := range names {
// 		if strings.HasPrefix(name, "A") {
// 			filteredNames = append(filteredNames, name)
// 		}
// 	}
// 	fmt.Printf("Names starting with 'A': %v\n", filteredNames)
// }

// // Exercise 5: Advanced Collections
// func exerciseAdvancedCollections() {
// 	fmt.Println("\n=== Exercise 5: Advanced Collections ===")

// 	// TODO: Create a 2D slice (slice of slices)
// 	// YOUR CODE HERE

// 	matrix := [][]int{}
// 	matrix = append(matrix, []int{1, 2, 3})
// 	matrix = append(matrix, []int{4, 5, 6})
// 	matrix = append(matrix, []int{7, 8, 9})
// 	fmt.Printf("Matrix: %v\n", matrix)

// 	// TODO: Access element at row 1, column 2
// 	// YOUR CODE HERE
// 	element := matrix[1][2]
// 	fmt.Printf("Element at [1][2]: %d\n", element)

// 	// TODO: Create a map with slice values
// 	// YOUR CODE HERE

// 	studentScores := map[string][]int{}
// 	studentScores["Alicia"] = []int{100, 90, 80}
// 	studentScores["Bob"] = []int{90, 80, 70}
// 	studentScores["Charlie"] = []int{80, 70, 60}
// 	fmt.Printf("Student scores: %v\n", studentScores)

// 	// TODO: Calculate average score for each student
// 	// YOUR CODE HERE
// 	for student, scores := range studentScores {
// 		sum := 0
// 		for _, score := range scores {
// 			sum += score
// 		}
// 		average := float64(sum) / float64(len(scores))
// 		fmt.Printf("Average score for %s: %.2f\n", student, average)
// 	}

// 	// TODO: Create a map with struct-like data
// 	// YOUR CODE HERE
// 	type Person struct {
// 		Name string
// 		Age  int
// 		City string
// 	}

// 	// YOUR CODE HERE
// 	people := map[int]Person{}
// 	people[1] = Person{Name: "Alice", Age: 25, City: "New York"}
// 	people[2] = Person{Name: "Bob", Age: 30, City: "Los Angeles"}
// 	people[3] = Person{Name: "Charlie", Age: 35, City: "Chicago"}
// 	fmt.Printf("People: %v\n", people)

// 	// TODO: Find person by ID
// 	// YOUR CODE HERE
// 	if person, exists := people[2]; exists {
// 		fmt.Printf("Person with ID 2: %s, %d years old, from %s\n",
// 			person.Name, person.Age, person.City)
// 	}
// }

// // Exercise 6: Collection Utilities
// func exerciseCollectionUtilities() {
// 	fmt.Println("\n=== Exercise 6: Collection Utilities ===")

// 	// TODO: Create a slice and demonstrate various operations
// 	// YOUR CODE HERE
// 	numbers := []int{1, 2, 3, 4, 5}
// 	fmt.Printf("Original: %v\n", numbers)

// 	// TODO: Find the length and capacity
// 	// YOUR CODE HERE
// 	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

// 	// TODO: Check if slice is nil
// 	// YOUR CODE HERE
// 	nilSlice := []int{}
// 	fmt.Printf("Nil slice is nil: %v\n", nilSlice == nil)

// 	// TODO: Create a slice with specific capacity
// 	// Hint: make([]int, 0, 10)
// 	// YOUR CODE HERE

// 	preallocated := make([]int, 0, 10)
// 	fmt.Printf("Preallocated: %v, Length: %d, Capacity: %d\n",
// 		preallocated, len(preallocated), cap(preallocated))

// 	// TODO: Demonstrate slice growth
// 	// YOUR CODE HERE
// 	preallocated = make([]int, 0, 2)
// 	for i := 0; i < 15; i++ {
// 		preallocated = append(preallocated, i)
// 		if i%5 == 0 {
// 			fmt.Printf("After %d appends: Length=%d, Capacity=%d\n",
// 				i+1, len(preallocated), cap(preallocated))
// 		}
// 	}

// 	// TODO: Create a map and demonstrate operations
// 	// YOUR CODE HERE

// 	config := map[string]string{}
// 	config["env"] = "development"
// 	config["db_url"] = "localhost:5432"
// 	config["db_user"] = "postgres"
// 	config["db_password"] = "password"
// 	config["db_name"] = "postgres"

// 	fmt.Printf("Config: %v\n", config)

// 	// TODO: Iterate over map keys and values
// 	// YOUR CODE HERE
// 	fmt.Println("Config entries:")
// 	for key, value := range config {
// 		fmt.Printf("  %s: %v (%T)\n", key, value, value)
// 	}

// 	// TODO: Get all keys from map
// 	// YOUR CODE HERE

// 	keys := make([]string, 0, len(config))
// 	for key := range config {
// 		keys = append(keys, key)
// 	}
// 	slices.Sort(keys)
// 	fmt.Printf("Sorted keys: %v\n", keys)
// }

// // Exercise 7: Python vs Go Collections Comparison
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
	// YOUR CODE HERE

	pythonStyleList := []int{1, 2, 3}
	pythonStyleDict := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Printf("\nPython-style list: %v\n", pythonStyleList)
	fmt.Printf("Python-style dict: %v\n", pythonStyleDict)

	// Go-style operations
	// YOUR CODE HERE
	pythonStyleList = append(pythonStyleList, 4)
	pythonStyleDict["new_key"] = 4

	if value, exists := pythonStyleDict["a"]; exists {
		fmt.Printf("Key 'a' exists with value: %d\n", value)
	}
}

// RunCollectionsExercises runs all collection exercises
func main() {
	fmt.Println("🎯 Go Collections Practice")
	fmt.Println("========================\n")

	// exerciseArraysVsSlices()
	// exerciseSliceOperations()
	// exerciseMaps()
	// exerciseStringSliceOperations()
	// exerciseAdvancedCollections()
	// exerciseCollectionUtilities()
	exercisePythonVsGoComparison()

	fmt.Println("\n✅ Collections exercises completed!")
	fmt.Println("\n💡 Key Takeaways:")
	fmt.Println("- Arrays have fixed size, slices are dynamic")
	fmt.Println("- Maps are Go's equivalent of Python dictionaries")
	fmt.Println("- Use make() to preallocate slices and maps")
	fmt.Println("- Slices can grow automatically with append()")
	fmt.Println("- Maps use comma-ok idiom for safe access")
}
