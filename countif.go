package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int
	for _, ch := range tab {
		if f(ch) == true {
			count++
		}
	}
	return count
}
