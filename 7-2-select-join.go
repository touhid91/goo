package main

import "fmt"

func foo2(proceed, halt chan bool) {
	for {
		fmt.Println("next iter")
		select {
		case <-proceed:
			fmt.Println("working hard man!")

		case m := <-halt:
			if !m {
				fmt.Println("false positive, you joking mate?")
				continue
			}
			fmt.Println("all done")
			return
		}
	}
}

func main() {
	a, b := make(chan bool), make(chan bool)
	go foo2(a, b)
	a <- true
	b <- false
	b <- true
}
