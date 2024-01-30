package main

import "fmt"

func main() {
	nums := []int{4, 5, 3, 2, 1}

	fmt.Println(quickSort(nums))
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right, pivot := 0, len(a)-1, 0

	a[pivot], a[right] = a[right], a[pivot]

	for i := 0; i < right; i++ {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
