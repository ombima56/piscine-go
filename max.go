package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	max := a[0]
	for _, q := range a[1:] {
		if q > max {
			max = q
		}
	}
	return max
}
