package piscine

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	for i := 0; i < len(arg); i++ {
		for j := i + 1; j < len(arg); j++ {
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}
	if len(arg)%2 == 0 {
		num := (len(arg) / 2) - 1
		return arg[num]
	} else {
		num := len(arg) / 2
		return arg[num]
	}
}
