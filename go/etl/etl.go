package etl

import "strings"

func Transform(scores map[int][]string) map[string]int {
	out := make(map[string]int, len(scores))
	for score, letters := range scores {
		for _, letter := range letters {
			out[strings.ToLower(letter)] = score
		}
	}
	return out
}
