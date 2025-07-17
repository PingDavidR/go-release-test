// Package helpers provides utility functions for the application.
package helpers

import (
	"fmt"
	"os"
	"strings"
)

// FormatNumber formats a number with comma separators for thousands.
func FormatNumber(n float64) string {
	parts := strings.Split(fmt.Sprintf("%.2f", n), ".")
	integerPart := parts[0]
	decimalPart := parts[1]

	var result []byte
	for i, c := range integerPart {
		if i > 0 && (len(integerPart)-i)%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, byte(c))
	}

	return string(result) + "." + decimalPart
}

// PrintError prints an error message to stderr and exits with code 1.
func PrintError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

// EnsureDir ensures that a directory exists.
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}
