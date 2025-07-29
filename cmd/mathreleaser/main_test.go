package main

import (
	"bytes"
	"flag"
	"io"
	"os"
	"strings"
	"testing"
)

// Save original values to restore after tests
// This ensures that tests are isolated and don't affect each other
var (
	originalArgs    []string
	originalStdout  = os.Stdout
	originalStderr  = os.Stderr
	originalFlagSet = flag.CommandLine
	exitCode        int
)

// Mock exit function for testing
func mockExit(code int) {
	exitCode = code
	// Don't exit, just record the exit code
}

// Reset flags and arguments for each test
func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = originalArgs
	exitCode = 0
}

// Setup function to run before each test
func setup() (*bytes.Buffer, *bytes.Buffer, *os.File, *os.File, *os.File, *os.File) {
	// Capture stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	// Capture stderr
	rErr, wErr, _ := os.Pipe()
	os.Stderr = wErr

	// Create buffers to store output
	var outBuf, errBuf bytes.Buffer

	return &outBuf, &errBuf, rOut, wOut, rErr, wErr
}

// Tear down function to run after each test
func teardown() {
	os.Stdout = originalStdout
	os.Stderr = originalStderr
	flag.CommandLine = originalFlagSet
}

// Helper function to get captured output
func getOutput(outBuf, errBuf *bytes.Buffer, rOut, wOut, rErr, wErr *os.File) (string, string) {
	// Close write ends of pipes to flush all data
	wOut.Close()
	wErr.Close()

	// Read from the read ends of the pipes into the buffers
	_, _ = io.Copy(outBuf, rOut)
	_, _ = io.Copy(errBuf, rErr)

	// Return the captured output
	return outBuf.String(), errBuf.String()
}

// Save original args before running tests
func TestMain(m *testing.M) {
	originalArgs = os.Args

	// Store the original exit function and replace it with our mock
	originalOsExit := osExit
	osExit = mockExit

	code := m.Run()

	// Restore the original exit function
	osExit = originalOsExit

	os.Exit(code)
}

