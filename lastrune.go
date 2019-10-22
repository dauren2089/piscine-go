package piscine

func LastRune(s string) rune {

	lastrune := []rune(s)

	lastindex := 0

	for index := range lastrune {

		lastindex = index
	}

	return lastrune[lastindex]

}
