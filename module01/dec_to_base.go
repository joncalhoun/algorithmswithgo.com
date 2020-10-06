package module01

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
	// John's solution
	const charset = "0123456789ABCDEF"
	var b string
	for dec > 0 {
		r := dec % base
		b = string(charset[r]) + b
		dec = dec / base
	}

	return b
}

func myTryDecToBase(dec, base int) string {
	var b string
	r := dec % base

	// what is faster, a switch or map?
	basem := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "A",
		11: "B",
		12: "C",
		13: "D",
		14: "E",
		15: "F",
	}
	b = basem[r] + b

	// switch r {
	// case 0:
	// 	b = "0" + b
	// case 1:
	// 	b = "1" + b
	// case 2:
	// 	b = "2" + b
	// case 3:
	// 	b = "3" + b
	// case 4:
	// 	b = "4" + b
	// case 5:
	// 	b = "5" + b
	// case 6:
	// 	b = "6" + b
	// case 7:
	// 	b = "7" + b
	// case 8:
	// 	b = "8" + b
	// case 9:
	// 	b = "9" + b
	// case 10:
	// 	b = "A" + b
	// case 11:
	// 	b = "B" + b
	// case 12:
	// 	b = "C" + b
	// case 13:
	// 	b = "D" + b
	// case 14:
	// 	b = "E" + b
	// case 15:
	// 	b = "F" + b
	// }

	q := dec / base
	//fmt.Printf("dec %d, base %d, r %d, b %s\n", dec, base, r, b)

	if q > 0 {
		b = DecToBase(q, base) + b
	}

	return b
}
