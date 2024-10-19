package presentation

func TrimAtoi(s string) int {
	var finres int
	neg := 1
	for i := range s {
		if s[i] == '-' && finres == 0 {
			neg = -1
		}
		if s[i] < 48 || s[i] > 57 {
			continue
		} else {
			finres += int(s[i] - 48)
			finres = finres * 10
		}
		if i == len(s) {
			return neg
		}
		if i+1 == len(s) && finres == 0 {
			neg = 0
		}
	}
	neg *= finres
	return neg / 10
}
