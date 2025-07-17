// Package calculator provides basic arithmetic operations.
package calculator

import "errors"

// Add returns the sum of two numbers.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference between two numbers.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers.
// Returns an error if the divisor is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
