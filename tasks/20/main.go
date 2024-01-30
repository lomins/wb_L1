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
	strSlice := strings.Split(str, " ")

	slices.Reverse(strSlice)

	res := strings.Join(strSlice, " ")

	fmt.Println(res)
}
