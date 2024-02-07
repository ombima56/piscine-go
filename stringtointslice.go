package piscine

func StringToIntSlice(str string) []int {
	var answer []int
	for _, ch := range str {
		answer = append(answer, int(ch))
	}
	return answer
}
