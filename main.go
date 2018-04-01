package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	fmt.Println("arg", a)
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	fmt.Println(<-c, <-c)
}
