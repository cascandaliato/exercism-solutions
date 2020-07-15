package grains

import (
	"errors"
)

// Square returns the number of grains on the n-th square of a chess board, assuming that the first square holds one grain and the number of grains doubles on each successive square.
func Square(sq int) (uint64, error) {
	if sq <= 0 || sq > 64 {
		return 0, errors.New("The square number must be between 1 and 64")
	}
	return 1 << (sq - 1), nil
}

// Total returns the total number of grains on a chess board, assuming that the first square holds one grain and the number of grains doubles on each successive square.
func Total() uint64 {
	return 1<<64 - 1
}
