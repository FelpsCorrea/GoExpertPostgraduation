package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go task("A")
	go task("B")

	time.Sleep(time.Second * 11)
}
