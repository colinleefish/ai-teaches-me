// Week 9: JSON Encoding and Decoding
// This file demonstrates encoding/json package for data serialization

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
)

// Sample data structures for JSON examples
type Person struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Email    string    `json:"email,omitempty"`
	Active   bool      `json:"active"`
	Created  time.Time `json:"created"`
	Tags     []string  `json:"tags,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Company struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Founded   int      `json:"founded"`
	Employees []Person `json:"employees"`
	Address   Address  `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}

// TODO: Demonstrate basic JSON marshaling
func demonstrateJSONMarshaling() {
	fmt.Println("=== JSON Marshaling (Go to JSON) ===")
	
	// TODO: Create sample data
	person := Person{
		Name:    "Alice Johnson",
		Age:     28,
		Email:   "alice@example.com",
		Active:  true,
		Created: time.Now(),
		Tags:    []string{"developer", "golang", "backend"},
		Metadata: map[string]interface{}{
			"department": "engineering",
			"level":      "senior",
			"salary":     75000,
		},
	}
	
	// TODO: Basic marshaling with json.Marshal
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}
	
	fmt.Printf("Marshaled JSON: %s\n", string(jsonData))
	
	// TODO: Pretty printing with json.MarshalIndent
	prettyJSON, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Printf("Error with indent: %v\n", err)
		return
	}
	
	fmt.Printf("Pretty JSON:\n%s\n", string(prettyJSON))
	
	// TODO: Marshaling different data types
	// Numbers, booleans, arrays, slices, maps
	
	// TODO: Handling nil values and zero values
	var emptyPerson Person
	emptyJSON, _ := json.Marshal(emptyPerson)
	fmt.Printf("Empty struct JSON: %s\n", string(emptyJSON))
}

// TODO: Demonstrate basic JSON unmarshaling
func demonstrateJSONUnmarshaling() {
	fmt.Println("\n=== JSON Unmarshaling (JSON to Go) ===")
	
	// TODO: Sample JSON data
	jsonStr := `{
		"name": "Bob Smith",
		"age": 32,
		"email": "bob@company.com",
		"active": true,
		"created": "2023-01-15T10:30:00Z",
		"tags": ["manager", "team-lead"],
		"metadata": {
			"department": "sales",
			"region": "west-coast"
		}
	}`
	
	// TODO: Unmarshal into struct
	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}
	
	fmt.Printf("Unmarshaled person: %+v\n", person)
	
	// TODO: Unmarshal into map[string]interface{}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Printf("Error unmarshaling to map: %v\n", err)
		return
	}
	
	fmt.Printf("Unmarshaled map: %v\n", data)
	
	// TODO: Type assertions for map values
	if name, ok := data["name"].(string); ok {
		fmt.Printf("Name from map: %s\n", name)
	}
	
	if age, ok := data["age"].(float64); ok { // JSON numbers are float64
		fmt.Printf("Age from map: %.0f\n", age)
	}
}

// TODO: Demonstrate JSON struct tags
func demonstrateJSONTags() {
	fmt.Println("\n=== JSON Struct Tags ===")
	
	// TODO: Different tag options
	type Product struct {
		ID          int     `json:"id"`
		Name        string  `json:"product_name"`        // Custom field name
		Price       float64 `json:"price,string"`        // Convert to string
		InStock     bool    `json:"in_stock,omitempty"`  // Omit if zero value
		Description string  `json:"-"`                   // Never include
		Internal    string  `json:"internal,omitempty"`  // Omit if empty
		CreatedAt   time.Time `json:"created_at"`
	}
	
	// TODO: Test with different values
	products := []Product{
		{
			ID:          1,
			Name:        "Laptop",
			Price:       999.99,
			InStock:     true,
			Description: "This should not appear in JSON",
			Internal:    "internal-code-123",
			CreatedAt:   time.Now(),
		},
		{
			ID:        2,
			Name:      "Mouse",
			Price:     29.99,
			InStock:   false, // This will be omitted due to omitempty
			CreatedAt: time.Now(),
		},
	}
	
	for i, product := range products {
		jsonData, _ := json.MarshalIndent(product, "", "  ")
		fmt.Printf("Product %d JSON:\n%s\n", i+1, string(jsonData))
	}
	
	// TODO: Explain tag options:
	// - Custom field names
	// - omitempty for zero values
	// - string option for type conversion
	// - "-" to exclude fields
}

