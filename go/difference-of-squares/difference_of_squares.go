package diffsquares

// SquareOfSum returns the square of the sum of every number 1..n
func SquareOfSum(n int) int {
	sum := (n*n + n) / 2
	return sum * sum
}

// SumOfSquares returns the sum of every number 1..n, squared
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference of the sum of squares and the square of sums
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
