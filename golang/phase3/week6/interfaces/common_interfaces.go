package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
	"time"
)

// ===== fmt.Stringer INTERFACE =====

// TODO: Understand fmt.Stringer interface
// The fmt.Stringer interface has one method: String() string
// When you implement this, fmt.Print, fmt.Printf %s, %v will use your String() method

// TODO: Define a PersonStringer struct with Name (string), Age (int), City (string) fields
// YOUR CODE HERE
type PersonStringer struct {
	Name string
	Age  int
	City string
}

// TODO: Implement String() method for PersonStringer
// Should return "[Name] ([Age] years old) from [City]"
// YOUR CODE HERE
func (p PersonStringer) String() string {
	return fmt.Sprintf("%s (%d years old) from %s", p.Name, p.Age, p.City)
}

// TODO: Define a BankAccount struct with AccountNumber (string), Balance (float64), Owner (string)
// YOUR CODE HERE
type BankAccount struct {
	AccountNumber string
	Balance       float64
	Owner         string
}

// TODO: Implement String() method for BankAccount
// Mask account number for security (show only last 4 digits)
// Should return "Account ****[last4] - Balance: $[balance] (Owner: [owner])"
// Hint: Use strings.Repeat("*", count) to create asterisks
// YOUR CODE HERE
func (b BankAccount) String() string {
	return fmt.Sprintf(
		"Account %s - Balance: $%.2f (Owner: %s)",
		b.AccountNumber[len(b.AccountNumber)-4:],
		b.Balance,
		b.Owner,
	)
}
func demonstrateStringer() {
	fmt.Println("=== fmt.Stringer Interface ===")

	// TODO: Create a PersonStringer instance with sample data
	// YOUR CODE HERE

	timCook := PersonStringer{
		Name: "Tim Cook",
		Age:  62,
		City: "Cupertino",
	}

	// TODO: Create a BankAccount instance with sample data
	// YOUR CODE HERE
	timCookAccount := BankAccount{
		AccountNumber: "1234567890",
		Balance:       1000000,
		Owner:         "Tim Cook",
	}

	// TODO: Print both using %s (this will call String() method automatically)
	// YOUR CODE HERE

	fmt.Printf("Person: %s\n", timCook)
	fmt.Printf("Bank Account: %s\n", timCookAccount)

	// TODO: Print both using %v (this also calls String() method)
	// YOUR CODE HERE
	fmt.Printf("Person: %v\n", timCook)
	fmt.Printf("Bank Account: %v\n", timCookAccount)

	// TODO: Create a slice of fmt.Stringer containing both person and account
	// Loop through and print each item
	// YOUR CODE HERE

	peopleAndAccounts := []fmt.Stringer{timCook, timCookAccount}

	for _, v := range peopleAndAccounts {
		fmt.Printf("%s\n", v)
	}
}

// ===== io.Reader INTERFACE =====

// TODO: Understand io.Reader interface
// The io.Reader interface has one method: Read([]byte) (n int, err error)
// It reads data into the provided byte slice and returns bytes read and error
// Return io.EOF when no more data to read

// TODO: Define a StringReader struct with data (string) and position (int) fields
// YOUR CODE HERE
type StringReader struct {
	data     string
	position int
}

// TODO: Implement NewStringReader constructor function
// Should return *StringReader with given data and position 0
// YOUR CODE HERE
func NewStringReader(data string) *StringReader {
	return &StringReader{
		data:     data,
		position: 0,
	}
}

// TODO: Implement Read method for StringReader (use pointer receiver)
// - Return 0, io.EOF if position >= len(data)
// - Use copy(p, sr.data[sr.position:]) to copy data
// - Update position and return bytes copied
// YOUR CODE HERE
func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.position >= len(sr.data) {
		return 0, io.EOF
	}
	n = copy(p, sr.data[sr.position:])
	sr.position += n
	return n, nil
}

// TODO: Define UpperCaseReader struct that wraps another io.Reader
// Should have reader field of type io.Reader
// YOUR CODE HERE
type UpperCaseReader struct {
	reader io.Reader
}

// TODO: Implement NewUpperCaseReader constructor
// YOUR CODE HERE
func NewUpperCaseReader(reader io.Reader) *UpperCaseReader {
	return &UpperCaseReader{
		reader: reader,
	}
}

