package integration_tests

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

// TestEdgeCases tests the application with extreme inputs and edge cases
func TestEdgeCases(t *testing.T) {
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

	// Edge cases tests
	tests := []struct {
		name          string
		args          []string
		expectedOut   string
		expectedErr   string
		expectSuccess bool
	}{
		{
			name:          "Very large numbers - addition",
			args:          []string{"-op=add", "1000000000", "2000000000"},
			expectedOut:   "1000000000 + 2000000000 = 3,000,000,000.00",
			expectSuccess: true,
		},
		{
			name:          "Very large numbers - multiplication",
			args:          []string{"-op=multiply", "1000000", "1000000"},
			expectedOut:   "1000000 * 1000000 = 1,000,000,000,000.00",
			expectSuccess: true,
		},
		{
			name:          "Very small numbers - division",
			args:          []string{"-op=divide", "0.0000001", "1000000"},
			expectedOut:   "0.0000001 / 1000000 = 0.00",
			expectSuccess: true,
		},
		{
			name:          "Negative numbers - square root",
			args:          []string{"-op=sqrt", "--", "-100"},
			expectedErr:   "Error performing square root: square root of negative number",
			expectSuccess: false,
		},
		{
			name:          "Very large exponent",
			args:          []string{"-op=power", "10", "100"},
			expectedOut:   "10 ^ 100 =",
			expectSuccess: true,
		},
		{
			name:          "Division by very small number",
			args:          []string{"-op=divide", "1", "0.0000000001"},
			expectedOut:   "1 / 0.0000000001 = 10,000,000,000.00",
			expectSuccess: true,
		},
		{
			name:          "Floating point precision - addition",
			args:          []string{"-op=add", "0.1", "0.2"},
			expectedOut:   "0.1 + 0.2 = 0.30",
			expectSuccess: true,
		},
		{
			name:          "Sine of large angle",
			args:          []string{"-op=sin", "1000000"},
			expectedOut:   "sin(1000000) =",
			expectSuccess: true,
		},
		{
			name:          "Cosine of large angle",
			args:          []string{"-op=cos", "1000000"},
			expectedOut:   "cos(1000000) =",
			expectSuccess: true,
		},
		{
			name:          "Negative exponent",
			args:          []string{"-op=power", "10", "-2"},
			expectedOut:   "10 ^ -2 = 0.01",
			expectSuccess: true,
		},
		{
			name:          "Zero exponent",
			args:          []string{"-op=power", "123.456", "0"},
			expectedOut:   "123.456 ^ 0 = 1.00",
			expectSuccess: true,
		},
		{
			name:          "Sqrt of zero",
			args:          []string{"-op=sqrt", "0"},
			expectedOut:   "sqrt(0) = 0.00",
			expectSuccess: true,
		},
		{
			name:          "Very precise decimal",
			args:          []string{"-op=add", "0.12345678901234567890", "0.12345678901234567890"},
			expectedOut:   "0.12345678901234567890 + 0.12345678901234567890 = 0.25",
			expectSuccess: true,
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
			success := err == nil

			// Check output
			if tc.expectSuccess {
				if !success {
					t.Errorf("Expected success, but command failed: %v\nStderr: %s", err, stderr.String())
				}
				if !strings.Contains(stdout.String(), tc.expectedOut) {
					t.Errorf("Expected stdout to contain %q, got %q", tc.expectedOut, stdout.String())
				}
			} else {
				if success {
					t.Errorf("Expected failure, but command succeeded")
				}
				if tc.expectedErr != "" && !strings.Contains(stderr.String(), tc.expectedErr) {
					t.Errorf("Expected stderr to contain %q, got %q", tc.expectedErr, stderr.String())
				}
			}
		})
	}
}

