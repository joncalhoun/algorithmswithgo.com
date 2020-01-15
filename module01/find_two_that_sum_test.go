package module01

import (
	"fmt"
	"testing"
)

func TestFindTwoThatSum(t *testing.T) {
	tests := []struct {
		numbers  []int
		sum      int
		possible bool
	}{
		{[]int{1, 2, 3, 4}, 7, true},
		{[]int{0, -1, 1}, 0, true},
		{[]int{0, 1, 1}, 0, false},
		{[]int{10, 1, 12, 3, 7, 2, 2, 1}, 4, true},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v with sum %v", tc.numbers, tc.sum), func(t *testing.T) {
			// append is used to copy the numbers so that even if the
			// FindTwoThatSum code alters the slice (which it shouldn't!)
			// a solution like the following isn't accepted:
			//
			//   numbers[0] = 0
			//   numbers[1] = sum
			//   return 0, 1
			//
			// If this doesn't make sense to you, don't worry about it
			// for now :)
			numbersCopy := append([]int{}, tc.numbers...)
			i, j := FindTwoThatSum(numbersCopy, tc.sum)
			if err := intSlicesEqual(numbersCopy, tc.numbers); err != nil {
				t.Fatalf("FindTwoThatSum() altered the numbers list; %v", err)
			}
			if !tc.possible {
				if i != -1 && j != -1 {
					t.Fatalf("FindTwoThatSum() = (%v, %v); want (%v, %v)", i, j, -1, -1)
				}
				return
			}
			got := tc.numbers[i] + tc.numbers[j]
			if got != tc.sum {
				t.Fatalf("FindTwoThatSum() = (%v, %v); sum = %v; want sum %v", i, j, got, tc.sum)
			}
		})
	}
}

func intSlicesEqual(got, want []int) error {
	if len(got) != len(want) {
		return fmt.Errorf("len(got) = %v; len(want) = %v", len(got), len(want))
	}
	for i, gv := range got {
		wv := want[i]
		if gv != wv {
			return fmt.Errorf("got[%v] = %v; want[%v] = %v", i, gv, i, wv)
		}
	}
	return nil
}
