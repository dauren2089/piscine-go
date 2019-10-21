package piscine

func IterativePower(nb int, power int) int {

	//var result int

	if nb != 0 && power > 0 {

		result := 1

		for i := 1; i <= power; i++ {

			result = result * nb

		}

		return result

	} else if nb == 0 || power == 0 {

		return 1
	}

	return 0

}