// TODO: Implement Read method for UpperCaseReader
// - Call wrapped reader's Read method first
// - Convert lowercase letters to uppercase in the byte slice
// - Check if p[i] >= 'a' && p[i] <= 'z', then p[i] = p[i] - 'a' + 'A'
// YOUR CODE HERE
func (ur *UpperCaseReader) Read(p []byte) (n int, err error) {
	n, err = ur.reader.Read(p)
	for i := 0; i < n; i++ {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] = p[i] - 'a' + 'A'
		}
	}
	return n, err
}

func demonstrateReader() {
	fmt.Println("\n=== io.Reader Interface ===")

	// TODO: Create a StringReader with some sample text
	// YOUR CODE HERE
	sr := NewStringReader("Shit, world!")

	// TODO: Wrap it with UpperCaseReader
	// YOUR CODE HERE
	ucr := NewUpperCaseReader(sr)

	// TODO: Use io.ReadAll to read all data and handle error
	// Print original text and uppercase result
	// YOUR CODE HERE
	ucrAll, err := io.ReadAll(ucr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(ucrAll))

	// TODO: Demonstrate with standard library strings.NewReader
	// Use io.ReadAll to read the data
	// YOUR CODE HERE
	sns := strings.NewReader("Shit, another world!")
	snsAll, err := io.ReadAll(sns)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(snsAll))

	// TODO: Chain readers - wrap strings.NewReader with UpperCaseReader
	// Read and print the result
	// YOUR CODE HERE
	snsUpper := strings.NewReader("Shit, three body world!")
	upperThreeBodyWorld := NewUpperCaseReader(snsUpper)
	upperThreeBodyWorldAll, err := io.ReadAll(upperThreeBodyWorld)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(upperThreeBodyWorldAll))
}

// ===== io.Writer INTERFACE =====

// TODO: Understand io.Writer interface
// The io.Writer interface has one method: Write([]byte) (n int, err error)
// It writes data from byte slice and returns bytes written and error

// TODO: Define LogWriter struct with prefix (string) field
// YOUR CODE HERE
type LogWriter struct {
	prefix string
}

// TODO: Implement NewLogWriter constructor
// YOUR CODE HERE
func NewLogWriter(prefix string) *LogWriter {
	return &LogWriter{
		prefix: prefix,
	}
}

// TODO: Implement Write method for LogWriter
// - Get timestamp with time.Now().Format("15:04:05")
// - Format message as "[timestamp] prefix: content"
// - Print the message and return len(p), nil
// YOUR CODE HERE
func (lw *LogWriter) Write(p []byte) (n int, err error) {
	timestamp := time.Now().Format("15:04:05")
	message := fmt.Sprintf("[%s] %s: %s", timestamp, lw.prefix, string(p))
	fmt.Println(message)
	return len(p), nil
}

// TODO: Define MemoryWriter struct with buffer ([]byte) field
// YOUR CODE HERE
type MemoryWriter struct {
	buffer []byte
}

// TODO: Implement NewMemoryWriter constructor
// Initialize with empty byte slice using make([]byte, 0)
// YOUR CODE HERE
func NewMemoryWriter() *MemoryWriter {
	return &MemoryWriter{
		buffer: make([]byte, 0),
	}
}

// TODO: Implement Write method for MemoryWriter
// Append p to buffer using append(mw.buffer, p...)
// Return len(p), nil
// YOUR CODE HERE
func (mw *MemoryWriter) Write(p []byte) (n int, err error) {
	mw.buffer = append(mw.buffer, p...)
	return len(p), nil
}

// TODO: Implement String() method for MemoryWriter
// Convert buffer to string and return
// YOUR CODE HERE
func (mw *MemoryWriter) String() string {
	return string(mw.buffer)
}

// TODO: Implement Reset() method for MemoryWriter
// Reset buffer to empty: mw.buffer = mw.buffer[:0]
// YOUR CODE HERE
func (mw *MemoryWriter) Reset() {
	mw.buffer = mw.buffer[:0]
}