// TODO: Demonstrate custom JSON marshaling
func demonstrateCustomMarshaling() {
	fmt.Println("\n=== Custom JSON Marshaling ===")
	
	// TODO: Custom time format
	type Event struct {
		Name      string    `json:"name"`
		Timestamp CustomTime `json:"timestamp"`
		Duration  CustomDuration `json:"duration"`
	}
	
	// TODO: Custom time type with MarshalJSON method
	type CustomTime time.Time
	
	func (ct CustomTime) MarshalJSON() ([]byte, error) {
		// TODO: Custom time format
		formatted := time.Time(ct).Format("2006-01-02 15:04:05")
		return json.Marshal(formatted)
	}
	
	func (ct *CustomTime) UnmarshalJSON(data []byte) error {
		// TODO: Custom time parsing
		var timeStr string
		if err := json.Unmarshal(data, &timeStr); err != nil {
			return err
		}
		
		parsed, err := time.Parse("2006-01-02 15:04:05", timeStr)
		if err != nil {
			return err
		}
		
		*ct = CustomTime(parsed)
		return nil
	}
	
	// TODO: Custom duration type
	type CustomDuration time.Duration
	
	func (cd CustomDuration) MarshalJSON() ([]byte, error) {
		// TODO: Convert duration to human-readable string
		return json.Marshal(time.Duration(cd).String())
	}
	
	func (cd *CustomDuration) UnmarshalJSON(data []byte) error {
		// TODO: Parse duration string
		var durationStr string
		if err := json.Unmarshal(data, &durationStr); err != nil {
			return err
		}
		
		parsed, err := time.ParseDuration(durationStr)
		if err != nil {
			return err
		}
		
		*cd = CustomDuration(parsed)
		return nil
	}
	
	// TODO: Test custom marshaling
	event := Event{
		Name:      "Conference",
		Timestamp: CustomTime(time.Now()),
		Duration:  CustomDuration(2 * time.Hour),
	}
	
	jsonData, _ := json.MarshalIndent(event, "", "  ")
	fmt.Printf("Custom marshaled event:\n%s\n", string(jsonData))
	
	// TODO: Test unmarshaling
	var unmarshaledEvent Event
	json.Unmarshal(jsonData, &unmarshaledEvent)
	fmt.Printf("Unmarshaled event: %+v\n", unmarshaledEvent)
}

// TODO: Demonstrate JSON streaming
func demonstrateJSONStreaming() {
	fmt.Println("\n=== JSON Streaming ===")
	
	// TODO: Create sample data
	people := []Person{
		{Name: "Alice", Age: 25, Email: "alice@example.com", Active: true},
		{Name: "Bob", Age: 30, Email: "bob@example.com", Active: false},
		{Name: "Charlie", Age: 35, Email: "charlie@example.com", Active: true},
	}
	
	// TODO: Streaming encoding with json.NewEncoder
	fmt.Println("Streaming encoding to stdout:")
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ") // Pretty printing
	
	for _, person := range people {
		if err := encoder.Encode(person); err != nil {
			fmt.Printf("Encoding error: %v\n", err)
		}
	}
	
	// TODO: Streaming encoding to string builder
	var jsonOutput strings.Builder
	streamEncoder := json.NewEncoder(&jsonOutput)
	
	for _, person := range people {
		streamEncoder.Encode(person)
	}
	
	fmt.Printf("Streamed JSON:\n%s\n", jsonOutput.String())
	
	// TODO: Streaming decoding with json.NewDecoder
	jsonInput := `{"name":"Dave","age":28,"email":"dave@example.com","active":true}
{"name":"Eve","age":26,"email":"eve@example.com","active":false}`
	
	decoder := json.NewDecoder(strings.NewReader(jsonInput))
	
	fmt.Println("Streaming decoding:")
	for {
		var person Person
		err := decoder.Decode(&person)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Decoding error: %v\n", err)
			break
		}
		fmt.Printf("Decoded: %+v\n", person)
	}
}

