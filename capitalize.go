package piscine

func Capitalize(s string) string {
	sentence := []rune(s)

	for index, value := range sentence {
		if value > 64 && value < 91 {
			sentence[index] = value + 32
		}
	}

	for index, value := range sentence {
		if value > 96 && value < 123 {
			if index == 0 {
				sentence[index] = value - 32
			} else if (sentence[index-1] < 'a' || sentence[index-1] > 'z') && (sentence[index-1] < 'A' || sentence[index-1] > 'Z') {
				if sentence[index-1] < '0' || sentence[index-1] > '9' {
					sentence[index] = value - 32
				}
			}
		}
	}
	return string(sentence)
}

//func prim(a rune) bool {
//
//	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
//		return true
//	}
//	return false
//}
//
//func Capitalize(s string) string {
//	ar := []rune(s)
//
//	letra := true
//
//	for i := 0; i < len(s); i++ {
//		if prim(ar[i]) == true && letra {
//			if ar[i] >= 'a' && ar[i] <= 'z' {
//				ar[i] = 'A' - 'a' + ar[i]
//
//			}
//			letra = false
//		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
//			ar[i] = 'a' - 'A' + ar[i]
//		} else if prim(ar[i]) == false {
//			letra = true
//		}
//	}
//	return string(ar)
//}