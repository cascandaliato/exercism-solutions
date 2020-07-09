package prime

import "math/big"

// Nth returns the n-th prime number.
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
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

	return p, true
}
