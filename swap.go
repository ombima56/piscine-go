package piscine

func Swap(a *int, b *int) {
	tempc := *a
	*a = *b
	*b = tempc
}
