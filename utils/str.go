package utils

import "unicode"

func ContainsLetterAndDigit(str string) bool {
	hasLetter := false
	hasDigit := false

	for _, char := range str {
		if unicode.IsLetter(char) {
			hasLetter = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}

		if hasLetter && hasDigit {
			return true
		}
	}

	return false
}
