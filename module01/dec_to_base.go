package module01

import (
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
	const s = "0123456789ABCDEF"
	var sb strings.Builder
	for {
		r := dec % base
		sb.WriteByte(s[r])
		if dec /= base; dec == 0 {
			break
		}
	}
	return Reverse(sb.String())
}
