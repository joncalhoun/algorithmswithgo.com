# Module 2: The many ways to sort

*NOTE: This README is a collection of my notes for this course. It is NOT a written version of the course, and as such it might not always make sense to you. If anyone wants to submit PRs to help turn this into a written version of the course I'm willing to work with you to do that, but for now it is just my notes.*

## 01 - Why sorting algorithms?

Module 02 is all about sorting algorithms

Before we can discuss why, let's first talk about what a sorting algorithm is.

### What is a sorting algorithm?

An algorithm that rearranges a list into a predefined order.

- A,2,3,4,5 for cards
- By last name, then first
- By age, then name
- Can be in reverse order as well

What is interesting is that almost all sorting algorithms DO NOT care about this so long as we have a way to campare two specific items:

`is A-hearts greater than 2-clubs in our final sort?`

### Why do we learn sorting algorithms?

Nearly everyone is familiar with sorting algorithms

Everyone who has sorted by `Price: Low to High` knows their value

But doesn't std lib offer these?
Yes, but...

- Many future algorithms build upon techniques introduced here.
  - Recursion
  - Divide and conquer
  - Using randomness
- There are many ways to sort, so this gives us a chance to explore many algos for one problem.
  - This means less time understanding the "problem" and more time exploring unique solutions.
  - More time exploring pros/cons of each (speed, memory, etc)
- Practice!
  - Crucial to learning
  - Sorting tends to be challenging & interesting while not so hard that people give up.
- Also demonstrates how theoretical speed != actual speed
  - All of this helps introduce Big O Notation
    - <https://www.interviewcake.com/article/java/big-o-notation-time-and-space-complexity>
  - Some algorithms are "slower" in theory, but faster in practice for smaller lists
  - Others are the same speed in theory, but in practice one almost always outperform another


### How do these algorithms differ?

Many algos have entirely different approaches.
Eg "bubble sort" and "quicksort"

We can also look at variations on the same algo.
Eg "quicksort" can use randomness, or a predefined item, but this has an effect on how it performs.

Some algos use less memory.

Some algos preserve the original list order when two items are the same; others do not.


### Increasing vs non-decreasing order

I'm lenient, but you may hear "non-decreasing" instead of "increasing".

Why?
Duplicates

1,1,1 is technically in non-decreasing order, NOT increasing order.

May feel nit-picky, but worth knowing the terms for future reading.



## 02 - Big O

Big-O notation is a rough way of estimating how long a function takes to run based on the inputs you pass in.


### Some functions are constant time

In some cases functions are constant time in Big O.

When we add two numbers, the "time" it takes to perform this doesn't change:

```go
Add(a, b int) int { return a + b }
```

We always do one operation then return it, so this is `O(1)`, or a constant time function.

That constant can be really big - maybe we ALWAYS do 1 billion operations regardless of the input.

The focus of Big O is how inputs CHANGE the time a function takes, not on evaluating the exact time of an algorithm.


### Big O based on input value

Let's say we had printed out every number from 1 to `max`:

```go
func SumToMax(max int) {
  for i := 1; i <= max; i++ {
    fmt.Println(i)
  }
}
```

This code will be slower with a larger `max` value because that var dictates how many things get printed out.

We would say this function is `O(N)` where N is the value of the argument. ie `N` is equal to the value of `max`.

Many functions exist like this:

```go
rand.Perm(123) // Generates a list with the numbers 0 to 122 in a random order
```

*SEE BENCHMARKS HERE*


### Big O can be based on input size

In other cases we can represent the function's runtime by the SIZE of the input

```go
func SumVals(vals []int) int {
  	var sum int
	for _, val := range vals {
		sum += val
	}
	return sum
}
```

Like before we are printing out values, but this time we are printing out every value in a slice provided.

In this case our function is still `O(N)` in big O, but `N` is equal to the size of the `vals` input, not the values in it.


### Multiple inputs

There are times when we have multiple inputs, and how they affect the runtime depends on each input.

```go
// returns the index of x in the list
func find(x int, list []int) int {
  for i, val := range list {
    if val == x {
      return i
    }
  }
  return -1
}
```

We have two inputs here, but only the size of the list affects the overall time this algorithm takes.

We have `O(N)` where `N == len(list)`


On the other hand we might print out an `X * Y` grid:

```go
func Grid(x, y int) {
  for yi := 0; yi < y; yi++ {
    for xi := 0; xi < x; x++ {
      fmt.Print('X')
    }
    fmt.Println()
  }
}
```

In this case our algorithm is dependent on the value of both inputs, so it is `O(xy)`

Pretty much any combo is possible.

### Sometimes functions grow much faster as the input size increases

Let's say we have a function that prints out all the coordinates of a cube:

