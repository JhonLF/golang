package piscine

func Capitalize(s string) string {
	sLower := ToLower(s)
	sByte := []byte(sLower)
	result := ""
	for i, v := range sByte {
		if i == 0 && IsUpper(s) == true {
			result += string(v - 32)
			continue
		} else if i == 0 && IsLower(s) == true {
			result += string(v - 32)
			continue
		}
		if i >= 1 {
			if PreviousCheck(sByte[i-1]) {
				if v >= 97 && v <= 122 {
					result += string(v - 32)
					continue
				}
			}
		}
		result += string(v)
	}
	return result
}

func PreviousCheck(b byte) bool {
	if b >= 32 && b <= 47 || b >= 58 && b <= 64 || b >= 91 && b <= 96 || b >= 123 && b <= 126 {
		return true
	}
	return false
}
