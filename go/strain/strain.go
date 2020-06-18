package strain

// Ints is a collection of integers.
type Ints []int

// Lists is a matrix of integers.
type Lists [][]int

// Strings is a collection of strings.
type Strings []string

// Keep returns a new collection containing only those elements where the predicate is true.
func (ints Ints) Keep(predicate func(int) bool) Ints {
	var kept Ints
	for _, v := range ints {
		if predicate(v) {
			kept = append(kept, v)
		}
	}
	return kept
}

// Discard returns a new collection containing  only those elements where the predicate is false.
func (ints Ints) Discard(predicate func(int) bool) Ints {
	var discarded Ints
	for _, v := range ints {
		if !predicate(v) {
			discarded = append(discarded, v)
		}
	}
	return discarded
}

// Keep returns a new collection containing only those elements where the predicate is true.
func (lists Lists) Keep(predicate func([]int) bool) Lists {
	var kept Lists = Lists{}
	for _, v := range lists {
		if predicate(v) {
			kept = append(kept, v)
		}
	}
	return kept
}

// Keep returns a new collection containing only those elements where the predicate is true.
func (strings Strings) Keep(predicate func(string) bool) Strings {
	var kept Strings
	for _, v := range strings {
		if predicate(v) {
			kept = append(kept, v)
		}
	}
	return kept
}
