package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var aRune rune
	for aRune = 122; aRune > 96; aRune-- {
		z01.PrintRune(aRune)
	}
	z01.PrintRune(10)
}
