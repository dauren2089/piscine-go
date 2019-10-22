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

	var answer []int64

	answer = IntToSlice(n, answer)

	SortInteger(answer)

	for i := range answer {

		for j := '0'; j <= '9'; j++ {

			if rune(answer[i]+48) == j {

				z01.PrintRune(j)

			}

		}
	}
}

func IntToSlice(n int64, sequence []int64) []int64 {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sequence = append([]int64{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}

func SortInteger(table []int64) {
	L := len(table)

	for i := 0; i < L; i++ {

		for j := 0; j < (L - 1 - i); j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
