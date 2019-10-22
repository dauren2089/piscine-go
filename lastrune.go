package piscine

func LastRune(s string) rune {

	lastrune := []rune(s)

	lastindex := 0

	for index := range lastrune {
		lastindex = index
	}

	return lastrune[lastindex]

}

//runes := []rune(s)
//len := 0
//for index := range runes {
//len = index
//}
//
//return runes[len]
//}