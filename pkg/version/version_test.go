package version

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	// Save original values
	origVersion := Version
	origGitCommit := GitCommit
	origBuildDate := BuildDate
	origGoVersion := GoVersion
	origPlatform := Platform

	// Restore original values after test
	defer func() {
		Version = origVersion
		GitCommit = origGitCommit
		BuildDate = origBuildDate
		GoVersion = origGoVersion
		Platform = origPlatform
	}()

	// Set test values
	Version = "1.2.3"
	GitCommit = "abcdef123456"
	BuildDate = "2025-07-29"
	GoVersion = "go1.19"
	Platform = "darwin/amd64"

	// Call the function
	info := Info()

	// Check that all values are present in the output
	if !strings.Contains(info, "Version: 1.2.3") {
		t.Errorf("Info() did not contain expected Version. Got: %s", info)
	}
	if !strings.Contains(info, "Git Commit: abcdef123456") {
		t.Errorf("Info() did not contain expected GitCommit. Got: %s", info)
	}
	if !strings.Contains(info, "Build Date: 2025-07-29") {
		t.Errorf("Info() did not contain expected BuildDate. Got: %s", info)
	}
	if !strings.Contains(info, "Go Version: go1.19") {
		t.Errorf("Info() did not contain expected GoVersion. Got: %s", info)
	}
	if !strings.Contains(info, "Platform: darwin/amd64") {
		t.Errorf("Info() did not contain expected Platform. Got: %s", info)
	}

	// Check the exact output format
	expected := fmt.Sprintf(
		"Version: %s\nGit Commit: %s\nBuild Date: %s\nGo Version: %s\nPlatform: %s",
		Version,
		GitCommit,
		BuildDate,
		GoVersion,
		Platform,
	)
	if info != expected {
		t.Errorf("Info() returned incorrect format.\nExpected: %s\nGot: %s", expected, info)
	}
}

func TestShortInfo(t *testing.T) {
	// Save original values
	origVersion := Version
	origGitCommit := GitCommit

	// Restore original values after test
	defer func() {
		Version = origVersion
		GitCommit = origGitCommit
	}()

	// Set test values
	Version = "1.2.3"
	GitCommit = "abcdef123456"

	// Call the function
	shortInfo := ShortInfo()

	// Check output
	expected := "v1.2.3 (abcdef123456)"
	if shortInfo != expected {
		t.Errorf("ShortInfo() returned incorrect value.\nExpected: %s\nGot: %s", expected, shortInfo)
	}
}

func TestRuntimeValues(t *testing.T) {
	// Test that runtime values are properly set
	if GoVersion != runtime.Version() {
		t.Errorf("GoVersion is not set to runtime.Version(). Expected: %s, Got: %s",
			runtime.Version(), GoVersion)
	}

	expectedPlatform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	if Platform != expectedPlatform {
		t.Errorf("Platform is not properly set. Expected: %s, Got: %s",
			expectedPlatform, Platform)
	}
}
