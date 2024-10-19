package presentation

import (
	"github.com/01-edu/z01"
)

func BefNeg(n int) {
	i := '0'
	for j := 1; j <= n%10; j++ {
		i++
	}
	for j := -1; j >= n%10; j-- {
		i++
	}
	if n/10 != 0 {
		BefNeg(n / 10)
	}
	z01.PrintRune(i)
}

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	BefNeg(n)
}
