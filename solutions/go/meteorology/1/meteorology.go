package meteorology

import "fmt"

// TemperatureUnit represents the unit of temperature.
type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// String method for TemperatureUnit
func (tu TemperatureUnit) String() string {
	switch tu {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	default:
		return "unknown"
	}
}

// Temperature holds a degree value and its corresponding unit.
type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// String method for Temperature (e.g., "21°C")
func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

// SpeedUnit represents the unit of speed.
type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// String method for SpeedUnit
func (su SpeedUnit) String() string {
	switch su {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	default:
		return "unknown"
	}
}

// Speed holds a magnitude and its corresponding unit.
type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// String method for Speed (e.g., "12 km/h")
func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

// MeteorologyData aggregates weather metrics for a specific location.
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// String method for MeteorologyData
// Example output: "London: 21°C, Wind NW at 12 km/h, 65% humidity"
func (md MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity",
		md.location,
		md.temperature, // This will now automatically include the space from above
		md.windDirection,
		md.windSpeed,
		md.humidity,
	)
}