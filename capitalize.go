package piscine

func Capitalize(s string) string {
	result := ""
	CapitalizeNext := true

	for _,char := range s {
		if IsAlphanumeric(char) {
			if CapitalizeNext {
				result += string(toUpper(char))
				CapitalizeNext = false
			} else {
				result += string(toLower(char))
			}
		} else {
			result += string(char)
			CapitalizeNext = true
		}
	}
	return result
}

func IsAlphanumeric(s rune) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s >= '9')
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