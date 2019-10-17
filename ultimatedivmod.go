package piscine

func DivMod(a *int, b *int) {

	A := new(int)
	//B := new(int)

	*A = *a
	//B = *b

	*a = *A / *b

	*b = *A % *b
}
