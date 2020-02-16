package accumulate

func Accumulate(s []string, f func(string) string) []string {
	for k, v := range s {
		s[k] = f(v)
	}
	return s
}
