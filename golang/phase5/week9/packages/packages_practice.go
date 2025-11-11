// Week 9: Essential Packages Practice
// Complete these exercises to master Go's standard library

package main

import (
	"fmt"
	"time"
)

// TODO: Exercise 1 - Text Processing CLI Tool
func exercise1_TextProcessor() {
	fmt.Println("=== Exercise 1: Text Processing CLI Tool ===")

	// TODO: Build a text processing tool that:
	// 1. Reads text from command line arguments
	// 2. Provides operations: uppercase, lowercase, reverse, word count
	// 3. Uses fmt for output formatting
	// 4. Uses strings package for manipulations
	// 5. Uses strconv for numeric conversions

	// Example usage: go run main.go -text "Hello World" -op uppercase

	fmt.Println("Exercise 1 completed!")
}

// TODO: Exercise 2 - JSON API Client
func exercise2_JSONAPIClient() {
	fmt.Println("\n=== Exercise 2: JSON API Client ===")

	// TODO: Build a GitHub API client that:
	// 1. Fetches user information from GitHub API
	// 2. Handles JSON marshaling/unmarshaling
	// 3. Implements proper error handling
	// 4. Uses context for timeouts
	// 5. Formats output nicely

	type GitHubUser struct {
		// TODO: Define struct with JSON tags
	}

	fetchUser := func(username string) (*GitHubUser, error) {
		// TODO: Implement API call with context timeout
		return nil, nil
	}

	// TODO: Test with different usernames

	fmt.Println("Exercise 2 completed!")
}

// TODO: Exercise 3 - Time Zone Converter
func exercise3_TimeZoneConverter() {
	fmt.Println("\n=== Exercise 3: Time Zone Converter ===")

	// TODO: Build a timezone converter that:
	// 1. Parses time strings in various formats
	// 2. Converts between different timezones
	// 3. Calculates time differences
	// 4. Formats output in multiple formats
	// 5. Handles invalid inputs gracefully

	convertTime := func(timeStr, fromTZ, toTZ string) (string, error) {
		// TODO: Implement timezone conversion
		return "", nil
	}

	// TODO: Test with various timezone conversions

	fmt.Println("Exercise 3 completed!")
}

// TODO: Exercise 4 - HTTP Load Tester
func exercise4_HTTPLoadTester() {
	fmt.Println("\n=== Exercise 4: HTTP Load Tester ===")

	// TODO: Build a simple load tester that:
	// 1. Makes concurrent HTTP requests
	// 2. Measures response times
	// 3. Tracks success/failure rates
	// 4. Uses context for cancellation
	// 5. Reports statistics

	type LoadTestResult struct {
		// TODO: Define result structure
	}

	runLoadTest := func(url string, concurrent int, duration time.Duration) *LoadTestResult {
		// TODO: Implement load testing
		return nil
	}

	// TODO: Test with different configurations

	fmt.Println("Exercise 4 completed!")
}

// TODO: Exercise 5 - Configuration Manager
func exercise5_ConfigurationManager() {
	fmt.Println("\n=== Exercise 5: Configuration Manager ===")

	// TODO: Build a configuration manager that:
	// 1. Reads config from JSON files
	// 2. Supports environment variable overrides
	// 3. Validates configuration values
	// 4. Provides default values
	// 5. Hot-reloads configuration

	type Config struct {
		// TODO: Define configuration structure
	}

	loadConfig := func(filename string) (*Config, error) {
		// TODO: Implement configuration loading
		return nil, nil
	}

	// TODO: Test configuration loading and validation

	fmt.Println("Exercise 5 completed!")
}

func main() {
	fmt.Println("📦 Welcome to Essential Packages Practice! 📦")

	exercise1_TextProcessor()
	// exercise2_JSONAPIClient()
	// exercise3_TimeZoneConverter()
	// exercise4_HTTPLoadTester()
	// exercise5_ConfigurationManager()

	fmt.Println("\n🎉 Congratulations! You've mastered essential packages!")
}
