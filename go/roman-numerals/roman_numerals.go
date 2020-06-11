// Package romannumerals provides utility functions for Roman numberals.
package romannumerals

import (
	"errors"
	"strconv"
)

var letters []rune = []rune{'I', 'V', 'X', 'L', 'C', 'D', 'M'}

// ToRomanNumeral converts any positive number up to 3,000 into a Roman numeral.
func ToRomanNumeral(arabic int) (string, error) {
	var roman []rune
	if arabic <= 0 || arabic > 3000 {
		return string(roman), errors.New("cannot convert numbers <= 0 or > 3000")
	}

	a := []rune(strconv.Itoa(arabic))
	for i, j := len(a)-1, 0; i >= 0; i, j = i-1, j+2 {
		if a[i] == '0' {
			continue
		}

		var r []rune
		switch a[i] {
		case '1':
			r = []rune{letters[j]}
		case '2':
			r = []rune{letters[j], letters[j]}
		case '3':
			r = []rune{letters[j], letters[j], letters[j]}
		case '4':
			r = []rune{letters[j], letters[j+1]}
		case '5':
			r = []rune{letters[j+1]}
		case '6':
			r = []rune{letters[j+1], letters[j]}
		case '7':
			r = []rune{letters[j+1], letters[j], letters[j]}
		case '8':
			r = []rune{letters[j+1], letters[j], letters[j], letters[j]}
		case '9':
			r = []rune{letters[j], letters[j+2]}
		}
		roman = append(r, roman...)
	}

	return string(roman), nil
}
