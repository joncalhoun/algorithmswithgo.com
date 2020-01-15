package module01solutions

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var res string
	for _, r := range word {
		res = string(r) + res
	}
	return res
}
