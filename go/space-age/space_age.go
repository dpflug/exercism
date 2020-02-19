// Package space implements an Exercism.io Space Age puzzle solution
package space

// Planet is a planet name
type Planet string

// Age takes a time in seconds and a planet name and
// returns the planetary year-equivalent
func Age(t float64, planet Planet) float64 {
	years := t / 31557600.0
	switch planet {
	case "Mercury":
		return years / 0.2408467
	case "Venus":
		return years / 0.61519726
	case "Earth":
		return years
	case "Mars":
		return years / 1.8808158
	case "Jupiter":
		return years / 11.862615
	case "Saturn":
		return years / 29.447498
	case "Uranus":
		return years / 84.016846
	case "Neptune":
		return years / 164.79132
	default:
		return -1
	}
}
