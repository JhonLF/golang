package piscine

func Sqrt(nb int) int {
	result := 0
	for i := 1; i < nb+1; i++ {
		if i*i == nb {
			result = i
		}
	}
	return result
}
