package presentation

func AlphaCount(s string) int {
	j := 0
	k := 0

	for range []rune(s) {
		if s[j] >= 65 && s[j] <= 90 || s[j] >= 97 && s[j] <= 122 {
			k++
		}
		if j < StrLen(s) {
			j++
		} else {
			continue
		}
	}
	return k
}
