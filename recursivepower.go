package piscine

func RecursivePower(nb int, power int) int {

	if nb > 0 {
		if power > 1 {
			return nb * RecursivePower(nb, power-1)
		}
		return nb
	} else if nb == 1{
		return 1
	}
	return 0
}
