package isogram

import "unicode"

// IsIsogram determines if a string is an isogram
// (a word or phrase without a repeating letter)
// https://en.wikipedia.org/wiki/Isogram
func IsIsogram(s string) bool {
	var seen = make(map[rune]bool)
	for _, r := range s {
		if r == '-' || r == ' ' {
			continue
		}

		rl := unicode.ToLower(r)

		if seen[rl] {
			return false
		}

		seen[rl] = true
	}
	return true
}
