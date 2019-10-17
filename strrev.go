package piscine

func StrRev(s string) string {

	var word string

	for i := len(s) - 1; i >= 0; i-- {
		word += string(s[i])
	}
	return word
}
