// Package hamming implements functions to calculate the Hamming distance between DNA strands.
package hamming

import "errors"

// Distance calculates the Hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("the two DNA strands should have the same length")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
