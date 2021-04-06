package module01

func GCD(a, b int) int {
	// Given two numbers, A and B:

	// Step 1: If B == 0, return A
	if b == 0 {
		return a
	}

	// Step 2: A becomes B, and B becomes the remainder of dividing A by B
	// `a, b = b, a % b`
	// tmp := a
	// a = b
	// b = tmp % b
	a, b = b, a%b

	// Step 3: Go to step 1
	return GCD(a, b)

	//iterative
	// for {

	// 	if b == 0 {
	// 		return a
	// 	}

	// 	a, b = b, a%b
	// }
}
