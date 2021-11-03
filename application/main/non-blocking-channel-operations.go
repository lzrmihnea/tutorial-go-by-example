package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message 1", msg)
	default:
		fmt.Println("no message received 1")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message 2", msg)
	default:
		fmt.Println("no message sent 2")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message 3", msg)
	case sig := <-signals:
		fmt.Println("received signal 3", sig)
	default:
		fmt.Println("no activity 3")
	}
}
