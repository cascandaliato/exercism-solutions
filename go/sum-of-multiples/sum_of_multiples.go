package summultiples

// SumMultiples calculates the sum of the unique multiples of the given divisors, up to but not including the given limit.
func SumMultiples(limit int, divisors ...int) int {
	multiples := map[int]bool{}
	for _, d := range divisors {
		if d == 0 {
			continue
		}
		for i := 1; i*d < limit; i++ {
			multiples[i*d] = true
		}
	}

	var sum int
	for m := range multiples {
		sum += m
	}
	return sum
}
