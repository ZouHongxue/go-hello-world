package tools

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(time.Second * 1)

	<-time1.C
	fmt.Println("Time 1 expired")

	time2 := time.NewTimer(time.Second)

	go func() {
		<-time2.C
		fmt.Println("time 2 expired")
	}()

	stop := time2.Stop()
	if stop {
		fmt.Println("time 2 stopped")
	}
}
