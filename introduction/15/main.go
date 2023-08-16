package main

import "fmt"

func main() {

	a := 10

	var t *int = &a

	*t = 20

	b := &a

	*b = 30
	fmt.Println(t)
	fmt.Println(a)
	fmt.Println(*b)

}
