package piscine

func Rot14(s string) string {
	var result string
	for _, v := range s {
		for i := 0; i < 14; i++ {
			if v >= 97 && v <= 122 {
				if v < 122 {
					v += 1
				} else {
					v = 97
				}
			}
			if v >= 65 && v <= 90 {
				if v < 90 {
					v += 1
				} else {
					v = 65
				}
			}
		}
		result = result + string(v)
	}
	return result
}
