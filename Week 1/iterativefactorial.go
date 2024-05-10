package piscine

func RecursiveFactorial(nb int) int {
	if nb > 20 {
		return 0
	}
	result := 0
	finalResult := 0
	if nb == 1 || nb == 0 {
		result = 1
	}
	if nb > 1 {
		result = nb * RecursiveFactorial(nb-1)
	}
	if result < 0 {
		finalResult = 0
	} else {
		finalResult = result
	}
	return finalResult
}
