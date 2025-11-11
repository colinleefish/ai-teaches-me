package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Exercise 1: Basic Type Conversions
func exerciseBasicConversions() {
	fmt.Println("=== Exercise 1: Basic Type Conversions ===")

	// TODO: Convert int to float64
	// Hint: Use float64() function
	integer := 42
	floatValue := float64(integer)
	fmt.Printf("Int %d -> Float64: %f (type: %T)\n", integer, floatValue, floatValue)

	// TODO: Convert float64 to int (truncates decimal)
	// Hint: Use int() function
	pi := 3.14159
	intValue := int(pi)
	fmt.Printf("Float64 %f -> Int: %d (type: %T)\n", pi, intValue, intValue)

	// TODO: Convert int to string using strconv
	// Hint: Use strconv.Itoa() or strconv.FormatInt()
	number := 123
	stringNumber := strconv.Itoa(number)
	fmt.Printf("Int %d -> String: %s (type: %T)\n", number, stringNumber, stringNumber)

	// TODO: Convert string to int using strconv
	// Hint: Use strconv.Atoi() and handle error
	stringInput := "456"
	if parsedInt, err := strconv.Atoi(stringInput); err == nil {
		fmt.Printf("String %s -> Int: %d (type: %T)\n", stringInput, parsedInt, parsedInt)
	} else {
		fmt.Printf("Error parsing %s: %v\n", stringInput, err)
	}

	// TODO: Convert float64 to string
	// Hint: Use strconv.FormatFloat()
	price := 19.99
	priceString := strconv.FormatFloat(price, 'f', 2, 64)
	fmt.Printf("Float64 %f -> String: %s (type: %T)\n", price, priceString, priceString)
}

// Exercise 2: String Conversions
func exerciseStringConversions() {
	fmt.Println("\n=== Exercise 2: String Conversions ===")

	// TODO: Convert string to float64
	// Hint: Use strconv.ParseFloat()
	heightStr := "5.75"
	if height, err := strconv.ParseFloat(heightStr, 64); err == nil {
		fmt.Printf("String %s -> Float64: %f (type: %T)\n", heightStr, height, height)
	} else {
		fmt.Printf("Error parsing %s: %v\n", heightStr, err)
	}

	// TODO: Convert bool to string
	// Hint: Use strconv.FormatBool()
	isActive := true
	boolString := strconv.FormatBool(isActive)
	fmt.Printf("Bool %t -> String: %s (type: %T)\n", isActive, boolString, boolString)

	// TODO: Convert string to bool
	// Hint: Use strconv.ParseBool()
	boolStr := "false"
	if boolValue, err := strconv.ParseBool(boolStr); err == nil {
		fmt.Printf("String %s -> Bool: %t (type: %T)\n", boolStr, boolValue, boolValue)
	} else {
		fmt.Printf("Error parsing %s: %v\n", boolStr, err)
	}

	// TODO: Convert rune to string
	// Hint: Use string() function
	char := 'A'
	charString := string(char)
	fmt.Printf("Rune %c -> String: %s (type: %T)\n", char, charString, charString)

	// TODO: Convert string to slice of bytes
	// Hint: Use []byte() conversion
	text := "Hello"
	bytes := []byte(text)
	fmt.Printf("String %s -> Bytes: %v (type: %T)\n", text, bytes, bytes)
}

// Exercise 3: Numeric Type Conversions
func exerciseNumericConversions() {
	fmt.Println("\n=== Exercise 3: Numeric Type Conversions ===")

	// TODO: Convert between different int types
	// Hint: Use explicit type conversion
	var smallInt int8 = 100
	var bigInt int64 = int64(smallInt)
	fmt.Printf("Int8 %d -> Int64: %d (type: %T)\n", smallInt, bigInt, bigInt)

	// TODO: Convert int to uint
	// Hint: Be careful with negative numbers
	positiveInt := 42
	unsignedInt := uint(positiveInt)
	fmt.Printf("Int %d -> Uint: %d (type: %T)\n", positiveInt, unsignedInt, unsignedInt)

	// TODO: Convert float32 to float64
	// Hint: Use float64() function
	var smallFloat float32 = 3.14
	bigFloat := float64(smallFloat)
	fmt.Printf("Float32 %f -> Float64: %f (type: %T)\n", smallFloat, bigFloat, bigFloat)

	// TODO: Demonstrate overflow behavior
	// Hint: Try converting a large number to a smaller type
	largeInt := 1000
	var smallInt8 int8 = int8(largeInt)
	fmt.Printf("Large int %d -> Int8: %d (overflow occurred!)\n", largeInt, smallInt8)

	// TODO: Convert complex numbers
	// Hint: Use complex() function
	realPart := 3.0
	imagPart := 4.0
	complexNum := complex(realPart, imagPart)
	fmt.Printf("Real: %f, Imag: %f -> Complex: %v (type: %T)\n", realPart, imagPart, complexNum, complexNum)
}

