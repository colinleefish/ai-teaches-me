// Week 10: File Operations
// This file demonstrates reading, writing, and manipulating files in Go

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// TODO: Demonstrate basic file reading operations
func demonstrateFileReading() {
	fmt.Println("=== File Reading Operations ===")

	// TODO: Create a sample file first
	sampleContent := `Hello, World!
This is a sample file.
It contains multiple lines.
Each line has different content.
End of file.`

	filename := "sample.txt"
	err := os.WriteFile(filename, []byte(sampleContent), 0644)
	if err != nil {
		fmt.Printf("Error creating sample file: %v\n", err)
		return
	}
	defer os.Remove(filename) // Clean up

	// TODO: Method 1: Read entire file with os.ReadFile
	fmt.Println("Method 1: os.ReadFile (entire file)")
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("File content (%d bytes):\n%s\n", len(content), string(content))

	// TODO: Method 2: Open file and read with io.ReadAll
	fmt.Println("\nMethod 2: os.Open + io.ReadAll")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	content2, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("Content read: %d bytes\n", len(content2))

	// TODO: Method 3: Buffered reading line by line
	fmt.Println("\nMethod 3: Line-by-line with bufio.Scanner")
	file2, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file2.Close()

	scanner := bufio.NewScanner(file2)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
	}
}

// TODO: Demonstrate basic file writing operations
func demonstrateFileWriting() {
	fmt.Println("\n=== File Writing Operations ===")

	// TODO: Method 1: Write entire content with os.WriteFile
	fmt.Println("Method 1: os.WriteFile")
	content := "Hello from Go!\nThis is written with os.WriteFile.\n"
	filename := "output1.txt"

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	defer os.Remove(filename)
	fmt.Printf("Written to %s\n", filename)

	// TODO: Method 2: Create file and write with io.WriteString
	fmt.Println("\nMethod 2: os.Create + io.WriteString")
	filename2 := "output2.txt"
	file, err := os.Create(filename2)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()
	defer os.Remove(filename2)

	lines := []string{
		"First line\n",
		"Second line\n",
		"Third line\n",
	}

	for i, line := range lines {
		n, err := io.WriteString(file, line)
		if err != nil {
			fmt.Printf("Error writing line %d: %v\n", i+1, err)
			return
		}
		fmt.Printf("Wrote %d bytes for line %d\n", n, i+1)
	}

	// TODO: Method 3: Buffered writing
	fmt.Println("\nMethod 3: Buffered writing with bufio.Writer")
	filename3 := "output3.txt"
	file3, err := os.Create(filename3)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file3.Close()
	defer os.Remove(filename3)

	writer := bufio.NewWriter(file3)
	defer writer.Flush() // Important: flush buffer before closing

	for i := 0; i < 1000; i++ {
		_, err := fmt.Fprintf(writer, "Line %d: Some data here\n", i+1)
		if err != nil {
			fmt.Printf("Error writing line %d: %v\n", i+1, err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing buffer: %v\n", err)
		return
	}
	fmt.Println("Buffered writing completed")
}

// TODO: Demonstrate file appending
func demonstrateFileAppending() {
	fmt.Println("\n=== File Appending ===")

	filename := "append_test.txt"

	// TODO: Create initial file
	err := os.WriteFile(filename, []byte("Initial content\n"), 0644)
	if err != nil {
		fmt.Printf("Error creating initial file: %v\n", err)
		return
	}
	defer os.Remove(filename)

	// TODO: Open file for appending
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file for append: %v\n", err)
		return
	}
	defer file.Close()

	// TODO: Append multiple lines
	appendLines := []string{
		"Appended line 1\n",
		"Appended line 2\n",
		"Appended line 3\n",
	}

	for _, line := range appendLines {
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Printf("Error appending: %v\n", err)
			return
		}
	}

	// TODO: Read and display final content
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading final content: %v\n", err)
		return
	}

	fmt.Printf("Final file content:\n%s", string(content))
}

