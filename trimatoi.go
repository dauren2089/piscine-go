package piscine

func TrimAtoi(s string) int {
	nbr := 0
	sign := true
	start := false

	for _, value := range s {
		if (value < '0') || (value > '9') {
			if start == false {
				if value == '-' {
					sign = false
				} else if value == '+' {
					sign = true
				}
			}

		} else {
			start = true
			for n := 0; n <= 9; n++ {
				if rune(n+48) == value {
					nbr = nbr*10 + n
					break
				}
			}
		}
	}
	if sign == false {
		nbr *= -1
	}
	return nbr
}
