// Package weather lear what is your location's weather conditions.
package weather

//CurrentCondition string.
var CurrentCondition string

//CurrentLocation string.
var CurrentLocation string

//Forecast return your weather condition for your's location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
