package encode

import (
	"strconv"
	"strings"
)

// RunLengthEncode encodes the input string using run-length encoding.
func RunLengthEncode(input string) string {
	var (
		enc  strings.Builder
		prev rune
		c    int
	)

	for _, r := range input + "#" {
		if r == prev {
			c++
			continue
		}
		if c > 1 {
			enc.WriteString(strconv.Itoa(c))
		}
		if c > 0 {
			enc.WriteRune(prev)
		}
		prev, c = r, 1
	}

	return enc.String()
}

// RunLengthDecode decodes a run-length encoded string.
func RunLengthDecode(enc string) string {
	var (
		dec strings.Builder
		c   int
	)

	for _, r := range enc {
		if d, err := strconv.Atoi(string(r)); err == nil {
			c = c*10 + d
			continue
		}
		if c == 0 {
			c = 1
		}
		dec.WriteString(strings.Repeat(string(r), c))
		c = 0
	}

	return dec.String()
}
