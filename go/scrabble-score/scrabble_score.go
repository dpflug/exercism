package scrabble

import "strings"

// Score a word like Scrabble's base case
func Score(s string) int {
	var score int
	if s == "" {
		return 0
	}
	sl := strings.ToLower(s)
	for i := 0; i < len(s); i++ {
		switch sl[i] {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score++
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}
	return score
}
