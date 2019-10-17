package piscine

func Swap(a *int, b *int) {

	var c = *a

	*a = *b
	*b = c

}
