// Package proverb forms a poem using a slice of provided words
package proverb

import "fmt"

// Proverb takes a slice of words and returns the poem
func Proverb(rhyme []string) []string {
	var poem []string
	var lastLine string
	var line string
	if len(rhyme) > 1 {
		for i, w := range rhyme {
			if i < len(rhyme)-1 {
				line = fmt.Sprintf("For want of a %s the %s was lost.", w, rhyme[i+1])
				poem = append(poem, line)
			}
		}
	} else if len(rhyme) == 0 {
		return poem
	}
	lastLine = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	poem = append(poem, lastLine)
	return poem
}
