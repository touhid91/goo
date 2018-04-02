package main

import (
	"sync"
	"fmt"
	"runtime"
)

var (
	wg2 sync.WaitGroup
)

func baz(char rune) {
	defer wg2.Done()

	for c := char; c < char+26; c++ {
		fmt.Printf("%c ", c)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	wg2.Add(2)
	go baz('a')
	go baz('A')
	wg2.Wait()
}
