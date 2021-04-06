package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// var res string
// for _, r := range word {
// 	res = string(r) + res
// }
func Reverse(word string) string {
	// i := len(word)
	// revword := make([]rune, i)
	// for _, c := range word {
	// 	revword[i-1] = c
	// 	i--
	// }
	// return string(revword)
	var revword string
	for _, c := range word {
		revword = string(c) + revword
	}
	return revword

}
