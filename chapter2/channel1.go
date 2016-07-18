package main

import "fmt"

func main()  {
	messages := make(chan string)
	//message2 := "test text"

	go func() { messages <- "ping" }()

	fmt.Println(<- messages)

}
