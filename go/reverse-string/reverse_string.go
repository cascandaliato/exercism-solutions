// Package reverse reverses things.
package reverse

// Reverse reverses the input string.
func Reverse(s string) string {
	var r []rune
	for _, c := range s {
		r = append([]rune{c}, r...)
	}
	return string(r)
}
