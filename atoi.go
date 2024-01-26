package piscine

func Atoi(s string) int {
	var num int
	var sign int = 1
	for q, char := range s {
		if char == '-' && q == 0 {
			sign = -1
		} else if char == '+' && q == 0 {
			sign = 1
		} else if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return sign * num
}