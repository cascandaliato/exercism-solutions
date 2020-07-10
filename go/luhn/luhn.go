package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

// Valid determines whether a number is valid according to Luhn formula (https://en.wikipedia.org/wiki/Luhn_algorithm).
func Valid(input string) bool {
	inputClean := strings.ReplaceAll(input, " ", "")

	// the input should contain two or more digits
	re := regexp.MustCompile(`^\d{2,}$`)
	if !re.MatchString(inputClean) {
		return false
	}

	sum := 0
	parity := len(inputClean)%2 == 0
	for _, c := range inputClean {
		n, _ := strconv.Atoi(string(c))

		if parity {
			n *= 2
			if n >= 10 {
				n -= 9
			}
		}

		sum += n
		parity = !parity
	}

	return sum%10 == 0
}
