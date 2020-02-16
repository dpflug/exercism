// This package implements the Exercism.io Gigasecond puzzle solution
package gigasecond

import "time"
import "log"

// AddGigasecond takes a time and add one BILLION seconds. Mwahaha
func AddGigasecond(t time.Time) time.Time {
	newtime, err := time.ParseDuration("1000000000s")
	if err == nil {
		return t.Add(newtime)
	} else {
		log.Fatal(err)
		return t
	}
}
