package piscine

func BasicAtoi2(s string) int {
	v_number := 0
	var num int
	result := []rune(s)
	for _, word := range result {
		if word < '0' || word > '9' {
			return 0
		}
		for i := '0'; i < word; i++ {
			num++
		}
		v_number = v_number*10 + num
		num = 0
	}
	return v_number
}
