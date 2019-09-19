// Package scrabble is used to determine the scrabble score of a word
package scrabble

// Score takes a string and returns the scrabble score as an integer
func Score(word string) int {
	onePointLetters := []string{"A", "a", "E", "e", "I", "i", "O", "o", "U", "u", "L", "l", "N", "n", "R", "r", "s", "S", "T", "t"}
	twoPointLetters := []string{"D", "d", "G", "g"}
	threePointLetters := []string{"B", "b", "C", "c", "M", "m", "P", "p"}
	fourPointLetters := []string{"F", "f", "H", "h", "V", "v", "W", "w", "Y", "y"}
	fivePointLetters := []string{"K", "k"}
	eightPointLetters := []string{"J", "j", "X", "x"}
	tenPointLetters := []string{"Q", "q", "Z", "z"}
	var score int
	for _, c := range word {
		for _, c2 := range onePointLetters {
			if string(c) == c2 {
				score++
			}
		}
		for _, c2 := range twoPointLetters {
			if string(c) == c2 {
				score += 2
			}
		}
		for _, c2 := range threePointLetters {
			if string(c) == c2 {
				score += 3
			}
		}
		for _, c2 := range fourPointLetters {
			if string(c) == c2 {
				score += 4
			}
		}
		for _, c2 := range fivePointLetters {
			if string(c) == c2 {
				score += 5
			}
		}
		for _, c2 := range eightPointLetters {
			if string(c) == c2 {
				score += 8
			}
		}
		for _, c2 := range tenPointLetters {
			if string(c) == c2 {
				score += 10
			}
		}
	}
	return score
}
