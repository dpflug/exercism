package raindrops

import (
	"strconv"
)

// Convert an int to a "raindrop"
// Just your basic FizzBuzz, really
func Convert(i int) string {
	var out = ""
	if i%3 == 0 {
		out +="Pling"
	}
	if i%5 == 0 {
		out += "Plang"
	}
	if i%7 == 0 {
		out += "Plong"
	}
	if out == "" {
		return strconv.Itoa(i)
	}
	return out
}
