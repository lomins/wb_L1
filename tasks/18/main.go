package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	val int
	mu  sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	c := &Counter{val: 0}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			c.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(c.val)
}
