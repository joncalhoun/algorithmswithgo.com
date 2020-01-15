# Why sorting algorithms?

## What is a sorting algorithm?

A sorting algorithm is an algorithm designed to take a list and re-arrange it into a predefined order. When initially learning, this almost always means taking a list of numbers and sorting it into increasing order. Eg you might start with the list:

```
6, 5, 3, 4, 2, 1
```

And then sort it into the following order:

```
0, 1, 2, 3, 4, 5
```

As you start to learn more about sorting algorithms you will start to understand that they can do much more than sort a list of numbers and that there are a multitude of variations we can explore. For instance, you can use sorting algorithms to sort a list of students by their last name then first name.

```
Beesly, Pam
Bernard, Andy
Halpert, Jim
Schrute, Dwight
Scott, Michael
```

You could even sort by two variables, such a person's age and then name:

```
8 - Jane
84 - Alice
84 - Bob
84 - Mary
88 - Jane
89 - Austin
```


## Why do we learn sorting algorithms?

Sorting algorithms are a good starting point because nearly everyone is familiar with sorting, and sorting algorithms are easy to visualize. Pretty much everyone has looked at an online store and understands what it means to sort by `Price: Low to High`, and they can also appreciate the value of this.

Sorting algorithms also make a great starting point because many future algorithms we will look at build upon them. For instance, binary search assumes that you have a sorted list, so understanding how to sort a list is valuable. Or when we learn about heaps we will see that we are solving a very similar, but slightly different problem and understanding sorting algorithms make it easier to understand how a heap data structure works.

Sorting is also a unique subject because there are a plethora of algorithms all designed to do the same thing. In this course we will learn about at least five different sorting algorithms, and there are plenty that we won't cover in this course. This is helpful because it allows us to use the same base problem (and practice problems) while experimenting with different solutions and learning from each.

Another advantage to sorting algorithms is the wide variety of topics they cover. In this module alone we will cover recursion, divide and conquer, and even how randomness can be used in algorithms.

Practice is also crucial to learning, and finding a good balance of challenging but not insurmountable is a key part of practice. Sorting algorithms tend to be challenging enough to keep students interested and excited, while not being so hard that a majority of students get stumped or give up.

Finally, another interesting aspect of sorting algorithms is that ones that are technically slower can perform better in very specific scenarios, while algorithms that appear to have the same overall runtime will turn out to perform very different in reality. All of this helps make sorting a great time to discuss [Big O Notation](https://www.interviewcake.com/article/java/big-o-notation-time-and-space-complexity) while also demonstrating that this is just a guideline and not an exact analysis of how fast your code will be.


## How do these algorithms differ?

Many of the algorithms we will explore have entirely different approaches to solving the problem. Typically when this happens each algorithm has a unique name like "bubble sort" or "quicksort".

At other times we will look at variations of the same algorithm. For instance, when learning quicksort we will see how one value we use in the algorithm can be chosen both deterministically and randomly, and we will talk about how this might affect the overall speed of the algorithm.

Other variations are more subtle. For instance, we might be interested in sorting algorithms that use as little memory as possible when dealing with tiny hardware.

Or we might need an algorithm that preserves the original order of the list as much as possible in some scenarios. Eg if we started with the list below:

```
Diapers - $10.99
Baby Oil - $4.99
Shampoo - $5.99
Crackers - $2.99
Pens - $4.99
```

And we wanted to sort this list from cheapest to most expensive products we would find that two of the items have the same price. In some sorting algorithms there is no guarantee that the original order will be preserved, while others may make this promise. Whether or not that matters to you might shape which algorithm you use.


## Increasing vs non-decreasing order

While I have been more lenient in this lesson, typically when we talk about sorting we will use the terms "non-decreasing order" or "non-increasing order" rather than "increasing order" and "decreasing order.

For example, the list `1, 2, 3, 4` is in `non-decreasing order`. Many of you will wonder why I don't just say it is in "increasing order", and this mostly stems from the fact that some lists will have duplicate items. Imagine we added another `3` to the list above:

```
1, 2, 3, 3, 4
```

Now this list is no longer in "increasing order" because the second `3` isn't increasing in value compared to the previous `3`. On the hand if we say the list is in "non-decreasing order" this implies that every value simply has to be equal to or greater than the previous value.

This distinction might feel like nit-picking, and truthfully I don't mind what terms you use, but you are very likely going to run across both terms and I want you to be familiar with the difference and why many people will try to be more precise with their terminology.
