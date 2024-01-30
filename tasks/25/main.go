package main

import (
	"fmt"
	"time"
)

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println("Спим...")

	Sleep(5)

	fmt.Println("Проснулись!")
}
