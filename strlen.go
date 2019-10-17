package piscine

import "strings"

func StrLen(str string) int {

	n := strings.Count(str, "")

	return n - 1
}
