package wordfreq

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			name:  "empty string",
			input: "",
			want:  map[string]int{},
		},
		{
			name:  "single word",
			input: "hello",
			want:  map[string]int{"hello": 1},
		},
		{
			name:  "case insensitive",
			input: "Go go GO",
			want:  map[string]int{"go": 3},
		},
		{
			name:  "trailing and leading punctuation",
			input: "hello, world! hello.",
			want:  map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			name:  "repeated words",
			input: "the quick brown fox jumps over the lazy dog the",
			want: map[string]int{
				"the":   3,
				"quick": 1,
				"brown": 1,
				"fox":   1,
				"jumps": 1,
				"over":  1,
				"lazy":  1,
				"dog":   1,
			},
		},
		{
			name:  "punctuation only",
			input: "... !! hello ?? world",
			want: map[string]int{
				"hello": 1,
				"world": 1,
			},
		},
		{
			name:  "numbers are kept",
			input: "version 2 version 2",
			want: map[string]int{
				"version": 2,
				"2":       2,
			},
		},
		{
			name:  "hyphenated words",
			input: "well-known well-known",
			want: map[string]int{
				"well-known": 2,
			},
		},
		{
			name:  "apostrophes",
			input: "it's it's",
			want: map[string]int{
				"it's": 2,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := WordFrequency(tc.input)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf(
					"WordFrequency(%q)\n  got  %v\n  want %v",
					tc.input,
					got,
					tc.want,
				)
			}
		})
	}
}