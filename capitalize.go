package piscine

func IsAlphanumeric(s rune) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s >= '9')
}

func Capitalize(s string) string {
	var answer string
	var CapitalizeNext = true

	for _, ch := range s {
		if IsAlphanumeric(ch) {
			if CapitalizeNext {
				answer += string(toUpper(ch))
				CapitalizeNext = false
			} else {
				answer += string(toLower(ch))
				CapitalizeNext = false
			}
		} else {
			answer += string(ch)
			CapitalizeNext = true
		}
	}
	return answer
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}
