package accumulate

// Accumulate returns a new collection containing the result of applying a given operation to each element of a given collection.
func Accumulate(xs []string, f func(string) string) []string {
	acc := make([]string, len(xs))
	for i := 0; i < len(xs); i++ {
		acc[i] = f(xs[i])
	}
	return acc
}
