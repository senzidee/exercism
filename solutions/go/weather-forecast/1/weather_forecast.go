// Package weather provides a Forecast function to obtain a formatted states of
// conditions.
package weather

// CurrentCondition represents the current condition.
var CurrentCondition string
// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast function is usefull to obtain the formatted conditions states.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
