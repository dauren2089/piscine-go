package piscine

func ToLower(s string) string {

	var result string

	slice := []rune(s)

	for i := 0; i <= strLen(s)-1; i++ {

		if (slice[i] >= 'A') && (slice[i] <= 'Z') {

			slice[i] = slice[i] + 32

		}
		result += string(slice[i])

	}
	return result
}

// func strLen(str string) int {

// 	count := 0

// 	for index := range str {
// 		count = index + 1
// 	}

// 	return count
// }
