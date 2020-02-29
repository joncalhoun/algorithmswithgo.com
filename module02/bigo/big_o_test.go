package bigo

import (
	"fmt"
	"math/rand"
	"testing"
)

// All of the code in this file is used when explaining Big O and viewing
// benchmarks to help illustrate a few points.

// Add is O(1) - constant time.
func TestAdd(t *testing.T) {
	for _, tc := range []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{10, 20, 30},
		{-5, 6, 1},
	} {
		t.Run(fmt.Sprintf("%d+%d", tc.a, tc.b), func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("Add() = %d; want %d", got, tc.want)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for _, tc := range []struct {
		a, b int
	}{
		{1, 2},
		{300, 250},
		{10000, 20000},
	} {
		b.Run(fmt.Sprintf("%d+%d", tc.a, tc.b), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Add(tc.a, tc.b)
			}
		})
	}
}

func TestSumToMax(t *testing.T) {
	testSumToMax(t, SumToMax)
}

func benchmarkSumFn(b *testing.B, sumFn func(int) int) {
	for _, max := range []int{
		10, 100, 1000, 10000,
	} {
		b.Run(fmt.Sprintf("%d", max), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				sumFn(max)
			}
		})
	}
}

func BenchmarkSumToMax(b *testing.B) { benchmarkSumFn(b, SumToMax) }

func TestSumToMaxV2(t *testing.T) {
	testSumToMax(t, SumToMaxV2)
}

func BenchmarkSumToMaxV2(b *testing.B) { benchmarkSumFn(b, SumToMaxV2) }

func testSumToMax(t *testing.T, sumFn func(n int) int) {
	for _, tc := range []struct {
		n, want int
	}{
		{1, 1},
		{5, 15},
		{8, 36},
		{100, 5050},
	} {
		t.Run(fmt.Sprintf("%d", tc.n), func(t *testing.T) {
			got := sumFn(tc.n)
			if got != tc.want {
				t.Fatalf("got %d; want %d", got, tc.want)
			}
		})
	}
}

func BenchmarkPerm(b *testing.B) {
	for _, n := range []int{
		10, 100, 1000, 10000,
	} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rand.Perm(n)
			}
		})
	}
}

func TestSumVals(t *testing.T) {
	for i, tc := range []struct {
		vals []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{10, -2, 23, -19}, 12},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 36},
		{rand.Perm(100), 4950},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := SumVals(tc.vals)
			if got != tc.want {
				t.Fatalf("SumVals() = %v; want %v", got, tc.want)
			}
		})
	}
}

func BenchmarkSumVals(b *testing.B) {
	for _, size := range []int{
		4, 16, 64, 256, 1024, 4096, 16384,
	} {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			var list []int
			for i := 0; i < size; i++ {
				list = append(list, rand.Int())
			}
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				SumVals(list)
			}
		})
	}
}

func BenchmarkFind(b *testing.B) {
	for _, size := range []int{
		4, 40, 400,
	} {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			seek := rand.Intn(size)
			list := rand.Perm(size)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				Find(list, seek)
			}
		})
	}
}

func TestGrid(t *testing.T) {
	for _, tc := range []struct {
		x, y int
		want string
	}{
		{2, 2, "xo\nxo\n"},
		{3, 3, "xox\noxo\nxox\n"},
		{5, 10, "xoxox\noxoxo\nxoxox\noxoxo\nxoxox\noxoxo\nxoxox\noxoxo\nxoxox\noxoxo\n"},
	} {
		t.Run(fmt.Sprintf("%dx%d", tc.x, tc.y), func(t *testing.T) {
			got := Grid(tc.x, tc.y)
			if got != tc.want {
				t.Fatalf("got %q; want %q", got, tc.want)
			}
		})
	}
}

func BenchmarkGrid(b *testing.B) {
	for _, tc := range []struct {
		x, y int
	}{
		{2, 2},
		{10, 10},
		{25, 25},
		{100, 100},
		// The following three all have the same result for X*Y
		{12500, 4000},
		{100, 500000},
		{500000, 100},
	} {
		b.Run(fmt.Sprintf("%dx%d", tc.x, tc.y), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Grid(tc.x, tc.y)
			}
		})
	}
}

func TestCube(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want string
	}{
		{1, "(0, 0, 0)\n"},
		{2, "(0, 0, 0)\n(0, 0, 1)\n(0, 1, 0)\n(0, 1, 1)\n(1, 0, 0)\n(1, 0, 1)\n(1, 1, 0)\n(1, 1, 1)\n"},
		{3, "(0, 0, 0)\n(0, 0, 1)\n(0, 0, 2)\n(0, 1, 0)\n(0, 1, 1)\n(0, 1, 2)\n(0, 2, 0)\n(0, 2, 1)\n(0, 2, 2)\n(1, 0, 0)\n(1, 0, 1)\n(1, 0, 2)\n(1, 1, 0)\n(1, 1, 1)\n(1, 1, 2)\n(1, 2, 0)\n(1, 2, 1)\n(1, 2, 2)\n(2, 0, 0)\n(2, 0, 1)\n(2, 0, 2)\n(2, 1, 0)\n(2, 1, 1)\n(2, 1, 2)\n(2, 2, 0)\n(2, 2, 1)\n(2, 2, 2)\n"},
	} {
		t.Run(fmt.Sprintf("n=%d", tc.n), func(t *testing.T) {
			got := Cube(tc.n)
			if got != tc.want {
				t.Fatalf("Cube() = %q; want %q", got, tc.want)
			}
		})
	}
}

func BenchmarkCube(b *testing.B) {
	for _, n := range []int{
		1, 2, 4, 8, 16, 32, 64,
	} {
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Cube(n)
			}
		})
	}
}
