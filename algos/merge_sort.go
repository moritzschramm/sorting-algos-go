package algos

func MergeSort(a []int) {

	mergeSortRec(a, 0, len(a))
}

func mergeSortRec(a []int, left, right int) {

	mid := (left+right)/2			// find middle
	if left < mid {
		mergeSortRec(a, left, mid)		// sort left slice recursively
		mergeSortRec(a, mid, right)	// sort right slice recursively
		merge(a, left, mid, right)		// merge both slices
	}
}

func merge(a []int, left, mid, right int) {

	n1 := mid - left	 				// length of left slice
	n2 := right - mid					// length of right slice

	l := make([]int, n1)				// create left and right slices
	r := make([]int, n2)

	for i := 0; i < n1; i++ {			// fill slices
		l[i] = a[left + i]				// left slice: [left - mid)
	}
	for i := 0; i < n2; i++ {
		r[i] = a[mid + i]				// right slice: [mid - right]
	}

	i, j := 0, 0
	for k := left; k < right; k++ {

		if i >= n1 {
			a[k] = r[j]
			j++
			continue
		} else if j >= n2 {
			a[k] = l[i]
			i++
			continue
		}

		if l[i] <= r[j] {
			a[k] = l[i]	
			i++
		} else {
			a[k] = r[j]
			j++
		}
	}
}