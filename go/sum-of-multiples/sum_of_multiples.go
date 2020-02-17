package summultiples

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
	for k, _ := range seen {
		sum += k
	}
	return sum
}
