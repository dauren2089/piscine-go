package piscine

func StrLen(str string) int {
	slice := []rune(str)

	count := 0

	for index, _ := range slice {
		count = index + 1
	}

	return count
}
