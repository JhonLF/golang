package piscine

func AlphaCount(s string) int {
	count := 0
	sByte := []byte(s)
	for i := range sByte {
		if sByte[i] >= 65 && sByte[i] <= 90 {
			count += 1
		}
		if sByte[i] >= 97 && sByte[i] <= 122 {
			count += 1
		}
	}
	return count
}
