package piscine

func IsLower(str string) bool {
	sentence := []rune(str)

	for index, value := range sentence {

		if !((value >= 'a' && value <= 'z') || value == ' ') {

			index = index + 0

			return false
		}

	}
	return true
}
