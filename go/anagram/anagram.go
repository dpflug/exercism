package anagram

import (
	"reflect"
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	var results = []string{}
	var sub_chars = make(map[rune]int)
	for _, char := range subject {
		sub_chars[unicode.ToLower(char)] += 1
	}
	for _, candidate := range candidates {
		var cand_chars = make(map[rune]int)
		if strings.ToLower(candidate) == strings.ToLower(subject) {
			continue
		}
		for _, char := range candidate {
			cand_chars[unicode.ToLower(char)] += 1
		}
		if reflect.DeepEqual(sub_chars, cand_chars) {
			results = append(results, candidate)
		}
	}
	return results
}
