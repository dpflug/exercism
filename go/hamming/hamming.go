package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("String lengths differ.")
	} else {
		if a == b {
			return 0, nil
		} else {
			var count int
			for i := 0; i < len(a); i++ {
				if a[i] != b[i] {
					count++
				}
			}
			return count, nil
		}
	}
}
