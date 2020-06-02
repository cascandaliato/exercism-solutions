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
	if isValidTriangle(a, b, c) {
		if a == b && b == c {
			return Equ
		} else if a == b || a == c || b == c {
			return Iso
		} else {
			return Sca
		}
	} else {
		return NaT
	}
}

func isValidSide(side float64) bool {
	return side > 0 && !math.IsNaN(side) && !math.IsInf(side, 0)
}

func isValidTriangle(a, b, c float64) bool {
	return isValidSide(a) && isValidSide(b) && isValidSide(c) && a+b >= c && a+c >= b && b+c >= a
}
