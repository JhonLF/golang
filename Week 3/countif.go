package piscine

func CountIf(f func(string) bool, tab []string) int {
	result := 0
	for _, v := range tab {
		if f(v) == true {
			result += 1
		}
	}
	return result
}
