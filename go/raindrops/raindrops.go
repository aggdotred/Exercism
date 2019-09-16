package raindrops

import "strconv"

// Convert takes an integer and returns a string based on the factors of the integer provided
func Convert(number int) string {
	var raindrop string
	var sounds = map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}

	for i := 3; i < 8; i += 2 {
		if number%i == 0 {
			raindrop += sounds[i]
		}
	}
	if raindrop == "" {
		raindrop = strconv.Itoa(number)
	}
	return raindrop
}
