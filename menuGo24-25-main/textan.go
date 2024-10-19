package presentation

import "github.com/01-edu/z01"

func isVowel(r rune) bool {
	vowels := "aeiouAEIOUéèêëàâäîïôöùûü"
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}

func isConsonant(r rune) bool {
	return !isVowel(r) && ((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'))
}

func containsDigit(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

func isPalindrome(s string) bool {
	if containsDigit(s) {
		return false
	}
	runes := []rune(s)
	if len(s) == 1 {
		return false
	}
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

func countOccurrences(str, word string) int {
	count := 0
	wordLength := len(word)
	for i := 0; i <= len(str)-wordLength; i++ {
		if str[i:i+wordLength] == word {
			count++
		}
	}
	return count
}

func TextAn(str string, word string) {
	wordCount := 0
	charCount := 0
	vowelCount := 0
	consonantCount := 0

	for _, r := range str {
		charCount++

		if r == ' ' || r == '\n' || r == '\t' {
			wordCount++
		} else {
			if isVowel(r) {
				vowelCount++
			} else if isConsonant(r) {
				consonantCount++
			}
		}
	}

	PrintStr("Amount of words : ")
	PrintNbr(wordCount)
	z01.PrintRune('\n')
	PrintStr("Amount of characters : ")
	PrintNbr(charCount)
	z01.PrintRune('\n')

	PrintStr("Amount of vowels : ")
	PrintNbr(vowelCount)
	z01.PrintRune('\n')

	PrintStr("Amount of consonants : ")
	PrintNbr(consonantCount)
	z01.PrintRune('\n')

	if len(word) == 0 {
		PrintStr("No word passed as parameter\n")
	} else {
		occurrences := countOccurrences(str, word)
		PrintStr(word + " occurs ")
		PrintNbr(occurrences)
		PrintStr(" times.\n")

		if isPalindrome(word) {
			PrintStr(word + " is a palindrome.\n")
		} else {
			PrintStr(word + " is not a palindrome.\n")
		}
	}
}
