package module02

import (
	"algo/module02/sorttest"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	sorttest.TestInt(t, InsertionStortInt)
}
