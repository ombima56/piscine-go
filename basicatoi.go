package piscine

func BasicAtoi(s string) int {
	var answer int
	for _, char := range s {
		answer = answer*10 + int(char-'0')
	}
	return answer
}