func demonstrateWriter() {
	fmt.Println("\n=== io.Writer Interface ===")

	// TODO: Create a LogWriter with prefix "APP"
	// YOUR CODE HERE
	lw := NewLogWriter("APP")
	// TODO: Write some log messages using Write method
	// Convert strings to []byte for writing
	// YOUR CODE HERE

	lw.Write([]byte("Shit, world!"))
	lw.Write([]byte("Shit, three body fleet coming!"))
	lw.Write([]byte("Shit, the solar system is flattening!"))
	// TODO: Create a MemoryWriter and write some data to it
	// YOUR CODE HERE
	mw := NewMemoryWriter()
	mw.Write([]byte("Eason, you have a new message!"))
	mw.Write([]byte("The world is ending and you are the last person on earth!"))
	mw.Write([]byte("This message is self destructing in 10 seconds!"))
	// TODO: Print the memory buffer contents using String() method
	// YOUR CODE HERE
	fmt.Println(mw)
	// TODO: Demonstrate fmt.Fprintf with both writers
	// fmt.Fprintf accepts io.Writer as first argument
	// YOUR CODE HERE
	mw.Reset()

	fmt.Fprintf(mw, "%s %s", time.Now().Format("15:04:05"), "Shit, world!")

	fmt.Println(mw)
}

// ===== io.ReadWriter INTERFACE =====

// TODO: Define InMemoryFile struct with content ([]byte) and position (int) fields
// This will implement both io.Reader and io.Writer (so it's an io.ReadWriter)
// YOUR CODE HERE
type InMemoryFile struct {
	content  []byte
	position int
}

// TODO: Implement NewInMemoryFile constructor
// Initialize with empty content slice and position 0
// YOUR CODE HERE
func NewInMemoryFile() *InMemoryFile {
	return &InMemoryFile{
		content:  make([]byte, 0),
		position: 0,
	}
}

// TODO: Implement Read method for InMemoryFile (similar to StringReader)
// - Return 0, io.EOF if position >= len(content)
// - Copy data from content[position:] to p
// - Update position and return bytes copied
// YOUR CODE HERE
func (imf *InMemoryFile) Read(p []byte) (n int, err error) {
	if imf.position >= len(imf.content) {
		return 0, io.EOF
	}
	n = copy(p, imf.content[imf.position:])
	imf.position += n
	return n, nil
}

// TODO: Implement Write method for InMemoryFile
// - Append p to content using append(imf.content, p...)
// - Return len(p), nil
// YOUR CODE HERE
func (imf *InMemoryFile) Write(p []byte) (n int, err error) {
	imf.content = append(imf.content, p...)
	return len(p), nil
}

// TODO: Implement Seek method for InMemoryFile
// Parameters: offset (int64), whence (int)
// Handle io.SeekStart, io.SeekCurrent, io.SeekEnd
// Clamp position between 0 and len(content)
// Return final position as int64
// YOUR CODE HERE
func (imf *InMemoryFile) Seek(offset int64, whence int) int64 {
	switch whence {
	case io.SeekStart:
		imf.position = int(offset)
	case io.SeekCurrent:
		imf.position += int(offset)
	case io.SeekEnd:
		imf.position = len(imf.content) + int(offset)
	}
	if imf.position < 0 {
		imf.position = 0
	}
	if imf.position > len(imf.content) {
		imf.position = len(imf.content)
	}
	return int64(imf.position)
}

func demonstrateReadWriter() {
	fmt.Println("\n=== io.ReadWriter Interface ===")

	// TODO: Create a new InMemoryFile
	// YOUR CODE HERE
	imf := NewInMemoryFile()

	// TODO: Write several lines of data to the file
	// YOUR CODE HERE
	imf.Write([]byte("Shit, world!"))
	imf.Write([]byte("Shit, three body fleet coming!"))
	imf.Write([]byte("Shit, the solar system is flattening!"))

	// TODO: Reset position to beginning using Seek(0, io.SeekStart)
	// YOUR CODE HERE
	imf.Seek(0, io.SeekStart)
	// TODO: Read all data using io.ReadAll and print it
	// Handle the error properly
	// YOUR CODE HERE
	imfAll, err := io.ReadAll(imf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(imfAll))

	fmt.Println("Reaching the end of Reader so no more data")
	imfAll, err = io.ReadAll(imf)
	fmt.Println(string(imfAll))

	// TODO: Append more data to the file
	// YOUR CODE HERE
	fmt.Fprintf(imf, "New civilization is born")

	// TODO: Reset to beginning and read all data again
	// Show that new data was appended
	// YOUR CODE HERE
	imf.Seek(0, io.SeekStart)
	imfAll, err = io.ReadAll(imf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(imfAll))
}

