// Package etl implements Extract-Transform-Load functions for the Scrabble game.
package etl

import "strings"

// Transform transforms Scrabble scores from a legacy "list of letters per score" format to the shiny new "score per letter" format.
func Transform(legacy map[int][]string) map[string]int {
	scorePerLetter := map[string]int{}
	for s, ll := range legacy {
		for _, l := range ll {
			scorePerLetter[strings.ToLower(l)] = s
		}
	}
	return scorePerLetter
}