// TODO: Demonstrate working with dynamic JSON
func demonstrateDynamicJSON() {
	fmt.Println("\n=== Dynamic JSON Handling ===")
	
	// TODO: Unknown structure JSON
	jsonData := `{
		"users": [
			{"name": "Alice", "age": 25, "skills": ["Go", "Python"]},
			{"name": "Bob", "age": 30, "location": {"city": "NYC", "country": "USA"}}
		],
		"total": 2,
		"metadata": {
			"version": "1.0",
			"generated": "2023-12-01T10:00:00Z"
		}
	}`
	
	// TODO: Parse into interface{}
	var data interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Printf("Error parsing dynamic JSON: %v\n", err)
		return
	}
	
	// TODO: Navigate dynamic JSON with type assertions
	navigateJSON := func(data interface{}, path string) interface{} {
		// TODO: Implement JSON path navigation
		// This is a simplified example
		if dataMap, ok := data.(map[string]interface{}); ok {
			return dataMap[path]
		}
		return nil
	}
	
	// TODO: Extract values safely
	if rootMap, ok := data.(map[string]interface{}); ok {
		// Get total
		if total, ok := rootMap["total"].(float64); ok {
			fmt.Printf("Total users: %.0f\n", total)
		}
		
		// Get users array
		if users, ok := rootMap["users"].([]interface{}); ok {
			fmt.Printf("Users count: %d\n", len(users))
			
			// Process each user
			for i, user := range users {
				if userMap, ok := user.(map[string]interface{}); ok {
					if name, ok := userMap["name"].(string); ok {
						fmt.Printf("User %d: %s\n", i+1, name)
					}
				}
			}
		}
	}
	
	// TODO: Convert back to JSON
	prettyJSON, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("Reformatted JSON:\n%s\n", string(prettyJSON))
}

// TODO: Demonstrate JSON validation and error handling
func demonstrateJSONValidation() {
	fmt.Println("\n=== JSON Validation and Error Handling ===")
	
	// TODO: Invalid JSON examples
	invalidJSONs := []string{
		`{"name": "Alice", "age": }`,           // Missing value
		`{"name": "Alice", "age": 25,}`,        // Trailing comma
		`{"name": "Alice" "age": 25}`,          // Missing comma
		`{"name": "Alice", "age": "invalid"}`,  // Type mismatch
	}
	
	for i, invalidJSON := range invalidJSONs {
		fmt.Printf("Testing invalid JSON %d:\n", i+1)
		var person Person
		err := json.Unmarshal([]byte(invalidJSON), &person)
		if err != nil {
			fmt.Printf("  Error: %v\n", err)
			
			// TODO: Check for specific error types
			if syntaxErr, ok := err.(*json.SyntaxError); ok {
				fmt.Printf("  Syntax error at offset %d\n", syntaxErr.Offset)
			}
			if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
				fmt.Printf("  Type error: cannot unmarshal %s into %s\n", 
					typeErr.Value, typeErr.Type)
			}
		}
	}
	
	// TODO: JSON validation function
	isValidJSON := func(data []byte) bool {
		// TODO: Implement JSON validation
		var result interface{}
		return json.Unmarshal(data, &result) == nil
	}
	
	// TODO: Test validation
	testData := `{"valid": true}`
	fmt.Printf("Is valid JSON: %v\n", isValidJSON([]byte(testData)))
}

// TODO: Demonstrate JSON performance optimizations
func demonstrateJSONPerformance() {
	fmt.Println("\n=== JSON Performance Considerations ===")
	
	// TODO: Large dataset for testing
	generateLargeDataset := func(size int) []Person {
		// TODO: Generate test data
		people := make([]Person, size)
		for i := 0; i < size; i++ {
			people[i] = Person{
				Name:   fmt.Sprintf("Person_%d", i),
				Age:    20 + (i % 50),
				Email:  fmt.Sprintf("person_%d@example.com", i),
				Active: i%2 == 0,
				Tags:   []string{"tag1", "tag2"},
			}
		}
		return people
	}
	
	// TODO: Compare marshaling methods
	data := generateLargeDataset(1000)
	
	// Method 1: json.Marshal
	timeOperation("json.Marshal", func() {
		_, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Marshal error: %v\n", err)
		}
	})
	
	// Method 2: json.NewEncoder with buffer
	timeOperation("json.NewEncoder", func() {
		var buf strings.Builder
		encoder := json.NewEncoder(&buf)
		encoder.Encode(data)
	})
	
	// TODO: Memory usage considerations
	// TODO: Streaming vs batch processing
	// TODO: Custom optimizations
}

