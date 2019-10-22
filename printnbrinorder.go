package piscine

import (
	"github.com/01-edu/z01"
	"strconv"
)

func PrintNbrInOrder(n int) {
	sentence := strconv.Itoa(n)

	for i := 48; i <= 57; i++ {

		for _, value := range sentence {
			if value == rune(i) {
				z01.PrintRune(value)
			}
		}
	}

}
