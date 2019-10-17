package piscine

func StrRev(s string) string {

	word := []rune(s)

	var reverse string

	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(word[i])
	}
	return reverse
}
