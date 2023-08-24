package main

import "fmt"

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}

	return soma
}

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

type Number interface {
	int | float64
}

func Soma2[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}
func main() {
	m := map[string]int{"A": 10, "B": 100, "C": 1000}
	fmt.Println(SomaInteiro(m))

	m2 := map[string]float64{"A": 10.5, "B": 100.5, "C": 1000.5}
	fmt.Println(SomaFloat(m2))

	fmt.Println(Soma(m))
	fmt.Println(Soma(m2))

	// m3 := map[string]string{"A": "10.5", "B": "100.5", "C": "1000.5"}
}
