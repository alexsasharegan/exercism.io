package gigasecond

import "time"

const testVersion = 4

// AddGigasecond calculates the moment when someone has lived for 10^9 seconds.
// A gigasecond is 10^9 (1,000,000,000) seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
