// Package weather provides tools for getting a forecast.
package weather

// CurrentCondition represents a current weather situation.
var CurrentCondition string

// CurrentLocation represents a current location.
var CurrentLocation string

// Forecast returns an string value of a weater condition in the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
