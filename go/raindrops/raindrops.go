package raindrops

import (
	"fmt"
	"strings"
)

// Convert an int to a "raindrop"
// Just your basic FizzBuzz, really
func Convert(i int) string {
	var out strings.Builder
	if i%3 == 0 {
		out.WriteString("Pling")
	}
	if i%5 == 0 {
		out.WriteString("Plang")
	}
	if i%7 == 0 {
		out.WriteString("Plong")
	}
	if out.String() == "" {
		return fmt.Sprintf("%d", i)
	}
	return out.String()
}
