package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args
	asciitable := string("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	//for i := 32; i <= 125; i++ {
	//z01.PrintRune(rune(i))
	//z01.PrintRune(10)

	for _, value := range asciitable {
		strvalue := string(value)
		for j := 1; j < strLenght(arguments); j++ {
			//printRune(arguments[j])
			//z01.PrintRune(10)

			if strvalue == arguments[j] {
				printRune(arguments[j])
				z01.PrintRune(10)
			}
		}
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
