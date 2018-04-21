package main

import (
	"testing"
	"math/rand"
	"time"
)


func BenchmarkInsertionSort(b *testing.B) {

	rand.Seed(time.Now().UnixNano())
	a := make([][]int, b.N)

	for i := 0; i < b.N; i++ {
		a[i] = make([]int, LENGTH)
		for k := 0; k < LENGTH; k++ {
			a[i][k] = rand.Intn(MAX_RAND)
		}
	}

	b.ResetTimer()
	
    for i := 0; i < b.N; i++ {

        insertionSort(a[i])
    }
}