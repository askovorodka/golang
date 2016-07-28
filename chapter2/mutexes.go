package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"math/rand"
	"runtime"
	"time"
)

func main()  {

	var state = make(map[int]int) // for our example the state will be a map
	var mutex = &sync.Mutex{} // this mutex will synchronize access to state
	var ops uint64 = 0

	for i:=1; i<=100; i++  {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&ops,1)
				runtime.Gosched()
			}
		}()
	}

	for w := 1; w <= 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("total:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

}