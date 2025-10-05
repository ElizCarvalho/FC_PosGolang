package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	channel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publish(channel)     // Thread 2
	go reader(channel, &wg) // Thread 3
	wg.Wait()
}

func publish(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func reader(channel chan int, wg *sync.WaitGroup) {
	for i := range channel {
		fmt.Printf("Received: %d\n", i)
		wg.Done()
	}
}
