package algos

import (
	"testing"
	"github.com/moritzschramm/sorting-algos-go/generator"
	"sort"
	"fmt"
)

const (
	TESTING_LENGTH = 10
	TESTING_RAND = 20
)

func TestAlgos(t *testing.T) {

	algoMap := map[string]func([] int) {
		"Insertion sort": InsertionSort,
		"Merge sort": MergeSort,
		"Quicksort": QuickSort,
		"Quicksort (with random pivot)": QuickSortRand,
		"Heapsort": HeapSort,
		//"Timsort:" Timsort,
	}		

	for name, algo := range algoMap {

		for i := 1; i < 11; i++ {

			if i == 1 {

				fmt.Printf("Testing %s...\n", name)
			}

			a := generator.Generate(TESTING_LENGTH*i, TESTING_RAND*i)
			expected := make([]int, len(a))
			copy(expected, a)
			sort.Ints(expected)
			algo(a)

			if !slicesEqual(a, expected) {

				t.Errorf("%s: expected %v, actual: %v", name, expected, a)
			}
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