package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, ch := range a {
		result[i] = f(ch)
	}
	return result
}
