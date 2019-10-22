package piscine

func Capitalize(s string) string {

	var checkbool bool

	sentence := []rune(s)

	for i := 0; i < strLenght(s); i++ {

		if isItWord(sentence[i]) == true && checkbool {

			if sentence[i] >= 'a' && sentence[i] <= 'z' {
				sentence[i] = 'A' - 'a' + sentence[i]

			}
			checkbool = false
		} else if sentence[i] >= 'A' && sentence[i] <= 'Z' {
			sentence[i] = 'a' - 'A' + sentence[i]
		} else if isItWord(sentence[i]) == false {
			checkbool = true
		}
	}
	return string(sentence)
}

func strLenght(str string) int {

	count := 0

	for index := range str {
		count = index + 1
	}

	return count
}

func isItWord(str rune) bool {

	if (str >= 'A' && str <= 'Z') || (str >= 'a' && str <= 'z') || (str >= '0' && str <= '9') {
		return true
	}
	return false
}
