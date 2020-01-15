package module02

import (
	"algo/module02/sorttest"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	sorttest.TestInt(t, BubbleSortInt)
}
