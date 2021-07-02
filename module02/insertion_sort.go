package module02

import (
	"sort"
)

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	var sorted []int
	for _, item := range list {
		sorted = insert(sorted, item)
	}

	for i, val := range sorted {
		list[i] = val
	}
	// newList := make([]int, len(list))
	// for _, v := range list {
	// 	fmt.Println(v)
	// 	newList = append(newList, v)
	// }
	// fmt.Printf("%v", newList)
}

func insert(sorted []int, item int) []int {
	for i, sortItem := range sorted {
		if item < sortItem {
			// sorted [:i] + item + sorted[i:]
			// endList := append([]int{item}, sorted[i:]...)
			// sorted = append(sorted[:i], endList...)
			// these two lines above in one below
			//sorted = append(sorted[:i], append([]int{item}, sorted[i:]...)...)
			return append(sorted[:i], append([]int{item}, sorted[i:]...)...)
		}
	}
	// since it is not > add it to the end
	return append(sorted, item)
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
}
