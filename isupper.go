package piscine

func IsUpper(str string) bool {
	sentence := []rune(str)

	for index, value := range sentence {

		if !((value >= 'A' && value <= 'Z') || value == ' ') {
			index = index + 0
			return false
		}

	}
	return true
}
