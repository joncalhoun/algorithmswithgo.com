package module01solutions

import (
	"fmt"
	"testing"
)

func TestBaseToBase(t *testing.T) {
	tests := []struct {
		val     string
		base    int
		newBase int
		want    string
	}{
		{
			"E", 16, 2,
			"1110",
		}, {
			"11011110101011011011111011101111", 2, 3,
			"100122100210211112102",
		}, {
			"3735928559", 10, 4,
			"3132223123323233",
		}, {
			"8831A383B", 12, 16,
			"DEADBEEF",
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v in base %v to base %v", tc.val, tc.base, tc.newBase), func(t *testing.T) {
			got := BaseToBase(tc.val, tc.base, tc.newBase)
			if got != tc.want {
				t.Fatalf("BaseToBase() = %v; want %v", got, tc.want)
			}
		})
	}
}
