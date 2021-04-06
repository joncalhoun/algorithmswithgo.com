package module01

// Sum will sum up all of the numbers passed
// in and return the result FOR LOOP
// func Sum(numbers []int) int {
// 	// var sum int
// 	sum := 0
// 	for _, n := range numbers {
// 		// fmt.Printf("sum %d n %d\n", sum, n)
// 		sum += n
// 		// sum = sum + n
// 		// fmt.Printf("sum %d\n", sum)
// 	}
// 	return sum
// }

// Sum will sum up all of the numbers passed
// in and return the result RECURSIVE
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	return numbers[0] + Sum(numbers[1:])
}
