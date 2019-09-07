// Package space takes an age in seconds and returns how old someone would be on earth
package space

// Planet string is the variable that will indicate what planet the age would be in earth years
type Planet string

// Age takes an age in seconds and planet and returns how many earth-years old the person would be.
func Age(seconds float64, planet Planet) float64 {
	var years float64
	earthYearSeconds := 31557600.
	switch planet {
	case "Earth":
		years = seconds / earthYearSeconds
	case "Mercury":
		years = seconds / (earthYearSeconds * 0.2408467)
	case "Venus":
		years = seconds / (earthYearSeconds * 0.61519726)
	case "Mars":
		years = seconds / (earthYearSeconds * 1.8808158)
	case "Jupiter":
		years = seconds / (earthYearSeconds * 11.862615)
	case "Saturn":
		years = seconds / (earthYearSeconds * 29.447498)
	case "Uranus":
		years = seconds / (earthYearSeconds * 84.016846)
	case "Neptune":
		years = seconds / (earthYearSeconds * 164.79132)
	}

	return years
}
