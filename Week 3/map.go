package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, v := range a {
		x := f(v)
		result = append(result, x)
	}
	return result
}
