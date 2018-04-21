package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	LENGTH = 1e2
	MAX_RAND = 1e3
)

func main() {

	rand.Seed(time.Now().UnixNano())
	a := make([]int, LENGTH)

	for i := 0; i < LENGTH; i++ {			// initialize slice with random values
		a[i] = rand.Intn(MAX_RAND)
	}


	fmt.Println(a)
	fmt.Println("---------------------------------------------------")

	start := time.Now()
	insertionSort(a)
	elapsed := time.Since(start)

	fmt.Println(a)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func insertionSort(a []int) {

	for i, elem := range a[1:] {			// loop through every element in array, skip first element

		for i >= 0 && a[i] > elem {			// shift (sorted) array elements to find position to insert key

			a[i+1] = a[i]					// a[i+1] is valid (no index out of bounds), since max i is 98 (max index of _sliced_ array a)
			i--
		}

		a[i+1] = elem 						// insert element
	}
}