package isogram

import "strings"

func IsIsogram(s string) bool {
	var count = make(map[byte]int)
	if s == "" {
		return true
	}
	sl := strings.ToLower(s)
	for i := 0; i < len(sl); i++ {
		c := sl[i]
		if c != '-' && c != ' ' {
			if count[c] == 0 {
				count[c] = 1
			} else {
				return false
			}
		}
	}
	return true
}
