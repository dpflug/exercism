// Implements a Exercism.io Go Bob puzzle solution
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
		} else {
			return "Sure."
		}
	case strings.ToUpper(remark) == remark && strings.ContainsAny(remark, alphabet):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
