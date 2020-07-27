package secret

var actions = []string{"wink", "double blink", "close your eyes", "jump"}

// Handshake returns the handshake sequence corresponding to the given code.
func Handshake(c uint) []string {
	var h []string
	for _, a := range actions {
		if c&1 == 1 {
			h = append(h, a)
		}
		c = c >> 1
	}
	if c&1 == 1 {
		for i, j := 0, len(h)-1; i < j; i, j = i+1, j-1 {
			h[i], h[j] = h[j], h[i]
		}
	}
	return h
}
