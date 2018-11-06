package channel

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		messages <- "hi"
	}()

	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")
	}

	msg := <-messages
	select {
	case messages <- msg:
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message received")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
