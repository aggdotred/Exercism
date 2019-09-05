// Package space takes an age in seconds and returns how old someone would be on another planet
package space

var (
	Planet  string = ""
	Seconds int
)

// Age takes an age in seconds and planet and returns how many earth-years old the person would be.
func Age(Seconds int) years int {
	years = Seconds
	return years
}
