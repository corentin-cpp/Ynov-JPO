package main

import (
	"bufio"
	"fmt"
	"os"
	"presentation"

	"github.com/01-edu/z01"
)

func main() {
	var choice int
	reader := bufio.NewReader(os.Stdin)

	for {
		z01.PrintRune(1)
		fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
		fmt.Println("┃             \033[1;32mMENU\033[0m              ┃")
		fmt.Println("┃                               ┃")
		fmt.Println("┃ 1. Capitalize                 ┃ \033[1;30mPermet de mettre la première lettre de chaque mot en majuscule, et le reste est mis en minuscules.\033[0m")
		fmt.Println("┃ 2. TrimAtoi                   ┃ \033[1;30mPermet d'extraire un nombre contenu dans une chaîne de caractères.\033[0m")
		fmt.Println("┃ 3. PrintCombN                 ┃ \033[1;30mPermet de mettre un chiffre entre 1 et 9 en paramètre, et d'afficher toutes les combinaisons différentes possibles dans l'ordre croissant.\033[0m")
		fmt.Println("┃ 4. IsPrime                    ┃ \033[1;30mPermet de savoir si le nombre passé en paramètre est premier ou non.\033[0m")
		fmt.Println("┃ 5. IsUpper                    ┃ \033[1;30mAnalyse d'une chaîne de caractères pour savoir si tous les caractères de celle-ci sont en majuscules ou non.\033[0m")
		fmt.Println("┃ 6. StrRev                     ┃ \033[1;30mInverse tous les caractères d'une chaîne de caractère.\033[0m")
		fmt.Println("┃ 7. Fibonacci                  ┃ \033[1;30mPermet de passer en paramètre un entier et renvoie le nombre correspondant dans la suite de Fibonacci\033[0m")
		fmt.Println("┃ 8. PrintNbrInOrder            ┃ \033[1;30mPermet de mettre dans l'ordre croissant les chiffres du nombre passé en paramètre.\033[0m")
		fmt.Println("┃ 9. Compare                    ┃ \033[1;30mPermet de comparer deux chaînes de caractères.\033[0m")
		fmt.Println("┃ 10. AlphaCount                ┃ \033[1;30mPermet de compter le nombre de lettres dans une chaîne de caractères.\033[0m")
		fmt.Println("┃ 11. ToLower                   ┃ \033[1;30mPermet de mettre toutes les lettres de la chaîne de caractères en minuscules.\033[0m")
		fmt.Println("┃                               ┃")
		fmt.Println("┃ \033[1;31mRAIDS (Exams du Weekend)\033[0m      ┃")
		fmt.Println("┃                               ┃")
		fmt.Println("┃ 12. TextAn                    ┃ \033[1;30mAnalyse d'une chaîne de caractères, avec possibilité d'avoir des informations sur un mot en paramètre.\033[0m")
		fmt.Println("┃ 13. Word Search Solver   	┃ \033[1;30mPermet de résoudre une grille de mots croisés.\033[0m")
		fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
		fmt.Print("\033[1;32mChoisissez une option : \033[0m")
		fmt.Scan(&choice)

		switch choice {
		case 1: //capitalize
			fmt.Println("Écrivez une phrase à 'Capitalizer' :")
			userInput, _ := reader.ReadString('\n')
			fmt.Println("Résultat : " + presentation.Capitalize(userInput))
		case 2: //trimatoi
			fmt.Println("Écrivez une phrase avec des chiffres à extraire à l'intérieur :")
			userInput, _ := reader.ReadString('\n')
			fmt.Printf("Résultat : %d\n", presentation.TrimAtoi(userInput))
		case 3: //printcombn
			fmt.Println("Choisissez un chiffre entre 1 et 9 (inclus) : ")
			var intUserInput int
			fmt.Scan(&intUserInput)
			fmt.Println("Résultat :")
			presentation.PrintCombN(intUserInput)
		case 4: //isprime
			var intUserInput int
			fmt.Println("Choisissez un nombre pour vérifier s'il est premier ou non : ")
			fmt.Scan(&intUserInput)
			fmt.Printf("Résultat : %t\n", presentation.IsPrime(intUserInput))
		case 5: //isupper
			fmt.Println("Écrivez une phrase pour vérifier si elle ne contient que des majuscules ou non :")
			userInput, _ := reader.ReadString('\n')
			fmt.Print("Résultat : ")
			fmt.Println(presentation.IsUpper(userInput))
		case 6: //strrev
			fmt.Println("Écrivez une phrase à renverser :")
			userInput, _ := reader.ReadString('\n')
			fmt.Println("Résultat : " + presentation.StrRev(userInput))
		case 7: //fibonacci
			fmt.Println("Écrivez un nombre (pas trop grand de préférence) :")
			var intUserInput int
			fmt.Scan(&intUserInput)
			fmt.Printf("Résultat : %d\n", presentation.Fibonacci(intUserInput))
		case 8: //printnbrinorder
			fmt.Println("Écrivez un nombre à ranger dans l'ordre croissant :")
			var intUserInput int
			fmt.Scan(&intUserInput)
			fmt.Println("Résultat : ")
			presentation.PrintNbrInOrder(intUserInput)
			fmt.Println()
		case 9: //compare
			fmt.Println("Écrivez la première string à comparer :")
			str1, _ := reader.ReadString('\n')
			fmt.Println("Écrivez la deuxième string à comparer :")
			str2, _ := reader.ReadString('\n')
			fmt.Printf("Résultat : %d\n", presentation.Compare(str1, str2))
		case 10: //alphacount
			fmt.Println("Écrivez une phrase pour compter le nombre de lettres présentes :")
			userInput, _ := reader.ReadString('\n')
			fmt.Printf("Résultat : %d\n", presentation.AlphaCount(userInput))
		case 11: //tolower
			fmt.Println("Écrivez une phrase à mettre en minuscules :")
			userInput, _ := reader.ReadString('\n')
			fmt.Println("Résultat : " + presentation.ToLower(userInput))
		case 12: //TextAn
			var word string
			fmt.Println("Écrivez la string à analyser :")
			str, _ := reader.ReadString('\n')
			fmt.Println("Écrivez le mot en argument :")
			fmt.Scanln(&word)
			fmt.Println("Résultat : ")
			presentation.TextAn(str, word)
		case 13:
			fmt.Println("Choisissez un chiffre entre 1 et 3 (par défaut : 1) : ")
			var intUserInput int
			fmt.Scan(&intUserInput)
			fmt.Println("Résultat :")
			presentation.WordSearch(intUserInput)
		default:
			fmt.Println("Invalid choice, please try again.")
		}
		println()
	}
}
