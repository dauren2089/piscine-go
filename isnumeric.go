package piscine

func IsNumeric(str string) bool {
	sentence := []rune(str)

	for index, value := range sentence {
		if !((value >= '0' && value <= '9') || value == ' ') {

			index = index + 0

			return false
		}
	}
	return true
}


//func IsNumeric(str string) bool {
//	h := []rune(str)
//
//	for i := 0; i <= lent(h)-1; i++ {
//		if (h[i] >= 0) && (h[i] <= 47) || (h[i] >= 58) && (h[i] <= 127){
//			return false
//		}
//
//	}
//	return true
//}