```go
func Cube(n int) string {
	var sb strings.Builder
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			for z := 0; z < n; z++ {
				sb.WriteString(fmt.Sprintf("(%d, %d, %d)\n", x, y, z))
			}
		}
	}
	return sb.String()
}
```

Even tho we have one input, this is `O(N^3)` and the time required will grow at a much faster rate as N increases.

You can have functions that grow even faster! Eg `O(2^N)`

### But sometimes my function runs faster...

See the `Find(list, x)` func...

Sometimes list will start with X, so isn't my code faster here?

Big O usually refers to the worse case scenario. That is, if we didn't find X, or if it was in the worst possible spot in the list - the last one we check.

Because of this, Big O is a guideline and not an absolute. In practice you might find an algorithm performing better than Big O suggests because real data isn't alway sin the worst case format.

Sorting is a good example - we are going to dive into Bubble Sort next and in theory it is one fo the slowest sorting algos we will explore, but in practice it does very well for small lists or partially sorted lists.

### Calculating Big O

I'm not going to get into this too much, but in most basic algorithms you can guess this pretty easily.

In more complex code it can be trickier and takes practice.

I'm teaching you roughly what Big O is so you can understand it when it comes up, and so we can talk a little about how these sorting algorithms compare to one another speed wise.

It is also neat to see how some differ despite being the same Big O.


### Additional resources

- <https://www.interviewcake.com/article/java/big-o-notation-time-and-space-complexity>


## 03 - Bubble sort

One of the slowest sorting algorithms we will look at.

`O(N^2)`

At a high level, we are "bubbling" large values to the top of a list.



Works by walking through a list and comparing consecutive pairs.

If next item goes before the current in the list we swap.

*Show example*

One pass will always move the largest unsorted item to the correct spot
So at most N passes to sort
Each pass takes N checks (roughly)
N^2 (more on big O later)
Or in code...

```go
for i := 0; i < len(list); i++ {
  for j := 0; j < len(list); j++ {
    if itemAtJ < itemAtJ+1 {
      swap
    }
  }
}
```

### Optimizations

There are two major optimizations you can make with bubble sort in your code.

1. Stop looking at sorted numbers

Each pass we sort an item
So on the Nth pass, N items should be sorted at the end.
We can code to skip these checks and it cuts our checks roughly in half.
Super minor optimization, but worth looking at.

2. Stop once our list is sorted

If we never make a swap, that means our list is sorted
No need to keep checking things if we have a sorted list!
This is part of why this algorithm can perform faster than others - it is better at recognizing an already sorted list.
We will learn more about where this algo is used later.



## 04 - Bubble sort [code]


## 05 - Bubble sort [bench]



## 06 - Insertion sort

You have probably come up with this on your own when picking up cards and placing them in a hand.

*Show high lvl*

This works well because our "hand" is always sorted. We only have to find where one new card goes.

Programmatically we have to figure out WHERE to put the new card.

*Pull card, go from left to right looking at 3 cards*

We can simplify this by saying, "is the new card less than the current card?"
If so, insert card here and move rest of cards "right".

The fastest version of insertion sort uses binary search to find the new cards spot. We will learn that soon after learning binary search, but for now we can code this without it to get the basics down.

W/out bin search: `O(N^2)`
W/: `O(N log N)` - more on this later.

## 07 - Insertion sort without binary search [code]

The simplest way to write insertion sort is to move sorted items to a new list.

Our function doesn't return a sorted slice, so we need to copy them back into the original.

*Code this...*

Now we can move items to a `sorted` list and then copy them back over when done.

Doesn't affect Big O because that copy is O(N), so doing `O(N) + O(N^2)` is basically the same as `O(N^2)`


SliceTricks: <https://github.com/golang/go/wiki/SliceTricks>


```go
func InsertionSortInt(list []int) {
	var sorted []int
	for _, item := range list {
		sorted = insert(sorted, item)
	}
	for i, val := range sorted {
		list[i] = val
	}
}

func insert(sorted []int, item int) []int {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			return append(sorted[:i], append([]int{item}, sorted[i:]...)...)
		}
	}
	return append(sorted, item)
}
```

## 08 - Practice Problems



## 09 - Binary search

## 10 - Binary search [code]

## 11 - Insertion sort with binary search [code]

## 12 - Recursion

We can go back over many problems from Module 01 for this.

- Sum all numbers in a list
- Convert a number from decimal to another base
- Fib
- GCD
- More?

## 13 - Divide and conquer

Often taught around recursion because the two mentally go hand in hand.

- Binary search

We will see more of these as we get into sorting algorithms, but I'll leave those for those sections and mention then that they are divide and conquery.

## 14 - Merge sort

## 15 - Merge sort [code]

## 16 - Merge sort with goroutines [code]

## 17 - Quicksort explained

## 18 - Quicksort [code]

## 19 - Combining algorithms [code]

## 20 - Benchmarking and comparing
