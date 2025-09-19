package gigasecond

import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
    duration, _ := time.ParseDuration("1000000000s")
	t = t.Add(duration)
	return t
}
