// Package leap checks to see if a year is a leap
// year or not.
package leap

// IsLeapYear takes a year as an int and returns true or
// false if the year is a leap year.
func IsLeapYear(year int) bool {
	if year%4 == 0 && year%100 == 0 && year%400 == 0 {
		return true
	} else if year%4 == 0 && year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
