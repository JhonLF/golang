package piscine

func IsPrime(nb int) bool {
	result := false
	if nb > 2 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				result = false
				break
			} else {
				result = true
			}
		}
	} else if nb <= 1 {
		result = false
	} else {
		result = true
	}
	return result
}
