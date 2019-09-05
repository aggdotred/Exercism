// Package hamming takes 2 string variables and returns the hamming distance or an error if invalid strings are provided
package hamming

import "errors"

// Distance takes 2 string variables and returns the hamming distance or an error if invalid strings are provided
func Distance(a, b string) (int, error) {
	var distance int
	if len(a) != len(b) {
		return 0, errors.New("the provided DNA strand string lengths do not match")
	}
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
