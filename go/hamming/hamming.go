package hamming

import "errors"

// Distance computes the Hamming Distance between two strings
// https://en.wikipedia.org/wiki/Hamming_distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("string lengths differ")
	}

	if a == b {
		return 0, nil
	}

	var count int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
