package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Phrase palindrome",
			input:    "A man, a plan, a canal: Panama!",
			expected: true,
		},
		{
			name:     "Single word mixed case",
			input:    "RaceCar",
			expected: true,
		},
		{
			name:     "Not a palindrome",
			input:    "Hello World",
			expected: false,
		},
		{
			name:     "Numeric palindrome",
			input:    "12321",
			expected: true,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}