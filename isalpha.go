package piscine

func IsAlpha(str string) bool {

	sentence := []rune(str)

	for i := 0; i <= strLenghts(str)-1; i++ {
		if (sentence[i] >= 0) && (sentence[i] <= 47) || (sentence[i] >= 58) && (sentence[i] <= 64) || (sentence[i] >= 91) && (sentence[i] <= 96) || (sentence[i] >= 123) && (sentence[i] <= 127) {
			return false
		}

	}
	return true
}

func strLenghts(str string) int {
	slice := []rune(str)

	count := 0

	for index := range slice {
		count = index + 1
	}

	return count
}
