package piscine

func Compare(a, b string) int {
	firstrune := []rune(a)
	secondrune := []rune(b)
	//compr := 0
	var compr int
	firstl := 0
	secondl := 0

	for index := range secondrune {
		secondl = index
	}

	for index := range firstrune {
		firstl = index
	}

	for index := range firstrune {

		if index > secondl {
			break
		}

		if firstrune[index] > secondrune[index] {
			return 1

		} else if firstrune[index] < secondrune[index] {
			return -1
		}
	}

	if firstl > secondl {
		compr = 1

	} else if firstl < secondl {
		compr = -1
	}

	return compr
}


//func Compare(a, b string) int {
//
//	if a == b {
//		return 0
//	} else if a > b {
//		return +1
//
//
//	}
//	return -1
//}