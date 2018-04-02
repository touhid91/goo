package main

import (
	"sync"
	"runtime"
	"fmt"
)

// package level variables
var (
	counter int64
	wg3     sync.WaitGroup
)

func inc(id int) {
	defer wg3.Done()

	for i := 0; i < 2; i++ {
		c := counter
		runtime.Gosched()
		c++
		counter = c
		//atomic.AddInt64(&counter, 1)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	wg3.Add(2)

	go inc(1)
	go inc(2)

	wg3.Wait()
	fmt.Println("final:", counter)
}
