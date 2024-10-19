package presentation

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 || nb == 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	for j := 2; j < nb; j++ {
		if nb%j == 0 {
			return false
		}
	}
	return true
}
