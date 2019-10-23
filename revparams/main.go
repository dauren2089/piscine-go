package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args

	for i := strLenght(arguments) - 1; i >= 1; i-- {

		printRune(arguments[i])
		z01.PrintRune(10)
	}

}

func printRune(str string) {

	runes := []rune(str)

	for _, value := range runes {

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
