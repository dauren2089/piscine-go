package piscine

func StrRev(s string) string {

	slice := []rune(s)

	count := 0

	for index := range slice {
		count = index + 1
	}
	for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return string(slice)
}
