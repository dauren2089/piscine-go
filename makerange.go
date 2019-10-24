package piscine

func MakeRange(min, max int) []int {

	var size int

	if min < max {

		size = max - min

		rangeNbr := make([]int, size)

		i := 0

		for j := min; j < max; j++ {

			rangeNbr[i] = j

			i++

		}

		return rangeNbr

	}

	var rangeNbr []int = nil

	return rangeNbr
}
