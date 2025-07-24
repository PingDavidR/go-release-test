// Package version provides version information about the application.
package version

import (
	"fmt"
	"runtime"
)

// Version variables. These will be populated at build time.

var (
	// Version is the current version of the application.
	Version = "0.4.0"

	// GitCommit is the git commit hash.
	GitCommit = "unknown"

	// BuildDate is the date when the binary was built.
	BuildDate = "unknown"

	// GoVersion is the version of Go used to build the binary.
	GoVersion = runtime.Version()

	// Platform is the operating system and architecture.
	Platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

// Info returns a string with version information.
func Info() string {
	return fmt.Sprintf(
		"Version: %s\nGit Commit: %s\nBuild Date: %s\nGo Version: %s\nPlatform: %s",
		Version,
		GitCommit,
		BuildDate,
		GoVersion,
		Platform,
	)
}

// ShortInfo returns a condensed version string.
func ShortInfo() string {
	return fmt.Sprintf("v%s (%s)", Version, GitCommit)
}
