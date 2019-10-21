package piscine

func IterativeFactorial(nb int) int {
	var result1 int
	if nb > 0 {
		result1 = nb

		for i := 1; i < nb; i++ {
			result1 = result1 * (i)
		}
		return result1

	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
}
