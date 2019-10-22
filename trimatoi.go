package piscine

func TrimAtoi(s string) int {
	var nbr int

	for index, value := range s {

		if value < '0' || value > '9' {

				if value == '-' {

					nbr = nbr * (-1)
				}
			index ++
		}

		for i := 0; i <= 9; i++ {

			if rune(i+48) == value {

				nbr = nbr*10 + i
			}
		}
	}
	return nbr
}
