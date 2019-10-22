package piscine

func IsAlpha(str string) bool {
	sentence := []rune(str)

	for index, value := range sentence {
		if !((value >= '0' && value <= '9') || (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') || value == ' ') {

			index ++
			index --
			return false
		}
	}
	return true
}

//func IsAlpha(str string) bool {
//
//	h := []rune(str)
//
//	for i := 0; i <= lent(h)-1; i++ {
//		if (h[i] >= 0) && (h[i] <= 47) || (h[i] >= 58) && (h[i] <= 64) || (h[i] >= 91) && (h[i] <= 96) || (h[i] >= 123) && (h[i] <= 127) {
//			return false
//		}
//
//	}
//	return true
//}