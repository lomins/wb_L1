package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //массив для перебора
	chW := make(chan int)                        //канал который принимает значения массива
	chR := make(chan int)

	go writer(chW, chR)
	go reader(chR)

	for _, val := range nums {
		chW <- val
	}
}

func writer(chW, chR chan int) {
	for {
		select {
		case val := <-chW:
			chR <- val * val
		default:
			fmt.Println("No values in the array")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func reader(chR chan int) {
	for {
		select {
		case val := <-chR:
			fmt.Println(val, " ")
		default:
			fmt.Println("No values in read channel")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
