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


## 02 - Bubble sort

One of the slowest sorting algorithms we will look at.

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


## 03 - Bubble sort [code]



## 04 - Insertion sort

You have probably come up with this on your own

We are going to explain what you would probably do on your own, then we will explore how you would repeat that with code.

- Show impl

First impl will be slow because we need a "fast" way to find where the new cards go.

After we learn binary search we will come back to this sorting algorithm to do it again. Maybe benchmark the two.

## 05 - Insertion sort without binary search [code]

## 06 - Binary search

## 07 - Binary search [code]

## 08 - Insertion sort with binary search [code]

## 09 - Recursion

We can go back over many problems from Module 01 for this.

- Sum all numbers in a list
- Convert a number from decimal to another base
- Fib
- GCD
- More?

## 10 - Divide and conquer

Often taught around recursion because the two mentally go hand in hand.

- Binary search

We will see more of these as we get into sorting algorithms, but I'll leave those for those sections and mention then that they are divide and conquery.

## 11 - Merge sort

## 12 - Merge sort [code]

## 13 - Merge sort with goroutines [code]

## 14 - Quicksort explained

## 15 - Quicksort [code]

## 16 - Big O

## 17 - Benchmarking some sorting algorithms

## 18 - Combining algorithms [code]
