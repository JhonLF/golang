package piscine

func AppendRange(min, max int) []int {
	var r []int
	if min > max {
		return r
	} else {
		for i := min - 1; i < max-1; i++ {
			r = append(r, i+1)
		}
	}
	return r
}
