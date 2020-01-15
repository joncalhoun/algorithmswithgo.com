package module01solutions

import "strings"

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
	const charset = "0123456789ABCDEF"
	var res string
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec = dec / base
	}
	return res
}

// DecToBaseAlt is an alternative solution to DecToBase
func DecToBaseAlt(dec, base int) string {
	const charset = "0123456789ABCDEF"
	// strings.Builder is more efficient than using strings
	// but it really doesn't matter in 90% of apps. I'm just
	// showing you how to use it here in case you are curious
	// and because it allows us to write bytes.
	var res strings.Builder
	for dec > 0 {
		res.WriteByte(charset[dec%base])
		dec /= base
	}
	return Reverse(res.String())
}
