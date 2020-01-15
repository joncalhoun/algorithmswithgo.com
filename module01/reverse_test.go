package module01

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"cat", "tac"},
		{"alphabet", "tebahpla"},
		// {"日本語", "語本日"},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.word), func(t *testing.T) {
			got := Reverse(tc.word)
			if got != tc.want {
				t.Fatalf("Reverse() = %v; want %v", got, tc.want)
			}
		})
	}
}
