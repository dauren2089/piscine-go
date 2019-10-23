package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args

	for i := 32; i <= 125; i++ {
		//z01.PrintRune(rune(i))
		//z01.PrintRune(10)

		for j := 1; j < strLenght(arguments); j++ {
			//printRune(arguments[j])
			//z01.PrintRune(10)
			if string(i) == arguments[j] {
				printRune(arguments[j])
				z01.PrintRune(10)
			}
		}
	}

}

func printRune(str string) {

	Runes := []rune(str)

	for _, value := range Runes {

		z01.PrintRune(value)

	}
}

func strLenght(str []string) int {

	count := 0

	for index := range str {
		count = index + 1
	}

	return count
}
