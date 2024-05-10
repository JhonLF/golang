package piscine

func Itoa(n int) string {
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	if len(digits) == 0 {
		digits = append(digits, 0)
	}

	runes := make([]rune, len(digits))
	if isNegative {
		runes = append([]rune{'-'}, runes...)
	}

	for i := range digits {
		runes[len(runes)-1-i] = rune(digits[i] + '0')
	}

	return string(runes)
}