// Exercise 4: Interface{} and Type Assertions
func exerciseInterfaceConversions() {
	fmt.Println("\n=== Exercise 4: Interface{} and Type Assertions ===")

	// TODO: Create interface{} variables
	// Hint: interface{} can hold any type
	var anything interface{} = 42
	fmt.Printf("Interface{} value: %v (type: %T)\n", anything, anything)

	anything = "hello"
	fmt.Printf("Interface{} value: %v (type: %T)\n", anything, anything)

	anything = 3.14
	fmt.Printf("Interface{} value: %v (type: %T)\n", anything, anything)

	// TODO: Type assertion (safe)
	// Hint: Use value, ok := interface{}.(Type) syntax
	var value interface{} = "test string"
	if str, ok := value.(string); ok {
		fmt.Printf("Type assertion successful: %s (type: %T)\n", str, str)
	} else {
		fmt.Println("Type assertion failed")
	}

	// TODO: Type assertion (unsafe - can panic)
	// Hint: Use value.(Type) syntax
	var number interface{} = 123
	intValue := number.(int)
	fmt.Printf("Unsafe assertion: %d (type: %T)\n", intValue, intValue)

	// TODO: Type switch
	// Hint: Use switch v := value.(type) syntax
	values := []interface{}{42, "hello", 3.14, true}
	for _, v := range values {
		switch val := v.(type) {
		case int:
			fmt.Printf("Int: %d\n", val)
		case string:
			fmt.Printf("String: %s\n", val)
		case float64:
			fmt.Printf("Float64: %f\n", val)
		case bool:
			fmt.Printf("Bool: %t\n", val)
		default:
			fmt.Printf("Unknown type: %T\n", val)
		}
	}
}

// Exercise 5: Custom Type Conversions
func exerciseCustomConversions() {
	fmt.Println("\n=== Exercise 5: Custom Type Conversions ===")

	// TODO: Define custom types
	type Celsius float64
	type Fahrenheit float64
	type Kelvin float64

	// TODO: Create conversion functions
	// Hint: Write functions to convert between temperature scales
	celsiusToFahrenheit := func(c Celsius) Fahrenheit {
		return Fahrenheit(c*9/5 + 32)
	}

	fahrenheitToCelsius := func(f Fahrenheit) Celsius {
		return Celsius((f - 32) * 5 / 9)
	}

	celsiusToKelvin := func(c Celsius) Kelvin {
		return Kelvin(c + 273.15)
	}

	// TODO: Test conversions
	roomTemp := Celsius(25.0)
	fahrenheit := celsiusToFahrenheit(roomTemp)
	kelvin := celsiusToKelvin(roomTemp)

	fmt.Printf("Room temperature: %.1f°C = %.1f°F = %.1fK\n", roomTemp, fahrenheit, kelvin)

	// TODO: Convert back
	backToCelsius := fahrenheitToCelsius(fahrenheit)
	fmt.Printf("Converted back: %.1f°F = %.1f°C\n", fahrenheit, backToCelsius)

	// TODO: Demonstrate type safety
	// Hint: Show that you can't mix types without explicit conversion
	// This would cause a compile error:
	// sum := roomTemp + fahrenheit  // Error: cannot add Celsius and Fahrenheit
	sum := float64(roomTemp) + float64(fahrenheit)
	fmt.Printf("Sum as float64: %.1f\n", sum)
}

// Exercise 6: Error Handling in Conversions
func exerciseErrorHandling() {
	fmt.Println("\n=== Exercise 6: Error Handling in Conversions ===")

	// TODO: Handle conversion errors properly
	// Hint: Always check errors from strconv functions
	testCases := []string{"123", "abc", "12.34", "-5", "99999999999999999999"}

	for _, test := range testCases {
		// Try to convert to int
		if intValue, err := strconv.Atoi(test); err == nil {
			fmt.Printf("✓ '%s' -> int: %d\n", test, intValue)
		} else {
			fmt.Printf("✗ '%s' -> error: %v\n", test, err)
		}

		// Try to convert to float64
		if floatValue, err := strconv.ParseFloat(test, 64); err == nil {
			fmt.Printf("✓ '%s' -> float64: %f\n", test, floatValue)
		} else {
			fmt.Printf("✗ '%s' -> error: %v\n", test, err)
		}
	}

	// TODO: Handle type assertion errors
	// Hint: Use the comma-ok idiom for safe type assertions
	interfaceValues := []interface{}{"hello", 42, 3.14, true, nil}

	for _, v := range interfaceValues {
		if v == nil {
			fmt.Printf("Value is nil\n")
			continue
		}

		if str, ok := v.(string); ok {
			fmt.Printf("✓ String: %s\n", str)
		} else if num, ok := v.(int); ok {
			fmt.Printf("✓ Int: %d\n", num)
		} else if flt, ok := v.(float64); ok {
			fmt.Printf("✓ Float64: %f\n", flt)
		} else if b, ok := v.(bool); ok {
			fmt.Printf("✓ Bool: %t\n", b)
		} else {
			fmt.Printf("✗ Unknown type: %T\n", v)
		}
	}
}

