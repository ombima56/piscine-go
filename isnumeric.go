package piscine

func IsNumeric(s string) bool {
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
