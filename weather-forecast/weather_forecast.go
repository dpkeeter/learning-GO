// Package weather is used  find the forecast
// of the weather for the provided city.
package weather

// CurrentCondition is the current condition.
var CurrentCondition string

// CurrentLocation is the current location.
var CurrentLocation string

// Forecast() function provides the weather forcast.
func Forecast(city, condition string) string {

	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
