package main

func Compare(a, b string) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return -1
		}

	}
	if len(a) == len(b) {
		return 0
	} else if len(a) < len(b) {
		return -1
	}
	return 1
}
