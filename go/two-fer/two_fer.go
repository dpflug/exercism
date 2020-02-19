// Package twofer implements a Exercism.io Two-Fer puzzle solution
package twofer

import "strings"

// ShareWith shares something with its argument, defaulting to "you"
func ShareWith(name string) string {
	var out strings.Builder

	out.WriteString("One for ")
	if name == "" {
		out.WriteString("you")
	} else {
		out.WriteString(name)
	}
	out.WriteString(", one for me.")

	return out.String()
}
