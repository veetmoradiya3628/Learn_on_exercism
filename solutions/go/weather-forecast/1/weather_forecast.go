// Package weather provides information on the weather details.
package weather

var (
	// CurrentCondition represents current condition.
    CurrentCondition string
    
	// CurrentLocation represents current location.
    CurrentLocation  string
)

// Forecast returns forecast details for the city & condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
