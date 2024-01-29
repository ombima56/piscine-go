package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb >= 12 {
		return 0
	}

	var number int = 1
	for q := 1; q <= nb; q++ {
		number *= q
	}
	return number
}
