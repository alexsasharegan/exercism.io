package clock

import (
	"fmt"
)

const testVersion = 4

const (
	hour = 60
	day  = 24 * hour
)

// Clock is a type identity around minutes.
type Clock int

// New returns a new Clock.
func New(hour, minute int) Clock {
	var c Clock
	return c.Add(hour*60 + minute)
}

// String converts a Clock to readable string.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add will increment the minutes of the Clock.
func (c Clock) Add(minutes int) Clock {
	c += Clock(minutes)
	c %= day
	if c < 0 {
		c += day
	}
	return c
}
