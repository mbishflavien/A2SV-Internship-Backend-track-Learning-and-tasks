package main

import (
	"strings"
	"unicode"
)

// WordFrequency counts the frequency of each word in a string,
// ignoring punctuation marks and treating words case-insensitively.
func WordFrequency(input string) map[string]int {
	counts := make(map[string]int)

	var sb strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(' ')
		}
	}

	words := strings.Fields(sb.String())
	for _, word := range words {
		counts[word]++
	}

	return counts
}