package romannumerals

import (
	"bytes"
	"errors"
	"fmt"
)

// ToRomanNumeral takes an integer and returns the roman numeral equivalent.
func ToRomanNumeral(num int) (string, error) {
	rBuf := bytes.Buffer{}
	if num <= 0 || num > 3000 {
		return "", errors.New("not a valid number for conversion")
	}
	thousands := num / 1000
	for i := 0; i < thousands; i++ {
		rBuf.WriteString("M")
	}
	rBuf.WriteString("D")
	rBuf.WriteString("C")
	rBuf.WriteString("L")
	rBuf.WriteString("X")
	rBuf.WriteString("V")
	rBuf.WriteString("I")
	fmt.Sprintln(rBuf)
	return fmt.Sprint(rBuf), nil
}
