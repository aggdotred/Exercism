package raindrops

import "strconv"

// Convert takes an integer and returns a string based on the factors of the integer provided
func Convert(number int) string {
	var raindrop string
	for i := 2; i <= number; i++ {
		if number%i == 0 {
			switch i {
			case 3:
				raindrop += "Pling"
			case 5:
				raindrop += "Plang"
			case 7:
				raindrop += "Plong"
			}
		}
	}
	if raindrop == "" {
		raindrop = strconv.Itoa(number)
	}
	return raindrop
}
