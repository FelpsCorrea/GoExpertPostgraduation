package main

import (
	"curso-go-19/matematica"
	"fmt"
)

func main() {
	m := map[string]int{"A": 10, "B": 100, "C": 1000}
	m2 := map[string]float64{"A": 10.5, "B": 100.5, "C": 1000.5}

	fmt.Println(matematica.Soma(m))
	fmt.Println(matematica.Soma(m2))

	// m3 := map[string]string{"A": "10.5", "B": "100.5", "C": "1000.5"}
}
