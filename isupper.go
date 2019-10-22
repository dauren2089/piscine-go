package piscine

func IsUpper(str string) bool {
	sentence := []rune(str)

	for index, value := range sentence {

		if !((value >= 'A' && value <= 'Z') || value == ' ') {
			index = index + 0
			return false
		}

	}
	return true
}


//func IsUpper(str string) bool {
//	h := []rune(str)
//
//	for i := 0; i <= lent(h)-1; i++ {
//		if (h[i] >= 0) && (h[i] <= 64) || (h[i] >= 91) && (h[i] <= 127){
//			return false
//		}
//
//	}
//	return true
//}