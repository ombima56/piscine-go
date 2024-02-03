package piscine

func SplitWhiteSpaces(s string) []string {
	var answer []string
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' {
			str += string(s[i])
		} else if str != "" {
			answer = append(answer, str)
			str = ""
		}
	}
	if str != "" {
		answer = append(answer, str)
	}
	return answer
}
