package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}
	return s
}