// ===== sort.Interface =====

// sort.Interface enables custom sorting
// type Interface interface {
//     Len() int
//     Less(i, j int) bool
//     Swap(i, j int)
// }

// TODO: Define Student struct with Name (string), Grade (float64), Age (int) fields
// YOUR CODE HERE
type Student struct {
	Name  string
	Grade float64
	Age   int
}

// TODO: Define ByGrade type as []Student and implement sort.Interface
// - Len() should return length of slice
// - Less(i, j) should sort by Grade in descending order
// - Swap(i, j) should swap elements i and j
// YOUR CODE HERE
type ByGrade []Student

func (bg ByGrade) Len() int {
	return len(bg)
}

func (bg ByGrade) Less(i, j int) bool {
	return bg[i].Grade > bg[j].Grade
}

func (bg ByGrade) Swap(i, j int) {
	bg[i], bg[j] = bg[j], bg[i]
}

// TODO: Define ByAge type as []Student and implement sort.Interface
// - Len() should return length of slice
// - Less(i, j) should sort by Age in ascending order
// - Swap(i, j) should swap elements i and j
// YOUR CODE HERE
type ByAge []Student

func (ba ByAge) Len() int {
	return len(ba)
}

func (ba ByAge) Less(i, j int) bool {
	return ba[i].Age < ba[j].Age
}

func (ba ByAge) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

// TODO: Define ByName type as []Student and implement sort.Interface
// - Len() should return length of slice
// - Less(i, j) should sort by Name alphabetically
// - Swap(i, j) should swap elements i and j
// YOUR CODE HERE
type ByName []Student

func (bn ByName) Len() int {
	return len(bn)
}

func (bn ByName) Less(i, j int) bool {
	return bn[i].Name < bn[j].Name
}

func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

func demonstrateSortInterface() {
	fmt.Println("\n=== sort.Interface ===")

	// TODO: Create a slice of students with sample data
	// YOUR CODE HERE
	students := []Student{
		{Name: "John", Grade: 85.5, Age: 20},
		{Name: "Jane", Grade: 90.0, Age: 22},
		{Name: "Jim", Grade: 78.5, Age: 21},
		{Name: "Jill", Grade: 88.0, Age: 20},
	}
	// TODO: Print original students
	// YOUR CODE HERE
	printStudents(students)
	// TODO: Sort by grade (descending) and print
	// YOUR CODE HERE
	sort.Sort(ByGrade(students))
	printStudents(students)
	// TODO: Sort by age (ascending) and print
	// YOUR CODE HERE
	sort.Sort(ByAge(students))
	printStudents(students)
	// TODO: Sort by name (alphabetical) and print
	// YOUR CODE HERE
	sort.Sort(ByName(students))
	printStudents(students)
}

// Helper function to print students
// TODO: Implement printStudents function
// YOUR CODE HERE
func printStudents(students []Student) {
	for _, s := range students {
		fmt.Printf("Name: %s; Age: %d; Grade: %.2f\n", s.Name, s.Age, s.Grade)
	}
	fmt.Println()
}

// ===== error INTERFACE =====

// error is a built-in interface
// type error interface {
//     Error() string
// }

// TODO: Define ValidationErrorCustom struct with Field (string), Value (interface{}), Message (string) fields
// YOUR CODE HERE
type ValidationErrorCustom struct {
	Field   string
	Value   any
	Message string
}

// TODO: Implement Error() method for ValidationErrorCustom
// Should return "validation failed for field '[Field]' with value '[Value]': [Message]"
// YOUR CODE HERE
func (vec ValidationErrorCustom) Error() (err string) {
	err = fmt.Sprintf("validation failed for field '%s' with value '%v': %s", vec.Field, vec.Value, vec.Message)
	return
}

// TODO: Define NetworkErrorCustom struct with Operation (string), URL (string), Code (int), Timestamp (time.Time) fields
// YOUR CODE HERE
type NetworkErrorCustom struct {
	Operation string
	URL       string
	Code      int
	Timestamp time.Time
}

