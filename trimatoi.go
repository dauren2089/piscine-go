package piscine

func TrimAtoi(s string) int {
	nbr := 0
	sign := true

	for _, value := range s {
		if (value < '0') || (value > '9') {
			if value == '-' {
				sign = false
			} else if value == '+' {
				sign = true
			}

		} else {
			for n := 0; n <= 9; n++ {
				if rune(n+48) == value {
					nbr = nbr*10 + n
					break
				}
			}
		}
	}
	if sign == false {
		nbr = nbr * -1
	}
	return nbr
}
