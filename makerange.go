package piscine

import "fmt"

func MakeRange(min, max int) []int {

	var size int

	if min < max {
		size = max - min
		fmt.Println("size: ", size)

	} else {
		size = 0
		fmt.Println("size: ", size)
	}

	rangeNbr := make([]int, size)

	//fmt.Println("RangeNumber before: ", rangeNbr)

	if min < max {

		i := 0

		for j := min; j < max; j++ {

			fmt.Println(i)

			rangeNbr[i] = j

			i++

		}

		//fmt.Println("RangeNumber after: ", rangeNbr)
	}

	return rangeNbr
}
