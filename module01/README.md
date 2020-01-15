# Module 1: Getting started with Algorithms

*NOTE: This README is a collection of my notes for this course. It is NOT a written version of the course, and as such it might not always make sense to you. If anyone wants to submit PRs to help turn this into a written version of the course I'm willing to work with you to do that, but for now it is just my notes.*

## 01 - Intro

Welcome!

We will be learning about algorithms, and coding them using Go.


### Who is this course for?

EVERYONE!

- Beginners can pair this with introductory books.
- Self-taught devs can finally get an understanding of algorithms like their CS degree peers.
- College students can refresh or get a new take on algorithms they are learning in school.
- Professional devs can dive in for fun for for interview prep!

### How will we learn?

We will not focus on the math.

- Proofs can be confusing
- Big-O is useful, and we will explore it, but it won't be our primary focus.

Math isn't unimportant, but it can be overwhelming or distracting.
Proofs NEVER help me understand an algorithm.

Example: <https://en.wikipedia.org/wiki/Quicksort#Formal_analysis>

Is that true for everyone? Probably not, but for me a heavy math focus is the wrong way to learn algorithms.

**What will we be doing instead?**

- Using simple language
- Using examples and sketching how things change as an algorithm runs
- Learning how to code each algorithm (very important!)
- Exploring when and how to apply each algorithm
- Practicing!


## Why should I learn algorithms?

Truthfully, you likely won't ever need to use half (or more) of the algorithms you learn about here in the real world.

The real goal here is to practice taking an idea and translating it into code. In my experience, algorithms are great for this.


**Other reasons include...**

You WILL need to come up with your algorithms to tackle your own custom problems, and having a basic understanding of how algorithms work can help.

It is useful to understand time complexity, and this is usually taught with algorithms.

Interviews often use quiz problems that are much easier after diving into algorithms like the ones we cover in this course.

Algorithms really can be fun if we stop focusing on the boring parts!


### Course format (roughly)

In order to learn, we typically follow the following format:

1. Learn about an algorithm
2. Look at a practice problem and code a solution

Sometimes we will look at multiple practice problems. Sometimes we will look at problems that look nothing like the algorithm we learned so you can see how it isn't always obvious what algorithms can be applied to solve a problem.

But in this first module we will have a few videos where I cover some basic background information that I think might be useful before diving in.


## 02 - What is an algorithm?

> a process or set of rules to be followed in calculations or other problem-solving operations, especially by a computer
> - Google Dictionary

Algorithms don't have to be super complex mathematical equations.

They are just some rules to solve a problem.

The hard part is defining rules that are not ambiguous, when we are so used to instructions being ambiguous because we humans can infer a lot.

**Ex: Recipe to cook something**
- What is a "pinch of salt"?
- What exactly does sauteing something entail?
- What is "tender" in "cook until tender"?
- Do we have to inform people to get things out of the fridge? To crack their eggs? How fast to stir the pot?

Its like the office when Michael turns into a lake: <https://www.youtube.com/watch?v=DOW_kPzY_JY>
If you tell a computer to turn right into a lake, it will.


## 03 - How to improve

Many people will "get" algorithms conceptually, but will struggle with coding them.
This is usually a sign that they struggle with creating a set of unambiguous steps; they aren't skilled in translating algorithms into code.

The best way to master this skill is to practice, a LOT!

When I was in college and I wanted to get really good at TopCoder, Google Code Jam, etc, I started doing practice problems all the time.
Like, 4+hrs a day, every day.
I would code between classes at the comp lab.
I would code after work.
I would code before bed.
My roommate, who didn't code, probably knew more about algorithms than any sane non-programmer should.
As a result, I improved a lot over a short period of time.
If you want to get skilled at this, you are going to need to practice a lot.

This whole first module focuses on simpler algorithms, but the actual algorithms are often silly in their simplicity.

The goal of this first module is to teach you to practice so that you can improve quickly and efficiently. This will be necessary to get through all the later modules.

## 04 - How to use the provided materials

*NOTE: This README is a collection of my notes for this course. It is NOT a written version of the course, and as such it might not always make sense to you. If anyone wants to submit PRs to help turn this into a written version of the course I'm willing to work with you to do that, but for now it is just my notes.*

Github repo: <https://github.com/joncalhoun/algorithmswithgo.com>

