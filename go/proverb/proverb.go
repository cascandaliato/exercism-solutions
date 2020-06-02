// Package proverb provides generators of useful proverbs.
package proverb

import "fmt"

// Proverb generates variants of the proverb "For want of a horseshoe nail, a kingdom was lost".
func Proverb(rhyme []string) []string {
	var proverb []string

	for i := 0; i < len(rhyme)-1; i++ {
		sentence := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		proverb = append(proverb, sentence)
	}

	if len(rhyme) >= 1 {
		ending := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		proverb = append(proverb, ending)
	}

	return proverb
}
