package module01

import (
	"fmt"
	"strings"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	var sb strings.Builder

	for {
		r := dec % base

		sb.WriteString(digit(r).String())

		if dec /= base; dec == 0 {
			break
		}
	}

	return Reverse(sb.String())
}

type digit int

func (n digit) String() string {
	switch {
	case n <= 9:
		return fmt.Sprintf("%d", n)
	case n == 10:
		return "A"
	case n == 11:
		return "B"
	case n == 12:
		return "C"
	case n == 13:
		return "D"
	case n == 14:
		return "E"
	case n == 15:
		return "F"
	default:
		return ""
	}
}
