package pangram

import "unicode"

// IsPangram determines if a string is a pangram
// (containing every letter of the alphabet)
func IsPangram(s string) bool {
	letters := make(map[rune]bool, 26)

	for _, v := range s {
		if unicode.IsLetter(v) {
			letters[unicode.ToLower(v)] = true
		}
	}

	if len(letters) == 26 {
		return true
	}

	return false
}
