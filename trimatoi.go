package piscine

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func TrimAtoi(s string) int {
	var number int
	var sign int = 1
	var found bool

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			number = number*10 + int(ch-'0')
			found = true
		} else if ch == '-' && !found {
			sign = -1
		}
	}
	return number * sign
}
