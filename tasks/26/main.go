package main

import (
	"fmt"
	"strings"
)

func isUniqueChars(str string) bool {
	str = strings.ToLower(str)
	hSet := make(map[rune]bool)

	for _, c := range str {
		if hSet[c] {
			return false
		}
		hSet[c] = true
	}

	return true
}

func main() {
	fmt.Println(isUniqueChars("abcd"))
	fmt.Println(isUniqueChars("abCdefAaf"))
	fmt.Println(isUniqueChars("aabcd"))

}
