package piscine

func Index(s string, toFind string) int {

	sentence := strLen(s)
	letter := strLen(toFind)

	if toFind == "" {
		return 0
	}

	for i := 0; i <= sentence-letter; i++ {

		if toFind == s[i:i+letter] {

			return (i)
		}
	}

	return -1

}

func strLen(str string) int {

	count := 0

	for index := range str {
		count = index + 1
	}

	return count
}
