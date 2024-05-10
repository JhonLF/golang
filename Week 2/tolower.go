package piscine

func ToLower(s string) string {
	sByte := []byte(s)
	result := ""
	for _, v := range sByte {
		if v >= 65 && v <= 90 {
			result += string(v + 32)
		} else {
			result += string(v)
		}
	}
	return result
}
