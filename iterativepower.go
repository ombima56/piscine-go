package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	number := 1
	for q := 0; q < power; q++ {
		number *= nb
	}
	return number
}
