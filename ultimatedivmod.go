package piscine

func UltimateDivMod(a *int, b *int) {
	A := new(int)

	*A = *a
	//*B = *b

	*a = *A / *b

	*b = *A % *b
}
