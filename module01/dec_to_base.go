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

		sb.WriteString(fmt.Sprintf("%X", r))

		if dec /= base; dec == 0 {
			break
		}
	}

	return Reverse(sb.String())
}
