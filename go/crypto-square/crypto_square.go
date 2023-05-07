package cryptosquare

import (
	"math"
	"strings"
)

// Encode encodes a plaintext using the square code technique.
func Encode(plaintext string) string {
	plaintext = normalize(plaintext)

	numColumns := int(math.Ceil(math.Sqrt(float64(len(plaintext)))))
	columns := make([]string, 0, numColumns)

	var column strings.Builder
	for c := 0; c < numColumns; c++ {
		column.Reset()
		for r := 0; r < numColumns; r++ {
			if position := r*numColumns + c; position < len(plaintext) {
				column.WriteByte(plaintext[position])
			} else if len(plaintext) > r*numColumns {
				// it's a square but not a perfect square, i.e. the last row is incomplete but not empty
				// in this case some columns will need padding
				column.WriteByte(' ')
			}
		}
		columns = append(columns, column.String())
	}

	return strings.Join(columns, " ")
}

func normalize(text string) string {
	var sb strings.Builder
	for _, ch := range text {
		if (ch < 'a' || ch > 'z') &&
			(ch < 'A' || ch > 'Z') &&
			(ch < '0' || ch > '9') {
			continue
		}
		sb.WriteRune(ch)
	}
	return strings.ToLower(sb.String())
}
