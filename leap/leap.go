// Leap stub file

// The package name is expected by the test program.
package leap

import "time"

// testVersion should match the targetTestVersion in the test file.
const testVersion = 3

// Returns true or false depending on if the year is a leap year.
func IsLeapYear(year int) bool {
	t := time.Date(year, time.December, 31, 23, 0, 0, 0, time.UTC)
	isLeapYear := t.YearDay() == 366
	return isLeapYear
}
