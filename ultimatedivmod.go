package piscine

func UltimateDivMod(a *int, b *int) {

	A := new(int)
	B := new(int)

	*A = *a
	*B = *b

	*a = *A / *B

	*b = *A % *B
}
