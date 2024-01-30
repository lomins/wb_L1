package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(makeSetArr(strs))
	fmt.Println(makeSetMap(strs))
}

// Создание сета без помощи Map
func makeSetArr(strs []string) []string {
	res := []string{}
	slices.Sort(strs)

	var prevStr string
	for _, curStr := range strs {
		if curStr != prevStr {
			res = append(res, curStr)
			prevStr = curStr
		}
	}

	return res
}

// Создание сета с помощью Map
func makeSetMap(strs []string) []string {
	res := []string{}
	m := make(map[string]bool)

	for _, str := range strs {
		m[str] = true
	}

	for str := range m {
		res = append(res, str)
	}

	return res
}
