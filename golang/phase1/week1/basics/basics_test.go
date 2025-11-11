// Test file for Week 1 Basics exercises
// Run with: go test

package main

import (
	"os"
	"runtime"
	"testing"
)

// Test for setup.go functionality
func TestGoInstallation(t *testing.T) {
	// Check if Go is properly installed
	version := runtime.Version()
	if version == "" {
		t.Error("Go version not found")
	}
	
	// Check basic Go environment
	goos := runtime.GOOS
	if goos == "" {
		t.Error("GOOS not set")
	}
	
	goarch := runtime.GOARCH
	if goarch == "" {
		t.Error("GOARCH not set")
	}
	
	t.Logf("Go version: %s, OS: %s, Arch: %s", version, goos, goarch)
}

// Test for hello.go functionality
func TestHelloWorldFunctions(t *testing.T) {
	// Test that functions exist and can be called
	// This is more of a compilation test
	t.Log("Testing hello world functions compilation")
	
	// You can extend this to test actual output if needed
	// For now, we're just ensuring the code compiles
}

// Test for commands.go functionality
func TestGoCommands(t *testing.T) {
	// Test basic runtime information
	version := runtime.Version()
	if len(version) == 0 {
		t.Error("Unable to get Go version")
	}
	
	// Test basic file operations
	_, err := os.Getwd()
	if err != nil {
		t.Error("Unable to get current working directory")
	}
	
	t.Log("Go commands test passed")
}

// Test for packages.go functionality
func TestPackageSystem(t *testing.T) {
	// Test basic package imports are working
	
	// Test math package
	result := calculateArea(1.0)
	expected := 3.141592653589793
	if result != expected {
		t.Errorf("calculateArea(1.0) = %f, expected %f", result, expected)
	}
	
	// Test string formatting
	message := formatMessage("Hello", "World")
	expected2 := "Hello, World!"
	if message != expected2 {
		t.Errorf("formatMessage('Hello', 'World') = %s, expected %s", message, expected2)
	}
	
	t.Log("Package system test passed")
}

// Benchmark test for performance demonstration
func BenchmarkSumCalculation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		for j := 0; j < 1000; j++ {
			sum += j
		}
	}
}

// Test helper function to check if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// Test module system basics
func TestModuleSystem(t *testing.T) {
	// Check if we're in a Go module
	if !fileExists("go.mod") {
		t.Log("Warning: Not in a Go module. Run 'go mod init phase1-foundation'")
	} else {
		t.Log("Running in a Go module ✓")
	}
} 