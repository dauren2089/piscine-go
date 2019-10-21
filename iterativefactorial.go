package piscine

func IterativeFactorial(nb int) int {
	var result int
	if nb > 0 {
		result = nb

		for i := 1; i < nb; i++ {
			result = result * (i)
		}
		return result
	}else if nb == 0 {
		return 1
	} else {
		return 0
	}
}