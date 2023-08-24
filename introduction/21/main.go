package main

func main() {

	// For básico
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"1", "2", "3", "4", "5"}

	// Percorrer listas
	for indice, valor := range numeros {
		println(indice, valor)
	}

	// While disfarçado
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// Loop infinito
	for {

	}

}
