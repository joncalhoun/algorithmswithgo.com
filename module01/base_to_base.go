package module01

import (
	"math"
	"strings"
)

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//   BaseToBase("E", 16, 2) => "1110"
//

const charset = "0123456789ABCDEF"

func BaseToBase(value string, base, newBase int) string {
	dec := baseToDec(value, base)

	return decToBase(dec, newBase)
}

func baseToDec(value string, base int) int {
	var result int

	p := len(value) - 1
	for _, c := range value {
		num := strings.Index(charset, string(c))
		result = result + num*int(math.Pow(float64(base), float64(p)))
		p--
	}
	return result
}

func decToBase(dec, base int) string {
	var b string
	for dec > 0 {
		r := dec % base
		b = string(charset[r]) + b
		dec = dec / base
	}

	return b
}
