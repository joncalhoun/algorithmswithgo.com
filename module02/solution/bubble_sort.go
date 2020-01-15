package module02

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
func BubbleSortInt(list []int) {
	for i := 0; i < len(list); i++ {
		swapped := false
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	} // 3. Go to (1), do this N times where N == len(list)
}

// func sweep(list []int) bool {
// 	swapped := false
// 	for i := 0; i < len(list)-1; i++ { // 100000
// 		if list[i] > list[i+1] {
// 			list[i], list[i+1] = list[i+1], list[i]
// 			swapped = true
// 		}
// 	}
// 	return swapped
// 	// Compare all pairs in the list. If a pair is out of order, swap them.
// }
