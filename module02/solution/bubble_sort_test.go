package module02_test

import (
	"algo/module02/sorttest"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	sorttest.TestInt(t, BubbleSortInt)
}

func BenchmarkBubbleSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, BubbleSortInt)

	// Below is the code used in the video. I have opted to use a single function
	// in the sorttest package for all benchmarks moving forward to make it easier
	// to benchmark all of the sort functions we cover.
	//
	// for _, size := range []int{
	// 	100, 200, 400, 800, 1600,
	// } {
	// 	b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
	// 		for n := 0; n < b.N; n++ {
	// 			b.StopTimer()
	// 			list := rand.Perm(size)
	// 			b.StartTimer()
	// 			BubbleSortInt(list)
	// 		}
	// 	})
	// }
}

func TestBubbleSortString(t *testing.T) {
	sorttest.TestString(t, BubbleSortString)
}

func TestBubbleSortInterface(t *testing.T) {
	sorttest.TestInterface(t, BubbleSort)
}

func BenchmarkBubbleSortInterface(b *testing.B) {
	sorttest.BenchmarkInterface(b, BubbleSort)
}

func TestBubbleSortPerson(t *testing.T) {
	testPeople(t, BubbleSortPerson)
}
