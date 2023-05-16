// Package gigasecond provides utilities to manipulate time in decimal increments.
package gigasecond

import (
	"time"
)

const gigasecond = 1e9 * time.Second

// AddGigasecond returns a new time set one gigasecond (~30 years) in the future of the provided time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
