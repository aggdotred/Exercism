// Package isogram determines if a string is an isogram
package isogram

import "strings"

// IsIsogram takes a string and returns true if it is an isogram
func IsIsogram(word string) bool {
	wordMap := make(map[rune]int)
	wordLowerCase := strings.ToLower(word)
	for _, r := range wordLowerCase {
		if string(r) != "-" && string(r) != " " {
			if wordMap[r] == 0 {
				wordMap[r] = 1
			} else {
				wordMap[r]++
			}
		}
		if wordMap[r] > 1 {
			return false
		}
	}
	return true
}
