package main

import (
	"unicode"
)

// IsPalindrome checks whether a string is a palindrome, ignoring spaces,
// punctuation, and capitalization.
func IsPalindrome(input string) bool {
	var cleaned []rune
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	left, right := 0, len(cleaned)-1
	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}