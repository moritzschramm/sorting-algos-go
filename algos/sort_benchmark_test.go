package algos

import (
	"github.com/moritzschramm/sorting-algos-go/generator"
	"testing"
)

const (
	LENGTH = 1e4
	RAND   = 1e5
)

func BenchmarkInsertionSort(b *testing.B) {

	a := make([][]int, b.N)

	for i := 0; i < b.N; i++ {
		a[i] = generator.Generate(LENGTH, RAND)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		InsertionSort(a[i])
	}
}

func BenchmarkMergeSort(b *testing.B) {

	a := make([][]int, b.N)

	for i := 0; i < b.N; i++ {
		a[i] = generator.Generate(LENGTH, RAND)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		MergeSort(a[i])
	}
}


func BenchmarkQuickSort(b *testing.B) {

	a := make([][]int, b.N)

	for i := 0; i < b.N; i++ {
		a[i] = generator.Generate(LENGTH, RAND)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		QuickSort(a[i])
	}
}

func BenchmarkQuickSortRand(b *testing.B) {

	a := make([][]int, b.N)

	for i := 0; i < b.N; i++ {
		a[i] = generator.Generate(LENGTH, RAND)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		QuickSortRand(a[i])
	}
}