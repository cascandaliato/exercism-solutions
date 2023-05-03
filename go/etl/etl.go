// Package etl implements Extract-Transform-Load functions for the Scrabble game.
package etl

import "strings"

// Transform transforms Scrabble scores from a legacy "list of letters per score" format to the shiny new "score per letter" format.
func Transform(legacy map[int][]string) map[string]int {
	scorePerLetter := make(map[string]int)
	for score, letters := range legacy {
		for _, letter := range letters {
			scorePerLetter[strings.ToLower(letter)] = score
		}
	}
	return scorePerLetter
}
