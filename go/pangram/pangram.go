package pangram

import "unicode"

func IsPangram(s string) bool {
	letters := make(map[rune]bool, 26)
	for _, v := range s {
		if unicode.IsLetter(v) {
			letters[unicode.ToLower(v)] = true
		}
	}
	if len(letters) == 26 {
		return true
	} else {
		return false
	}
}
