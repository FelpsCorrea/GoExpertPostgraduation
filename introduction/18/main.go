package main

import "fmt"

func main() {
	a := "ABACATE"
	b := 30
	showType(a)
	showType(b)
}

func showType(x interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", x, x)
}