I have tried to provide tests that make it easy for you to verify your code as you write it.

To test, navigate to the code and run `go test`

```bash
go test
```

*Sidenote: Want color? Try something like this: <https://github.com/rakyll/gotest>. I use a bash function that does something similar. It is provided below.*

```bash
go_test() {
  go test $* | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/SKIP/s//$(printf "\033[34mSKIP\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | GREP_COLOR="01;33" egrep --color=always '\s*[a-zA-Z0-9\-_.]+[:][0-9]+[:]|^'
}
```

*Bash script stems from something shared by Dave Cheney on Twitter iirc*

Generally you don't want to run ALL of the tests.
We can use the `-run` flag to run a tests for a specific function.

```bash
go test -run=NumInList
```

It is also worth noting that these tests aren't a great example of how you should necessarily write tests. Many of them are written specifically to make it easier for beginners. Eg rather than passing an `io.Writer` into the `FizzBuzz` function, I instead opt to change what `os.Stdout` is in the test. I generally wouldn't recommend this.

If that is all confusing to you, don't worry about it. It is a technical detail that doesn't matter right now. I'm just saying don't use the algorithm tests as a way to learn to testing in Go. Instead, start here: <https://www.calhoun.io/how-to-test-with-go/>

I also have a paid course on testing here: <https://testwithgo.com>
If you like this course, you may enjoy that one too.

And one final warning - these tests are NOT exhaustive. They may miss corner cases. The idea is to give you a decent starting point to verify your code but that's it. They will probably suffice for the first few modules, but in later, more complex algorithms, they probably won't. See the last few sections of this module for some suggestions on other ways to practice and test your code using sites like Code Jam, HackerRank, etc.

As the course progresses we will probably start to use benchmarks too. I'm not going to teach those, but there are tons of great resources including the std lib docs, Dave Cheney's blog post (<https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go>), and I even touch on them briefly in my testing course.

Lastly, the `solutions` directory will have code solutions (along with a copy of the tests) in case you need them. I'll never check solutions into the base directory so you have a good starting point with tests.


## 05 - Determine if a number is in a list [code]

Source file: `num_in_list.go`
Function def: `NumInList(list []int, num int) bool`

Given a list of numbers, determine if a specific number is in that list.

Ex:

```go
NumInList([]int{1,2,3,4,5}, 5) // true
NumInList([]int{3,3,3,3,3}, 5) // false
NumInList([]int{3,5,3,5,3}, 5) // true
NumInList([]int{4,2,22,-10,8}, -10) // true

// empty lists!
NumInList(nil, 5) // false
NumInList([]int{}, 5) // false
```

## 06 - Sum all the numbers in a list [code]

Source file: `sum.go`
Function def: `Sum(numbers []int) int`

Given a list of numbers, sum them all up and return the sum.


Ex:

```go
Sum([]int{1,2,3,4,5}) // 15
Sum([]int{3,3}) // 6
Sum([]int{3,5,3,5,3}) // 19
Sum([]int{4,2,22,-10,8}) // 26

// let's just return 0 for empty lists
Sum(nil) // 0
Sum([]int{}) // 0
```

## 07 - Reverse a string [code]

Source file: `reverse.go`
Function def: `Reverse(word string) string`

Given a string, return its reverse.

Ex

```go
Reverse("cat") // "tac"
Reverse("ca t") // "t ac"
Reverse("alphabet") // "tebahpla"
```


## 08 - FizzBuzz [code]

Source file: `fizz_buzz.go`
Function def: `FizzBuzz(n int)`

Given a number N, print out all the numbers from 1 to N but when a number is divisible by 3 print out "Fizz" instead of the number, and when a number is divisible by 5 print out "Buzz" instead of the number. For numbers divisible by both 3 and 5 print "Fizz Buzz".

Ex

```go
FizzBuzz(1)
// 1
FizzBuzz(5)
// 1, 2, Fizz, 4, Buzz
FizzBuzz(15)
// 1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz
```


## 09 - Decimal to another base [code]

Source file: `dec_to_base.go`
Function def: `DecToBase(dec, base int) string`

You can convert a number from one base to another by:

1. Take the number you have in decimal and get the remainder when it is divided by the new base. The remainder is the next "digit" in the new base. You add this digit to the LEFT of the new number as you create it.

