package module02

import (
	"algo/module02/sorttest"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	sorttest.TestInt(t, InsertionSortInt)
}

func BenchmarkInsertionSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, InsertionSortInt)
}

func TestInsertionSortString(t *testing.T) {
	sorttest.TestString(t, InsertionSortString)
}

func TestInsertionSortInterface(t *testing.T) {
	sorttest.TestInterface(t, InsertionSort)
}

func BenchmarkInsertionSortInterface(b *testing.B) {
	sorttest.BenchmarkInterface(b, InsertionSort)
}

func TestInsertionSortPerson(t *testing.T) {
	testPeople(t, InsertionSortPerson)
}
