package etl

import "strings"

// Transform a map of scores (ints) to a list of upper case letter (strings) to
// a map of lower case letter (strings) to scores (ints).
func Transform(scores map[int][]string) map[string]int {
	out := make(map[string]int, len(scores))
	for score, letters := range scores {
		for _, letter := range letters {
			out[strings.ToLower(letter)] = score
		}
	}
	return out
}
