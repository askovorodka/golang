package main

import (
	"fmt"
	"time"
	"math/rand"
)

func faa(from int) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func f(n int) {
	for i:=0; i< 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main()  {

	/*fmt.Println("test", 1, 2)
	for i := 0; i < 10; i++ {
		go faa(i)
	}*/

	/*go faa("direct")
	go faa("goroutine")
	go func(from string) {
		fmt.Println(from)
	}("going")*/

	for i := 0; i < 10; i++  {
		go faa(i)
	}

	var input string
	fmt.Scanln(&input)

	fmt.Println("done")

}