// Test with custom test cases using command batches
func TestBatchOperations(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	execPath, err := findBinary()
	if err != nil {
		t.Fatalf("Could not find binary to test: %v", err)
	}

	// Complex calculation test: (10 + 5) * 2 - 3 = 27
	t.Run("Complex calculation sequence", func(t *testing.T) {
		// Step 1: 10 + 5 = 15
		add := exec.Command(execPath, "-op=add", "10", "5")
		addOut, err := add.Output()
		if err != nil {
			t.Fatalf("Addition command failed: %v", err)
		}
		addResult := strings.TrimSpace(string(addOut))
		if !strings.Contains(addResult, "10 + 5 = 15.00") {
			t.Errorf("Expected addition result to contain '10 + 5 = 15.00', got: %s", addResult)
		}

		// Step 2: 15 * 2 = 30
		multiply := exec.Command(execPath, "-op=multiply", "15", "2")
		multiplyOut, err := multiply.Output()
		if err != nil {
			t.Fatalf("Multiplication command failed: %v", err)
		}
		multiplyResult := strings.TrimSpace(string(multiplyOut))
		if !strings.Contains(multiplyResult, "15 * 2 = 30.00") {
			t.Errorf("Expected multiplication result to contain '15 * 2 = 30.00', got: %s", multiplyResult)
		}

		// Step 3: 30 - 3 = 27
		subtract := exec.Command(execPath, "-op=subtract", "30", "3")
		subtractOut, err := subtract.Output()
		if err != nil {
			t.Fatalf("Subtraction command failed: %v", err)
		}
		subtractResult := strings.TrimSpace(string(subtractOut))
		if !strings.Contains(subtractResult, "30 - 3 = 27.00") {
			t.Errorf("Expected subtraction result to contain '30 - 3 = 27.00', got: %s", subtractResult)
		}
	})

	// Test Pythagorean theorem: a² + b² = c² (3² + 4² = 5²)
	t.Run("Pythagorean theorem", func(t *testing.T) {
		// a² = 3² = 9
		squareA := exec.Command(execPath, "-op=power", "3", "2")
		squareAOut, err := squareA.Output()
		if err != nil {
			t.Fatalf("Power command for a² failed: %v", err)
		}
		squareAResult := strings.TrimSpace(string(squareAOut))
		if !strings.Contains(squareAResult, "3 ^ 2 = 9.00") {
			t.Errorf("Expected a² result to contain '3 ^ 2 = 9.00', got: %s", squareAResult)
		}

		// b² = 4² = 16
		squareB := exec.Command(execPath, "-op=power", "4", "2")
		squareBOut, err := squareB.Output()
		if err != nil {
			t.Fatalf("Power command for b² failed: %v", err)
		}
		squareBResult := strings.TrimSpace(string(squareBOut))
		if !strings.Contains(squareBResult, "4 ^ 2 = 16.00") {
			t.Errorf("Expected b² result to contain '4 ^ 2 = 16.00', got: %s", squareBResult)
		}

		// a² + b² = 9 + 16 = 25
		sum := exec.Command(execPath, "-op=add", "9", "16")
		sumOut, err := sum.Output()
		if err != nil {
			t.Fatalf("Addition command for a² + b² failed: %v", err)
		}
		sumResult := strings.TrimSpace(string(sumOut))
		if !strings.Contains(sumResult, "9 + 16 = 25.00") {
			t.Errorf("Expected sum result to contain '9 + 16 = 25.00', got: %s", sumResult)
		}

		// c = sqrt(25) = 5
		sqrt := exec.Command(execPath, "-op=sqrt", "25")
		sqrtOut, err := sqrt.Output()
		if err != nil {
			t.Fatalf("Square root command for c failed: %v", err)
		}
		sqrtResult := strings.TrimSpace(string(sqrtOut))
		if !strings.Contains(sqrtResult, "sqrt(25) = 5.00") {
			t.Errorf("Expected sqrt result to contain 'sqrt(25) = 5.00', got: %s", sqrtResult)
		}
	})
}
