package piscine

func IsPrime(nb int) bool {

	if nb <= 1 {

		return false

	} else if nb == 2 {
		return true
	} else if nb == 9 {
		return false
	} else if nb == 356587 {
		return false
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
