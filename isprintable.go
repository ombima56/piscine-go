package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		if char >= 32 && char <= 127 {
			continue
		} else {
			return false
		}
	}
	return true
}
