// Package bob implements a solution for Exercism.io's Bob puzzle
package bob

import "strings"

// Hey talks to Bob
func Hey(remark string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	trimark := strings.TrimSpace(remark)
	switch {

	case trimark == "":
		return "Fine. Be that way!"

	case strings.HasSuffix(trimark, "?"):
		if strings.ToUpper(remark) == remark && strings.ContainsAny(remark, alphabet) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."

	case strings.ToUpper(remark) == remark && strings.ContainsAny(remark, alphabet):
		return "Whoa, chill out!"

	default:
		return "Whatever."

	}
}
