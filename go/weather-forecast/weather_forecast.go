// Package weather implements functions related to weather forecasting.
package weather

// CurrentCondition holds the current weather condition in CurrentLocation.
var CurrentCondition string

// CurrentLocation is the current location for which the weather is being forecast.
var CurrentLocation string

// Forecast builds a weather forecast message for the given city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
