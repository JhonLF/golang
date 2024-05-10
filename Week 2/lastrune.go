package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	l := len([]rune(s))
	return runes[l-1]
}
