package piscine

func IsPrintable(str string) bool {
	sentence := []rune(str)
	for index, value := range sentence {
		if !(value >= ' ') {
			index = index + 0

			return false
		}
	}
	return true
}


//func IsPrintable(str string) bool {
//	h := []rune(str)
//
//	for i := 0; i <= len(h)-1; i++ {
//		if (h[i] >= 0) && (h[i] <= 31) {
//			return false
//		}
//
//	}
//	return true
//}
