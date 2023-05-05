package anagram

import (
	"strings"
)

// Detect returns the sublist of candidates that are anagrams of the given text (case-insensitive).
func Detect(text string, candidates []string) []string {
	anagrams := make([]string, 0)
	textLC := strings.ToLower(text)
	counter := countLetters(textLC)

	for _, candidate := range candidates {
		candidateLC := strings.ToLower(candidate)
		if textLC != candidateLC && areCountersEqual(counter, countLetters(candidateLC)) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

func countLetters(s string) map[rune]int {
	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}
	return count
}

func areCountersEqual(a map[rune]int, b map[rune]int) bool {
	for keyA, valueA := range a {
		if valueB, ok := b[keyA]; !ok || valueA != valueB {
			return false
		}
	}

	for keyB, valueB := range b {
		if valueA, ok := a[keyB]; !ok || valueA != valueB {
			return false
		}
	}

	return true
}
