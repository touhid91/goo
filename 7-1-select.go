package main

import (
	"time"
	"fmt"
)

func scold(c chan string) {
	time.Sleep(time.Duration(2) * time.Second)
	c <- "stupid duck"
}

func main() {
	c := make(chan string)
	go scold(c)
	select {
	case m := <-c:
		fmt.Println(m)
	}
}
