package integration_tests

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestIntegration runs a series of integration tests against the built binary
func TestIntegration(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Find the binary to test
	execPath, err := findBinary()
	if err != nil {
		t.Fatalf("Could not find binary to test: %v", err)
	}

	t.Logf("Testing binary at: %s", execPath)

	// Table-driven tests according to the style guide
	tests := []struct {
		name          string
		args          []string
		expectedOut   string
		expectedErr   string
		expectedExit  int
		expectSuccess bool
	}{
		{
			name:          "Version flag",
			args:          []string{"-version"},
			expectedOut:   "Version:",
			expectSuccess: true,
		},
		{
			name:          "Addition",
			args:          []string{"-op=add", "5", "3"},
			expectedOut:   "5 + 3 = 8.00",
			expectSuccess: true,
		},
		{
			name:          "Subtraction",
			args:          []string{"-op=subtract", "10", "4"},
			expectedOut:   "10 - 4 = 6.00",
			expectSuccess: true,
		},
		{
			name:          "Multiplication",
			args:          []string{"-op=multiply", "6", "7"},
			expectedOut:   "6 * 7 = 42.00",
			expectSuccess: true,
		},
		{
			name:          "Division",
			args:          []string{"-op=divide", "20", "5"},
			expectedOut:   "20 / 5 = 4.00",
			expectSuccess: true,
		},
		{
			name:          "Division by zero",
			args:          []string{"-op=divide", "10", "0"},
			expectedErr:   "Error: Error performing division: division by zero",
			expectSuccess: false,
		},
		{
			name:          "Power",
			args:          []string{"-op=power", "2", "3"},
			expectedOut:   "2 ^ 3 = 8.00",
			expectSuccess: true,
		},
		{
			name:          "Square root",
			args:          []string{"-op=sqrt", "16"},
			expectedOut:   "sqrt(16) = 4.00",
			expectSuccess: true,
		},
		{
			name:          "Square root of negative",
			args:          []string{"-op=sqrt", "--", "-16"},
			expectedErr:   "Error performing square root: square root of negative number",
			expectSuccess: false,
		},
		{
			name:          "Sine",
			args:          []string{"-op=sin", "0"},
			expectedOut:   "sin(0) = 0.00",
			expectSuccess: true,
		},
		{
			name:          "Cosine",
			args:          []string{"-op=cos", "0"},
			expectedOut:   "cos(0) = 1.00",
			expectSuccess: true,
		},
		{
			name:          "Invalid operation",
			args:          []string{"-op=invalid", "5", "3"},
			expectedErr:   "Error: Unknown operation: invalid",
			expectSuccess: false,
		},
		{
			name:          "Missing arguments for add",
			args:          []string{"-op=add", "5"},
			expectedOut:   "Usage: mathreleaser -op=[add|subtract|multiply|divide|power] <number1> <number2>",
			expectSuccess: false,
		},
		{
			name:          "Invalid number format",
			args:          []string{"-op=add", "five", "3"},
			expectedErr:   "Error: Error parsing first number",
			expectSuccess: false,
		},
		{
			name:          "Missing arguments for sqrt",
			args:          []string{"-op=sqrt"},
			expectedOut:   "Usage: mathreleaser -op=[sqrt|sin|cos] <number>",
			expectSuccess: false,
		},
		{
			name:          "Default operation with arguments",
			args:          []string{"5", "3"},
			expectedOut:   "5 + 3 = 8.00", // Default operation is add
			expectSuccess: true,
		},
		{
			name:          "No arguments",
			args:          []string{},
			expectedOut:   "Usage:",
			expectSuccess: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command(execPath, tc.args...)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()

			// Check exit status
			exitCode := 0
			if err != nil {
				if exitError, ok := err.(*exec.ExitError); ok {
					exitCode = exitError.ExitCode()
				} else {
					t.Fatalf("Failed to execute command: %v", err)
				}
			}

			// Check output
			if tc.expectSuccess {
				if exitCode != 0 {
					t.Errorf("Expected success (exit code 0), but got %d", exitCode)
				}
				if !strings.Contains(stdout.String(), tc.expectedOut) {
					t.Errorf("Expected stdout to contain %q, got %q", tc.expectedOut, stdout.String())
				}
			} else {
				if exitCode == 0 {
					t.Errorf("Expected failure (non-zero exit code), but got 0")
				}
				if tc.expectedErr != "" && !strings.Contains(stderr.String(), tc.expectedErr) {
					t.Errorf("Expected stderr to contain %q, got %q", tc.expectedErr, stderr.String())
				}
				if tc.expectedOut != "" && !strings.Contains(stdout.String(), tc.expectedOut) {
					t.Errorf("Expected stdout to contain %q, got %q", tc.expectedOut, stdout.String())
				}
			}
		})
	}
}

// findBinary looks for the mathreleaser binary in common locations
func findBinary() (string, error) {
	// Common places to look for the binary
	searchPaths := []string{
		"../bin/mathreleaser-test",
		"../bin/mathreleaser",
		"../bin/mathreleaser-darwin-amd64",
		"../bin/mathreleaser-darwin-arm64",
	}

	for _, path := range searchPaths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			continue
		}
		if _, err := os.Stat(absPath); err == nil {
			return absPath, nil
		}
	}

	// As a fallback, build the binary if we couldn't find it
	cmd := exec.Command("go", "build", "-o", "../bin/mathreleaser-test", "../cmd/mathreleaser")
	if err := cmd.Run(); err != nil {
		return "", err
	}

	absPath, err := filepath.Abs("../bin/mathreleaser-test")
	if err != nil {
		return "", err
	}
	return absPath, nil
}
