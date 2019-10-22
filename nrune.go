package piscine

func NRune(s string, n int) rune {

	firstrune := []rune(s)

	return firstrune[n-1]

}

	//len := 0
	//for index := range runes {
	//	len = index
	//}
	//
	//if n-1 >= 0 && n-1 <= len {
	//	return runes[n-1]
	//}
	//	return 0
	//}