// Package romannumerals provides utility functions for Roman numerals.
package romannumerals

import (
	"errors"
	"strconv"
	"strings"
)

var letters []byte = []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}

// ToRomanNumeral converts any positive number up to 3,999 into a Roman numeral.
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3999 {
		return "", errors.New("cannot convert numbers <= 0 or > 3999")
	}

	var numeral strings.Builder
	digits := []byte(strconv.Itoa(arabic))
	for i, j := 0, 2*(len(digits)-1); i < len(digits); i, j = i+1, j-2 {
		switch digits[i] {
		case '0':
		case '1':
			numeral.WriteByte(letters[j])
		case '2':
			numeral.Write([]byte{letters[j], letters[j]})
		case '3':
			numeral.Write([]byte{letters[j], letters[j], letters[j]})
		case '4':
			numeral.Write([]byte{letters[j], letters[j+1]})
		case '5':
			numeral.WriteByte(letters[j+1])
		case '6':
			numeral.Write([]byte{letters[j+1], letters[j]})
		case '7':
			numeral.Write([]byte{letters[j+1], letters[j], letters[j]})
		case '8':
			numeral.Write([]byte{letters[j+1], letters[j], letters[j], letters[j]})
		case '9':
			numeral.Write([]byte{letters[j], letters[j+2]})
		}
	}
	return numeral.String(), nil
}
