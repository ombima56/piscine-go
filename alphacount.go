package piscine

func AlphaCount(s string) int {
	var count int
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			count++
		}
	}
	return count
}
