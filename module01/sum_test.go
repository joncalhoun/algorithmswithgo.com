package module01

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		numbers []int
		want    int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{}, 0},
		{[]int{-1, -1}, -2},
		{[]int{-1, -1, 0, 0, 0}, -2},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, 0, 1}, 0},
		{[]int{1, 0, -1}, 0},
		{[]int{43, 67, 65, 42, 74, 95, 100, 21}, 507},
		{[]int{26, 26, 94, 55, 71, 57, 7, 72}, 408},
		{[]int{78, 11, 48, 91, 13, 26, 74, 99}, 440},
		{[]int{31, 96, 10, 55, 41, 13, 92, 5}, 343},
		{[]int{14, 39, 18, 39, 3, 29, 8, 29}, 179},
		{[]int{-95, -46, -65, -63, 10}, -259},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("(%v)", tc.numbers), func(t *testing.T) {
			got := Sum(tc.numbers)
			if got != tc.want {
				t.Fatalf("Sum() = %v; want %v", got, tc.want)
			}
		})
	}
}
