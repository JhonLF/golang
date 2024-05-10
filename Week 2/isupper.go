package piscine

func IsUpper(s string) bool {
	sByte := []byte(s)
	result := true
	for _, v := range sByte {
		if v < 65 || v > 90 {
			result = false
			break
		}
	}
	return result
}
