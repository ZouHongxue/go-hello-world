package tools

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 3000)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
