package presentation

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}

// écrire la fnction qui calcule la fibonacci puis comparer à
// l'int en paramètre
//
//
