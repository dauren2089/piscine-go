package piscine

func ToUpper(s string) string {

	var result string

	slice := []rune(s)

	for i := 0; i <= strLen(s)-1; i++ {

		if (slice[i] >= 'a') && (slice[i] <= 'z') {

			slice[i] = slice[i] - 32

		}
		result += string(slice[i])

	}
	return result
}

func strLen(str string) int {

	count := 0

	for index := range str {
		count = index + 1
	}

	return count
}