// Test version flag
func TestVersionFlag(t *testing.T) {
	// Setup
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	// Set args to test version flag
	os.Args = []string{"mathreleaser", "-version"}

	// Run main
	mainInternal()

	// Get captured output
	stdout, _ := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	// Check output
	if !strings.Contains(stdout, "Version:") {
		t.Errorf("Expected version information, got %s", stdout)
	}

	// Check exit code
	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

// Test add operation with valid arguments
func TestAddValidArgs(t *testing.T) {
	// Setup
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	// Set args to test add operation
	os.Args = []string{"mathreleaser", "-op=add", "5", "3"}

	// Run main
	mainInternal()

	// Get captured output
	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	// Check output
	if !strings.Contains(stdout, "5 + 3 = 8.00") {
		t.Errorf("Expected '5 + 3 = 8.00', got stdout: %s, stderr: %s", stdout, stderr)
	}

	// Check exit code
	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

// Test subtract operation with valid arguments
func TestSubtractValidArgs(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser", "-op=subtract", "10", "4"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stdout, "10 - 4 = 6.00") {
		t.Errorf("Expected '10 - 4 = 6.00', got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

// Test multiply operation with valid arguments
func TestMultiplyValidArgs(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser", "-op=multiply", "6", "7"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stdout, "6 * 7 = 42.00") {
		t.Errorf("Expected '6 * 7 = 42.00', got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

// Test divide operation with valid arguments
func TestDivideValidArgs(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser", "-op=divide", "20", "5"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stdout, "20 / 5 = 4.00") {
		t.Errorf("Expected '20 / 5 = 4.00', got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

// Test power operation with valid arguments
func TestPowerValidArgs(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser", "-op=power", "2", "3"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stdout, "2 ^ 3 = 8.00") {
		t.Errorf("Expected '2 ^ 3 = 8.00', got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

// Test sqrt operation with valid arguments
func TestSqrtValidArgs(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser", "-op=sqrt", "16"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stdout, "sqrt(16) = 4.00") {
		t.Errorf("Expected 'sqrt(16) = 4.00', got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

// Test sin operation with valid arguments
func TestSinValidArgs(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser", "-op=sin", "0"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stdout, "sin(0) = 0.00") {
		t.Errorf("Expected 'sin(0) = 0.00', got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

// Test cos operation with valid arguments
func TestCosValidArgs(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser", "-op=cos", "0"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stdout, "cos(0) = 1.00") {
		t.Errorf("Expected 'cos(0) = 1.00', got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

// Test invalid operation
func TestInvalidOperation(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser", "-op=invalid", "5", "3"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stderr, "Unknown operation") {
		t.Errorf("Expected error about unknown operation, got stdout: %s, stderr: %s", stdout, stderr)
	}
}

// Test binary operations with missing arguments
func TestBinaryOpMissingArgs(t *testing.T) {
	tests := []struct {
		name string
		op   string
		args []string
	}{
		{"add_no_args", "add", []string{}},
		{"add_one_arg", "add", []string{"5"}},
		{"add_too_many_args", "add", []string{"5", "3", "1"}},
		{"subtract_no_args", "subtract", []string{}},
		{"multiply_no_args", "multiply", []string{}},
		{"divide_no_args", "divide", []string{}},
		{"power_no_args", "power", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
			defer teardown()

			args := []string{"mathreleaser", "-op=" + tt.op}
			args = append(args, tt.args...)
			os.Args = args
			mainInternal()

			stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

			if !strings.Contains(stdout, "Usage:") {
				t.Errorf("Expected usage information for missing arguments, got stdout: %s, stderr: %s", stdout, stderr)
			}

			if exitCode != 1 {
				t.Errorf("Expected exit code 1 for missing arguments, got %d", exitCode)
			}
		})
	}
}

// Test unary operations with missing arguments
func TestUnaryOpMissingArgs(t *testing.T) {
	tests := []struct {
		name string
		op   string
		args []string
	}{
		{"sqrt_no_args", "sqrt", []string{}},
		{"sqrt_too_many_args", "sqrt", []string{"16", "4"}},
		{"sin_no_args", "sin", []string{}},
		{"cos_no_args", "cos", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
			defer teardown()

			args := []string{"mathreleaser", "-op=" + tt.op}
			args = append(args, tt.args...)
			os.Args = args
			mainInternal()

			stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

			if !strings.Contains(stdout, "Usage:") {
				t.Errorf("Expected usage information for missing arguments, got stdout: %s, stderr: %s", stdout, stderr)
			}

			if exitCode != 1 {
				t.Errorf("Expected exit code 1 for missing arguments, got %d", exitCode)
			}
		})
	}
}

// Test invalid number arguments
func TestInvalidNumberArgs(t *testing.T) {
	tests := []struct {
		name     string
		op       string
		args     []string
		errorMsg string
	}{
		{"add_invalid_first", "add", []string{"abc", "3"}, "Error parsing first number"},
		{"add_invalid_second", "add", []string{"5", "xyz"}, "Error parsing second number"},
		{"sqrt_invalid", "sqrt", []string{"invalid"}, "Error parsing number"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetFlags()
			outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
			defer teardown()

			args := []string{"mathreleaser", "-op=" + tt.op}
			args = append(args, tt.args...)
			os.Args = args
			mainInternal()

			stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

			if !strings.Contains(stderr, tt.errorMsg) {
				t.Errorf("Expected error message '%s', got stdout: %s, stderr: %s", tt.errorMsg, stdout, stderr)
			}

			if exitCode != 1 {
				t.Errorf("Expected exit code 1 for invalid number argument, got %d", exitCode)
			}
		})
	}
}

// Test divide by zero
func TestDivideByZero(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser", "-op=divide", "10", "0"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stderr, "Error performing division") {
		t.Errorf("Expected division by zero error, got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 1 {
		t.Errorf("Expected exit code 1 for division by zero, got %d", exitCode)
	}
}

// Test negative square root
func TestNegativeSqrt(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser", "-op=sqrt", "--", "-16"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stderr, "Error performing square root") {
		t.Errorf("Expected negative square root error, got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 1 {
		t.Errorf("Expected exit code 1 for negative square root, got %d", exitCode)
	}
}

// Test no operation specified with no arguments
func TestNoOpNoArgs(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	os.Args = []string{"mathreleaser"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stdout, "Usage:") {
		t.Errorf("Expected usage information for no operation and no arguments, got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 1 {
		t.Errorf("Expected exit code 1 for no operation and no arguments, got %d", exitCode)
	}
}
