package main

import "fmt"

func main() {
	a, b := 5, 10

	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(num1, num2 *int) {
	*num1, *num2 = *num2, *num1
}
