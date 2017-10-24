// Package twofer is short for two for one. One for you and one for me.
package twofer

import (
	"fmt"
)

// ShareWith is a function that if the given name is "Alice",
// the result should be "One for Alice, one for me."
// If no name is given, the result should be "One for you, one for me."
func ShareWith(s string) string {
	if s == "" {
		s = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", s)
}
