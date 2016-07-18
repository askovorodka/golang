package main

import (
	"fmt"
	"time"
)

func main()  {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "result 1"
	}()

	select {
	case res1 := <-c1:
		fmt.Println(res1)
	case <- time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "result 2"
	}()

	select {
	case res2 := <- c2:
		fmt.Println(res2)
	case <- time.After(time.Second * 4):
		fmt.Println("timeout 2")
	}

}
