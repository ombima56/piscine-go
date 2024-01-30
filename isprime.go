package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%2 == 0 {
			return false
		}
	}
	return true
}
