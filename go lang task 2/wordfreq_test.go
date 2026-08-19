package main

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "Basic sentence with punctuation",
			input: "Hello, world! Hello Go, world.",
			expected: map[string]int{
				"hello": 2,
				"world": 2,
				"go":    1,
			},
		},
		{
			name:  "Mixed case and extra spaces",
			input: "  Go go  GO!  ",
			expected: map[string]int{
				"go": 3,
			},
		},
		{
			name:     "Empty input",
			input:    "",
			expected: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordFrequency(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("WordFrequency(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}