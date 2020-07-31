// Package triangle provides functions to classify triangles.
package triangle

import "math"

// Kind is the type of triangle.
type Kind string

const (
	// NaT represents a shape that are not a triangle.
	NaT = "NaT"

	// Equ represents an equilateral triangle.
	Equ = "Equ"

	// Iso represents an isosceles triangle.
	Iso = "Iso"

	// Sca represents a scalene triangle.
	Sca = "Sca"
)

// KindFromSides classifies a triangle given its sides.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case !isValidTriangle(a, b, c):
		return NaT
	case a == b && b == c:
		return Equ
	case a == b, a == c, b == c:
		return Iso
	default:
		return Sca
	}
}

func isValidNumber(n float64) bool {
	return n != 0 && !math.IsInf(n, 0)
}

func isValidTriangle(a, b, c float64) bool {
	return isValidNumber(a*b*c) && a+b >= c && a+c >= b && b+c >= a
}
