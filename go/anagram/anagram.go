package anagram

import "strings"

type lettersCount map[rune]int

func (lc lettersCount) le(o lettersCount) bool {
	for r, n := range lc {
		if n > o[r] {
			return false
		}
	}
	return true
}

// Detect returns the sublist of candidates that are anagrams of the given text (case-insensitive).
func Detect(text string, candidates []string) []string {
	var anagrams []string

	for _, candidate := range candidates {
		if isAnagram(strings.ToLower(text), strings.ToLower(candidate)) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

func isAnagram(a, b string) bool {
	lcA := getLettersCount(a)
	lcB := getLettersCount(b)

	return a != b && lcA.le(lcB) && lcB.le(lcA)
}

func getLettersCount(s string) lettersCount {
	lc := make(map[rune]int)
	for _, ch := range s {
		lc[ch]++
	}
	return lc
}
