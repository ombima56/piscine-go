package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 || nb > 63 {
		return 0
	}
	number := 1
	for q := 0; q < power; q++ {
		number = number * nb
	}
	return number
}
