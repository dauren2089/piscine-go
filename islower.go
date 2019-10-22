package piscine

func IsLower(str string) bool {
	sentence := []rune(str)

	for index, value := range sentence {

		if !((value >= 'a' && value <= 'z') || value == ' ') {

			index = index + 0

			return false
		}

	}
	return true
}

//func IsLower(str string) bool {
//	h := []rune(str)
//
//	for i := 0; i <= lent(h)-1; i++ {
//		if (h[i] >= 0) && (h[i] <= 96) || (h[i] >= 123) && (h[i] <= 127){
//			return false
//		}
//
//	}
//	return true
//}
