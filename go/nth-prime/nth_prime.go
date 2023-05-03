package prime

import (
	"errors"
	"math/big"
)

// Nth returns the n-th prime number.
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("parameter n must be greater than zero")
	}

	p := 2
	for ; ; p++ {
		if big.NewInt(int64(p)).ProbablyPrime(0) {
			n--
			if n == 0 {
				break
			}
		}
	}

	return p, nil
}