// TODO: Implement Error() method for NetworkErrorCustom
// Should return "network error during [Operation] to [URL]: HTTP [Code] at [Timestamp]"
// Format timestamp using Format("15:04:05")
// YOUR CODE HERE
func (nec NetworkErrorCustom) Error() (err string) {
	err = fmt.Sprintf("network error during '%s' to '%s': HTTP %d at %s", nec.Operation, nec.URL, nec.Code, nec.Timestamp.Format("15:04:05"))
	return
}

// TODO: Implement String() method for NetworkErrorCustom (fmt.Stringer)
// Should return "NetworkError{Op: [Operation], URL: [URL], Code: [Code]}"
// YOUR CODE HERE
func (nec NetworkErrorCustom) String() (desc string) {
	desc = fmt.Sprintf("NetworkError{Op: %s, URL: %s, Code: %d}", nec.Operation, nec.URL, nec.Code)
	return
}

func demonstrateErrorInterface() {
	fmt.Println("\n=== error Interface ===")

	// TODO: Create sample validation and network errors
	// YOUR CODE HERE
	vec := ValidationErrorCustom{
		Field:   "Name",
		Value:   "John",
		Message: "Name is required",
	}
	nec := NetworkErrorCustom{
		Operation: "GET",
		URL:       "https://example.com",
		Code:      500,
		Timestamp: time.Now(),
	}
	// TODO: Create slice of error interface containing both errors
	// YOUR CODE HERE
	errors := []error{vec, nec}
	// TODO: Range over errors and print each one
	// YOUR CODE HERE
	for _, err := range errors {
		fmt.Println(err)
	}
	// TODO: Use type assertion to handle each error type specifically
	// For ValidationErrorCustom: Print field name
	// For NetworkErrorCustom: Print status code and use String() method
	// YOUR CODE HERE
	for _, err := range errors {
		switch err := err.(type) {
		case ValidationErrorCustom:
			fmt.Printf("Validation Error: %s\n", err.Field)
		case NetworkErrorCustom:
			fmt.Printf("Network Error: %s\n", err.String())
		}
	}
}

// ===== COMBINING INTERFACES =====

// TODO: Define Logger interface with Log(message string) method
// YOUR CODE HERE
type Logger interface {
	Log(message string)
}

// TODO: Define FileSystemLogger interface that combines Logger and io.Writer
// Also add SetLogLevel(level string) method
// YOUR CODE HERE
type FileSystemLogger interface {
	Logger
	io.Writer
	SetLogLevel(level string)
}

// TODO: Define ConsoleFileLogger struct with logLevel (string) and buffer (*bytes.Buffer) fields
// YOUR CODE HERE
type ConsoleFileLogger struct {
	buffer   *bytes.Buffer
	logLevel string
}

// TODO: Implement NewConsoleFileLogger constructor
// Initialize with "INFO" level and new buffer
// YOUR CODE HERE
func NewConsoleFileLogger() *ConsoleFileLogger {
	return &ConsoleFileLogger{
		buffer:   &bytes.Buffer{},
		logLevel: "INFO",
	}
}

// TODO: Implement Log method for ConsoleFileLogger
// - Format with timestamp, level, message
// - Write to console and buffer
// YOUR CODE HERE

func (cf *ConsoleFileLogger) Log(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	level := strings.ToUpper(cf.logLevel)
	content := fmt.Sprintf("[%s] %s: %s", timestamp, level, message)
	cf.buffer.WriteString(content)
	cf.buffer.WriteString("\n")
}

// TODO: Implement Write method for ConsoleFileLogger (io.Writer)
// - Convert bytes to string, trim space
// - Call Log method
// - Return length and nil error
// YOUR CODE HERE
func (cf *ConsoleFileLogger) Write(p []byte) (n int, err error) {
	message := strings.TrimSpace(string(p))
	if message == "" {
		return 0, nil
	}
	cf.Log(message)
	return len(p), nil
}

// TODO: Implement SetLogLevel method
// YOUR CODE HERE

func (cf *ConsoleFileLogger) SetLogLevel(level string) {
	cf.logLevel = strings.ToUpper(level)
}

// TODO: Implement GetBufferedLogs method to return buffer contents
// YOUR CODE HERE
func (cf *ConsoleFileLogger) GetBufferedLogs() string {
	return cf.buffer.String()
}

