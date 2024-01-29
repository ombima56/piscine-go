package piscine

func IterativeFactorial(nb int) int {
	var number int = 1
	for q := 1; q <= nb; q++ {
		number *= q
	}
	return number
}
