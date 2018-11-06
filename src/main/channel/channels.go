package channel

import "fmt"

func main() {
	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message

	fmt.Println(msg)

	message2 := make(chan string, 2)
	message2 <- "buffered"
	message2 <- "channel"
	fmt.Println(<-message2)
	fmt.Println(<-message2)
}
