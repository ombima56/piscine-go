package piscine

func ActiveBits(n int) int {
	var answer int
	for n > 0 {
		if n%2 == 1 {
			answer++
		}
		n = n >> 1
	}
	return answer
}
