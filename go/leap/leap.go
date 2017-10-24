package leap

const testVersion = 3

// IsLeapYear given a year, report if it is a leap year.
func IsLeapYear(y int) bool {
	return y%4 == 0 && (y%100 != 0 || y%400 == 0)
}
