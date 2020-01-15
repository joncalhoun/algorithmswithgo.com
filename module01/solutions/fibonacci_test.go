package module01solutions

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{14, 377},
		{22, 17711},
		{25, 75025},
		// This test case may be much slower depending on your
		// solution. We will look at how to speed it up in a future
		// module. Feel free to comment it out if it is too slow.
		{41, 165580141},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("n=%v", tc.n), func(t *testing.T) {
			got := Fibonacci(tc.n)
			if got != tc.want {
				t.Fatalf("Fibonacci() = %v; want %v", got, tc.want)
			}
		})
	}
}
