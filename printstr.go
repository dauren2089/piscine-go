package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {

	slice := []rune(str)
	for _, word := range slice {
		z01.PrintRune(word)
	}
	z01.PrintRune(10)
}
