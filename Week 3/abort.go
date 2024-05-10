package piscine

func Abort(a, b, c, d, e int) int {
	myslice := []int{a, b, c, d, e}
	for i := 0; i < len(myslice); i++ {
		for j := 0; j < len(myslice)-i-1; j++ {
			if myslice[j] > myslice[j+1] {
				myslice[j], myslice[j+1] = myslice[j+1], myslice[j]
			}
		}
	}
	n := (len(myslice) + 1) / 2
	result := myslice[n-1]
	return result
}
