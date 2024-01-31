package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
