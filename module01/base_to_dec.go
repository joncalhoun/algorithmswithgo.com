package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	var result int

	p := len(value) - 1
	for _, c := range value {
		num := strings.Index(charset, string(c))
		result = result + num*int(math.Pow(float64(base), float64(p)))
		p--
	}
	return result
}
