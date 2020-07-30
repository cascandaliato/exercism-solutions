package sieve

// Sieve uses the Sieve of Eratosthenes to return all primes up to (and including) the given number.
func Sieve(lim int) []int {
	marks := make([]int, lim+1)

	for i := 2; i*i <= lim; i++ {
		for j := 2; i*j <= lim; j++ {
			marks[i*j] = 1
		}
	}

	var primes []int
	for i := 2; i <= lim; i++ {
		if marks[i] == 0 {
			primes = append(primes, i)
		}
	}
	return primes
}
