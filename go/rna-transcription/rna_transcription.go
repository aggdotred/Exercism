// Package strand takes a DNA strand and returns it's RNA complement (per RNA transcription)
package strand

import "fmt"

var (
	dnaToRNA = map[string]string{"A": "U", "C": "G", "G": "C", "T": "A"}
)

// ToRNA takes a string of dna and returns a string of RNA
func ToRNA(dna string) string {
	var rna string
	if dna == "" {
		return rna
	}
	for _, c := range dna {
		rna += fmt.Sprintf(dnaToRNA[string(c)])
	}
	return rna
}
