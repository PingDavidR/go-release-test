// Package calculator provides basic arithmetic operations.
package calculator

import (
	"errors"
	"math"
)

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

// Power returns the result of raising a to the power of b.
func Power(a, b float64) float64 {
	return math.Pow(a, b)
}

// SquareRoot returns the square root of a number.
// Returns an error if the number is negative.
func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative number")
	}
	return math.Sqrt(a), nil
}

// Sin returns the sine of an angle (in radians).
func Sin(a float64) float64 {
	return math.Sin(a)
}

// Cos returns the cosine of an angle (in radians).
func Cos(a float64) float64 {
	return math.Cos(a)
}

// Tan returns the tangent of an angle (in radians).
// Note: Returns ±Inf when the angle is close to ±π/2 + nπ
// (where the tangent is undefined).
func Tan(a float64) float64 {
	return math.Tan(a)
}
