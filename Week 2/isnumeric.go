package piscine

func IsNumeric(s string) bool {
	sByte := []byte(s)
	result := true
	for _, v := range sByte {
		if !(v >= 48 && v <= 57) {
			result = false
			break
		}
	}
	return result
}
