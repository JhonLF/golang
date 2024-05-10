package piscine

func ToUpper(s string) string {
	sByte := []byte(s)
	result := ""

	for _, v := range sByte {
		if v >= 97 && v <= 122 {
			result += string(v - 32)
		} else {
			result += string(v)
		}
	}
	return result
}
