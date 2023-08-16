package main

import (
	"fmt"
)

func main() {

	total := func() int {
		return sum(10, 20, 30, 40)
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0

	for _, valor := range numeros {
		total += valor
	}

	return total
}
