// Implements a Exercism.io Triangle puzzle solution
package triangle

import "math"

type Kind int

const (
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

// KindFromSides takes 3 side lengths of a triangle and returns what type it is
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a + b + c) {
		// arguments are numeric,
		return NaT
	}
	if math.IsInf(a+b+c, 0) {
		// not infinite,
		return NaT
	}
	if a <= 0 || b <= 0 || c <= 0 {
		// and make sense
		return NaT
	}
	if a+b < c || b+c < a || c+a < b {
		// Triangle inequality
		return NaT
	}
	switch {
	case a == b && b == c && c == a:
		return Equ
	case a == b || b == c || c == a:
		return Iso
	default:
		return Sca
	}
}
