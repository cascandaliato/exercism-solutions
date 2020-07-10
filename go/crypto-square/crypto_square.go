package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// Encode encodes a plaintext using the square code technique.
func Encode(plaintext string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	plaintext = re.ReplaceAllLiteralString(strings.ToLower(plaintext), "")
	l := len(plaintext)

	var transposed []string
	c := int(math.Ceil(math.Sqrt(float64(l))))

	for i := 0; i < c; i++ {
		var s []byte
		for j := 0; j < c; j++ {
			// last row of the square could be empty
			if j == c-1 && l <= c*(c-1) {
				break
			}
			ch := byte(' ')
			if j*c+i < l {
				ch = plaintext[j*c+i]
			}
			s = append(s, ch)
			if ch == ' ' {
				break
			}
		}
		transposed = append(transposed, string(s))
	}

	return strings.Join(transposed, " ")
}
