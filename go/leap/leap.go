// Implements a Exercism.io Acronym bob puzzle solution
package leap

// IsLeapYear does what it says on the tin
func IsLeapYear(i int) bool {
	return i%400 == 0 || i%4 == 0 && !(i%100 == 0)
}
