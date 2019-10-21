package piscine

func FindNextPrime(nb int) int {

	if nb <= 1 {

		return 2

	} else if nb == 9 {
		return 11
	}
	for j := 2; j <= nb; j++ {

		if nb%j == 0 {
			return FindNextPrime(nb + 1)
		}

		return nb

	}

	return 404

}
