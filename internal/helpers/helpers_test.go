package helpers

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFormatNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"zero", 0, "0.00"},
		{"small_positive", 42.5, "42.50"},
		{"small_negative", -42.5, "-42.50"},
		{"thousands", 1234.56, "1,234.56"},
		{"millions", 1234567.89, "1,234,567.89"},
		{"negative_thousands", -1234.56, "-1,234.56"},
		{"large_number", 1234567890.12, "1,234,567,890.12"},
		{"integer", 42, "42.00"},
		{"small_decimal", 0.42, "0.42"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatNumber(tt.input)
			if got != tt.expected {
				t.Errorf("FormatNumber(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestEnsureDir(t *testing.T) {
	// Create temporary directory for testing
	tmpDir := filepath.Join(os.TempDir(), "mathreleaser-test")
	defer os.RemoveAll(tmpDir) // Clean up after test
	
	tests := []struct {
		name      string
		path      string
		shouldErr bool
	}{
		{"simple_dir", filepath.Join(tmpDir, "test1"), false},
		{"nested_dir", filepath.Join(tmpDir, "test2", "nested"), false},
		{"deep_nested_dir", filepath.Join(tmpDir, "test3", "nested1", "nested2", "nested3"), false},
		{"existing_dir", filepath.Join(tmpDir, "test4"), false}, // Will create first then test again
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EnsureDir(tt.path)
			if (err != nil) != tt.shouldErr {
				t.Errorf("EnsureDir(%v) error = %v, shouldErr %v", tt.path, err, tt.shouldErr)
				return
			}
			
			// Verify directory exists
			if !tt.shouldErr {
				if _, err := os.Stat(tt.path); os.IsNotExist(err) {
					t.Errorf("EnsureDir(%v) did not create directory", tt.path)
				}
			}
			
			// Test idempotence - calling again should not error
			err = EnsureDir(tt.path)
			if err != nil {
				t.Errorf("EnsureDir(%v) second call error = %v", tt.path, err)
			}
		})
	}
}

// We don't test PrintError since it calls os.Exit(1), which would terminate the test.
// To test this properly, we'd need to refactor the function to make it more testable,
// for example by accepting an io.Writer for output and a function for exiting.
// Here's a comment explaining this decision:
/*
PrintError is not tested because it calls os.Exit(1), which would terminate the test process.
To make this function testable, we could refactor it to:
1. Accept an io.Writer parameter for the output
2. Accept a function parameter for the exit behavior
3. Return an error instead of exiting

Example of a more testable version:
```go
func PrintError(w io.Writer, exitFunc func(int), format string, args ...interface{}) {
    fmt.Fprintf(w, format+"\n", args...)
    exitFunc(1)
}
```
*/
