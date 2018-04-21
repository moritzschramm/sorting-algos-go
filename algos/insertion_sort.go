package algos

func InsertionSort(a []int) {

	for i, elem := range a[1:] {			// loop through every element in array, skip first element

		for i >= 0 && a[i] > elem {			// shift (sorted) array elements to find position to insert key

			a[i+1] = a[i]					// a[i+1] is valid (no index out of bounds), since max i is 98 (max index of _sliced_ array a)
			i--
		}

		a[i+1] = elem 						// insert element
	}
}