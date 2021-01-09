package module01

import "fmt"

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		v := getFizzBuzzValue(i)
		fmt.Print(v)
		if i != n {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func getFizzBuzzValue(n int) string {
	by3 := n%3 == 0
	by5 := n%5 == 0

	switch {
	case by3 && by5:
		return "Fizz Buzz"
	case by3:
		return "Fizz"
	case by5:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", n)
	}
}
