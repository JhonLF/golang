package piscine

func ActiveBits(n int) int {
	count := 0
	for n > 0 {
		if n%2 != 0 {
			count += 1
			n = n / 2
		} else {
			n = n / 2
		}
	}
	return count
}
