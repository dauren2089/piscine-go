package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int64) {
	if n < 0 {
		return
	}
	if n == 0 {

		z01.PrintRune('0')

	}

	var sentense []int64

	sentense = IntToSlice(n, sentense)

	SortInteger(sentense)

	for i := range sentense {

		for j := '0'; j <= '9'; j++ {

			if rune(sentense[i]+48) == j {

				z01.PrintRune(j)

			}

		}
	}
}

func IntToSlice(n int64, sentense []int64) []int64 {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sentense = append([]int64{i}, sentense...)
		return IntToSlice(n/10, sentense)
	}
	return sentense
}

func SortInteger(str []int64) {
	L := len(str)

	for i := 0; i < L; i++ {

		for j := 0; j < (L - 1 - i); j++ {
			if str[j] > str[j+1] {
				str[j], str[j+1] = str[j+1], str[j]
			}
		}
	}
}
