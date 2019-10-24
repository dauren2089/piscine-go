package piscine

func MakeRange(min, max int) []int {

	var size int

	if min < max {
		size = max - min

	} else {
		size = 0
	}

	rangeNbr := make([]int, size)

	if min < max {

		i := 0

		for j := min; j < max; j++ {

			rangeNbr[i] = j

			i++

		}

	}

	return rangeNbr
}
