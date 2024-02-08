package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	r := ""
	count := 0
	for i, v := range str {
		if v != ' ' && count != 5 {
			r += string(v)
			count++
		} else if count == 5 {
			r += " "
			count = 0
		}
		if i == len(str)-1 && len(r) > 0 && r[len(r)-1] == ' ' {
			r = r[:len(r)-1]
		}
	}
	r += "\n"
	return r
}
