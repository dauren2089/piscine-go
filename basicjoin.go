package piscine

func BasicJoin(strs []string) string {
	word := ""
	for value := range strs {
		word += strs[value]
	}
	return word
}

//func BasicJoin(strs []string) string {
//	var ret string
//	for _, str := range strs {
//		ret += str
//	}
//	return ret
//}

