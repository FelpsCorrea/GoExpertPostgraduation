package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"` // ou `json:"-"` para não disponibilizar como json essa propriedade
}

func main() {

	conta := Conta{Numero: 1, Saldo: 100}

	// Transformar struct para json
	res, err := json.Marshal(conta)

	if err != nil {
		println(err)
	}

	println(string(res))

	// encoder := json.NewEncoder(os.Stdout)
	// encoder.Encode(conta)

	// Transformar struct em json
	err = json.NewEncoder(os.Stdout).Encode(conta)

	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"numero": 2,"saldo":200}`)
	var contaX Conta

	// Transformar json em struct
	err = json.Unmarshal(jsonPuro, &contaX)

	if err != nil {
		println(err)
	}

	println(contaX.Saldo)
}
