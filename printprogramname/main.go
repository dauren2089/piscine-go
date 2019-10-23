package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args

	printArgs(arguments[0])

	z01.PrintRune('\n')
}

func printArgs(str string) {

	runes := []rune(str)

	for _, value := range runes {

		z01.PrintRune(value)

	}
}
