package presentation

func Compare(a, b string) int {
	var c int
	if a == b {
		c = 0
	}
	if a < b {
		c = -1
	}
	if a > b {
		c = 1
	}
	return (c)
}
