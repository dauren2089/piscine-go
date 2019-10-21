package piscine

func RecursivePower(nb int, power int) int {

	if nb != 0 && power > 0 {

		if power > 1 {

			return nb * RecursivePower(nb, power-1)

		}

		return nb

	} else if power == 0 {

		return 1

	} else if nb == 0 {

		return 0
	}

	return 0

}
