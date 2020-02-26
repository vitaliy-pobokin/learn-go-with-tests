package iteration

// Repeat returns a new string consisting of times copies of the string s.
func Repeat(char string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += char
	}
	return repeated
}
