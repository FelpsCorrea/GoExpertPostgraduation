package main

import "fmt"

func MudarValor(a *int) {
	*a = 10
}

func main() {

	a := 10

	var t *int = &a

	*t = 20

	b := &a

	*b = 30

	MudarValor(&a)
	fmt.Println(t)
	fmt.Println(a)
	fmt.Println(*b)

}