// TODO: Demonstrate JSON schema validation (conceptual)
func demonstrateJSONSchema() {
	fmt.Println("\n=== JSON Schema Validation (Conceptual) ===")
	
	// TODO: Define expected schema structure
	type Schema struct {
		Type       string            `json:"type"`
		Properties map[string]Schema `json:"properties,omitempty"`
		Required   []string          `json:"required,omitempty"`
		Items      *Schema           `json:"items,omitempty"`
	}
	
	// TODO: Simple validation function
	validatePersonJSON := func(data []byte) error {
		// TODO: Implement basic JSON validation
		// This is a simplified example - real JSON schema validation
		// would use libraries like github.com/xeipuuv/gojsonschema
		
		var person Person
		err := json.Unmarshal(data, &person)
		if err != nil {
			return err
		}
		
		// Basic validation rules
		if person.Name == "" {
			return fmt.Errorf("name is required")
		}
		if person.Age < 0 || person.Age > 150 {
			return fmt.Errorf("age must be between 0 and 150")
		}
		
		return nil
	}
	
	// TODO: Test validation
	validJSON := `{"name": "Alice", "age": 25, "email": "alice@example.com"}`
	invalidJSON := `{"name": "", "age": -5}`
	
	fmt.Printf("Valid JSON validation: %v\n", validatePersonJSON([]byte(validJSON)))
	fmt.Printf("Invalid JSON validation: %v\n", validatePersonJSON([]byte(invalidJSON)))
}

// Helper function for timing operations
func timeOperation(name string, fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Printf("%s took: %v\n", name, elapsed)
}

// Helper function to pretty print any JSON-serializable data
func prettyPrint(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func main() {
	fmt.Println("🔄 Welcome to JSON Handling! 🔄")
	fmt.Println("This file teaches you JSON encoding and decoding in Go")
	
	// TODO: Implement each demonstration function
	// Start with basic marshaling and progress to advanced topics
	
	demonstrateJSONMarshaling()
	// demonstrateJSONUnmarshaling()
	// demonstrateJSONTags()
	// demonstrateCustomMarshaling()
	// demonstrateJSONStreaming()
	// demonstrateDynamicJSON()
	// demonstrateJSONValidation()
	// demonstrateJSONPerformance()
	// demonstrateJSONSchema()
	
	fmt.Println("\n🎉 Congratulations! You've mastered JSON handling in Go!")
	fmt.Println("Next: Learn HTTP client operations in http_client.go")
}

/*
🔍 Key Concepts to Remember:

1. **Marshaling**: Go structs → JSON strings (json.Marshal)
2. **Unmarshaling**: JSON strings → Go structs (json.Unmarshal)
3. **Struct Tags**: Control JSON field names and behavior
4. **Streaming**: Handle large JSON data efficiently
5. **Custom Types**: Implement MarshalJSON/UnmarshalJSON
6. **Error Handling**: Validate and handle JSON errors
7. **Performance**: Choose appropriate method for your use case

📋 Essential JSON Operations:
```go
// Basic marshaling/unmarshaling
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}

// Marshal to JSON
jsonData, err := json.Marshal(person)

// Unmarshal from JSON
var person Person
err := json.Unmarshal(jsonData, &person)

// Streaming
encoder := json.NewEncoder(writer)
decoder := json.NewDecoder(reader)
```

🏷️ Common Struct Tags:
- `json:"field_name"` - Custom field name
- `json:",omitempty"` - Omit zero values
- `json:"-"` - Exclude field
- `json:",string"` - Convert to/from string

🚨 Common Mistakes:
- Forgetting to export struct fields (lowercase)
- Not handling unmarshaling errors
- Incorrect struct tag syntax
- Type mismatches (JSON numbers are float64)
- Not using streaming for large datasets

🎯 Next Steps:
- Learn HTTP client operations for API integration
- Master context usage for timeouts and cancellation
- Practice with real-world JSON APIs
- Build data processing pipelines
*/
