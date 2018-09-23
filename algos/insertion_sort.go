package algos

func InsertionSort(a []int) {

	insertionSort(a, 0, len(a))
}

func insertionSort(a []int, left, right int) {

	start := left + 1	// increment left boundary by one since the first element can't be moved further to left

	for index, elem := range a[start:right] {

		i := left + index	// get the correct offset, index starts at 0...(right-start)

		for i >= left && a[i] > elem && i+1 < right { // shift (sorted) array elements to find position to insert key

			a[i+1] = a[i]
			i--
		}

		a[i+1] = elem // insert element
	}
}
