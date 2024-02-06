package utils

import (
	"fmt"
	"unicode"
)

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

type StringTool struct {
}

func (receiver StringTool) MakeNumStr(number, len int64) string {
	return fmt.Sprintf("%0*d", len, number)
}
