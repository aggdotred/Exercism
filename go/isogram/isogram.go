// Package isogram determines if a string is an isogram
package isogram

import "strings"

// IsIsogram takes a string and returns true if it is an isogram
func IsIsogram(word string) bool {
	wordMap := map[rune]bool{}
	wordLowerCase := strings.ToLower(word)
	for _, r := range wordLowerCase {
		if wordMap[r] {
			return false
		}
		if r == '-' || r == ' ' {
			continue
		}
		wordMap[r] = true
	}
	return true
}
