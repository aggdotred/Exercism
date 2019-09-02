// Package twofer takes a string and inserts it into a string or provides a default value.
package twofer

import "fmt"

// ShareWith generalizes the "One for you, one for me." string with you if no name is provided.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
