package piscine

func IsLower(s string) bool {
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