// TODO: Demonstrate file metadata and permissions
func demonstrateFileMetadata() {
	fmt.Println("\n=== File Metadata and Permissions ===")

	filename := "metadata_test.txt"
	content := "File for metadata testing"

	// TODO: Create file with specific permissions
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer os.Remove(filename)

	// TODO: Get file information
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}

	fmt.Printf("File: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Mode: %v\n", fileInfo.Mode())
	fmt.Printf("ModTime: %v\n", fileInfo.ModTime())
	fmt.Printf("IsDir: %v\n", fileInfo.IsDir())

	// TODO: Change file permissions
	fmt.Println("\nChanging file permissions...")
	err = os.Chmod(filename, 0600) // Owner read/write only
	if err != nil {
		fmt.Printf("Error changing permissions: %v\n", err)
		return
	}

	// TODO: Check new permissions
	fileInfo2, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Error getting updated file info: %v\n", err)
		return
	}

	fmt.Printf("New mode: %v\n", fileInfo2.Mode())

	// TODO: Change file timestamps
	newTime := time.Now().Add(-24 * time.Hour) // 24 hours ago
	err = os.Chtimes(filename, newTime, newTime)
	if err != nil {
		fmt.Printf("Error changing timestamps: %v\n", err)
		return
	}

	fileInfo3, _ := os.Stat(filename)
	fmt.Printf("New ModTime: %v\n", fileInfo3.ModTime())
}

// TODO: Demonstrate temporary files and directories
func demonstrateTemporaryFiles() {
	fmt.Println("\n=== Temporary Files and Directories ===")

	// TODO: Create temporary file
	tmpFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tmpFile.Name()) // Clean up
	defer tmpFile.Close()

	fmt.Printf("Created temp file: %s\n", tmpFile.Name())

	// TODO: Write to temporary file
	content := "This is temporary content\n"
	_, err = tmpFile.WriteString(content)
	if err != nil {
		fmt.Printf("Error writing to temp file: %v\n", err)
		return
	}

	// TODO: Create temporary directory
	tmpDir, err := os.MkdirTemp("", "example-dir-*")
	if err != nil {
		fmt.Printf("Error creating temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir) // Clean up entire directory

	fmt.Printf("Created temp directory: %s\n", tmpDir)

	// TODO: Create files in temporary directory
	for i := 0; i < 3; i++ {
		filename := filepath.Join(tmpDir, fmt.Sprintf("file_%d.txt", i+1))
		content := fmt.Sprintf("Content of file %d\n", i+1)
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error creating file in temp dir: %v\n", err)
			continue
		}
		fmt.Printf("Created: %s\n", filename)
	}
}

// TODO: Demonstrate file copying and moving
func demonstrateFileCopyMove() {
	fmt.Println("\n=== File Copy and Move Operations ===")

	// TODO: Create source file
	sourceFile := "source.txt"
	content := "This is the source file content.\nIt has multiple lines.\nFor testing copy operations.\n"
	err := os.WriteFile(sourceFile, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating source file: %v\n", err)
		return
	}
	defer os.Remove(sourceFile)

	// TODO: Copy file using io.Copy
	copyFile := func(src, dst string) error {
		sourceFile, err := os.Open(src)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		destFile, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, sourceFile)
		return err
	}

	// TODO: Test file copying
	destFile := "destination.txt"
	err = copyFile(sourceFile, destFile)
	if err != nil {
		fmt.Printf("Error copying file: %v\n", err)
		return
	}
	defer os.Remove(destFile)

	fmt.Printf("Copied %s to %s\n", sourceFile, destFile)

	// TODO: Verify copy
	originalInfo, _ := os.Stat(sourceFile)
	copyInfo, _ := os.Stat(destFile)

	fmt.Printf("Original size: %d bytes\n", originalInfo.Size())
	fmt.Printf("Copy size: %d bytes\n", copyInfo.Size())

	// TODO: Move file (rename)
	movedFile := "moved.txt"
	err = os.Rename(destFile, movedFile)
	if err != nil {
		fmt.Printf("Error moving file: %v\n", err)
		return
	}
	defer os.Remove(movedFile)

	fmt.Printf("Moved %s to %s\n", destFile, movedFile)

	// TODO: Verify original file still exists and moved file exists
	if _, err := os.Stat(destFile); os.IsNotExist(err) {
		fmt.Printf("Original %s no longer exists (moved successfully)\n", destFile)
	}

	if _, err := os.Stat(movedFile); err == nil {
		fmt.Printf("Moved file %s exists\n", movedFile)
	}
}

