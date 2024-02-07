package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	length := len(a)

	first_arg := true
	second_arg := false

	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			second_arg = false
			break
		}
	}
	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			first_arg = false
			break
		}
	}
	return first_arg || second_arg
}
