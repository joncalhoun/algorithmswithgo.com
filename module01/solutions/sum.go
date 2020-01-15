package module01solutions

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	res := 0
	for _, val := range numbers {
		res += val
	}
	return res

	// Recursive
	// if len(numbers) == 0 {
	// 	return 0
	// }
	// return numbers[0] + Sum(numbers[1:])
}