// TODO: Demonstrate file locking (basic approach)
func demonstrateFileLocking() {
	fmt.Println("\n=== File Locking (Basic) ===")

	// TODO: Create a lock file approach
	lockFile := "process.lock"

	// TODO: Function to acquire lock
	acquireLock := func() (*os.File, error) {
		// Try to create lock file exclusively
		file, err := os.OpenFile(lockFile, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
		if err != nil {
			if os.IsExist(err) {
				return nil, fmt.Errorf("process already running (lock file exists)")
			}
			return nil, err
		}

		// Write PID to lock file
		_, err = fmt.Fprintf(file, "%d\n", os.Getpid())
		if err != nil {
			file.Close()
			os.Remove(lockFile)
			return nil, err
		}

		return file, nil
	}

	// TODO: Function to release lock
	releaseLock := func(file *os.File) error {
		file.Close()
		return os.Remove(lockFile)
	}

	// TODO: Test locking
	fmt.Println("Acquiring lock...")
	lock, err := acquireLock()
	if err != nil {
		fmt.Printf("Error acquiring lock: %v\n", err)
		return
	}

	fmt.Printf("Lock acquired, PID: %d\n", os.Getpid())

	// TODO: Simulate work
	time.Sleep(1 * time.Second)

	// TODO: Try to acquire lock again (should fail)
	fmt.Println("Trying to acquire lock again...")
	_, err = acquireLock()
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// TODO: Release lock
	fmt.Println("Releasing lock...")
	err = releaseLock(lock)
	if err != nil {
		fmt.Printf("Error releasing lock: %v\n", err)
	} else {
		fmt.Println("Lock released successfully")
	}
}

// TODO: Demonstrate atomic file operations
func demonstrateAtomicOperations() {
	fmt.Println("\n=== Atomic File Operations ===")

	// TODO: Atomic write using temporary file and rename
	atomicWrite := func(filename string, data []byte) error {
		// Create temporary file in same directory
		dir := filepath.Dir(filename)
		tmpFile, err := os.CreateTemp(dir, "atomic-*.tmp")
		if err != nil {
			return err
		}

		// Write data to temporary file
		_, err = tmpFile.Write(data)
		if err != nil {
			tmpFile.Close()
			os.Remove(tmpFile.Name())
			return err
		}

		// Close temporary file
		err = tmpFile.Close()
		if err != nil {
			os.Remove(tmpFile.Name())
			return err
		}

		// Atomic rename
		return os.Rename(tmpFile.Name(), filename)
	}

	// TODO: Test atomic write
	filename := "atomic_test.txt"
	data := []byte("This data was written atomically\n")

	err := atomicWrite(filename, data)
	if err != nil {
		fmt.Printf("Error in atomic write: %v\n", err)
		return
	}
	defer os.Remove(filename)

	fmt.Printf("Atomic write completed to %s\n", filename)

	// TODO: Verify content
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading atomic file: %v\n", err)
		return
	}

	fmt.Printf("Content: %s", string(content))
}

