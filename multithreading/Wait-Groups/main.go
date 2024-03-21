package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second * 1)
		wg.Done()
	}
}

// Thread 1
func main() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)

	// Thread 2
	go task("A", &waitGroup)

	// Thread 3
	go task("B", &waitGroup)

	// Thread 4
	go func() {
		for i := 0; i < 25; i++ {
			fmt.Println("C", ":", i)
			time.Sleep(time.Second * 1)
			waitGroup.Done()
		}
	}()
}
