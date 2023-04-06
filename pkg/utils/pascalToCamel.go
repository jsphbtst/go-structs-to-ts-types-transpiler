package utils

import "unicode"

func PascalToCamelCase(s string) string {
	if len(s) == 0 {
		return s
	}

	// Convert the first character to lowercase
	firstRune := []rune(s)[0]
	if unicode.IsUpper(firstRune) {
		firstRune = unicode.ToLower(firstRune)
	}

	return string(firstRune) + s[1:]
}
