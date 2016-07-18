package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {

	/*messages := make(chan string, 1)

	for {
		go func() { messages <- "one" }()
		go func() { messages <- "two" }()

		time.Sleep(time.Second)

		fmt.Println(<-messages)
	}*/

	channel1 := make(chan string, 5)

	rand.Seed(42)

	go func() {
		for {
			channel1 <- "string1"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			channel1 <- "string2"
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			channel1 <- "string3"
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			channel1 <- "string4"
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			channel1 <- "string5"
			time.Sleep(time.Second * 1)
		}
	}()


	for {

		fmt.Println(<- channel1)
		time.Sleep(time.Second * 1)

	}

	var input string
	fmt.Scanln(&input)

}
