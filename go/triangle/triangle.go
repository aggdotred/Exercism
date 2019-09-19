// Package triangle takes the sides of a triangle and provides the triangle type
package triangle

import "math"

// Kind is the classification of triangle
type Kind string

const (
	// NaT is not a triangle
	NaT = "NaT"
	// Equ is an equilateral triangle
	Equ = "Equ"
	// Iso is an isoceles triangle
	Iso = "Iso"
	// Sca is a scalene triangle
	Sca = "Sca"
)

// KindFromSides takes 3 floats and returns if the numbers are not the sides of a triangle or what the type of triangle the sides are
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
	} else if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		k = NaT
	} else if (a <= 0) || (b <= 0) || (c <= 0) {
		k = NaT
	} else if (a+b < c) || (a+c < b) || (b+c < a) {
		k = NaT
	} else {
		if (a == b) && (b == c) {
			k = Equ
		} else if (a == b) || (b == c) || (a == c) {
			k = Iso
		} else {
			k = Sca
		}
	}
	return k
}
