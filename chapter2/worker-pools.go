package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan <- int)  {
	for j := range jobs  {
		fmt.Println("working:", id, "processing job:", j)
		time.Sleep(time.Second)
		results<- j * 2
	}
}

func main()  {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j:=1; j <= 7; j++ {
		jobs<- j
	}
	close(jobs)

	for r := 1; r <= 7; r++ {
		<- results
	}

}
