package piscine

func StrLen(s string) int {
	// unicode/utf8 utf8.RuneLen(s)
	l := len([]rune(s))
	return l
}
