package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//
func Fibonacci(n int) int {
	// F(N) = F(N-1) + F(n-2)
	// F(5) = F(4) + F(3)
	// 13 = 8 + 5
	// n > 1 ( 1 != 0 + -1)
	// if n < 1 {
	// 	return 0
	// }
	// if n == 1 {
	// 	return 1
	// }
	// John's
	if n <= 1 {
		return n
	}
	//

	return Fibonacci(n-1) + Fibonacci(n-2)
}
