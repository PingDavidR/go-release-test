package helpers

import (
	"os"
	"path/filepath"
	"testing"
)

// Original exit function - will be restored after tests
var originalOsExit = osExit

// Restore original exit function after tests
func restoreOsExit() {
	osExit = originalOsExit
}

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

// TestPrintError tests the PrintError function by temporarily replacing stderr and os.Exit
func TestPrintError(t *testing.T) {
	// Save original stderr and restore after test
	originalStderr := os.Stderr
	defer func() { os.Stderr = originalStderr }()

	// Save original exit function and restore after test
	defer restoreOsExit()

	// Create a pipe to capture stderr output
	r, w, _ := os.Pipe()
	os.Stderr = w

	// Track if exit was called and with what code
	var exitCode int
	osExit = func(code int) {
		exitCode = code
	}

	// Test cases
	tests := []struct {
		name         string
		format       string
		args         []interface{}
		expectedText string
		expectedCode int
	}{
		{
			name:         "simple_message",
			format:       "Error: %s",
			args:         []interface{}{"test error"},
			expectedText: "Error: test error\n",
			expectedCode: 1,
		},
		{
			name:         "multiple_args",
			format:       "Error: %s, code: %d",
			args:         []interface{}{"test error", 42},
			expectedText: "Error: test error, code: 42\n",
			expectedCode: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset exitCode
			exitCode = 0

			// Call the function
			PrintError(tt.format, tt.args...)

			// Close the write end of the pipe to flush all data
			w.Close()

			// Read output from pipe
			var buf = make([]byte, 1024)
			n, _ := r.Read(buf)
			output := string(buf[:n])

			// Check if output matches expected
			if output != tt.expectedText {
				t.Errorf("PrintError() output = %q, want %q", output, tt.expectedText)
			}

			// Check if exit code is correct
			if exitCode != tt.expectedCode {
				t.Errorf("PrintError() exit code = %d, want %d", exitCode, tt.expectedCode)
			}

			// Reset stderr for next test
			r, w, _ = os.Pipe()
			os.Stderr = w
		})
	}
}
