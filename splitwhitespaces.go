package piscine

import "fmt"

func SplitWhiteSpaces(str string) []string {

	size := strlenght(str)

	runesizes := getSize(str)

	splitstr := make([]string, runesizes)

	var newRunes []rune = make([]rune, size)

	var count int = int(0)

	fmt.Println(runesizes)

	i := 0
	j := 0

	for _, value := range str {

		//fmt.Println(value)

		if value != 32 {

			newRunes[i] = value

			splitstr[j] = string(newRunes)

			i++

			count++

		} else {

			clearRunes(newRunes, count)

			i = 0

			j++
		}
	}

	return splitstr
}

func clearRunes(clearing []rune, count int) []rune {

	for i := 0; i <= count; i++ {
		clearing[i] = 0
	}

	return clearing
}

func strlenght(str string) int {

	count := 0

	for range str {

		count++
	}

	return count
}

func getSize(str string) int {

	runecount := 0

	for _, value := range str {
		if value == 32 {
			runecount++
		}

	}
	fmt.Println(runecount)

	return runecount
}
