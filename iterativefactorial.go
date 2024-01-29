package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	var number int = 1
	for q := 1; q <= nb; q++ {
		if number < 0 {
			return 0
		}
		number *= q
	}
	return number
}
