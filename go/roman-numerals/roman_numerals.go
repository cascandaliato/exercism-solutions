// Package romannumerals provides utility functions for Roman numerals.
package romannumerals

import (
	"errors"
	"strconv"
)

var letters []rune = []rune{'I', 'V', 'X', 'L', 'C', 'D', 'M'}

// ToRomanNumeral converts any positive number up to 3,999 into a Roman numeral.
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3999 {
		return "", errors.New("cannot convert numbers <= 0 or > 3999")
	}

	roman := make([]rune, 0)
	digits := []rune(strconv.Itoa(arabic))
	for i, j := 0, 2*(len(digits)-1); i < len(digits); i, j = i+1, j-2 {
		var runes []rune
		switch digits[i] {
		case '0':
		case '1':
			runes = []rune{letters[j]}
		case '2':
			runes = []rune{letters[j], letters[j]}
		case '3':
			runes = []rune{letters[j], letters[j], letters[j]}
		case '4':
			runes = []rune{letters[j], letters[j+1]}
		case '5':
			runes = []rune{letters[j+1]}
		case '6':
			runes = []rune{letters[j+1], letters[j]}
		case '7':
			runes = []rune{letters[j+1], letters[j], letters[j]}
		case '8':
			runes = []rune{letters[j+1], letters[j], letters[j], letters[j]}
		case '9':
			runes = []rune{letters[j], letters[j+2]}
		}
		roman = append(roman, runes...)
	}

	return string(roman), nil
}
