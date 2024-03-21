package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int, 3)

	nWorkers := 10

	for i := 0; i < nWorkers; i++ {
		go worker(i, data)
	}

	// Escreve infos no canal
	for i := 0; i < 100; i++ {
		data <- i
	}
}
