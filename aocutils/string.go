package aocutils

func UpdateCharAt(s string, r rune, i int) string {
	runes := []rune(s)

	if i > 0 && i < len(runes) {
		runes[i] = r
	}

	return string(runes)
}
