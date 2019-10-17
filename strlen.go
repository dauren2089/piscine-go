package piscine

func StrLen(str string) int {

	len := []rune(str)

	count := 0

	for index, _ := range len {
		count = index + 1
	}

	return count
}
