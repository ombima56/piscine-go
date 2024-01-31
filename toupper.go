package piscine

func ToUpper(s string) string {
	answer := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			char = char - 32
		}
		answer += string(char)
	}

	return answer
}
