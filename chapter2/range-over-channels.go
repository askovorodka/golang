package main

import (
	"fmt"
)

func main()  {
	channel1 := make(chan string, 2)

	channel1<- "one"
	channel1<- "two"

	close(channel1)
	for value := range channel1{
		fmt.Println(value)
	}
}

