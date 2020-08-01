package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	old := []byte(word)
	if len(old) != 0{
		rev := make([]byte, len(old))

		for i,_ := range old {
			index := len(old)-(i+1)
			rev[i] = old[index]
		}

	return string(rev)
	}
	return ""

}
