// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of subtracting the second
// from the first.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of subtracting the second
// from the first.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("invalid inputs(%f, %f): division by zero error", a, b)
	}
	return a / b, nil
}

// Sqrt takes one numbers and returns the square root
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("invalid input(%f): division by zero error", a)
	}
	return math.Sqrt(a), nil
}
