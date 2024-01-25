package piscine

func UltimateDivMod(a *int, b *int) {
	tempc := *a % *b
	*a = *a / *b
	*b = tempc
}
