// Package raindrops provides utility functions to manipulate raindrops.
package raindrops

import "strconv"

// Convert returns a string representation of a number according to the "raindrops" rules.
func Convert(n int) string {
	s := ""

	if n%3 == 0 {
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}

	if s == "" {
		s = strconv.Itoa(n)
	}

	return s
}
