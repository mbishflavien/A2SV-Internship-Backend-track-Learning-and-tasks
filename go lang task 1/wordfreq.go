package wordfreq

import (
	"strings"
	"unicode"
)

func WordFrequency(s string) map[string]int {
	frequencies := make(map[string]int)

	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) &&
			!unicode.IsDigit(r) &&
			r != '-' &&
			r != '\''
		})

	for _, word := range words {
		word = strings.ToLower(word)
		frequencies[word]++
	}

	return frequencies
}