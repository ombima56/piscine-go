package piscine

func ToLower(s string) string {
	answer := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			char = char + 32
		}
		answer += string(char)
	}

	return answer
}
