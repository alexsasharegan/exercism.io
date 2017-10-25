package hamming

import (
	"errors"
)

const testVersion = 6

// Distance compares two DNA strands and counts how many of the nucleotides
// are different from their equivalent in the other string.
func Distance(a, b string) (int, error) {
	var d int
	if len(a) != len(b) {
		return d, errors.New("strings must be of equal length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}

	return d, nil
}
