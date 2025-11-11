// Week 9: Text Formatting and String Manipulation
// This file demonstrates fmt, strings, strconv, and unicode packages

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// TODO: Demonstrate fmt package for formatting
func demonstrateFmtFormatting() {
	fmt.Println("=== Fmt Package Formatting ===")
	
	// TODO: Basic print functions
	// fmt.Print(), fmt.Println(), fmt.Printf()
	name := "Alice"
	age := 30
	score := 95.7
	
	// TODO: Different print methods
	// Show differences between Print, Println, and Printf
	
	// TODO: Common format verbs
	// %v (default format), %+v (with field names), %#v (Go representation)
	// %T (type), %t (boolean), %d (decimal), %f (float), %s (string)
	// %q (quoted string), %x (hex), %o (octal), %b (binary)
	
	// TODO: Width and precision formatting
	// %5d (width), %.2f (precision), %5.2f (width and precision)
	// %-5d (left align), %05d (zero padding)
	
	// TODO: String formatting functions
	// fmt.Sprintf() for creating formatted strings
	// fmt.Errorf() for creating formatted errors
	
	fmt.Printf("Name: %s, Age: %d, Score: %.1f\n", name, age, score)
}

// TODO: Demonstrate advanced fmt formatting
func demonstrateAdvancedFormatting() {
	fmt.Println("\n=== Advanced Fmt Formatting ===")
	
	// TODO: Custom types and Stringer interface
	type Person struct {
		Name string
		Age  int
	}
	
	// TODO: Implement fmt.Stringer interface
	// func (p Person) String() string {
	//     return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
	// }
	
	// TODO: Formatting slices and maps
	numbers := []int{1, 2, 3, 4, 5}
	data := map[string]int{"apples": 5, "oranges": 3}
	
	// TODO: Show different formatting options for collections
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Data: %v\n", data)
	
	// TODO: Input formatting with fmt.Scanf family
	// fmt.Scanf(), fmt.Sscanf(), fmt.Fscanf()
	
	// TODO: Error handling in format operations
}

// TODO: Demonstrate strings package
func demonstrateStringsPackage() {
	fmt.Println("\n=== Strings Package ===")
	
	text := "Hello, World! Welcome to Go programming."
	
	// TODO: String inspection functions
	// strings.Contains(), strings.HasPrefix(), strings.HasSuffix()
	// strings.Index(), strings.LastIndex(), strings.Count()
	
	// TODO: String modification functions
	// strings.ToUpper(), strings.ToLower(), strings.Title()
	// strings.TrimSpace(), strings.Trim(), strings.TrimPrefix()
	
	// TODO: String splitting and joining
	// strings.Split(), strings.SplitN(), strings.Fields()
	// strings.Join()
	
	// TODO: String replacement
	// strings.Replace(), strings.ReplaceAll()
	// strings.Replacer for multiple replacements
	
	// TODO: String building for performance
	// strings.Builder for efficient string concatenation
	
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Upper: %s\n", strings.ToUpper(text))
	fmt.Printf("Words: %v\n", strings.Fields(text))
}

// TODO: Demonstrate string performance patterns
func demonstrateStringPerformance() {
	fmt.Println("\n=== String Performance Patterns ===")
	
	// TODO: Compare string concatenation methods
	// 1. Simple concatenation with +
	// 2. fmt.Sprintf()
	// 3. strings.Join()
	// 4. strings.Builder
	
	// TODO: Benchmark different approaches
	// Show when to use each method
	
	// TODO: Demonstrate strings.Builder
	var builder strings.Builder
	words := []string{"Hello", "World", "from", "Go"}
	
	for i, word := range words {
		builder.WriteString(word)
		if i < len(words)-1 {
			builder.WriteString(" ")
		}
	}
	
	result := builder.String()
	fmt.Printf("Built string: %s\n", result)
	
	// TODO: Show memory efficiency of Builder vs concatenation
}

