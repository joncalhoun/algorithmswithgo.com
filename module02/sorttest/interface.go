package sorttest

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// TestInterface is a helper for testing functions that sort using
// sort.Interface.
func TestInterface(t *testing.T, sortFn func(sort.Interface)) {
	seed := time.Now().UnixNano()
	t.Log("Seed for random cases:", seed)
	rand.Seed(seed)

	for name, list := range map[string][]int{
		"sorted":          []int{1, 2, 3, 4},
		"reverse":         []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		"duplicates":      []int{3, 5, 3, 5, 3, 5},
		"random-len10":    rand.Perm(10),
		"random-len20":    rand.Perm(20),
		"random-len50":    rand.Perm(50),
		"random-len100":   rand.Perm(100),
		"random-len1000":  rand.Perm(1000),
		"sorted-len10000": intList(10000),
	} {
		t.Run(name, func(t *testing.T) {
			want := make([]int, len(list))
			for i, val := range list {
				want[i] = val
			}
			sort.Ints(want)
			sortFn(sort.IntSlice(list))
			errorCount := 0
			if len(list) != len(want) {
				t.Fatalf("got len %d; want %d", len(list), len(want))
			}
			for i := 0; i < len(want); i++ {
				if errorCount >= 5 {
					t.Fatalf("additional errors omitted for brevity")
				}
				if list[i] != want[i] {
					errorCount++
					t.Errorf("list[%d] = %d; want %d", i, list[i], want[i])
				}
			}
		})
	}
}

// BenchmarkInterface is a helper for benchmarking sort functions that sort
// using the sort.Interface.
func BenchmarkInterface(b *testing.B, sortFn func(sort.Interface)) {
	for _, size := range []int{
		100, 200, 400, 800, 1600, 3200,
	} {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				b.StopTimer()
				list := rand.Perm(size)
				b.StartTimer()
				sortFn(sort.IntSlice(list))
			}
		})
	}
}
