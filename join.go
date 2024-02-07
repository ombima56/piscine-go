package piscine

func Join(strs []string, sep string) string {
	var answer string
	for i, ch := range strs {
		answer += ch
		if i < len(strs)-1 {
			answer += sep
		}
	}
	return answer
}
