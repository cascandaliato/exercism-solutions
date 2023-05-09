package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode encodes a plaintext using the square code technique.
func Encode(plaintext string) string {
	plaintext = normalize(plaintext)
	numColumns := int(math.Ceil(math.Sqrt(float64(len(plaintext)))))
	return squareAndTranspose(plaintext, numColumns)
}

func squareAndTranspose(text string, nCols int) string {
	var encoded strings.Builder

	for c := 0; c < nCols; c++ {
		if c > 0 {
			encoded.WriteByte(' ')
		}

		for r := 0; r < nCols; r++ {
			rStartPos := r * nCols

			if position := rStartPos + c; position < len(text) {
				encoded.WriteByte(text[position])
			} else if len(text) > rStartPos {
				// it's a square but not a perfect square
				// i.e. the last row is incomplete but not empty
				// in this case some columns will need padding
				encoded.WriteByte(' ')
			}
		}
	}

	return encoded.String()
}

func normalize(text string) string {
	var sb strings.Builder
	for _, ch := range text {
		if !isValid(ch) {
			continue
		}
		sb.WriteRune(unicode.ToLower(ch))
	}
	return sb.String()
}

func isValid(ch rune) bool {
	return ('a' <= ch && ch <= 'z') ||
		('A' <= ch && ch <= 'Z') ||
		('0' <= ch && ch <= '9')
}
