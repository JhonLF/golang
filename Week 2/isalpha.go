package piscine

func IsAlpha(s string) bool {
	sByte := []byte(s)
	result := true
	for _, v := range sByte {
		if !(v >= 97 && v <= 122 || v >= 65 && v <= 90 || v >= 48 && v <= 57) {
			result = false
			break
		}
	}
	return result
}
