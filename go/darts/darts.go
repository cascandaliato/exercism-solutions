package darts

import "math"

const (
	inner  = 1
	middle = 5
	outer  = 10
)

func Score(x, y float64) int {
	switch distance := math.Sqrt(x*x + y*y); {
	case distance <= inner:
		return 10
	case distance <= middle:
		return 5
	case distance <= outer:
		return 1
	default:
		return 0
	}
}
