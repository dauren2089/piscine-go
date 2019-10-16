package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var aRune rune
	for aRune = 48; aRune < 58; aRune++ {
		z01.PrintRune(aRune)
	}
	z01.PrintRune(10)
}
