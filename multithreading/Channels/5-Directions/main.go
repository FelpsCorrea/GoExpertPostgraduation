package main

import (
	"fmt"
)

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
}

// chan<- : canal somente para escrita (escrever no canal)
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// <-chan : canal somente para leitura (ler do canal)
func ler(data <-chan string) {
	fmt.Println("Recebido: ", <-data)
}
