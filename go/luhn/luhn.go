package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// Valid determines whether a number is valid according to Luhn formula (https://en.wikipedia.org/wiki/Luhn_algorithm).
func Valid(input string) bool {
	s := strings.ReplaceAll(input, " ", "")

	// the input should contain two or more digits
	re := regexp.MustCompile(`^\d{2,}$`)
	if !re.MatchString(s) {
		return false
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(string(s[len(s)-1-i]))

		if i%2 != 0 {
			n *= 2
			if n >= 10 {
				n -= 9
			}
		}

		sum += n
	}

	return sum%10 == 0
}
