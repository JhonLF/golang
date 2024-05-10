package piscine

func IsPrintable(s string) bool {
	sByte := []byte(s)
	result := true
	for _, v := range sByte {
		if !(v >= 32 && v <= 126) {
			result = false
			break
		}
	}
	return result
}
