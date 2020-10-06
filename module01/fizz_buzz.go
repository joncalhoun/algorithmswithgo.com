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

func FizzBuzzIf(n int) {
	var msg string

	for i := 1; i <= n; i++ {
		if msg != "" {
			msg += ", "
		}

		if i%3 == 0 && i%5 == 0 {
			msg += fmt.Sprintf("Fizz Buzz")
		} else if i%3 == 0 {
			msg += fmt.Sprintf("Fizz")
		} else if i%5 == 0 {
			msg += fmt.Sprintf("Buzz")
		} else {
			msg += fmt.Sprintf("%d", i)
		}
	}
	fmt.Println(msg)
}

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("Fizz Buzz")
		case i%5 == 0:
			fmt.Printf("Buzz")
		case i%3 == 0:
			fmt.Printf("Fizz")
		default:
			fmt.Printf("%d", i)
		}

		if i != n { //not the end
			fmt.Printf(", ")
		}

	}

	fmt.Println()
}
