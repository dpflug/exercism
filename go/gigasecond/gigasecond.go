// Package gigasecond implements a solution to the Exercism.io Gigasecond puzzle
package gigasecond

import "time"

// AddGigasecond takes a time and add one BILLION seconds. Mwahaha
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
