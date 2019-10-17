package piscine

func StrRev(s string) string {

	word := []rune(s)

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return string(word)
}
