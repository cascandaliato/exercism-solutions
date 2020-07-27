package clock

import "fmt"

// Time represents a time (hours and minutes) without a date.
type Time struct{ h, m int }

// New creates a new Time corresponding to the given hours and minutes. It normalizes the input to fall in the 24h range.
func New(h, m int) Time {
	t := Time{h, m}
	t.normalize()
	return t
}

// String returns a string representation of the time in HH:MM format.
func (t Time) String() string {
	return fmt.Sprintf("%02d:%02d", t.h, t.m)
}

// Add returns a new Time m minutes after the original Time.
func (t Time) Add(m int) Time {
	return New(t.h, t.m+m)
}

// Subtract returns a new Time m minutes before the original Time.
func (t Time) Subtract(m int) Time {
	return New(t.h, t.m-m)
}

func (t *Time) normalize() {
	t.h += t.m / 60
	t.m %= 60
	if t.m < 0 {
		t.h--
		t.m += 60
	}
	t.h = t.h % 24
	if t.h < 0 {
		t.h += 24
	}
}
