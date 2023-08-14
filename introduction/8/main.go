package main

import (
	"errors"
	"fmt"
)

func main() {

	valor, err := sum4(50, 1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)
}

// Base

// Caso sejam da mesma tipagem
func sum2(a, b int) int {
	return a + b
}

// Caso seja para retornar mais de um valor
func sum3(a, b int) (int, bool) {

	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}

// Caso seja para retornar um erro
func sum4(a, b int) (int, error) {

	if a+b >= 50 {
		return 0, errors.New("A soma é maior do que 50")
	}

	return a + b, nil
}
