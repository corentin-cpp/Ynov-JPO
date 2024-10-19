package presentation

func ToLower(s string) string {
	modstr := []rune(s)
	for j := 0; j < StrLen(s); j++ {
		if modstr[j] >= 65 && s[j] <= 90 {
			modstr[j] = modstr[j] + 32
		} else {
			continue
		}
	}
	s = string(modstr)
	return s
}
