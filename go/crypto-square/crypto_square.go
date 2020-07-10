package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// Encode encodes a plaintext using the square code technique.
func Encode(plaintext string) string {
	if len(plaintext) == 0 {
		return ""
	}

	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	plaintext = re.ReplaceAllLiteralString(strings.ToLower(plaintext), "")

	r, c := fitRowCol(plaintext)
	plaintext = spacePadRight(plaintext, r*c)
	rows := splitEvery(plaintext, c)

	var transposed []string
	for i := 0; i < c; i++ {
		var s []byte
		for _, row := range rows {
			ch := row[i]
			s = append(s, ch)
			if ch == ' ' {
				break
			}
		}
		transposed = append(transposed, string(s))
	}

	return strings.Join(transposed, " ")
}

func fitRowCol(str string) (int, int) {
	l := len(str)
	for r := 1; ; r++ {
		if (r-1)*r+1 <= l && l <= r*r {
			return r, r
		} else if (r-1)*(r+1)+1 <= l && l <= r*(r+1) {
			return r, r + 1
		}
	}
}

func spacePadRight(str string, length int) string {
	return str + strings.Repeat(" ", int(math.Max(0, float64(length-len(str)))))
}

func splitEvery(str string, n int) []string {
	var splits []string
	for i := 0; (i+1)*n <= len(str); i++ {
		splits = append(splits, str[i*n:(i+1)*n])
	}
	return splits
}