*NOTE: If the digit is >=10, then you need to use a letter like A, B, C, D, ...*

2. Divide your number by the new base, then repeat from step (1) if the new number is > 0.


Example 1:

```
14 to base 2:

Step 1: 14 % 2 = 0, so next digit is "0".
  Our new number is currently "0"

Step 2: 14 / 2 = 7, so we update our number to be 7
  7 is > 0, so we go to step 1...

Step 1: 7 % 2 = 1; next digit is "1"
  Our new number is currently "10" (we add the new digit to the left)

Step 2: 7 / 2 = 3 (round down always), so we update our number to be 3
  3 > 0, so go to step 1

Step 1: 3 % 2 = 1; next digit is "1"
  Our new number is "110"

Step 2: 3 / 2 = 1; Go to step 1

Step 1: 1 % 1 = 1; next digit is "1"
  Our new number is "1110"

Step 2: 1 / 2 = 0; stop here!

Final number is "1110", which is 14 in base 2!
```

Example 2:

```
15 to base 16:

Step 1: 15 % 16 = 15, so next digit is "F"
  Current new number: "F"

Step 2: 15 / 16 = 0. stop!

Final new number: "F"
```


*Note: If it is easier, you can use the `Reverse` function from earlier to help solve this. Breaking problems into smaller problems and reusing solutions is often a great way to make them easier.*

Ex

```go
DecToBase(1, 2) // "1"
DecToBase(2, 2) // "10"
DecToBase(7, 3) // "21"
DecToBase(14, 2) // "1110"
DecToBase(14, 16) // "E"
DecToBase(17, 16) // "11"
```

## 10 - Another base to decimal [code]

Source file: `base_to_dec.go`
Function def: `BaseToDec(value string, base int) int`


Each digit gets multiplied by the base raised to a power. Eg for 1110 in base 2:

```
1     2     3
10^2  10^1  10^0

1    1    1    0
2^3  2^2  2^1  2^0

so this is: 1*2^3 + 1*2^2 + 1*2^1 + 0*2^0 = 8 + 4 + 2 + 0 = 14
```

```go
BaseToDec("1", 2) // 1
BaseToDec("10", 2) // 2
BaseToDec("21", 3) // 7
BaseToDec("1110", 2) // 14
BaseToDec("E", 16) // 14
BaseToDec("11", 16) // 17
```


## 11 - Any base to any base [code]

Source file: `base_to_base.go`
Function def: `BaseToBase(value string, base, newBase int) string`

Break it into smaller problems we have already solved!

Ex:

```go
BaseToBase("E", 16, 2) // "1110"
BaseToBase("8831A383B", 12, 16) // "DEADBEEF"
```

## 12 - Find two that sum [code]

Source file: `find_two_that_sum.go`
Function def: `FindTwoThatSum(numbers []int, sum int) (int, int)`

Find two numbers in a list that sum to a given amount.

FindTwoThatSum will look for two numbers in the provided numbers list that sum up to the sum argument. It will then return the indices of each of these numbers.

If there are multiple correct answers, any of them may be used. If there are no correct answers, (-1, -1) should be returned.

Ex:

```go
FindTwoThatSum([]int{1, 2, 3, 4}, 7) // (2, 3)
FindTwoThatSum([]int{0, -1, 1}, 0) // (1, 2)
FindTwoThatSum([]int{0, 1, 1}, 0) // (-1, -1)
```

Or if we have duplicate answers any of them are okay. All of the following are correct.

```go
FindTwoThatSum([]int{10, 1, 12, 3, 7, 2, 2, 1}, 4) // (5, 6)
FindTwoThatSum([]int{10, 1, 12, 3, 7, 2, 2, 1}, 4) // (1, 3)
FindTwoThatSum([]int{10, 1, 12, 3, 7, 2, 2, 1}, 4) // (1, 7)
```

Each number will only be used once, and the original numbers list should not be rearranged or altered in any way.

If you want to challenge yourself further, try adding different criteria for which numbers will be used when there are multiple correct answers.
Eg:

1. Always return the lowest index possible for the first value.
2. Always return the index of lowest value possible for the first return value.
3. Always return the index of the two values who have a minimal difference. Eg prefer the values 2, 2 over 1, 3 over 0, 4 for the sum of 4.


## 13 - Prime factorization (primes provided) [code]

