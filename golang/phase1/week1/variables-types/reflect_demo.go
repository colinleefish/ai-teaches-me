package main

import (
	"fmt"
	"reflect"
)

func exerciseReflection() {
	fmt.Println("🎯 Go Reflection (reflect) Demo")
	fmt.Println("==============================\n")

	// Test data
	name := "Alice"
	age := 25
	price := 19.99
	isActive := true
	numbers := []int{1, 2, 3}
	config := map[string]string{"host": "localhost"}

	fmt.Println("=== Basic Type Inspection ===")
	fmt.Printf("Type of name: %s\n", reflect.TypeOf(name))
	fmt.Printf("Type of age: %s\n", reflect.TypeOf(age))
	fmt.Printf("Type of price: %s\n", reflect.TypeOf(price))
	fmt.Printf("Type of isActive: %s\n", reflect.TypeOf(isActive))
	fmt.Printf("Type of numbers: %s\n", reflect.TypeOf(numbers))
	fmt.Printf("Type of config: %s\n", reflect.TypeOf(config))

	fmt.Println("\n=== Value Inspection ===")
	fmt.Printf("Value of name: %v\n", reflect.ValueOf(name))
	fmt.Printf("Value of age: %v\n", reflect.ValueOf(age))
	fmt.Printf("Value of price: %v\n", reflect.ValueOf(price))
	fmt.Printf("Value of isActive: %v\n", reflect.ValueOf(isActive))

	fmt.Println("\n=== Kind (Underlying Type) ===")
	fmt.Printf("Kind of name: %s\n", reflect.TypeOf(name).Kind())
	fmt.Printf("Kind of age: %s\n", reflect.TypeOf(age).Kind())
	fmt.Printf("Kind of price: %s\n", reflect.TypeOf(price).Kind())
	fmt.Printf("Kind of isActive: %s\n", reflect.TypeOf(isActive).Kind())
	fmt.Printf("Kind of numbers: %s\n", reflect.TypeOf(numbers).Kind())
	fmt.Printf("Kind of config: %s\n", reflect.TypeOf(config).Kind())

	fmt.Println("\n=== Type Comparison ===")
	fmt.Printf("Is name a string? %t\n", reflect.TypeOf(name) == reflect.TypeOf(""))
	fmt.Printf("Is age an int? %t\n", reflect.TypeOf(age) == reflect.TypeOf(0))
	fmt.Printf("Is price a float64? %t\n", reflect.TypeOf(price) == reflect.TypeOf(0.0))

	fmt.Println("\n=== Struct Inspection ===")
	type Person struct {
		Name string
		Age  int
		City string
	}

	person := Person{Name: "Bob", Age: 30, City: "New York"}
	personType := reflect.TypeOf(person)
	personValue := reflect.ValueOf(person)

	fmt.Printf("Struct type: %s\n", personType)
	fmt.Printf("Number of fields: %d\n", personType.NumField())

	// Inspect each field
	for i := 0; i < personType.NumField(); i++ {
		field := personType.Field(i)
		value := personValue.Field(i)
		fmt.Printf("Field %d: %s (%s) = %v\n", i, field.Name, field.Type, value.Interface())
	}

	fmt.Println("\n=== Dynamic Function Calls ===")
	// Create a function dynamically
	funcType := reflect.TypeOf(func(a, b int) int { return 0 })
	funcValue := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		a := args[0].Int()
		b := args[1].Int()
		result := a + b
		return []reflect.Value{reflect.ValueOf(result)}
	})

	// Call the function
	args := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	results := funcValue.Call(args)
	fmt.Printf("Dynamic function result: %v\n", results[0])

	fmt.Println("\n=== Type Assertion with Reflection ===")
	var value interface{} = "hello"

	// Using reflection to check type
	if reflect.TypeOf(value) == reflect.TypeOf("") {
		fmt.Printf("Value is a string: %s\n", value)
	}

	// Using reflection to get value
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() == reflect.String {
		fmt.Printf("String value: %s\n", reflectValue.String())
	}

	fmt.Println("\n=== Slice and Map Inspection ===")
	sliceValue := reflect.ValueOf(numbers)
	fmt.Printf("Slice length: %d\n", sliceValue.Len())
	fmt.Printf("Slice capacity: %d\n", sliceValue.Cap())

	for i := 0; i < sliceValue.Len(); i++ {
		fmt.Printf("Element %d: %v\n", i, sliceValue.Index(i))
	}

	mapValue := reflect.ValueOf(config)
	fmt.Printf("Map keys: %v\n", mapValue.MapKeys())

	for _, key := range mapValue.MapKeys() {
		value := mapValue.MapIndex(key)
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	fmt.Println("\n=== When to Use Reflection ===")
	fmt.Println("✅ Use reflection when:")
	fmt.Println("   - You need to work with unknown types")
	fmt.Println("   - Building generic utilities")
	fmt.Println("   - Serialization/deserialization")
	fmt.Println("   - Testing and debugging")
	fmt.Println("   - Building frameworks or libraries")

	fmt.Println("\n❌ Avoid reflection when:")
	fmt.Println("   - You know the types at compile time")
	fmt.Println("   - Performance is critical")
	fmt.Println("   - Simple type assertions would work")
	fmt.Println("   - You can use generics instead")

	fmt.Println("\n=== Performance Note ===")
	fmt.Println("Reflection is slower than direct type operations.")
	fmt.Println("Use it only when you need dynamic behavior!")
}
