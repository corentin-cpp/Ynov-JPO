package presentation

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, str := range s {
		z01.PrintRune(rune(str))
	}
}
