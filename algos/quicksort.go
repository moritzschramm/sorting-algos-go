package algos

import (
	"math/rand"
	"time"
)

// uses rightmost element as pivot
func QuickSort(a []int) {

	quicksortRec(a, 0, len(a)-1)
}

// uses random element as pivot
func QuickSortRand(a []int) {

	rand.Seed(time.Now().UnixNano())
	quicksortRecRand(a, 0, len(a)-1)
}

func quicksortRec(a []int, left, right int) {

	if left < right {

		pivot := partition(a, left, right)
		quicksortRec(a, left, pivot -1)
		quicksortRec(a, pivot + 1, right)
	}
}

func quicksortRecRand(a []int, left, right int) {

	if left < right {

		pivot := randPartition(a, left, right)
		quicksortRec(a, left, pivot -1)
		quicksortRec(a, pivot + 1, right)
	}
}

func randPartition(a []int, left, right int) int {

	pivot := left + rand.Intn(right - left)
	tmp := a[pivot]
	a[pivot] = a[right]
	a[right] = tmp
	return partition(a, left, right)
}

func partition(a []int, left, right int) int {

	pivot := a[right]
	i := left

	for j := left; j < right; j++ {

		if(a[j] <= pivot) {
			if i != j {
	 			tmp := a[j]
				a[j] = a[i]
				a[i] = tmp
			}
			i++
		}
	}
	a[right] = a[i]
	a[i] = pivot

	return i
}