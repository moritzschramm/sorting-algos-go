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

	for i := 0; i < LENGTH; i++ {			// populate slice with random values
		a[i] = rand.Intn(MAX_RAND)
	}


	fmt.Println(a)
	fmt.Println("---------------------------------------------------")

	start := time.Now()
	insertion_sort(a)
	elapsed := time.Since(start)

	fmt.Println(a)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func insertion_sort(a []int) {

	for i, elem := range a {				// loop through every element in array

		if i == 0 {
			continue						// skip first element (A[0 - 0] already sorted)
		}

		j := i - 1
		for j >= 0 && a[j] > elem {

			a[j+1] = a[j]					// shift (sorted) array elements to find position to insert key
			j--
		}

		a[j+1] = elem 						// insert element
	}
}