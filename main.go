package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	// CANONICAL WAY
	// A WaitGroup waits for a collection of goroutines to finish.
	// The main goroutine calls Add to set the number of goroutines to wait for.
	// Then each of the goroutines runs and calls Done when finished.
	// At the same time, Wait can be used to block until all goroutines have finished.
)

func paralPrint(arg ...string) {
	for _, v := range arg {
		// defer wg.Done()
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(time.Now(), v)
		// A defer statement defers the execution of a function until the surrounding function returns.

		// The deferred call's arguments are evaluated immediately,
		// but the function call is not executed until
		// the surrounding function returns.
		// wg.Done()
	}
}

func main() {
	wg.Add(3)
	go paralPrint("The", "Phoenix", "Lament")
	time.Sleep(3001 * time.Millisecond)
	fmt.Println("Execution Complete")
	// wg.Wait()
}
