package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int)
	go publish(ch)

	// Fora de thread para o programa não morrer
	reader(ch)
}

func reader(ch chan int) {
	for msg := range ch {
		fmt.Printf("Received: %d\n", msg)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	// Sinaliza que o canal está fechado
	close(ch)
}
