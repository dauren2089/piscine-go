package piscine

func ToLower(s string) string {
	sentence := []rune(s)
	for index, value := range sentence {
		if value > 64 && value < 91 {
			sentence[index] = value + 32
		}
	}
	return string(sentence)
}

//func ToLower(s string) string {
//	h := []rune(s)
//	result := ""
//	for i := 0; i <= lent(h)-1; i++ {
//
//		if (h[i] >= 'A') && (h[i] <= 'Z') {
//
//			h[i] = h[i] + 32
//
//		}
//		result += string(h[i])
//
//	}
//	return result
//}
