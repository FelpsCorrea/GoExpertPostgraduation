package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // Canal Vazio

	// Thread 2
	go func() {
		channel <- "Olá Mundo!" // Canal Cheio
	}()

	// Thread 1
	msg := <-channel // Canal Vazio
	fmt.Println(msg)
}
