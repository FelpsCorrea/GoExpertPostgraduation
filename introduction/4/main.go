package main

import "fmt"

type ID int

var (
	b bool
	c int
	d string = "default"
	e ID     = 1
)

func main() {
	a := "abacate"
	println(a)
	println(e)
	fmt.Printf("O tipo de e é %T", e)
}
