package main

import "fmt"

func main() {
	sl1 := []int{1, 2, 3, 4, 5}
	sl2 := []int{1, 2, 3, 4, 5}

	res1 := removeInOrder(sl1, 2)
	res2 := removeWithoutOrder(sl2, 2)

	fmt.Println(res1)
	fmt.Println(res2)
}

func removeInOrder(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func removeWithoutOrder(arr []int, index int) []int {
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}
