package piscine

func StrLen(str string) int {

	len1 := []rune(str)

	count := 0

	for index, _ := range len1 {
		count = index + 1
	}

	return count
}
