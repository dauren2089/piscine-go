package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		return
	}
	for _, value := range sortnbr(converttoint(n)) {
		z01.PrintRune(rune(value) + '0')
	}
}

func sortnbr(sentence []int) []int {
	for value := true; value; {
		value = false
		for i := 1; i < len(sentence); i++ {
			if sentence[i] < sentence[i-1] {
				sentence[i], sentence[i-1] = sentence[i-1], sentence[i]
				value = true
			}
		}
	}
	return sentence
}

func converttoint(n int) []int {
	var nbr []int
	for n > 0 {
		if n == 0 {
			nbr = append(nbr, 0)
		} else {
			nbr = append(nbr, n%10)
		}
		n /= 10
	}
	return nbr
}
