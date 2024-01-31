package piscine

func IsUpper(s string) bool {
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
