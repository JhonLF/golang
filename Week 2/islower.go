package piscine

func IsLower(s string) bool {
	sByte := []byte(s)
	result := true
	for _, v := range sByte {
		if v < 97 || v > 122 {
			result = false
			break
		}
	}
	return result
}
