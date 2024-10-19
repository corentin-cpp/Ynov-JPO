package presentation

func IsUpper(s string) bool {
	j := 0
	var res bool

	for range []rune(s) {
		if s[j] < 65 && s[j] > 90 {
			j--
		} else if s[j] >= 65 && s[j] <= 90 {
			j++
		}
	}
	if j >= len(s) {
		res = true
	}
	if j < len(s) {
		res = false
	}
	return res
}
