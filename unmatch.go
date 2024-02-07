package piscine

func Unmatch(a []int) int {
	var count int
	for _, ch := range a {
		for _, num := range a {
			if ch == num {
				count++
			}
		}
		if count%2 != 0 {
			return ch
		}
	}
	return -1
}
