// Package diffsquares provides methods for diffing squares
package diffsquares

const testVersion = 1

// Difference returns the difference between the square of the sums and the sum of the squares of the first given number of natural numbers
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

// SumOfSquares returns the sum of the squares of the first given number of natural numbers
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// SquareOfSums returns the square of the sums of the first given number of natural numbers
func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}
