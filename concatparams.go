package piscine

func ConcatParams(args []string) string {

	var concatstr string

	size := stringLen(args) - 1

	newline := "\n"

	arguments := make([]string, size)

	arguments = args

	for index, word := range arguments {

		concatstr = concatstr + word

		if index < size {

			concatstr = concatstr + newline
		}
	}

	return concatstr
}

func stringLen(word []string) int {

	count := 0

	for range word {

		count++
	}

	return count
}
