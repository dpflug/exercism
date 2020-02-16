// Implements a Exercism.io Proverb puzzle solution
package proverb

import "fmt"

// Proverb outputs a jouncy little rhyme in a slice of strings
func Proverb(rhyme []string) []string {
	items := len(rhyme)
	if items == 0 {
		return []string{}
	}
	out := make([]string, items)
	for i := 0; i < items-1; i++ {
		out[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}
	out[items-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	return out
}
