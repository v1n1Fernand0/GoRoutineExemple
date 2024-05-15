package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d recived %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	canal := make(chan int)

	go worker(1, canal)
	go worker(2, canal)

	for i := 0; i <= 100; i++ {
		canal <- i
	}
}
