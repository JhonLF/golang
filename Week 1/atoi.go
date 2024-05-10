package piscine

func Atoi(s string) int {
	isNeg := false
	countN := 0
	count := 0
	n := 0
	for _, v := range s {
		if !(v >= 48 && v <= 57) && v != 45 && v != 43 {
			return 0
		}
		if v == 45 {
			countN++
			isNeg = true
			if countN > 1 || count == 1 || n > 0 {
				return 0
			}
		}
		if v == 43 {
			count++
			if count > 1 || countN == 1 || n > 0 {
				return 0
			}
		}
		if v >= 48 && v <= 57 {
			v -= '0'
			n = n*10 + int(v)
		}
	}
	if isNeg == true {
		return -n
	}
	return n
}
