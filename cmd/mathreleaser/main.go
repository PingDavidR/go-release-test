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
	operation := flag.String("op", "add", "Operation to perform: add, subtract, multiply, divide, power, sqrt, sin, cos")

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
	case "add", "subtract", "multiply", "divide", "power":
		if len(args) != 2 {
			fmt.Println("Usage: mathreleaser -op=[add|subtract|multiply|divide|power] <number1> <number2>")
			fmt.Println("       mathreleaser -version")
			os.Exit(1)
		}
		a, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			helpers.PrintError("Error parsing first number: %v", err)
		}
		b, err = strconv.ParseFloat(args[1], 64)
		if err != nil {
			helpers.PrintError("Error parsing second number: %v", err)
		}
	case "sqrt", "sin", "cos":
		if len(args) != 1 {
			fmt.Println("Usage: mathreleaser -op=[sqrt|sin|cos] <number>")
			fmt.Println("       mathreleaser -version")
			os.Exit(1)
		}
		a, err = strconv.ParseFloat(args[0], 64)
		if err != nil {
			helpers.PrintError("Error parsing number: %v", err)
		}
	default:
		if len(args) == 0 {
			fmt.Println("Usage: mathreleaser -op=[add|subtract|multiply|divide|power] <number1> <number2>")
			fmt.Println("       mathreleaser -op=[sqrt|sin|cos] <number>")
			fmt.Println("       mathreleaser -version")
			os.Exit(1)
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
			helpers.PrintError("Error performing division: %v", opErr)
		}
		fmt.Printf("%s / %s = %s\n", args[0], args[1], helpers.FormatNumber(result))
	case "power":
		result = calculator.Power(a, b)
		fmt.Printf("%s ^ %s = %s\n", args[0], args[1], helpers.FormatNumber(result))
	case "sqrt":
		result, opErr = calculator.SquareRoot(a)
		if opErr != nil {
			helpers.PrintError("Error performing square root: %v", opErr)
		}
		fmt.Printf("sqrt(%s) = %s\n", args[0], helpers.FormatNumber(result))
	case "sin":
		result = calculator.Sin(a)
		fmt.Printf("sin(%s) = %s\n", args[0], helpers.FormatNumber(result))
	case "cos":
		result = calculator.Cos(a)
		fmt.Printf("cos(%s) = %s\n", args[0], helpers.FormatNumber(result))
	default:
		helpers.PrintError("Unknown operation: %s", *operation)
	}
}
