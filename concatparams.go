package piscine

func ConcatParams(args []string) string {
	var answer string
	for i, arg := range args {
		answer += arg
		if i < len(args)-1 {
			answer += "\n"
		}
	}
	return answer
}
