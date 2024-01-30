package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	stopChannel1 := make(chan bool)
	go chanGoroutine(stopChannel1)
	time.Sleep(3 * time.Second)
	stopChannel1 <- true
	time.Sleep(time.Second)

	fmt.Println("chanGoroutine end work")

	stopChannel2 := make(chan struct{})
	go chanStructGoroutines(stopChannel2)
	time.Sleep(3 * time.Second)
	close(stopChannel2)
	time.Sleep(time.Second)

	fmt.Println("chanStructGoroutines end work")

	ctx, cancel := context.WithCancel(context.Background())
	go contextGoroutine(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)

	fmt.Println("contextGoroutine end work")

	fmt.Println("Main Goroutine exited")

}

func chanGoroutine(stopChannel chan bool) {
	for {
		select {
		case <-stopChannel:
			fmt.Println("Stopping chanG...")
			return
		default:
			fmt.Println("My Goroutine is running :( ")
			time.Sleep(time.Second)
		}
	}
}

func chanStructGoroutines(stopChannel chan struct{}) {
	for {
		select {
		case <-stopChannel:
			fmt.Println("Stopping chanStructG")
			return
		default:
			fmt.Println("My Goroutine is running :(")
			time.Sleep(time.Second)
		}
	}
}

func contextGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping ctxG")
			return
		default:
			fmt.Println("My Goroutine is running :(")
			time.Sleep(time.Second)
		}
	}
}
