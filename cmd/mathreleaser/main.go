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

// Variable to hold the exit function, allowing it to be mocked in tests
var osExit = os.Exit

// Variable to hold the mainInternal function, allowing it to be mocked in tests
var mainInternalFunc = mainInternal

// mainInternal is a version of main that allows for exit code testing
func mainInternal() {
	// Define command-line flags
	versionFlag := flag.Bool("version", false, "Print version information")
	operation := flag.String("op", "add", "Operation to perform: add, subtract, multiply, divide, power, sqrt, sin, cos, tan, random")

	// Parse command-line flags
	flag.Parse()

	// Print version information if requested
	if *versionFlag {
		fmt.Println(version.Info())
		return
	}

	// Check arguments based on operation
	args := flag.Args()
	var a, b float64
	var err error

	switch *operation {
	case "add", "subtract", "multiply", "divide", "power", "random":
		if len(args) != 2 {
			fmt.Println("Usage: mathreleaser -op=[add|subtract|multiply|divide|power|random] <number1> <number2>")
			fmt.Println("       mathreleaser -version")
			osExit(1)
			return
		}
		a, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Error parsing first number: %v\n", err)
			osExit(1)
			return
		}
		b, err = strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Error parsing second number: %v\n", err)
			osExit(1)
			return
		}
	case "sqrt", "sin", "cos", "tan":
		if len(args) != 1 {
			fmt.Println("Usage: mathreleaser -op=[sqrt|sin|cos|tan] <number>")
			fmt.Println("       mathreleaser -version")
			osExit(1)
			return
		}
		a, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Error parsing number: %v\n", err)
			osExit(1)
			return
		}
	default:
		if len(args) == 0 {
			fmt.Println("Usage: mathreleaser -op=[add|subtract|multiply|divide|power|random] <number1> <number2>")
			fmt.Println("       mathreleaser -op=[sqrt|sin|cos|tan] <number>")
			fmt.Println("       mathreleaser -version")
			osExit(1)
			return
		}
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
			fmt.Fprintf(os.Stderr, "Error: Error performing division: %v\n", opErr)
			osExit(1)
			return
		}
		fmt.Printf("%s / %s = %s\n", args[0], args[1], helpers.FormatNumber(result))
	case "power":
		result = calculator.Power(a, b)
		fmt.Printf("%s ^ %s = %s\n", args[0], args[1], helpers.FormatNumber(result))
	case "sqrt":
		result, opErr = calculator.SquareRoot(a)
		if opErr != nil {
			fmt.Fprintf(os.Stderr, "Error: Error performing square root: %v\n", opErr)
			osExit(1)
			return
		}
		fmt.Printf("sqrt(%s) = %s\n", args[0], helpers.FormatNumber(result))
	case "sin":
		result = calculator.Sin(a)
		fmt.Printf("sin(%s) = %s\n", args[0], helpers.FormatNumber(result))
	case "cos":
		result = calculator.Cos(a)
		fmt.Printf("cos(%s) = %s\n", args[0], helpers.FormatNumber(result))
	case "tan":
		result = calculator.Tan(a)
		fmt.Printf("tan(%s) = %s\n", args[0], helpers.FormatNumber(result))
	case "random":
		result = calculator.Random(a, b)
		fmt.Printf("random(%s, %s) = %s\n", args[0], args[1], helpers.FormatNumber(result))
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown operation: %s\n", *operation)
		osExit(1)
		return
	}
}

func main() {
	mainInternalFunc()
}
