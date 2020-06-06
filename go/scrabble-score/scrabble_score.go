// Package scrabble implements functions for the Scrabble game.
package scrabble

import "strings"

var letterValues map[rune]uint8 = map[rune]uint8{}

func init() {
	values := []struct {
		letters []string
		value   uint8
	}{
		{[]string{"a", "e", "i", "o", "u", "l", "n", "r", "s", "t"}, 1},
		{[]string{"d", "g"}, 2},
		{[]string{"b", "c", "m", "p"}, 3},
		{[]string{"f", "h", "v", "w", "y"}, 4},
		{[]string{"k"}, 5},
		{[]string{"j", "x"}, 8},
		{[]string{"q", "z"}, 10},
	}

	for _, v := range values {
		for _, l := range v.letters {
			letterValues[[]rune(l)[0]] = v.value
		}
	}

}

// Score calculates the Scrabble score of a word.
func Score(word string) int {
	score := 0
	for _, char := range strings.ToLower(word) {
		score += int(letterValues[char])
	}
	return score
}
