package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	sentence := []rune(s)
	find := []rune(toFind)
	len := 0
	for index := range find {
		len = index
	}
	str1 := ""
	str2 := ""
	for index := range sentence {
		str1 = string(sentence[index : index+len+1])
		str2 = string(find)
		if str1 == str2 {
			return index
		}

	}
	return -1

}

//j:=[]rune(s)
//l:=[]rune(toFind)
//n:=lent(j)
//k:=lent(l)
//
//for i:=0; i<= n-k; i++ {
//if toFind==s[i:i+k]{
//return (i)
//}
//}
//return -1
