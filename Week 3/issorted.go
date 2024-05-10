package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	desc := true
	asc := true

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) > 0 {
				asc = false
			}
		}
	}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) <= 0 {
				desc = false
			}
		}
	}
	if asc || desc {
		return true
	} else {
		return false
	}
}

func f(a, b int) int {
	if a < 0 && b < 0 {
		if a < b {
			return 1
		} else if a == b {
			return 0
		} else {
			return -1
		}
	} else {
		if a > b {
			return 1
		} else if a == b {
			return 0
		} else {
			return -1
		}
	}
}
