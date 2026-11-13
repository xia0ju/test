// Package weather provides tools for weather forecast.
package weather

var (
    // CurrentCondition represents the current weather condition.
	CurrentCondition string
// CurrentLocation represents the current location for the weather.
	CurrentLocation  string
)
// Forecast returns a string combining the current location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
