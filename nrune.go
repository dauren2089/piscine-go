package piscine

func NRune(s string, n int) rune {

	nrune := []rune(s)
	count := 0

	for value := range nrune {
		count = value
	}

	if n-1 >= 0 && n-1 <= count {

		return nrune[n-1]

	}
	return 0

}
