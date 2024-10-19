package presentation

import (
	"bufio"
	"log"
	"os"

	"github.com/01-edu/z01"
)

func WordSearch(input int) {
	var data [10][10]rune
	switch input {
	case 2:
		data = [10][10]rune{
			{'s', 'h', 'f', 'h', 'm', 'g', 'm', 't', 'i', 'f'},
			{'q', 'a', 'z', 'd', 't', 'u', 'f', 'c', 'n', 'u'},
			{'p', 'â', 't', 'i', 's', 's', 'e', 'r', 'i', 'e'},
			{'g', 'r', 'a', 'i', 'n', 's', 'é', 'o', 't', 'e'},
			{'r', 'j', 't', 'm', 't', 'f', 'r', 'i', 's', 'r'},
			{'o', 'u', 'e', 'e', 'a', 't', 'd', 's', 'r', 't'},
			{'e', 't', 'i', 'c', 'u', 'a', 'a', 's', 'i', 'a'},
			{'v', 'e', 'r', 'r', 'e', 't', 'i', 'a', 'o', 'i'},
			{'o', 'd', 'u', 'o', 'a', 't', 'l', 'n', 'n', 'c'},
			{'a', 'a', 's', 's', 'i', 'e', 't', 't', 'e', 'i'},
		}
	case 3:
		data = [10][10]rune{
			{'s', 'p', 'd', 'a', 'e', 'e', 'v', 'i', 'd', 'r'},
			{'p', 'a', 'g', 'r', 'm', 'f', 'v', 'e', 'n', 't'},
			{'t', 'o', 'n', 'm', 'i', 'd', 'n', 'u', 'd', 'w'},
			{'y', 't', 'e', 't', 'i', 'p', 'a', 'c', 'e', 'v'},
			{'n', 'f', 'e', 'u', 'é', 'e', 'i', 'r', 'h', 'i'},
			{'m', 'e', 'd', 'i', 't', 'a', 't', 'i', 'o', 'n'},
			{'c', 't', 'r', 'a', 'u', 'l', 'i', 'w', 'r', 't'},
			{'a', 'i', 'b', 'q', 'd', 'a', 'a', 'o', 's', 'r'},
			{'e', 'f', 'e', 'n', 'e', 'c', 'o', 'h', 'a', 'e'},
			{'e', 'e', 'f', 'l', 'e', 'a', 'o', 't', 'r', 'i'},
		}
	default:
		data = [10][10]rune{
			{'c', 'o', 't', 'd', 't', 'r', 's', 'n', 'e', 'c'},
			{'r', 'e', 'e', 'o', 't', 'e', 'o', 'h', 'u', 'c'},
			{'ê', 'u', 'h', 'h', 't', 'r', 'l', 'a', 'o', 'a'},
			{'p', 'f', 'w', 'a', 'e', 'r', 'e', 'a', 'a', 'f'},
			{'e', 's', 't', 'p', 'r', 't', 'a', 'e', 'r', 'é'},
			{'s', 'a', 'p', 'y', 'u', 'i', 'o', 'e', 's', 'd'},
			{'p', 'a', 'e', 'o', 'i', 't', 'c', 'd', 't', 'e'},
			{'n', 'n', 'c', 'n', 'c', 'w', 'b', 'o', 'b', 'n'},
			{'f', 'o', 'u', 'r', 'c', 'h', 'e', 't', 't', 'e'},
			{'b', 'a', 'i', 'e', 's', 't', 'h', 'n', 'w', 's'},
		}
	}
	PrintLetterTable(data)
	bigstring := transform(data)
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)
	for Scanner.Scan() {
		Solver(string(Scanner.Text()), bigstring)
	}
	if err := Scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func transform(data [10][10]rune) string {
	var data_str string
	for i := 0; i < 10; i++ { //GAUCHE A DROITE
		for j := 0; j < 10; j++ {
			data_str += (string(data[i][j]))
		}
	}

	for i := 0; i < 10; i++ { //HAUT EN BAS
		for j := 0; j < 10; j++ {
			data_str += (string(data[j][i]))
		}
	}

	// DIAGONALE HAUT VERS BAS DROITE

	xStart := 9
	yStart := 1
	for {
		var xLoop, yLoop int
		if xStart >= 0 {
			xLoop = xStart
			yLoop = 0
			xStart--
		} else if yStart < 10 {
			xLoop = 0
			yLoop = yStart
			yStart++
		} else {
			break
		}
		for ; xLoop < 10 && yLoop < 10; xLoop++ {
			data_str += string(data[xLoop][yLoop])
			yLoop++
		}
	}

	for i := 0; i < 10*2; i++ { // DIAGONALE BAS VERS HAUT DROITE
		for j := 0; j <= i; j++ {
			k := i - j
			if k < 10 && j < 10 {
				data_str += (string(data[k][j]))
			}
		}
	}

	return data_str
}

func Solver(Scannedword string, bigstring string) {
	for i := 0; i < len(bigstring)-len(Scannedword); i++ {
		if bigstring[i:i+len(Scannedword)] == Scannedword {
			PrintStr(Scannedword)
			PrintStr(" ")
			return
		} else {
			continue
		}
	}
}

func PrintLetterTable(data [10][10]rune) {
	for i := 0; i < 10; i++ {
		PrintStr("[ ")
		for j := 0; j < 10; j++ {
			z01.PrintRune(data[i][j])
			z01.PrintRune(' ')
		}
		PrintStr("]\n")
	}
	PrintStr("\n\n")
}
