// Package diffsquares provides mathematical functions for the squares of natural numbers.
package diffsquares

// SquareOfSum calculates the square of the sum of the first n natural numbers (starting from 1).
func SquareOfSum(n int) int {
	return sum(n) * sum(n)
}

func sum(n int) int {
	return n * (n + 1) / 2
}

// SumOfSquares calculates the sum of the square of the first n natural number (starting from 1).
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
