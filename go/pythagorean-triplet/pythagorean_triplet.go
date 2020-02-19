// Package pythagorean for working with Pythagorean triplets
package pythagorean

//import "log"

// Triplet is a 3-int representing a Pythagorean triplet
type Triplet [3]int

// Range finds Pythagorean triplets with side lenghts between 2 bounds
func Range(min, max int) []Triplet {
	maxsum := max + max + max - 2
	var triplets []Triplet
	for a := min; a < (maxsum / 3); a++ {
		for b := a + 1; b < (maxsum - 2*a - 1); b++ {
			for c := b + 1; c < (maxsum - a - b); c++ {
				if a*a+b*b == c*c {
					triplets = append(triplets, [3]int{a, b, c})
				}
			}
		}
	}
	return triplets
}

// Sum finds Pythagorean triplets whose side lengths sum to give number
func Sum(p int) []Triplet {
	var triplets []Triplet
	for a := 1; a < p/3; a++ {
		for b := a + 1; b < p-2*a-1; b++ {
			c := p - b - a
			if a*a+b*b == c*c && a+b+c == p {
				triplets = append(triplets, [3]int{a, b, c})
			}
		}
	}
	return triplets
}
