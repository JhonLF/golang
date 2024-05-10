package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	l := len([]rune(s))
	if n <= 0 || n > l {
		return '\x00'
	} else {
		return runes[n-1]
	}
}
