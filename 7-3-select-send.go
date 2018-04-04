package main

import "fmt"

func fib(next, quit chan int) {
	p, q := 0, 1
	for {
		select {
		case next <- p:
			p, q = q, p+q
		case <-quit:
			return
		}
	}
}

func main() {

	f, q := make(chan int), make(chan int)

	go fib(f, q)

	for i := 0; i < 10; i++ {
		fmt.Println("iter", i, <-f)
	}
	q <- 1e1
}
