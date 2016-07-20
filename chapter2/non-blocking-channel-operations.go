package main

import (
	"fmt"
)

func main()  {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("recieved message", msg)
		default :
		fmt.Println("no recieved message")
	}


	msg := "hi"
	select {
	case messages <-msg:
		fmt.Println("message send", msg)
	default:
		fmt.Println("no message send")
	}

	select {
	case msg := <-messages:
		fmt.Println("recoeved message", msg)
	case sig := <-signals:
		fmt.Println("recieved signal", sig)
	default:
		fmt.Println("no activity")
	}

}
