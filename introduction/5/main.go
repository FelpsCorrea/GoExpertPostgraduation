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
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	fmt.Println(len(meuArray) - 1)
	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O valor do %d é %d\n", i, v)
	}
}
