package main

import (
	"fmt"
	"time"
)

func main()  {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <- jobs
			if more {
				fmt.Println("recieved job", j, time.Now() )
			} else {
				fmt.Println("recieved all jobs")
				done <- true
				return
			}

		}
	}()

	for j:=1; j<=5; j++ {
		jobs<-j
		fmt.Println("send job", j, time.Now() )
	}

	close(jobs)
	fmt.Println("send all jobs")
	<-done

}
