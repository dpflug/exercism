package lsproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (int, error) {
	maxp := 0
	if span > len(digits) {
		return -1, errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}
	for i := span - 1; i < len(digits); i++ {
		thisp := 1
		for j := 0; j < span; j++ {
			cur := int(digits[i-j] - '0')
			if cur > 9 || cur < 0 {
				return -1, errors.New("digits input must only contain digits")
			}
			thisp = thisp * cur
		}
		if thisp > maxp {
			maxp = thisp
		}
	}
	return maxp, nil
}
