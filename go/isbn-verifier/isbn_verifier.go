package isbn

import (
	"regexp"
	"strings"
)

// IsValidISBN verifies whether the given string is a valid ISBN.
func IsValidISBN(s string) bool {
	validISBN := regexp.MustCompile(`^\d{9}[X\d]$`)

	s = strings.ReplaceAll(s, "-", "")

	if !validISBN.MatchString(s) {
		return false
	}

	var checksum int
	for i, r := range s {
		switch r {
		case 'X':
			checksum += 10
		default:
			checksum += int(r-'0') * (10 - i)
		}
	}

	return checksum%11 == 0
}
