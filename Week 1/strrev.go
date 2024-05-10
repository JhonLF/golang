package piscine

func StrRev(s string) string {
	var r string
	for i := len([]rune(s)) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}
