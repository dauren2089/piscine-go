package piscine

func IsPrime(nb int) bool {

	for i := 1; i <= nb; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				return false
			}
			return true
		}
	}
	return false
}
