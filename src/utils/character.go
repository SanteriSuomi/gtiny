package utils

import (
	"unicode"

	"github.com/SanteriSuomi/gtiny/src/entities"
)

// IsDigit checks if the given string is a single digit. Only regular digits (0-9) are considered.
//
// Parameters:
// - str: The string to be checked.
//
// Returns:
// - bool: True if the string is a single digit, false otherwise.
func IsDigit(str string) bool {
	if len(str) == 1 {
		c := []rune(str)[0]
		return c >= '0' && c <= '9'
	}
	return false
}

// IsAlpha checks if the given string contains only alphabetic characters.
//
// Parameters:
// - str: The string to be checked.
//
// Returns:
// - bool: True if the string contains only alphabetic characters, false otherwise.
func IsLetter(str string) bool {
	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// IsAlphaNumeric checks if the given string contains only alphabetic characters or is a single digit.
//
// Parameters:
// - str: The string to be checked.
// Return type: bool
func IsAlphaNumeric(str string) bool {
	return str != entities.EOF && IsLetter(str) || IsDigit(str)
}
