package summultiples

// SumMultiples takes an upper limit and some divisors and outputs all
// multiples of the divisors under the limit
func SumMultiples(limit int, divisors ...int) int {
	seen := map[int]bool{}
	for _, factor := range divisors {
		if factor > 0 {
			iter := factor
			for iter < limit {
				seen[iter] = true
				iter += factor
			}
		}
	}
	var sum int = 0
	for k := range seen {
		sum += k
	}
	return sum
}
