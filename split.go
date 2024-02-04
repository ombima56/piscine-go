package piscine

func Split(s, sep string) []string {
	var answer []string
	start := 0
	for i := 0; i < len(s)-len(sep)+1; i++ {
		if s[i:i+len(sep)] == sep {
			answer = append(answer, s[start:i])
			start = i + len(sep)
		}
	}
	answer = append(answer, s[start:])
	return answer
}
