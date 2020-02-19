package accumulate

// Accumulate approximates other languages' map() function. Given an array of
// strings and a function, return the array produced by running the function on
// each string.
func Accumulate(s []string, f func(string) string) []string {
	for k, v := range s {
		s[k] = f(v)
	}
	return s
}