func demonstrateInterfaceCombination() {
	fmt.Println("\n=== Combining Interfaces ===")

	// TODO: Create new logger
	// YOUR CODE HERE
	cf := NewConsoleFileLogger()

	// TODO: Use as Logger interface
	// YOUR CODE HERE
	cf.Log("This is a test message")

	// TODO: Use as io.Writer interface with fmt.Fprintf
	// YOUR CODE HERE
	fmt.Fprintf(cf, "This is a test message from fmt.Fprintf")

	// TODO: Use as FileSystemLogger interface
	// YOUR CODE HERE
	cf.SetLogLevel("DEBUG")

	// TODO: Show buffered logs
	// YOUR CODE HERE
	fmt.Println(cf.GetBufferedLogs())
}

// ===== INTERFACE COMPOSITION PATTERNS =====

// TODO: Define Processor interface with Process(data string) (string, error) method
// YOUR CODE HERE
type Processor interface {
	Process(data string) (string, error)
}

// TODO: Define Validator interface with Validate(data string) error method
// YOUR CODE HERE
type Validator interface {
	Validate(data string) error
}

// TODO: Define ProcessorValidator interface combining both interfaces
// YOUR CODE HERE
type ProcessorValidator interface {
	Processor
	Validator
}

// TODO: Define EmailProcessor struct with domain (string) field
// YOUR CODE HERE
type EmailProcessor struct {
	domain string
}

// TODO: Implement NewEmailProcessor constructor
// YOUR CODE HERE
func NewEmailProcessor(domain string) *EmailProcessor {
	return &EmailProcessor{
		domain: domain,
	}
}

// TODO: Implement Validate method for EmailProcessor
// - Check if email contains @
// - Check if email ends with correct domain
// - Return appropriate errors
// YOUR CODE HERE
func (ep *EmailProcessor) Validate(data string) error {
	if !strings.Contains(data, "@") {
		return fmt.Errorf("invalid email format")
	}
	if !strings.HasSuffix(data, ep.domain) {
		return fmt.Errorf("invalid domain")
	}
	return nil
}

// TODO: Implement Process method for EmailProcessor
// - Call Validate first
// - If valid, normalize email (lowercase, trim space)
// - Return result or error
// YOUR CODE HERE
func (ep *EmailProcessor) Process(data string) (string, error) {
	if err := ep.Validate(data); err != nil {
		return "", err
	}
	normalized := strings.ToLower(strings.TrimSpace(data))
	return normalized, nil
}

func demonstrateInterfaceCompositionPatterns() {
	fmt.Println("\n=== Interface Composition Patterns ===")

	// TODO: Create email processor with "@company.com" domain
	// YOUR CODE HERE
	ep := NewEmailProcessor("@company.com")

	// TODO: Create slice of test emails (valid and invalid)
	// YOUR CODE HERE
	emails := []string{
		"test@company.com",
		"ceo@company.com",
		"ceo@company.com.com",
		"ceo@company.com.com.com",
		"ceo@company.com.com.com.com",
		"ceo@company.com.com.com.com.com",
		"ceo@company.com.com.com.com.com.com",
		"ceo@company.com.com.com.com.com.com.com",
	}

	// TODO: Process each email:
	// - Use as ProcessorValidator interface
	// - Validate first
	// - If valid, process and show result
	// - Handle errors appropriately
	// YOUR CODE HERE
	for _, email := range emails {
		result, err := ep.Process(email)
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", email, err)
		} else {
			fmt.Printf("Processed %s -> %s\n", email, result)
		}
	}
}

func main() {
	fmt.Println("📚 Go Common Standard Library Interfaces Practice")
	fmt.Println("=================================================")

	demonstrateStringer()
	demonstrateReader()
	demonstrateWriter()
	demonstrateReadWriter()
	// TODO: Uncomment these as you implement them
	demonstrateSortInterface()
	demonstrateErrorInterface()
	demonstrateInterfaceCombination()
	demonstrateInterfaceCompositionPatterns()

	fmt.Println("\n✅ Common interfaces practice completed!")
	fmt.Println("\n🎯 Learning Goals:")
	fmt.Println("- Implement fmt.Stringer for custom string representation")
	fmt.Println("- Create io.Reader/Writer for flexible data handling")
	fmt.Println("- Use sort.Interface for custom sorting logic")
	fmt.Println("- Implement error interface for rich error information")
	fmt.Println("- Apply interface composition for powerful abstractions")
	fmt.Println("- Make types integrate with Go's standard library")
}
