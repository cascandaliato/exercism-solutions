// Package reverse reverses things.
package reverse

// Reverse reverses the input string.
func Reverse(input string) string {
	runes := []rune(input)
	reversed := make([]rune, len(runes))
	for i := 0; i < len(runes); i++ {
		reversed[len(runes)-1-i] = runes[i]
	}
	return string(reversed)
}
