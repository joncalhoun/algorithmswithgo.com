package module01
import (
	"fmt"
)

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
func FizzBuzz(number int) {
	for i := 1; i < number;i++ {
		if i % 3 == 0 && i%5 == 0 {
			fmt.Print("Fizz Buzz", ", ")
		}else if i % 3 ==0{
			fmt.Print("Fizz", ", ")
		}else if i%5 == 0{
			fmt.Print("Buzz", ", ")
		}else{
			fmt.Print(i, ", ")
		}
	}
	if number % 3 == 0 && number%5 == 0 {
			fmt.Print("Fizz Buzz", "\n")
		}else if number % 3 ==0{
			fmt.Print("Fizz", "\n")
		}else if number%5 == 0{
			fmt.Print("Buzz", "\n")
		}else{
			fmt.Print(number, "\n")
		}

	
}