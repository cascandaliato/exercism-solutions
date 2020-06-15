// Package isogram provides functions to detect isograms.
// An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter, however spaces and hyphens are allowed to appear multiple times.
// Examples of isograms:
// - lumberjacks
// - background
// - downstream
// - six-year-old
// The word 'isograms', however, is not an isogram, because the 's' repeats.
package isogram

import "strings"

// IsIsogram determines whether a word is an isogram.
func IsIsogram(s string) bool {
	seen := map[rune]bool{}
	for _, r := range strings.ToLower(s) {
		if r != ' ' && r != '-' {
			if _, ok := seen[r]; ok {
				return false
			}
			seen[r] = true
		}
	}
	return true
}
