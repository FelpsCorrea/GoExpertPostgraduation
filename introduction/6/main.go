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

	s := []int{1, 2, 3, 4, 5}
	// Capacidade = 5
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = append(s, 4)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// Capacidade = 10

	s = append(s, 10)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// Capacidade = 10

	itemsToAdd := []int{4, 5, 6, 7, 8}
	s = append(s, itemsToAdd...)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// Capacidade = 20

	s = append(s, itemsToAdd...)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// Capacidade = 20

	s = append(s, itemsToAdd...)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// Capacidade = 40

}
