package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var (
	wg5 sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg5.Add(3) // num of workers
	pipeline := make(chan string)

	go assign(pipeline, "Manna")
	go assign(pipeline, "Rubel")
	go assign(pipeline, "Joshim")

	pipeline <- "Baba keno chakor"
	pipeline <- "Dhil mari tor tiner chaale"
	pipeline <- "1 takar fokir"
	pipeline <- "Jomilar kanna"
	pipeline <- "Korimoner hashi(evil grin)"

	close(pipeline)

	wg5.Wait()
}

func assign(pipeline chan string, nick string) {
	defer wg5.Done()

	for {
		task, ok := <-pipeline

		if !ok {
			fmt.Println("KIL>", nick, "died in the process, RIP")
			return
		}

		fmt.Println("ACK>", nick, "got", task)

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		fmt.Println("END>", nick, "completed task '", task, "'")
	}

}
