package accumulate

const testVersion = 1

// Accumulate is a map implementation.
func Accumulate(s []string, f func(string) string) []string {
	ret := make([]string, len(s))
	for i, x := range s {
		ret[i] = f(x)
	}
	return ret
}