// Exercise 7: Advanced Conversions
func exerciseAdvancedConversions() {
	fmt.Println("\n=== Exercise 7: Advanced Conversions ===")

	// TODO: Convert between slices and arrays
	// Hint: Use copy() function or explicit conversion
	slice := []int{1, 2, 3, 4, 5}
	var array [5]int
	copy(array[:], slice)
	fmt.Printf("Slice %v -> Array: %v (type: %T)\n", slice, array, array)

	// TODO: Convert slice to array (Go 1.20+)
	// Hint: Use array conversion syntax
	if len(slice) == 5 {
		array2 := [5]int(slice)
		fmt.Printf("Slice %v -> Array: %v (type: %T)\n", slice, array2, array2)
	}

	// TODO: Convert between different slice types
	// Hint: Use reflection or manual conversion
	intSlice := []int{1, 2, 3, 4, 5}
	var floatSlice []float64 = make([]float64, len(intSlice))
	for i, v := range intSlice {
		floatSlice[i] = float64(v)
	}
	fmt.Printf("Int slice %v -> Float slice: %v\n", intSlice, floatSlice)

	// TODO: Convert struct to map using reflection
	// Hint: Use reflect package
	type Person struct {
		Name string
		Age  int
		City string
	}

	person := Person{Name: "Alice", Age: 30, City: "New York"}
	personMap := make(map[string]interface{})

	v := reflect.ValueOf(person)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		personMap[field.Name] = value.Interface()
	}

	fmt.Printf("Struct %+v -> Map: %v\n", person, personMap)
}

// Exercise 8: Python vs Go Conversions
func exercisePythonVsGoConversions() {
	fmt.Println("\n=== Exercise 8: Python vs Go Conversions ===")

	fmt.Println("Python vs Go Type Conversions:")
	fmt.Println("Python: int('123')")
	fmt.Println("Go:     strconv.Atoi('123')")

	fmt.Println("\nPython: str(42)")
	fmt.Println("Go:     strconv.Itoa(42)")

	fmt.Println("\nPython: float('3.14')")
	fmt.Println("Go:     strconv.ParseFloat('3.14', 64)")

	fmt.Println("\nPython: isinstance(value, str)")
	fmt.Println("Go:     value, ok := interface{}.(string)")

	// TODO: Demonstrate the differences
	pythonStyleString := "42"
	pythonStyleInt := 123

	// Go-style conversions
	if goInt, err := strconv.Atoi(pythonStyleString); err == nil {
		fmt.Printf("Python: int('%s') = %d\n", pythonStyleString, goInt)
	}

	goString := strconv.Itoa(pythonStyleInt)
	fmt.Printf("Python: str(%d) = '%s'\n", pythonStyleInt, goString)

	// Type checking
	var value interface{} = "hello"
	if str, ok := value.(string); ok {
		fmt.Printf("Python: isinstance('%s', str) = True\n", str)
	}
}

// RunConversionsExercises runs all conversion exercises
func RunConversionsExercises() {
	fmt.Println("🎯 Go Type Conversions Practice")
	fmt.Println("==============================\n")

	exerciseBasicConversions()
	exerciseStringConversions()
	exerciseNumericConversions()
	exerciseInterfaceConversions()
	exerciseCustomConversions()
	exerciseErrorHandling()
	exerciseAdvancedConversions()
	exercisePythonVsGoConversions()

	fmt.Println("\n✅ Type conversion exercises completed!")
	fmt.Println("\n💡 Key Takeaways:")
	fmt.Println("- Use strconv package for string conversions")
	fmt.Println("- Always handle conversion errors")
	fmt.Println("- Type assertions require comma-ok idiom")
	fmt.Println("- Custom types need explicit conversion")
	fmt.Println("- Reflection enables dynamic type operations")
}

func main() {
	RunConversionsExercises()
}
