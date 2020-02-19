package diffsquares

// SquareOfSum returns the square of the sum of every number 1..n
func SquareOfSum(n int) int {
	var sum int
	for i := n; i > 0; i-- {
		sum += i
	}
	return sum * sum
}

// SumOfSquares returns the sum of every number 1..n, squared
func SumOfSquares(n int) int {
	var sum int
	for i := n; i > 0; i-- {
		sum += i * i
	}
	return sum
}

// Difference of the sum of squares and the square of sums
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
