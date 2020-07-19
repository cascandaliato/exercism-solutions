package wordcount

import (
	"regexp"
	"strings"
)

var wordRegex = regexp.MustCompile(`\b(\d+|[a-z]|[a-z][a-z']*[a-z])\b`)

// Frequency stores how many times words appear in a phrase.
type Frequency map[string]int

// WordCount counts the number of occurrences of each word in the given phrase.
func WordCount(phrase string) Frequency {
	f := Frequency{}
	for _, w := range wordRegex.FindAllString(strings.ToLower(phrase), -1) {
		f[w]++
	}
	return f
}
