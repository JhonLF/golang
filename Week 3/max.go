package piscine

func Max(a []int) int {
	maxN := a[0]
	for _, v := range a {
		if v > maxN {
			maxN = v
		}
	}
	return maxN
}
