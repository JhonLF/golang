package piscine

func Any(f func(string) bool, a []string) bool {
	var result bool
	for _, v := range a {
		if f(v) == true {
			result = true
			break
		} else {
			result = false
		}
	}
	return result
}
