// Package pangram exposes functions to analyze pangrams.
package pangram

import (
	"strings"
)

// IsPangram determines whether a text is a pangram.
func IsPangram(text string) bool {
	letters := map[rune]bool{}

	for _, ch := range strings.ToLower(text) {
		letters[ch] = true
	}

	for ch := 'a'; ch <= 'z'; ch++ {
		if !letters[ch] {
			return false
		}
	}

	return true
}
