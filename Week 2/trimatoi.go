package piscine

func TrimAtoi(s string) int {
	result := 0
	for _, v := range s {
		num := v - '0'
		if num >= 0 && num <= 9 {
			result = result*10 + int(num)
		}
	}
	if CheckNegative(s) {
		result = -result
	}
	return result
}

func CheckNegative(s string) bool {
	result := false
	for i, v := range s {
		if v == 45 {
			result = true
			for _, v2 := range s[:i] {
				if v2 >= 48 && v2 <= 57 {
					return false
				}
			}
		}
	}
	return result
}
