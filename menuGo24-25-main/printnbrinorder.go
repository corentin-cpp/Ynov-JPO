package presentation

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n > 0 {
		var o int
		var reverse []rune
		for i := 0; n != 0; i++ {
			reverse = append(reverse, rune(n%10+48))
			n /= 10
			o++
		}
		for j := range reverse {
			for k := range reverse {
				if reverse[j] > reverse[k] {
					reverse[j], reverse[k] = reverse[k], reverse[j]
				}
			}
		}
		for j := o - 1; j >= 0; j-- {
			z01.PrintRune(reverse[j])
		}
	} else if n == 0 {
		z01.PrintRune('0')
	}
}
