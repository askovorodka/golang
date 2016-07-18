package main

import "fmt"

func main()  {

	messages := make(chan string, 2)

	messages <- "ping1"
	messages <- "ping2"

	fmt.Println(<- messages)
	fmt.Println(<- messages)


}
