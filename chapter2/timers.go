package main

import (
	"fmt"
	"time"
)

func main()  {
	timer1 := time.NewTimer(time.Second * 1)
	<-timer1.C
	fmt.Println("timer1 expired")

	timer2 := time.NewTimer(time.Second * 5)
	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stopped")
	}
}
