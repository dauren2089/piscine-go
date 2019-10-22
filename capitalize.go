package piscine

func Capitalize(s string) string {
	sentence := []rune(s)

	for index, value := range sentence {
		if value > 64 && value < 91 {
			sentence[index] = value + 32
		}
	}

	for index, value := range sentence {
		if value > 96 && value < 123 {
			if index == 0 {
				sentence[index] = value - 32
			} else if (sentence[index-1] < 'a' || sentence[index-1] > 'z') && (sentence[index-1] < 'A' || sentence[index-1] > 'Z') {
				if sentence[index-1] < '0' || sentence[index-1] > '9' {
					sentence[index] = value - 32
				}
			}
		}
	}

	return string(sentence)

}
