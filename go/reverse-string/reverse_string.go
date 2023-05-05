// Package reverse reverses things.
package reverse

// Reverse reverses the input string.
func Reverse(input string) string {
	runes := []rune(input)
	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - 1 - i
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
