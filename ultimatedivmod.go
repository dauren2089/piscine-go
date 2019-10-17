package piscine

func UltimateDivMod(a *int, b *int) {

	var c = *a
	var d = *b

	*a = c / d

	*b = c % d
}
