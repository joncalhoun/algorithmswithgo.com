package module01

import (
	"strings"
	"unicode/utf8"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	b := []byte(word)
	var sb strings.Builder
	for {
		if len(b) == 0 {
			break
		}
		r, n := utf8.DecodeLastRune(b)
		sb.WriteRune(r)
		b = b[:len(b)-n]
	}
	return sb.String()
}
