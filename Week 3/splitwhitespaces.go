package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	txt := ""
	for _, v := range s {
		if string(v) == " " {
			if txt != "" {
				result = append(result, txt)
				txt = ""
			}
		} else {
			txt += string(v)
		}
	}
	if txt != " " {
		result = append(result, txt)
		txt = ""
	}
	return result
}
