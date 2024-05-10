package piscine

func Compact(ptr *[]string) int {
	var s []string
	for i := 0; i < len(*ptr)-1; i++ {
		if (*ptr)[i] != "" {
			s = append(s, (*ptr)[i])
		}
	}
	*ptr = s
	return len(s)
}
