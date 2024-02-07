package piscine

func JumpOver(str string) string {
	if str == "" {
		return "\n"
	}
	var answer string
	for i, ch := range str {
		if (i+1)%3 == 0 {
			answer += string(ch)
		}
	}
	if answer == "" {
		return "\n"
	}
	return answer + "\n"
}
