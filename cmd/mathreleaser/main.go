// Package main is the entry point for the application.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/PingDavidR/go-release-test/internal/helpers"
	"github.com/PingDavidR/go-release-test/pkg/calculator"
	"github.com/PingDavidR/go-release-test/pkg/version"
)

func main() {
	// Define command-line flags
	versionFlag := flag.Bool("version", false, "Print version information")
	operation := flag.String("op", "add", "Operation to perform: add, subtract, multiply, divide")

	// Parse command-line flags
	flag.Parse()

	// Print version information if requested
	if *versionFlag {
		fmt.Println(version.Info())
		return
	}

	// Check if we have the required number of arguments
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: mathreleaser -op=[add|subtract|multiply|divide] <number1> <number2>")
		fmt.Println("       mathreleaser -version")
		os.Exit(1)
	}

	// Parse the input numbers
	a, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		helpers.PrintError("Error parsing first number: %v", err)
	}

	b, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		helpers.PrintError("Error parsing second number: %v", err)
	}

	// Perform the requested operation
	var result float64
	var opErr error

	switch *operation {
	case "add":
		result = calculator.Add(a, b)
		fmt.Printf("%s + %s = %s\n", args[0], args[1], helpers.FormatNumber(result))
	case "subtract":
		result = calculator.Subtract(a, b)
		fmt.Printf("%s - %s = %s\n", args[0], args[1], helpers.FormatNumber(result))
	case "multiply":
		result = calculator.Multiply(a, b)
		fmt.Printf("%s * %s = %s\n", args[0], args[1], helpers.FormatNumber(result))
	case "divide":
		result, opErr = calculator.Divide(a, b)
		if opErr != nil {
			helpers.PrintError("Error performing division: %v", opErr)
		}
		fmt.Printf("%s / %s = %s\n", args[0], args[1], helpers.FormatNumber(result))
	default:
		helpers.PrintError("Unknown operation: %s", *operation)
	}
}
