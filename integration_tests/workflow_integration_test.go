package integration_tests

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

// TestWorkflows tests realistic user workflows with sequences of operations
func TestWorkflows(t *testing.T) {
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

	// Test workflow 1: Basic calculation sequence
	t.Run("Basic calculation workflow", func(t *testing.T) {
		// Step 1: Start with adding two numbers
		addResult := runCommand(t, execPath, []string{"-op=add", "10", "5"})
		if !strings.Contains(addResult, "10 + 5 = 15.00") {
			t.Errorf("Addition failed: %s", addResult)
		}

		// Step 2: Take the result and multiply by another number
		multiplyResult := runCommand(t, execPath, []string{"-op=multiply", "15", "2"})
		if !strings.Contains(multiplyResult, "15 * 2 = 30.00") {
			t.Errorf("Multiplication failed: %s", multiplyResult)
		}

		// Step 3: Take the result and calculate square root
		sqrtResult := runCommand(t, execPath, []string{"-op=sqrt", "30"})
		if !strings.Contains(sqrtResult, "sqrt(30) = 5.48") {
			t.Errorf("Square root failed: %s", sqrtResult)
		}
	})

	// Test workflow 2: Trigonometric calculations
	t.Run("Trigonometric workflow", func(t *testing.T) {
		// Step 1: Calculate sine of an angle
		sinResult := runCommand(t, execPath, []string{"-op=sin", "0.5"})
		if !strings.Contains(sinResult, "sin(0.5) = 0.48") {
			t.Errorf("Sine calculation failed: %s", sinResult)
		}

		// Step 2: Calculate cosine of the same angle
		cosResult := runCommand(t, execPath, []string{"-op=cos", "0.5"})
		if !strings.Contains(cosResult, "cos(0.5) = 0.88") {
			t.Errorf("Cosine calculation failed: %s", cosResult)
		}

		// Step 3: Verify sin²(x) + cos²(x) = 1 by squaring and adding the results
		// First square the sine value
		sinSquared := runCommand(t, execPath, []string{"-op=power", "0.48", "2"})
		if !strings.Contains(sinSquared, "0.48 ^ 2 = 0.23") {
			t.Errorf("Sine squared calculation failed: %s", sinSquared)
		}

		// Then square the cosine value
		cosSquared := runCommand(t, execPath, []string{"-op=power", "0.88", "2"})
		if !strings.Contains(cosSquared, "0.88 ^ 2 = 0.77") {
			t.Errorf("Cosine squared calculation failed: %s", cosSquared)
		}

		// Finally add the squared values (should be approximately 1)
		sum := runCommand(t, execPath, []string{"-op=add", "0.23", "0.77"})
		if !strings.Contains(sum, "0.23 + 0.77 = 1.00") {
			t.Errorf("Sum of squares failed: %s", sum)
		}
	})

	// Test workflow 3: Error handling and recovery
	t.Run("Error handling workflow", func(t *testing.T) {
		// Step 1: Attempt to divide by zero
		divideByZeroResult := runCommandWithError(t, execPath, []string{"-op=divide", "10", "0"})
		if !strings.Contains(divideByZeroResult, "Error: Error performing division: division by zero") {
			t.Errorf("Divide by zero error not detected: %s", divideByZeroResult)
		}

		// Step 2: Recover by doing a valid calculation after the error
		validResult := runCommand(t, execPath, []string{"-op=add", "10", "5"})
		if !strings.Contains(validResult, "10 + 5 = 15.00") {
			t.Errorf("Recovery addition failed: %s", validResult)
		}
	})
}

// Helper function to run a command and return stdout
func runCommand(t *testing.T, execPath string, args []string) string {
	cmd := exec.Command(execPath, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to execute command %s %v: %v\nStderr: %s",
			execPath, args, err, stderr.String())
	}

	return stdout.String()
}

// Helper function to run a command expected to error and return stderr
func runCommandWithError(t *testing.T, execPath string, args []string) string {
	cmd := exec.Command(execPath, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err == nil {
		t.Fatalf("Expected command to fail, but it succeeded: %s %v", execPath, args)
	}

	return stderr.String()
}