Source file: `factor.go`
Function def: `Factor(primes []int, number int) []int`

Factor takes in a list of primes and a number and factors that number with the provided primes.

The returned numbers can be in any order; tests will sort them in increasing order to make checking easier.

Bonus: Any remainder (aside from 1) that can't be factored will be treated as a prime in the results.

Examples:

```go
Factor([]int{2,3,5,7}, 30) // []int{2,3,5}
Factor([]int{2,3,5,7}, 28) // []int{2,2,7}
Factor([]int{2,3,5,7}, 720) // []int{2,2,2,3,3,5}
```

Examples with remainders:

```go
Factor([2,5], 30) // []int{2,5,3}
Factor([3,5], 720) // []int{3,3,5,16}
Factor([], 4) // []int{4}
```

## 14 - Fibonacci sequence [code]

Source file: `fibonacci.go`
Function def: `Fibonacci(n int) int `

Fibonacci returns the `n`th fibonacci number.

The fibonacci series is:

A Fibonacci number N is defined as:

```
Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
```

Where the following base cases are used:

```go
Fibonacci(0) // 0
Fibonacci(1) // 1
```

Examples:

```go
Fibonacci(0) // 0
Fibonacci(1) // 1
Fibonacci(2) // 1
Fibonacci(3) // 2
Fibonacci(4) // 3
Fibonacci(5) // 5
Fibonacci(6) // 8
Fibonacci(7) // 13
Fibonacci(14) // 377
```

## 15 - Greatest common divisor [code]

Source file: `gcd.go`
Function def: `GCD(a, b int) int`

GCD stands for greatest common divisor

Given two numbers, GCD calculates the largest number you could divide both numbers by without getting a remainder.

Examples:

```
A = 10, B = 5, GCD = 5 (10%5 == 0, 5%5 == 0)
A = 25, B = 5, GCD = 5
A = 30, B = 15, GCD = 15
A = 30, B = 9, GCD = 3
A = 100, B = 9, GCD = 1
```

One way to solve this is to factor each number and then find the common factors from each:


```
A = 10, B = 5
  10 = [2,5]
  5 = [5]
  GCD = [5] = 5

A = 30, B = 15
  30 = [2,3,5]
  15 = [3,5]
  GCD = [3,5] = 15

A = 30, B = 9:
  30 = [2,3,5]
  9 = [3,3]
  GCD = [3] = 3

A = 100, B = 9:
  100 = [2,2,5,5]
  9 = [3,3]
  GCD = [] = 1


A = 100, B = 8:
  100 = [2,2,5,5]
  9 = [2,2,2]
  GCD = [2,2] = 4
```

If you wanted, you could get a big list of primes and use our `Factor` function to solve this problem that way.

Another way to solve GCD is using the Euclidean algorithm. We won't dive into the math of it, but what makes this algorithm really nice is both its simplicity and the fact that we don't actually need to factor numbers using primes.

Here is how it works:

Given two numbers, A and B:

Step 1: If B == 0, return A
Step 2: A becomes B, and B becomes the remainder of dividing A by B
`a, b = b, a % b`
Step 3: Go to step 1

*Note: It doesn't matter if `A > B` or `A < B`.*

This is a really good algorithm to know because GCD is often used in algorithmic contests. For instance, GCD was used in a qualifier for the most recent (2019) Google Code Jam. It is also a relatively simple algorithm that we can look at and then try to translate into code while exploring both recursive and iterative solutions (which we will do in a later module on recursion).

## 16 - stdin and stdout

Many problem solve sites have you reading from `stdin` and `stdout` or files.

`stdin` (standard input) is input that is provided by a user, often by typing from their keyboard, but this can also be piped in in most terminals.

`stdout` (standard output) is where most programs will write output. In Go you write here by default if you do something like `fmt.Println("hi")`.

Piping files to stdin and stdout is also pretty easy:

```bash
$ go run main.go < input.txt > output.txt
```

In this section we will cover how to more easily read from stdin without having to type every test case every single time, how to write to standard out and save the results in a file, and along the way we will also look at different ways to make it easier to read things like integers, floats, whole lines, single words, and more.

This is all meant to just make it easier if you want to go practice on sites like Google Code Jam and others that tend to use stdin and stdout for input/output of a program. I highly recommend practicing with sites like this, so I highly recommend this section 😜

