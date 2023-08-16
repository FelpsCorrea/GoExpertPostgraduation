package main

import "fmt"

type Pessoa interface {
	Desativar(x int) int
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
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

	felipe.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", felipe.Nome, felipe.Idade, felipe.Ativo)

}
