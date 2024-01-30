package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	fmt.Println(binarySearch(nums, 3))
	fmt.Println(binarySearch(nums, 8))
}

func binarySearch(a []int, target int) int {
	l, r := 0, len(a)-1

	for l <= r {
		mid := (r + l) / 2

		switch {
		case a[mid] == target:
			return mid
		case a[mid] > target:
			r = mid - 1
		case a[mid] < target:
			l = mid + 1

		}
	}
	return -1
}
