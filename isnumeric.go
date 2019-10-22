package piscine

func IsNumeric(str string) bool {

	sentence := []rune(str)

	for index, value := range sentence {
		if !((value >= '0' && value <= '9') || value == ' ') {

			index = index + 0

			return false
		}

	}
	return true
}
