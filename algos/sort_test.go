package algos

import (
	"testing"
	"github.com/moritzschramm/sorting-algos-go/generator"
	"sort"
)

const (
	TESTING_LENGTH = 20
	TESTING_RAND = 40
)

func TestInsertionSort(t *testing.T) {

	for i := 0; i < 10; i++ {

		a := generator.Generate(TESTING_LENGTH, TESTING_RAND)
		expected := make([]int, len(a))
		copy(expected, a)
		sort.Ints(expected)
		InsertionSort(a)

		if !slicesEqual(a, expected) {

			t.Errorf("Insertion sort: expected %v, actual: %v", expected, a)
		}
	}
}

func TestMergeSort(t *testing.T) {

	for i := 0; i < 10; i++ {

		a := generator.Generate(TESTING_LENGTH, TESTING_RAND)
		expected := make([]int, len(a))
		copy(expected, a)
		sort.Ints(expected)
		MergeSort(a)

		if !slicesEqual(a, expected) {

			t.Errorf("Merge sort: expected %v, actual: %v", expected, a)
		}
	}
}

func TestQuickSort(t *testing.T) {

	for i := 0; i < 10; i++ {

		a := generator.Generate(TESTING_LENGTH, TESTING_RAND)
		expected := make([]int, len(a))
		copy(expected, a)
		sort.Ints(expected)
		QuickSort(a)

		if !slicesEqual(a, expected) {

			t.Errorf("Quicksort: expected %v, actual: %v", expected, a)
		}
	}
}

func TestQuickSortRand(t *testing.T) {

	for i := 0; i < 10; i++ {

		a := generator.Generate(TESTING_LENGTH, TESTING_RAND)
		expected := make([]int, len(a))
		copy(expected, a)
		sort.Ints(expected)
		QuickSortRand(a)

		if !slicesEqual(a, expected) {

			t.Errorf("Quicksort (with random pivot): expected %v, actual: %v", expected, a)
		}
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}