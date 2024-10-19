package presentation

func StrRev(s string) string {
	var strf string
	for j := len(s) - 1; j >= 0; j-- {
		strf += string(s[j])
	}
	return strf
}
