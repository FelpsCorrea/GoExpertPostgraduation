package main

import (
	"fmt"
)

func main() {

	valor := sum(1, 2, 3, 4, 5, 5)

	fmt.Println(valor)
}

func sum(numeros ...int) int {
	total := 0

	for _, valor := range numeros {
		total += valor
	}

	return total
}