// TODO: Demonstrate strconv package
func demonstrateStrconvPackage() {
	fmt.Println("\n=== Strconv Package ===")
	
	// TODO: String to number conversions
	// strconv.Atoi(), strconv.Itoa()
	// strconv.ParseInt(), strconv.ParseFloat(), strconv.ParseBool()
	
	// TODO: Number to string conversions
	// strconv.FormatInt(), strconv.FormatFloat(), strconv.FormatBool()
	
	// TODO: Base conversions
	// Binary, octal, hexadecimal conversions
	
	// TODO: Error handling in conversions
	// Handle conversion errors properly
	
	// Examples
	numStr := "42"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Error converting %s to int: %v\n", numStr, err)
	} else {
		fmt.Printf("Converted %s to %d\n", numStr, num)
	}
	
	// TODO: Advanced parsing with ParseInt
	// strconv.ParseInt(s, base, bitSize)
	
	// TODO: Quoting and unquoting strings
	// strconv.Quote(), strconv.Unquote()
	// strconv.QuoteToASCII(), strconv.QuoteToGraphic()
}

// TODO: Demonstrate unicode package
func demonstrateUnicodePackage() {
	fmt.Println("\n=== Unicode Package ===")
	
	text := "Hello, 世界! 123 @#$"
	
	// TODO: Unicode character classification
	// unicode.IsLetter(), unicode.IsDigit(), unicode.IsSpace()
	// unicode.IsUpper(), unicode.IsLower(), unicode.IsPunct()
	
	// TODO: Unicode normalization and transformation
	// unicode.ToUpper(), unicode.ToLower(), unicode.ToTitle()
	
	// TODO: UTF-8 handling with utf8 package
	// utf8.ValidString(), utf8.RuneCount(), utf8.DecodeRune()
	
	fmt.Printf("Text: %s\n", text)
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(text))
	fmt.Printf("Byte count: %d\n", len(text))
	
	// TODO: Iterate over runes vs bytes
	fmt.Println("Character analysis:")
	for i, r := range text {
		fmt.Printf("  Index %d: '%c' - Letter: %v, Digit: %v\n", 
			i, r, unicode.IsLetter(r), unicode.IsDigit(r))
	}
}

// TODO: Demonstrate regular expressions
func demonstrateRegularExpressions() {
	fmt.Println("\n=== Regular Expressions ===")
	
	// TODO: Basic regexp operations
	// regexp.Match(), regexp.MatchString()
	// regexp.Compile(), regexp.MustCompile()
	
	text := "Contact us at support@example.com or sales@company.org"
	
	// TODO: Email extraction example
	emailPattern := `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`
	re := regexp.MustCompile(emailPattern)
	
	// TODO: Find operations
	// FindString(), FindAllString(), FindStringSubmatch()
	
	emails := re.FindAllString(text, -1)
	fmt.Printf("Found emails: %v\n", emails)
	
	// TODO: Replace operations
	// ReplaceAllString(), ReplaceAllStringFunc()
	
	// TODO: Compile vs MustCompile
	// Error handling in regexp compilation
	
	// TODO: Named capture groups
	// Submatch operations
}

// TODO: Demonstrate string validation patterns
func demonstrateStringValidation() {
	fmt.Println("\n=== String Validation Patterns ===")
	
	// TODO: Common validation functions
	// Email, phone, URL, credit card validation
	
	validateEmail := func(email string) bool {
		// TODO: Implement email validation using regexp
		return false
	}
	
	validatePhone := func(phone string) bool {
		// TODO: Implement phone number validation
		return false
	}
	
	validateURL := func(url string) bool {
		// TODO: Implement URL validation
		return false
	}
	
	// TODO: Custom validation using strings and unicode packages
	
	// TODO: Test validation functions
	testEmails := []string{
		"user@example.com",
		"invalid.email",
		"test@domain",
	}
	
	for _, email := range testEmails {
		valid := validateEmail(email)
		fmt.Printf("Email %s is valid: %v\n", email, valid)
	}
}

