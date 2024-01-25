package piscine

func StrRev(s string) string {
	var answer string
	for q := len(s) - 1; q >= 0; q-- {
		answer += string(s[q])
	}
	return answer
}
