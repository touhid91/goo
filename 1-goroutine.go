package main

import (
	"fmt"
)

func foo() {
	fmt.Println("ZZZ")
}

func main() {
	foo()
	go foo()
}
