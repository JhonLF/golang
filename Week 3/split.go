package piscine

func Split(s, sep string) []string {
	var result []string
	contador := 0

	for i := 0; i <= len(s)-len(sep); i++ {
		if i == len(s) || s[i:i+len(sep)] == sep {
			if contador < i {
				result = append(result, s[contador:i])
			}
			contador = i + len(sep)
		}
	}
	if contador < len(s) {
		result = append(result, s[contador:])
	}

	return result
}