// TODO: Demonstrate text processing utilities
func demonstrateTextProcessing() {
	fmt.Println("\n=== Text Processing Utilities ===")
	
	// TODO: Word counting
	countWords := func(text string) int {
		// TODO: Implement word counting
		return 0
	}
	
	// TODO: Line processing
	processLines := func(text string) []string {
		// TODO: Process text line by line
		return nil
	}
	
	// TODO: Text cleaning
	cleanText := func(text string) string {
		// TODO: Remove extra spaces, normalize whitespace
		return ""
	}
	
	// TODO: Case conversion utilities
	toCamelCase := func(text string) string {
		// TODO: Convert to camelCase
		return ""
	}
	
	toSnakeCase := func(text string) string {
		// TODO: Convert to snake_case
		return ""
	}
	
	// TODO: Test text processing functions
	sample := "  Hello   World  \n  Welcome to Go  "
	fmt.Printf("Original: %q\n", sample)
	fmt.Printf("Cleaned: %q\n", cleanText(sample))
}

// TODO: Demonstrate common string algorithms
func demonstrateStringAlgorithms() {
	fmt.Println("\n=== String Algorithms ===")
	
	// TODO: String searching algorithms
	// Boyer-Moore, KMP, naive search
	
	// TODO: String distance algorithms
	// Levenshtein distance, Hamming distance
	
	levenshteinDistance := func(s1, s2 string) int {
		// TODO: Implement Levenshtein distance
		return 0
	}
	
	// TODO: String similarity
	similarity := func(s1, s2 string) float64 {
		// TODO: Calculate string similarity percentage
		return 0.0
	}
	
	// TODO: Anagram detection
	isAnagram := func(s1, s2 string) bool {
		// TODO: Check if two strings are anagrams
		return false
	}
	
	// TODO: Palindrome detection
	isPalindrome := func(s string) bool {
		// TODO: Check if string is palindrome
		return false
	}
	
	// TODO: Test string algorithms
	fmt.Printf("Distance between 'hello' and 'world': %d\n", 
		levenshteinDistance("hello", "world"))
}

// Helper function for performance measurement
func measureStringOperation(name string, fn func()) {
	// TODO: Implement timing measurement for string operations
	fmt.Printf("Measuring %s...\n", name)
	fn()
}

func main() {
	fmt.Println("📝 Welcome to Text Formatting and Strings! 📝")
	fmt.Println("This file teaches you Go's text processing capabilities")
	
	// TODO: Implement each demonstration function
	// Start with basic formatting and progress to advanced text processing
	
	demonstrateFmtFormatting()
	// demonstrateAdvancedFormatting()
	// demonstrateStringsPackage()
	// demonstrateStringPerformance()
	// demonstrateStrconvPackage()
	// demonstrateUnicodePackage()
	// demonstrateRegularExpressions()
	// demonstrateStringValidation()
	// demonstrateTextProcessing()
	// demonstrateStringAlgorithms()
	
	fmt.Println("\n🎉 Congratulations! You've mastered text processing in Go!")
	fmt.Println("Next: Learn time and math operations in time_math.go")
}

/*
🔍 Key Concepts to Remember:

1. **fmt Package**: Formatted I/O with verbs like %v, %d, %s, %f
2. **strings Package**: Comprehensive string manipulation utilities
3. **strconv Package**: String and number conversions
4. **unicode Package**: Unicode character handling and UTF-8
5. **regexp Package**: Regular expression pattern matching
6. **Performance**: Use strings.Builder for concatenation
7. **Validation**: Common patterns for data validation

📋 Essential Functions:
```go
// Formatting
fmt.Printf("Name: %s, Age: %d", name, age)
formatted := fmt.Sprintf("Hello %s", name)

// String operations
upper := strings.ToUpper(text)
words := strings.Split(text, " ")
result := strings.Join(words, "-")

// Conversions
num, err := strconv.Atoi("123")
str := strconv.Itoa(456)

// Unicode
isLetter := unicode.IsLetter('A')
count := utf8.RuneCountInString(text)
```

🚨 Common Mistakes:
- Using + for string concatenation in loops (inefficient)
- Ignoring errors from strconv functions
- Not handling Unicode properly (runes vs bytes)
- Compiling regexp in loops instead of once
- Not using appropriate format verbs in Printf

🎯 Next Steps:
- Learn time and math operations
- Master JSON handling for data exchange
- Build HTTP clients for API integration
- Practice with real-world text processing tasks
*/
