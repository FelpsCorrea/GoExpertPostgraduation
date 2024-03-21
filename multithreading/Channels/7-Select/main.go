package main

import "time"

type Message struct {
	id  int
	Msg string
}

func main() {

	c1 := make(chan Message)
	c2 := make(chan Message)

	go func() {
		msg := Message{1, "Hello from RabbitMQ"}
		time.Sleep(time.Second)
		c1 <- msg
	}()
	go func() {
		msg := Message{1, "Hello from Kafka"}
		time.Sleep(time.Second * 2)
		c2 <- msg
	}()

	for {
		select {
		case msg1 := <-c1: //rabbitmq
			println("received", msg1)
		case msg2 := <-c2: //kafka
			println("received", msg2)
		case <-time.After(5 * time.Second):
		}
	}

}
