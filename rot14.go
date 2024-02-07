package piscine

func Rot14(s string) string {
	result := []rune(s)
	for i, ch := range result {
		if ch >= 'a' && ch <= 'z' {
			result[i] = 'a' + (ch-'a'+14)%26
		} else if ch >= 'A' && ch <= 'Z' {
			result[i] = 'A' + (ch-'A'+14)%26
		}
	}
	return string(result)
}
