// Package leap provides methods related to leap years
package leap

const testVersion = 2

// IsLeapYear determines whether a year is a leap year.
func IsLeapYear(y int) bool {
        switch {
		case y%4 != 0:
			return false
		case y%100 != 0:
			return true
		case y%400 == 0:
			return true
		default:
			return false
        }
}
