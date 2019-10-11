// Package diffsquares provides functions for determining the square of the sums
// the sume of the squares and the difference between the two
package diffsquares

// SquareOfSum takes an integer and returns the square of the sum of that integer
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return (sum * sum)
}

// SumOfSquares takes an integer nad returns the sum of the squares of that integer
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference takes an integer and returns the difference between the sum of the
// squares and the square of the sum
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
