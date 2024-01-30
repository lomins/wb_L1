package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Введите строку: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	str := strings.ReplaceAll(text, "\r\n", "")

	runes := []rune(str)
	slices.Reverse(runes)

	fmt.Println(string(runes))
}
