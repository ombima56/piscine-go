package piscine

func Capitalize(s string) string {
	result := ""
	capitalizeNext := true

	for _, char := range s {
		if isAlphanumeric(char) {
			if capitalizeNext {
				result += string(toUpper(char))
				capitalizeNext = false
			} else {
				result += string(toLower(char))
			}
		} else {
			result += string(char)
			capitalizeNext = true
		}
	}
	return result
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func toLower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + 'a' - 'A'
	}
	return char
}

func toUpper(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - 'a' + 'A'
	}
	return char
}
