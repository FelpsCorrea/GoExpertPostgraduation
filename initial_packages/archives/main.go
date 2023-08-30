package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Criar arquivo
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	// Escrever String
	// tamanho, err := f.WriteString("Hello, World!")

	// Escrever Bytes
	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)

	f.Close()

	// Abrir o arquivo
	// arquivo, err := os.Open("arquivo.txt")

	// Ler o arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Lê o conteúdo que está aberto no arquivo
	reader := bufio.NewReader(arquivo2)

	// De quanto em quanto o reader vai ler
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	arquivo2.Close()

	// Remover um arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
