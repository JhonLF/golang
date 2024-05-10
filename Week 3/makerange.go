package piscine

func MakeRange(min, max int) []int {
	if min < 0 && max < 0 {
		if max < min {
			return []int(nil)
		}
	}
	if min == 0 && max == 0 {
		return []int(nil)
	}
	if min > max {
		return []int(nil)
	} else {
		l := max - min
		r := make([]int, l)
		x := min
		for i := 0; i < l; i++ {
			r[i] += x
			x += 1
		}
		return r
	}
}
