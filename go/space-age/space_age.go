// Package space provdes utilities for space travellers
package space

import "math"

// Planet is the name of one of the planets in the Solar System.
// Valid values are: Mercury, Venus, Earth, Mars, Jupiter, Saturn, Uranus, Neptune.
// Sorry, Pluto!
type Planet string

var orbitalPeriods map[Planet]float64 = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age converts the given amount of seconds to the corresponding years on the specified planet.
func Age(seconds float64, planet Planet) float64 {
	var years float64

	secondsInYear := 60 * 60 * 24 * 365.25
	if orbitalPeriod, prs := orbitalPeriods[planet]; prs {
		years = math.Round(seconds/secondsInYear/orbitalPeriod*100) / 100
	}

	return years
}