// TODO: Demonstrate large file handling
func demonstrateLargeFileHandling() {
	fmt.Println("\n=== Large File Handling ===")

	// TODO: Create a large file for testing
	largeFile := "large_test.txt"

	fmt.Println("Creating large file...")
	file, err := os.Create(largeFile)
	if err != nil {
		fmt.Printf("Error creating large file: %v\n", err)
		return
	}
	defer file.Close()
	defer os.Remove(largeFile)

	// TODO: Write large amount of data efficiently
	writer := bufio.NewWriter(file)
	for i := 0; i < 100000; i++ {
		_, err := fmt.Fprintf(writer, "Line %d: This is a line with some content to make it longer\n", i+1)
		if err != nil {
			fmt.Printf("Error writing line %d: %v\n", i+1, err)
			return
		}

		// Flush periodically to avoid using too much memory
		if i%10000 == 0 {
			writer.Flush()
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing: %v\n", err)
		return
	}

	file.Close() // Close before reading

	// TODO: Read large file efficiently (streaming)
	fmt.Println("Reading large file efficiently...")

	file2, err := os.Open(largeFile)
	if err != nil {
		fmt.Printf("Error opening large file: %v\n", err)
		return
	}
	defer file2.Close()

	scanner := bufio.NewScanner(file2)
	lineCount := 0
	totalBytes := 0

	for scanner.Scan() {
		lineCount++
		totalBytes += len(scanner.Bytes())

		// Process every 10000th line to show progress
		if lineCount%10000 == 0 {
			fmt.Printf("Processed %d lines...\n", lineCount)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning large file: %v\n", err)
		return
	}

	fmt.Printf("Large file processing complete:\n")
	fmt.Printf("  Lines: %d\n", lineCount)
	fmt.Printf("  Total bytes processed: %d\n", totalBytes)

	// TODO: Get file size for comparison
	fileInfo, _ := os.Stat(largeFile)
	fmt.Printf("  File size on disk: %d bytes\n", fileInfo.Size())
}

func main() {
	fmt.Println("📁 Welcome to File Operations! 📁")
	fmt.Println("This file teaches you file reading, writing, and manipulation in Go")

	// TODO: Implement each demonstration function
	// Start with basic operations and progress to advanced techniques

	demonstrateFileReading()
	// demonstrateFileWriting()
	// demonstrateFileAppending()
	// demonstrateFileMetadata()
	// demonstrateTemporaryFiles()
	// demonstrateFileCopyMove()
	// demonstrateFileLocking()
	// demonstrateAtomicOperations()
	// demonstrateLargeFileHandling()

	fmt.Println("\n🎉 Congratulations! You've mastered file operations in Go!")
	fmt.Println("Next: Learn directory management in directory_management.go")
}

/*
🔍 Key Concepts to Remember:

1. **File Reading**: os.ReadFile(), os.Open(), bufio.Scanner()
2. **File Writing**: os.WriteFile(), os.Create(), bufio.Writer()
3. **Resource Management**: Always close files with defer
4. **Error Handling**: Check errors for all file operations
5. **Permissions**: Use appropriate file permissions (0644, 0600, etc.)
6. **Atomic Operations**: Use temp file + rename for atomic writes
7. **Large Files**: Use streaming with bufio for memory efficiency

📋 Essential File Operations:
```go
// Reading
content, err := os.ReadFile("file.txt")
file, err := os.Open("file.txt")
defer file.Close()

// Writing
err := os.WriteFile("file.txt", data, 0644)
file, err := os.Create("file.txt")
defer file.Close()

// Metadata
info, err := os.Stat("file.txt")
err := os.Chmod("file.txt", 0644)

// Temporary files
tmpFile, err := os.CreateTemp("", "prefix-*.txt")
defer os.Remove(tmpFile.Name())
```

🚨 Common Mistakes:
- Forgetting to close files (resource leak)
- Not checking errors from file operations
- Using wrong file permissions
- Not handling large files efficiently
- Not using atomic operations for critical writes
- Hardcoding file paths instead of using filepath.Join

🎯 Next Steps:
- Learn directory management and file system navigation
- Master environment configuration and CLI tools
- Practice with system interfaces and cross-platform code
- Build complete file processing applications
*/
