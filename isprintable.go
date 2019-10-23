package piscine

func IsPrintable(str string) bool {
	sentence := []rune(str)

	for i := 0; i <= strLen(str)-1; i++ {
		if (sentence[i] >= 0) && (sentence[i] <= 31) {
			return false
		}

	}
	return true

}
