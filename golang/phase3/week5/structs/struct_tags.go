package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ===== BASIC STRUCT TAGS =====

// UserAPI with JSON tags for demonstration
type UserAPI struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`             // Never serialize
	Age      int    `json:"age,omitempty"` // Omit if zero value
}

func demonstrateBasicJSONTags() {
	fmt.Println("=== Basic JSON Tags ===")

	user := UserAPI{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "secret123",
		Age:      0, // Will be omitted due to omitempty
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}

	fmt.Printf("JSON output: %s\n", string(jsonData))
	// Output: {"id":1,"name":"Alice","email":"alice@example.com"}
	// Note: Password is excluded, Age is omitted (zero value + omitempty)

	// Unmarshal from JSON
	jsonInput := `{"id":2,"name":"Bob","email":"bob@example.com","age":30}`
	var newUser User
	err = json.Unmarshal([]byte(jsonInput), &newUser)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}

	fmt.Printf("Unmarshaled user: %+v\n", newUser)
}

// ===== ADVANCED JSON TAGS =====

// type Product struct {
// 	ID        int                    `json:"id"`
// 	Name      string                 `json:"product_name"` // Different field name
// 	Price     float64                `json:"price,string"` // Convert to/from string
// 	InStock   bool                   `json:"in_stock"`
// 	Tags      []string               `json:"tags,omitempty"` // Omit empty slice
// 	Metadata  map[string]interface{} `json:"metadata,omitempty"`
// 	CreatedAt string                 `json:"created_at,omitempty"`
// 	UpdatedAt *string                `json:"updated_at,omitempty"` // Pointer for null handling
// }

func demonstrateAdvancedJSONTags() {
	fmt.Println("\n=== Advanced JSON Tags ===")

	product := Product{
		ID:      101,
		Name:    "Laptop",
		Price:   999.99,
		InStock: true,
		Tags:    []string{"electronics", "computers"},
		Metadata: map[string]interface{}{
			"brand":    "TechCorp",
			"warranty": 2,
		},
		CreatedAt: "2023-01-15T10:30:00Z",
		// UpdatedAt is nil (will be omitted)
	}

	jsonData, _ := json.Marshal(product)
	fmt.Printf("Product JSON: %s\n", string(jsonData))

	// Demonstrating string conversion for numbers
	jsonWithStringPrice := `{"id":102,"product_name":"Phone","price":"599.99","in_stock":true}`
	var newProduct Product
	json.Unmarshal([]byte(jsonWithStringPrice), &newProduct)
	fmt.Printf("Product with string price: %+v\n", newProduct)
}

// ===== MULTIPLE TAG TYPES =====

// User with multiple tag types
type EmployeeTags struct {
	ID          int     `json:"id" db:"employee_id" xml:"id" validate:"required"`
	FirstName   string  `json:"first_name" db:"first_name" xml:"firstName" validate:"required,min=2"`
	LastName    string  `json:"last_name" db:"last_name" xml:"lastName" validate:"required,min=2"`
	Email       string  `json:"email" db:"email" xml:"email" validate:"required,email"`
	Age         int     `json:"age" db:"age" xml:"age" validate:"gte=18,lte=65"`
	Department  string  `json:"department" db:"dept" xml:"department" validate:"required"`
	Salary      float64 `json:"-" db:"salary" xml:"-" validate:"gt=0"` // Hidden in JSON/XML
	PhoneNumber string  `json:"phone,omitempty" db:"phone" xml:"phone,omitempty"`
}

func demonstrateMultipleTags() {
	fmt.Println("\n=== Multiple Tag Types ===")

	emp := EmployeeTags{
		ID:          1,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@company.com",
		Age:         30,
		Department:  "Engineering",
		Salary:      75000.00,
		PhoneNumber: "+1-555-0123",
	}

	// JSON serialization
	jsonData, _ := json.Marshal(emp)
	fmt.Printf("JSON: %s\n", string(jsonData))

	// Show how we might use other tags (simulated)
	fmt.Println("\nTag inspection:")
	t := reflect.TypeOf(emp)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		dbTag := field.Tag.Get("db")
		validateTag := field.Tag.Get("validate")

		fmt.Printf("Field %s:\n", field.Name)
		if jsonTag != "" {
			fmt.Printf("  JSON: %s\n", jsonTag)
		}
		if dbTag != "" {
			fmt.Printf("  DB: %s\n", dbTag)
		}
		if validateTag != "" {
			fmt.Printf("  Validation: %s\n", validateTag)
		}
		fmt.Println()
	}
}

// ===== CUSTOM TAG PROCESSING =====

// Custom validation based on tags
func validateStruct(s interface{}) []string {
	var errors []string
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		validateTag := field.Tag.Get("validate")

		if validateTag == "" {
			continue
		}

		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			if err := validateField(field.Name, value, rule); err != "" {
				errors = append(errors, err)
			}
		}
	}

	return errors
}

func validateField(fieldName string, value reflect.Value, rule string) string {
	switch {
	case rule == "required":
		if isZeroValue(value) {
			return fmt.Sprintf("%s is required", fieldName)
		}

	case strings.HasPrefix(rule, "min="):
		minStr := strings.TrimPrefix(rule, "min=")
		min, _ := strconv.Atoi(minStr)
		if value.Kind() == reflect.String && len(value.String()) < min {
			return fmt.Sprintf("%s must be at least %d characters", fieldName, min)
		}

	case rule == "email":
		if value.Kind() == reflect.String && !strings.Contains(value.String(), "@") {
			return fmt.Sprintf("%s must be a valid email", fieldName)
		}

	case strings.HasPrefix(rule, "gte="):
		gteStr := strings.TrimPrefix(rule, "gte=")
		gte, _ := strconv.Atoi(gteStr)
		if value.Kind() == reflect.Int && int(value.Int()) < gte {
			return fmt.Sprintf("%s must be greater than or equal to %d", fieldName, gte)
		}
	}

	return ""
}

func isZeroValue(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Bool:
		return !value.Bool()
	default:
		return value.IsZero()
	}
}

func demonstrateCustomValidation() {
	fmt.Println("\n=== Custom Tag Processing ===")

	// Valid employee
	validEmp := EmployeeTags{
		ID:         1,
		FirstName:  "Alice",
		LastName:   "Smith",
		Email:      "alice@company.com",
		Age:        28,
		Department: "Marketing",
		Salary:     60000,
	}

	errors := validateStruct(validEmp)
	if len(errors) == 0 {
		fmt.Println("Valid employee: No errors")
	} else {
		fmt.Printf("Validation errors: %v\n", errors)
	}

	// Invalid employee
	invalidEmp := EmployeeTags{
		ID:        0,               // Required field is zero
		FirstName: "A",             // Too short
		LastName:  "",              // Required field is empty
		Email:     "invalid-email", // Invalid email format
		Age:       17,              // Below minimum age
		// Department missing (required)
		Salary: 0,
	}

	errors = validateStruct(invalidEmp)
	fmt.Printf("\nInvalid employee errors:\n")
	for _, err := range errors {
		fmt.Printf("- %s\n", err)
	}
}

// ===== TAG-BASED SERIALIZATION =====

type APIResponse struct {
	Success   bool        `json:"success" xml:"success"`
	Message   string      `json:"message,omitempty" xml:"message,omitempty"`
	Data      interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Error     string      `json:"error,omitempty" xml:"error,omitempty"`
	Timestamp string      `json:"timestamp" xml:"timestamp"`
}

type UserProfile struct {
	ID       int    `json:"id" xml:"id"`
	Username string `json:"username" xml:"username"`
	FullName string `json:"full_name" xml:"fullName"`
	Avatar   string `json:"avatar_url,omitempty" xml:"avatarUrl,omitempty"`
}

// Helper function to create field mapping from tags
func getFieldMapping(s interface{}, tagName string) map[string]string {
	mapping := make(map[string]string)
	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)

		if tag != "" && tag != "-" {
			// Handle complex tag values like "name,omitempty"
			tagParts := strings.Split(tag, ",")
			if len(tagParts) > 0 && tagParts[0] != "" {
				mapping[field.Name] = tagParts[0]
			}
		}
	}

	return mapping
}

func demonstrateTagBasedSerialization() {
	fmt.Println("\n=== Tag-Based Serialization ===")

	profile := UserProfile{
		ID:       123,
		Username: "alice_dev",
		FullName: "Alice Developer",
		Avatar:   "https://example.com/avatar.jpg",
	}

	response := APIResponse{
		Success:   true,
		Data:      profile,
		Timestamp: "2023-01-15T10:30:00Z",
	}

	// JSON serialization
	jsonData, _ := json.Marshal(response)
	fmt.Printf("JSON Response:\n%s\n", string(jsonData))

	// Show field mappings
	jsonMapping := getFieldMapping(UserProfile{}, "json")
	fmt.Printf("\nJSON field mappings for UserProfile:\n")
	for field, jsonName := range jsonMapping {
		fmt.Printf("  %s -> %s\n", field, jsonName)
	}

	xmlMapping := getFieldMapping(UserProfile{}, "xml")
	fmt.Printf("\nXML field mappings for UserProfile:\n")
	for field, xmlName := range xmlMapping {
		fmt.Printf("  %s -> %s\n", field, xmlName)
	}
}

// ===== DYNAMIC TAG INSPECTION =====

func inspectStructTags(s interface{}) {
	fmt.Println("\n=== Dynamic Tag Inspection ===")

	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fmt.Printf("Inspecting struct: %s\n", t.Name())
	fmt.Println(strings.Repeat("-", 50))

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s (Type: %s)\n", field.Name, field.Type)

		// Get all tags
		tag := field.Tag
		if tag == "" {
			fmt.Println("  No tags")
			continue
		}

		// Parse all tags
		tagStr := string(tag)
		fmt.Printf("  Raw tag: `%s`\n", tagStr)

		// Extract specific tag types
		commonTags := []string{"json", "xml", "db", "validate", "form"}
		for _, tagType := range commonTags {
			if value := tag.Get(tagType); value != "" {
				fmt.Printf("  %s: %s\n", tagType, value)
			}
		}
		fmt.Println()
	}
}

func demonstrateTagInspection() {
	fmt.Println("\n=== Tag Inspection Examples ===")

	// Inspect different structs
	inspectStructTags(EmployeeTags{})
	inspectStructTags(Product{})
}

// ===== MAIN DEMO FUNCTION =====

func runStructTagsDemo() {
	fmt.Println("🏷️ Go Struct Tags Tutorial")
	fmt.Println("============================")

	demonstrateBasicJSONTags()
	demonstrateAdvancedJSONTags()
	demonstrateMultipleTags()
	demonstrateCustomValidation()
	demonstrateTagBasedSerialization()
	demonstrateTagInspection()

	fmt.Println("\n✅ Struct tags concepts covered!")
	fmt.Println("\n🎯 Key Points:")
	fmt.Println("- Struct tags provide metadata for fields")
	fmt.Println("- JSON tags control serialization behavior")
	fmt.Println("- Use `omitempty` to skip zero values")
	fmt.Println("- Use `-` to exclude fields from serialization")
	fmt.Println("- Multiple tag types can coexist on the same field")
	fmt.Println("- Tags enable custom validation and processing")
	fmt.Println("- Reflection allows runtime tag inspection")
	fmt.Println("- Tags are the foundation for many Go libraries")
}
