package algos

func HeapSort(a []int) {
	
	// build max heap
	for i := len(a)-1; i >= 0; i-- {
		maxHeapify(a, i, len(a)-1)
	}

	heapSize := len(a) -1

	for i := len(a) -1; i >= 1; i-- {

		tmp := a[0]
		a[0] = a[i]
		a[i] = tmp

		heapSize--

		maxHeapify(a, 0, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {

	largest := i

	if left(i) <= heapSize && a[left(i)] > a[largest] {
		largest = left(i)
	}

	if right(i) <= heapSize && a[right(i)] > a[largest] {
		largest = right(i)
	}

	if largest != i {

		tmp := a[i]
		a[i] = a[largest]
		a[largest] = tmp

		maxHeapify(a, largest, heapSize)
	}
}


func parent(i int) int {

	return int(i/2)
}

func left(i int) int {

	return i*2
}

func right(i int) int {

	return i*2 + 1
}