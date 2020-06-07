// Package dna provides utility functions to work with DNA strands.
package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]uint32

// DNA is a sequence of nucleotides.
type DNA string

var nucleotides = map[rune]bool{
	'A': true,
	'C': true,
	'G': true,
	'T': true,
}

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = map[rune]uint32{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, n := range d {
		if _, ok := nucleotides[n]; !ok {
			return h, fmt.Errorf("DNA strand %s contains a non-valid nucleotide %c", d, n)
		}
		h[n]++
	}

	return h, nil
}
