package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

var (
	wg4 sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg4.Add(2)

	table := make(chan int)

	// goroutines are locked
	go spawn("Abinash", table)
	go spawn("Arup", table)

	// start and wait
	// unlock goroutines
	table <- 1
	wg4.Wait()
}

func spawn(nick string, table chan int) {
	defer wg4.Done()

	for {
		ball, ok := <-table

		// ball not received
		if !ok {
			fmt.Println(nick, "won")
			return
		}

		// unlucky miss

		if 0 == rand.Intn(100)%13 {
			fmt.Println(nick, "missed ball", ball)
			close(table)
			return
		}

		// hit the ball back
		ball++
		table <- ball
		fmt.Println(nick, "returned back", ball)
	}
}
