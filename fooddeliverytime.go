package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggest": {preptime: 12},
	}
	if i, ok := menu[order]; ok {
		return i.preptime
	}
	return 404
}
