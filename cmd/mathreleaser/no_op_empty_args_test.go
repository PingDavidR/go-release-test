package main

import (
	"os"
	"strings"
	"testing"
)

// TestNoOpEmptyArgs tests the default case in mainInternal with no operation and empty args
func TestNoOpEmptyArgs(t *testing.T) {
	resetFlags()
	outBuf, errBuf, rOut, wOut, rErr, wErr := setup()
	defer teardown()

	// Set args to test default case with no arguments (only program name and unknown op)
	os.Args = []string{"mathreleaser", "-op="}
	mainInternal()

	stdout, stderr := getOutput(outBuf, errBuf, rOut, wOut, rErr, wErr)

	if !strings.Contains(stdout, "Usage:") {
		t.Errorf("Expected usage information for no operation and empty args, got stdout: %s, stderr: %s", stdout, stderr)
	}

	if exitCode != 1 {
		t.Errorf("Expected exit code 1 for no operation and empty args, got %d", exitCode)
	}
}
