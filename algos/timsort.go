package algos

import "fmt"

const (
	RUN = 32
)

func TimSort(a []int) {

	length := len(a)

	fmt.Printf("LENGTH: %d\n", length)

	for i := 0; i < len(a); i += RUN {

		insertionSort(a, i, min((i + RUN), length))
	}

	for size := RUN; size < length; size = 2*size {

		for left := 0; left < length; left += 2*size {

			right := min((left + 2*size - 1), (length - 1))
			mid := left + (right - left)/2
			fmt.Printf("size: %d\n", size)
			fmt.Printf("len(a): %d\n", length)
			fmt.Printf("left: %d\n", left)
			fmt.Printf("mid: %d\n", mid)
			fmt.Printf("right: %d\n", right)
			fmt.Println("------------------")

			mergeTim(a, left, mid, right)
		}
	}
}

func mergeTim(a []int, l, m, r int) {

	len1 := m - l + 1 
	len2 := r - m

	left := make([]int, len1)
	right := make([]int, max(len2, 0))

	for i := range left {
		left[i] = a[l + i]
	}
	for i := range right {
		right[i] = a[m + 1 + i]
	}

	i := 0
	j := 0
	k := 0

	for i < len1 && j < len2 {


		if(left[i] <= right[j]) {

			a[k] = left[i]
			i++

		} else {

			a[k] = right[j]
			j++
		}
		k++
	}

	for i < len1 {

		a[k] = left[i]
		k++
		i++
	}
	for j < len2 {

		a[k] = right[j]
		k++
		j++
	}
}

func min(left, right int) int {

	if left <= right {
		return left
	} else {
		return right
	}
}
func max(left, right int) int {

	if left >= right {
		return left
	} else {
		return right
	}
}