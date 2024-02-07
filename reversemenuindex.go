package piscine

func ReverseMenuIndex(menu []string) []string {
	reversed := make([]string, len(menu))
	for i := len(menu) - 1; i >= 0; i-- {
		reversed[len(menu)-1-i] = menu[i]
	}
	return reversed
}
