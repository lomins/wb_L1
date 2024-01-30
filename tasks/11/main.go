package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 6}
	arr2 := []int{0, 2, 5, 6}

	fmt.Println(setIntersection(arr1, arr2))

	arr1 = []int{1, 2, 2, 1}
	arr2 = []int{2, 2}

	fmt.Println(setIntersection(arr1, arr2))

	arr1 = []int{4, 9, 5}
	arr2 = []int{9, 4, 9, 8, 4}

	fmt.Println(setIntersection(arr1, arr2))
}

func setIntersection(arr1, arr2 []int) []int {
	hSet := make(map[int]bool)
	res := []int{}

	for _, num := range arr1 {
		hSet[num] = true
	}

	for _, num := range arr2 {
		if hSet[num] {
			res = append(res, num)
			hSet[num] = false
		}
	}

	return res
}
