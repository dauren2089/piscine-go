package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args

	lenght := strLenght(arguments) - 1

	arguments = arguments[1 : lenght+1]

	sortedArguments := sortNumbers(arguments, lenght)

	for _, value := range sortedArguments {

		for _, runes := range value {
			z01.PrintRune(runes)
		}

		z01.PrintRune(10)
	}

}

func strLenght(str []string) int {

	count := 0

	for index := range str {
		count = index + 1
	}

	return count
}

func sortNumbers(result []string, lenght int) []string {
	for i := 0; i < lenght; i++ {
		for j := i + 1; j < lenght; j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}
