// Package weather provides functionalities for managing and reporting weather conditions for a specific location.
package weather

// CurrentCondition stores the current weather condition for a location.
var CurrentCondition string

// CurrentLocation stores the name of the current location.
var CurrentLocation string

// Forecast updates the weather condition and location, then returns a formatted weather report.
// Parameters:
// - city: The name of the location.
// - condition: The weather condition in the location.
// Returns:
// A string containing the location and its current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
