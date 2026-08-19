package main

import (
	"fmt"
)

func main() {
	//Word Frequency Count manual test
	text := "Hello, world! Hello Go, world."
	fmt.Println("--- Word Frequency Count ---")
	fmt.Printf("Input: %q\n", text)
	frequencies := WordFrequency(text)
	for word, count := range frequencies {
		fmt.Printf("  %s: %d\n", word, count)
	}

	//palindrome check manual test
	fmt.Println("\n--- Palindrome Check ---")
	samples := []string{
		"A man, a plan, a canal: Panama!",
		"RaceCar",
		"Hello World",
	}

	for _, sample := range samples {
		fmt.Printf("  %q -> Palindrome: %t\n", sample, IsPalindrome(sample))
	}
}