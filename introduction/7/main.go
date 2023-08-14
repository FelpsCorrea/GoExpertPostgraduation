package main

import "fmt"

func main() {

	salarios := map[string]int{
		"Felipe": 1000,
		"Maria":  2000,
		"João":   3000,
	}

	fmt.Println(salarios)

	delete(salarios, "Maria")

	salarios["Felipe2"] = 5000

	fmt.Println(salarios)

	salarios2 := make(map[string]int)
	salarios3 := map[string]int{}

	fmt.Println(salarios2)
	fmt.Println(salarios3)

	for nome, salario := range salarios {
		fmt.Println(nome)
		fmt.Println(salario)
	}

	for _, salario := range salarios {
		fmt.Println(salario)
	}
}
