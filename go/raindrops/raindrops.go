package raindrops

import (
	"fmt"
)

const testVersion = 3

// Convert is a fizz-buzz variant.
func Convert(i int) (s string) {
	if i%3 == 0 {
		s += "Pling"
	}
	if i%5 == 0 {
		s += "Plang"
	}
	if i%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		return fmt.Sprintf("%d", i)
	}
	return
}
