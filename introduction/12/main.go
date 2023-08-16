package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	// Endereco
	Adress Endereco
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

func main() {

	felipe := Cliente{
		Nome:  "Felipe",
		Idade: 25,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", felipe.Nome, felipe.Idade, felipe.Ativo)

}
