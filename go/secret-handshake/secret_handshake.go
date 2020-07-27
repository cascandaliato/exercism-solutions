package secret

var actions = []string{"wink", "double blink", "close your eyes", "jump"}

// Handshake returns the handshake sequence corresponding to the given code.
func Handshake(c uint) []string {
	var h []string
	for _, a := range actions {
		if c&1 == 1 {
			h = append(h, a)
		}
		c >>= 1
	}
	if c&1 == 1 {
		reverse(h)
	}
	return h
}

func reverse(s []string) {
	for i := range s[:len(s)/2] {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
