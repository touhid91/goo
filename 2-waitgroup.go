package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	wg sync.WaitGroup
)

func speak(args []string) {
	for _, v := range args {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(v)
		wg.Done()
	}
}

func main() {
	command := []string{"The", "Phoenix", "Lament"}
	wg.Add(len(command))
	go speak(command)
	wg.Wait()
}
