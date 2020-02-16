package diffsquares

func SquareOfSum(n int) int {
	var sum int
	for i := n; i > 0; i-- {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	var sum int
	for i := n; i > 0; i-- {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
