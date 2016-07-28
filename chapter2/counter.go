package main

import (
	"fmt"
	"time"
	"sync/atomic"
	"runtime"
)

func main()  {

	var ops uint64 = 0

		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("total:", opsFinal)

}
