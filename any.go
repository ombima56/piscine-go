package piscine

func Any(f func(string) bool, a []string) bool {
	for _, ch := range a {
		if f(ch) {
			return true
		}
	}
	return false
}
