package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	count := make(chan int)

	wg.Add(2)

	go player("one", count)
	go player("two", count)

	count <- 1

	wg.Wait()
	fmt.Printf("end")
}

func player(name string, count chan int) {
	defer wg.Done()
	for {
		ball, ok := <-count
		if !ok {
			fmt.Printf("Player %s won", name)
			return
		}

		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s miss\n", name)
			// 关闭通道
			close(count)
			return
		}

		fmt.Printf("Player %s hit %d \n", name, ball)

		ball++
		count <- ball
	}
}
