package module01

import (
	"golang.org/x/text/unicode/norm"
)

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	b := []byte(word)
	rb := make([]byte, len(b))

	current := 0
	for {
		if current == len(b) {
			break
		}
		next := norm.NFC.NextBoundary(b[current:], false)
		if next == -1 {
			next = len(b[current:])
		}

		for i := 0; i < next; i++ {
			ri := len(rb) - current - next + i
			rb[ri] = b[current+i]
		}

		current += next
	}

	reversedWord := string(rb)
	return reversedWord
}
