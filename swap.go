package piscine

func Swap(a *int, b *int) {

	A := new(int)

	*A = *a
	*a = *b
	*b = *A
}
