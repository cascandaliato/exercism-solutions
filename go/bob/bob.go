// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}

	switch {
	case strings.HasSuffix(remark, "?") && isUpper(remark):
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	case isUpper(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isUpper(remark string) bool {
	containsLetter := false
	for _, r := range remark {
		if unicode.IsLetter(r) {
			containsLetter = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return containsLetter
}
