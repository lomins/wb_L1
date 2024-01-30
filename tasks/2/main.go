package main

import (
	"fmt"
	"sync"
)

func main() {
	ints := []int{2, 4, 6, 8, 10}

	fmt.Println("Исходный слайс: ", ints)

	fmt.Println("Слайс возведённый в квадрат через mutex: ", SquareSliceMutex(ints))
	fmt.Println("Слайс возведённый в квадрат через channels: ", SquareSliceChannel(ints))
}

func SquareSliceMutex(ints []int) []int {
	var wg sync.WaitGroup
	var mu sync.RWMutex

	res := make([]int, len(ints))

	for i := 0; i < len(ints); i++ {
		wg.Add(1)
		go func(index int, n int) {
			defer wg.Done()
			square := n * n
			mu.Lock()
			res[index] = square
			mu.Unlock()
		}(i, ints[i])
	}
	wg.Wait()

	return res
}

func SquareSliceChannel(ints []int) []int {
	var wg sync.WaitGroup
	res := make([]int, len(ints))
	inputCh := make(chan int, len(ints))
	outputCh := make(chan int, len(ints))

	go func(numbers []int, ch chan<- int) {
		for _, num := range numbers {
			ch <- num
		}
	}(ints, inputCh)

	for i := 0; i < len(ints); i++ {
		wg.Add(1)
		go func(inputCh <-chan int, outputCh chan<- int) {
			num := <-inputCh
			square := num * num
			outputCh <- square
			wg.Done()
		}(inputCh, outputCh)
	}
	wg.Wait()

	close(inputCh)

	for i := 0; i < len(res); i++ {
		res[i] = <-outputCh
	}

	return res
}
