package piscine

//import "fmt"

func AlphaCount(str string) int {
	count := 0

	for i:= 'a'; i <= 'z'; i++ {
		for _, value := range str {
			if value == i {
				count++
			}
		}
	}

	for j:= 'A'; j <= 'Z'; j++ {

		for _, value := range str {
			if value == j {
				count++
			}
		}

	}
	//fmt.Println(str)
	return count
}
