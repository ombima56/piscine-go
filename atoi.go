package piscine

func Atoi(s string) int {
	var number int
	var sign int = 1
	for i, x := range s {
		if x == '-' && i == 0 {
			sign = -1
		} else if x == '+' && i == 0 {
			sign = 1
		} else if x >= '0' && x <= '9' {
			number = number*10 + int(x-'0')
		} else {
			return 0
		}
	}
	return sign * number
}
