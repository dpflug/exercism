package isogram

import "unicode"

func IsIsogram(s string) bool {
	var seen = make(map[rune]bool)
	for _, r := range s {
		if r == '-' || r == ' ' {
			continue
		}
		rl := unicode.ToLower(r)
		if seen[rl] {
			return false
		} else {
			seen[rl] = true
		}
	}
	return true
}
