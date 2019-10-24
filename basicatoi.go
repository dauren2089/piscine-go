package piscine

func BasicAtoi(s string) [10]int {

	//var number int
	var numberz [10]int

	runes := []rune(s)

	for i := 0; i <= findLenght(s)-1; i++ {

		if isItNbr(runes[i]) {

			numberz[i] = runeToInt(runes[i])

			// for j := 0; j <= findLenght(s)-1; j++ {
			//z01.PrintRune(runes[i])
			//numberz[j] = int(runes[i])
			// numberz[j] = runeToInt(runes[i])

			//	number = runeToInt(runes[i])

			// 	return number
			//}
		}
	}

	return numberz
}

func runeToInt(runes rune) int {

	nbr := 0

	for i := '0'; i < runes; i++ {
		nbr++
	}

	return nbr
}

func isItNbr(runes rune) bool {

	if runes < 47 || runes > 58 {
		return false
	}
	return true
}

func findLenght(str string) int {
	count := 0

	for index := range str {
		count = index + 1
	}

	return count
}
