package main

import (
	"os"
	"strings"
	"testing"
)

// TestDefaultWithArgs tests the default case in mainInternal with arguments
func TestDefaultWithArgs(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	// Set args to test default case with some arguments
	// This will trigger the "default" case but len(args) will not be 0
	os.Args = []string{"mathreleaser", "-op=unsupported", "some", "arguments"}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stderr, "Unknown operation") {
		t.Errorf("Expected 'Unknown operation' error, got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 1 {
		t.Errorf("Expected exit code 1 for unknown operation, got %d", exitCode)
	}
}
