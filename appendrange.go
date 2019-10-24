package piscine

func AppendRange(min, max int) []int {

	var rangeNbr []int

	if min < max {

		for i := min - 1; i < max-1; i++ {

			rangeNbr = append(rangeNbr, i+1)

		}
		return rangeNbr
	}

	return rangeNbr
}
