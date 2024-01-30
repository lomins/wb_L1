package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	mu      sync.RWMutex
	numbers map[int]int
}

func main() {
	m := MyMap{
		mu:      sync.RWMutex{},
		numbers: make(map[int]int),
	}

	var wg sync.WaitGroup

	numsOfGoroutines := 10

	for i := 0; i < numsOfGoroutines; i++ {
		wg.Add(1)
		go func(num int) {
			m.Add(num, num)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(m.numbers)
}

func (m *MyMap) Add(key int, val int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.numbers[key] = val
}
