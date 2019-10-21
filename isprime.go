package piscine

func IsPrime(nb int) bool {

	if nb <= 1 {

		return false

	} else if nb == 2 {
		return true
	}

	//for i := 1; i <= nb; i++ {

	for j := 2; j <= nb; j++ {

		if nb%j == 0 {

			return false
		}

		return true

	}
	//}

	return false

}
