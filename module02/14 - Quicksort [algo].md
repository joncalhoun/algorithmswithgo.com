<!-- Refs:

"What is the advantage of randomized Quicksort?"

  If the input array is distributed uniformly at random then (as you noted) there is no difference between always picking an element at a fixed position (for example the middle one as you suggest) or picking an element chosen at random.

  If however your input array is not really in random order (which happens to be the case in almost all practical scenarios) then one needs to either "preshufle" the array in order for the elements in it to be placed in random order, or (equivalently) always take a random element as a pivot. This ensures the partitioning phase of quicksort partitions the arrays into sub-arrays of almost equal size and hence that the expected running time remains 𝑂(𝑛log𝑛)

- https://cs.stackexchange.com/questions/7582/what-is-the-advantage-of-randomized-quicksort

-->
