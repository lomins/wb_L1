package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var resultCh = make(chan int)

	var result int
	numbers := []int{2, 4, 6, 8, 10}

	for i := 0; i < len(numbers); i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resultCh <- n * n
		}(numbers[i])
	}

	go func() {
		for n := range resultCh {
			result += n
		}
	}()

	wg.Wait()
	close(resultCh)

	fmt.Println(result)
}
