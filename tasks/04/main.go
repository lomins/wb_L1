package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var num int
	ch := make(chan struct{})

	flag.IntVar(&num, "w", 10, "Количество воркеров")
	flag.Parse()

	// При нажатии Ctrl + C сигнал о завершении будет
	// передан всем подписчикам контекста
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	for i := 0; i < num; i++ {
		go reader(ch, i)
	}

	writer(ch, ctx)
}

func writer(ch chan struct{}, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nGraceful Shutdown and Stop all goroutines...")
			close(ch)
			return
		default:
			ch <- struct{}{}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func reader(ch chan struct{}, id int) {
	for range ch {
		fmt.Print(id, " ")
	}
}
