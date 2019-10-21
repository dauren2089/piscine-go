package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if nb > 0 {
		for i := 1; i <= power; i++ {
			result = result * nb
		}
		return result
	}
	return 0

}
