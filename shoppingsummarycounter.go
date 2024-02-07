package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	list := make(map[string]int)
	var itemsName string
	for _, ch := range str {
		if ch == 32 {
			list[itemsName] += 1
			itemsName = ""
		} else if ch != 32 {
			itemsName += string(byte(ch))
		}
	}
	list[itemsName] += 1
	return list
}
