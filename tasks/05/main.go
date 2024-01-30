package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type MyChannel struct {
	ch       chan string
	isClosed bool
	mu       sync.Mutex
}

func main() {
	var seconds int
	flag.IntVar(&seconds, "s", 1, "количество секунд работы программы. дефолт: 10")
	flag.Parse()
	fmt.Println("Программа завершится через", seconds, "секунд.")

	var wg sync.WaitGroup

	mc := &MyChannel{ch: make(chan string)}

	wg.Add(4)

	go writer(mc, &wg)
	go writer(mc, &wg)

	go reader(mc, &wg)
	go reader(mc, &wg)

	time.Sleep(time.Second * time.Duration(seconds))
	fmt.Println("\n", "main: время вышло")

	mc.SafeClose()

	wg.Wait()
}

func writer(mc *MyChannel, wg *sync.WaitGroup) {
	fmt.Println("writer start write")
	defer wg.Done()
	for {
		if mc.isClosed {
			fmt.Println("канал закрыт, остановка writer...")
			return
		}
		mc.ch <- "-"
		time.Sleep(100 * time.Millisecond)
	}
}

func reader(mc *MyChannel, wg *sync.WaitGroup) {
	fmt.Println("reader start read")
	defer wg.Done()
	for data := range mc.ch {
		fmt.Print(data)
	}
	fmt.Println("из канала больше нечего читать, остановка reader...")
}

func (mc *MyChannel) SafeClose() {
	fmt.Println("Закрываю канал...")
	mc.mu.Lock()
	if !mc.isClosed {
		close(mc.ch)
		mc.isClosed = true
	}
	mc.mu.Unlock()
}
